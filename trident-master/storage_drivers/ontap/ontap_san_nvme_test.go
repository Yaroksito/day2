// Copyright 2023 NetApp, Inc. All Rights Reserved.

package ontap

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	tridentconfig "github.com/netapp/trident/config"
	mockapi "github.com/netapp/trident/mocks/mock_storage_drivers/mock_ontap"
	"github.com/netapp/trident/storage"
	sa "github.com/netapp/trident/storage_attribute"
	drivers "github.com/netapp/trident/storage_drivers"
	"github.com/netapp/trident/storage_drivers/ontap/api"
	"github.com/netapp/trident/storage_drivers/ontap/awsapi"
	"github.com/netapp/trident/utils"
	"github.com/netapp/trident/utils/errors"
)

var mockIPs = []string{"0.0.0.0", "1.1.1.1"}

func newNVMeDriver(apiOverride api.OntapAPI, awsApiOverride awsapi.AWSAPI, fsxId *string) *NVMeStorageDriver {
	sPrefix := "test_"

	config := &drivers.OntapStorageDriverConfig{}
	config.ManagementLIF = ONTAPTEST_LOCALHOST
	config.SVM = "svm"
	config.Aggregate = "data"
	config.Username = "ontap-san-user"
	config.Password = "password1!"
	config.UseREST = true
	config.CommonStorageDriverConfig = &drivers.CommonStorageDriverConfig{
		DebugTraceFlags:   debugTraceFlags,
		StoragePrefix:     &sPrefix,
		StorageDriverName: "ontap-san",
	}

	if fsxId != nil {
		config.AWSConfig = &drivers.AWSConfig{}
		config.AWSConfig.FSxFilesystemID = *fsxId
	}

	driver := &NVMeStorageDriver{Config: *config, API: apiOverride, AWSAPI: awsApiOverride}
	driver.telemetry = &Telemetry{
		Plugin:        driver.Name(),
		SVM:           config.SVM,
		StoragePrefix: *driver.GetConfig().StoragePrefix,
		Driver:        driver,
	}

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[FileSystemType] = tridentconfig.FsExt4
	driver.virtualPools = map[string]storage.Pool{"pool1": pool1}
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}

	return driver
}

func newNVMeDriverAndMockApiAndAwsApi(t *testing.T) (*NVMeStorageDriver, *mockapi.MockOntapAPI, *mockapi.MockAWSAPI) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	mockAWSAPI := mockapi.NewMockAWSAPI(mockCtrl)
	fsxId := FSX_ID

	return newNVMeDriver(mockAPI, mockAWSAPI, &fsxId), mockAPI, mockAWSAPI
}

func newNVMeDriverAndMockApi(t *testing.T) (*NVMeStorageDriver, *mockapi.MockOntapAPI) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)

	return newNVMeDriver(mockAPI, nil, nil), mockAPI
}

func TestNVMeBackendName(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)

	// Backend name non-empty.
	d.Config.BackendName = "san-nvme-backend"

	assert.Equal(t, d.BackendName(), "san-nvme-backend", "Backend name is not correct.")

	// Empty backend name and no dataLIFs.
	d.Config.BackendName = ""
	d.ips = []string{}

	assert.Equal(t, d.BackendName(), "ontapsan_noLIFs", "Backend name is not correct.")

	// Empty backend name with dataLIFs.
	d.ips = mockIPs

	assert.Equal(t, d.BackendName(), "ontapsan_0.0.0.0", "Backend name is not correct.")
}

func TestNVMeInitialize_ConfigUnmarshalError(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)
	commonConfig := &drivers.CommonStorageDriverConfig{
		// Version:           1,
		StorageDriverName: "ontap-san",
		DriverContext:     tridentconfig.ContextCSI,
		// DebugTraceFlags:   debugTraceFlags,
	}
	configJSON := `{"SANType": }`

	err := d.Initialize(ctx, tridentconfig.ContextCSI, configJSON, commonConfig, nil, BackendUUID)

	assert.ErrorContains(t, err, "could not decode JSON configuration")
}

func TestNVMeInitialize_NVMeNotSupported(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	commonConfig := &drivers.CommonStorageDriverConfig{
		StorageDriverName: "ontap-san",
		DriverContext:     tridentconfig.ContextCSI,
	}
	configJSON := `{"SANType": "nvme"}`
	mAPI.EXPECT().SupportsFeature(ctx, gomock.Any()).Return(false)

	err := d.Initialize(ctx, tridentconfig.ContextCSI, configJSON, commonConfig, nil, BackendUUID)

	assert.ErrorContains(t, err, "ontap doesn't support NVMe")
}

func TestNVMeInitialize_GetDataLifError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	commonConfig := &drivers.CommonStorageDriverConfig{
		StorageDriverName: "ontap-san",
		DriverContext:     tridentconfig.ContextCSI,
	}
	configJSON := `{"SANType": "nvme"}`
	mAPI.EXPECT().SupportsFeature(ctx, gomock.Any()).Return(true)
	mAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, sa.NVMeTransport).Return(nil, fmt.Errorf("error getting dataLifs"))

	err := d.Initialize(ctx, tridentconfig.ContextCSI, configJSON, commonConfig, nil, BackendUUID)

	assert.ErrorContains(t, err, "error getting dataLifs")
}

func TestNVMeInitialize_NoDataLifs(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	commonConfig := &drivers.CommonStorageDriverConfig{
		StorageDriverName: "ontap-san",
		DriverContext:     tridentconfig.ContextCSI,
	}
	configJSON := `{"SANType": "nvme"}`
	mAPI.EXPECT().SupportsFeature(ctx, gomock.Any()).Return(true)
	mAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, sa.NVMeTransport).Return([]string{}, nil)
	mAPI.EXPECT().SVMName().Return("svm")

	err := d.Initialize(ctx, tridentconfig.ContextCSI, configJSON, commonConfig, nil, BackendUUID)

	assert.ErrorContains(t, err, "no NVMe data LIFs found on SVM svm")
}

func TestNVMeInitialize_GetAggrNamesError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	commonConfig := &drivers.CommonStorageDriverConfig{
		StorageDriverName: "ontap-san",
		DriverContext:     tridentconfig.ContextCSI,
	}
	configJSON := `{"SANType": "nvme"}`
	mAPI.EXPECT().SupportsFeature(ctx, gomock.Any()).Return(true)
	mAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, sa.NVMeTransport).Return(mockIPs, nil)
	mAPI.EXPECT().IsSVMDRCapable(ctx).Return(true, nil)
	mAPI.EXPECT().GetSVMAggregateNames(ctx).Return(nil, fmt.Errorf("failed to get aggrs"))

	err := d.Initialize(ctx, tridentconfig.ContextCSI, configJSON, commonConfig, nil, BackendUUID)

	assert.ErrorContains(t, err, "failed to get aggrs")
}

func TestNVMeInitialize_ValidateStoragePrefixError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	badStoragePrefix := "abc&$#"
	commonConfig := &drivers.CommonStorageDriverConfig{
		StorageDriverName: "ontap-san",
		DriverContext:     tridentconfig.ContextCSI,
		StoragePrefix:     &badStoragePrefix,
	}
	configJSON := `{"SANType": "nvme"}`
	mAPI.EXPECT().SupportsFeature(ctx, gomock.Any()).Return(true)
	mAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, sa.NVMeTransport).Return(mockIPs, nil)
	mAPI.EXPECT().IsSVMDRCapable(ctx).Return(true, nil)
	mAPI.EXPECT().GetSVMAggregateNames(ctx).Return([]string{"data"}, nil)
	mAPI.EXPECT().GetSVMAggregateAttributes(ctx).Return(nil, nil)
	mAPI.EXPECT().SVMName().Return("svm")

	err := d.Initialize(ctx, tridentconfig.ContextCSI, configJSON, commonConfig, nil, BackendUUID)

	assert.ErrorContains(t, err, "storage prefix may only contain letters/digits/underscore/dash")
}

func TestNVMeInitialize_Success(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	commonConfig := &drivers.CommonStorageDriverConfig{
		StorageDriverName: "ontap-san",
		DriverContext:     tridentconfig.ContextCSI,
	}
	configJSON := `{"SANType": "nvme"}`
	mAPI.EXPECT().SupportsFeature(ctx, gomock.Any()).Return(true)
	mAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, sa.NVMeTransport).Return(mockIPs, nil)
	mAPI.EXPECT().IsSVMDRCapable(ctx).Return(true, nil)
	mAPI.EXPECT().GetSVMAggregateNames(ctx).Return([]string{"data"}, nil)
	mAPI.EXPECT().GetSVMAggregateAttributes(ctx).Return(nil, nil)
	mAPI.EXPECT().SVMName().Return("svm")
	mAPI.EXPECT().EmsAutosupportLog(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	mAPI.EXPECT().GetSVMUUID().Return("svm-uuid")

	err := d.Initialize(ctx, tridentconfig.ContextCSI, configJSON, commonConfig, nil, BackendUUID)

	assert.NoError(t, err, "Failed to initialize NVMe driver.")
	assert.True(t, d.Initialized(), "NVMe driver is not initialized.")
}

func TestNVMeTerminate_Success(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)
	d.telemetry = NewOntapTelemetry(ctx, d)

	d.Terminate(ctx, "")

	assert.False(t, d.Initialized(), "NVMe driver is still running.")
}

func TestNVMeValidate_ReplicationValidationError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	d.Config.ReplicationPolicy = "pol1"
	mAPI.EXPECT().SnapmirrorPolicyGet(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to get policy"))

	err := d.validate(ctx)

	assert.ErrorContains(t, err, "replication validation failed")
}

func TestNVMeValidate_StoragePoolError(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[SnapshotPolicy] = ""
	d.virtualPools = map[string]storage.Pool{"pool1": pool1}

	err := d.validate(ctx)

	assert.ErrorContains(t, err, "storage pool validation failed")
}

func TestNVMeGetStorageBackendSpecs(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)
	backend := storage.StorageBackend{}

	backend.SetStorage(map[string]storage.Pool{})

	assert.NoError(t, d.GetStorageBackendSpecs(ctx, &backend), "Backend specs not updated.")
}

func TestNVMeGetStorageBackendPhysicalPoolNames(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)
	assert.Equal(t, d.GetStorageBackendPhysicalPoolNames(ctx), []string{"pool1"}, "Physical pools are different.")
}

func TestNVMeGetStorageBackendPools(t *testing.T) {
	driver, mockAPI := newNVMeDriverAndMockApi(t)
	svmUUID := "SVM1-uuid"
	driver.physicalPools = map[string]storage.Pool{
		"pool1": storage.NewStoragePool(nil, "pool1"),
		"pool2": storage.NewStoragePool(nil, "pool2"),
	}
	mockAPI.EXPECT().GetSVMUUID().Return(svmUUID)

	pools := driver.getStorageBackendPools(ctx)

	assert.NotEmpty(t, pools)
	assert.Equal(t, len(driver.physicalPools), len(pools))

	pool := pools[0]
	assert.NotNil(t, driver.physicalPools[pool.Aggregate])
	assert.Equal(t, driver.physicalPools[pool.Aggregate].Name(), pool.Aggregate)
	assert.Equal(t, svmUUID, pool.SvmUUID)

	pool = pools[1]
	assert.NotNil(t, driver.physicalPools[pool.Aggregate])
	assert.Equal(t, driver.physicalPools[pool.Aggregate].Name(), pool.Aggregate)
	assert.Equal(t, svmUUID, pool.SvmUUID)
}

func TestNVMeGetVolumeOpts(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)
	volConfig := storage.VolumeConfig{}
	assert.NotNil(t, d.GetVolumeOpts(ctx, &volConfig, nil), "Couldn't get VolumeOpts.")
}

func TestNVMeGetInternalVolumeName(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)
	assert.Equal(t, d.GetInternalVolumeName(ctx, "vol1"), "test_vol1", "Got different volume.")
}

func TestNVMeGetProtocol(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)
	assert.Equal(t, d.GetProtocol(ctx), tridentconfig.Block, "Incorrect protocol.")
}

func TestNVMeStoreConfig(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)
	persistentConfig := &storage.PersistentStorageBackendConfig{}

	d.StoreConfig(ctx, persistentConfig)

	assert.Equal(t, json.RawMessage("{}"), d.Config.CommonStorageDriverConfig.StoragePrefixRaw,
		"Raw prefix mismatch.")
	assert.Equal(t, d.Config, *persistentConfig.OntapConfig, "Ontap config mismatch.")
}

func TestNVMeGetUpdateType_InvalidUpdate(t *testing.T) {
	d1 := newNVMeDriver(nil, nil, nil)
	_, d2 := newMockOntapNASDriver(t)

	bMap := d1.GetUpdateType(ctx, d2)

	assert.True(t, bMap.Contains(storage.InvalidUpdate), "Valid driver update.")
}

func TestNVMeGetUpdateType_OtherUpdates(t *testing.T) {
	d1 := newNVMeDriver(nil, nil, nil)
	d2 := newNVMeDriver(nil, nil, nil)

	sPrefix := "diff"
	d2.Config.DataLIF = "1.1.1.1"
	d2.Config.Password = "diff-password"
	d2.Config.Username = "diff-username"
	d2.Config.Credentials = map[string]string{"diff": "diff"}
	d2.Config.StoragePrefix = &sPrefix

	bMap := d1.GetUpdateType(ctx, d2)

	assert.True(t, bMap.Contains(storage.InvalidVolumeAccessInfoChange), "Unchanged dataLIFs.")
	assert.True(t, bMap.Contains(storage.PasswordChange), "Unchanged password.")
	assert.True(t, bMap.Contains(storage.UsernameChange), "Unchanged username.")
	assert.True(t, bMap.Contains(storage.CredentialsChange), "Unchanged credentials.")
	assert.True(t, bMap.Contains(storage.PrefixChange), "Unchanged prefix.")
}

func TestNVMeGetCommonConfig(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)
	assert.Equal(t, d.GetCommonConfig(ctx), d.Config.CommonStorageDriverConfig, "Driver configuration not found.")
}

func TestNVMeEstablishMirror_Errors(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	snapmirrorPolicy := &api.SnapmirrorPolicy{
		Type:             "async",
		CopyAllSnapshots: true,
	}

	// Empty replication policy and schedule trying out synchronously.
	mAPI.EXPECT().SVMName().Return("svm")

	err := d.EstablishMirror(ctx, "vol1", "vol1", "", "")

	assert.ErrorContains(t, err, "could not parse remoteVolumeHandle")

	// Empty replication policy and schedule trying out synchronously.
	d.Config.ReplicationPolicy = "pol1"

	mAPI.EXPECT().SnapmirrorPolicyGet(ctx, "pol1").Return(snapmirrorPolicy,
		fmt.Errorf("failed to get policy")).Times(2)
	mAPI.EXPECT().SVMName().Return("svm")

	err = d.EstablishMirror(ctx, "vol1", "vol1", "", "")

	assert.ErrorContains(t, err, "could not parse remoteVolumeHandle")

	// Empty replication policy and schedule trying out asynchronously.
	mAPI.EXPECT().SnapmirrorPolicyGet(ctx, "pol1").Return(snapmirrorPolicy, nil)
	mAPI.EXPECT().SVMName().Return("svm")

	err = d.EstablishMirror(ctx, "vol1", "vol1", "", "")

	assert.ErrorContains(t, err, "could not parse remoteVolumeHandle")

	// Non-empty replication schedule trying out asynchronously.
	mAPI.EXPECT().SnapmirrorPolicyGet(ctx, "pol1").Return(snapmirrorPolicy, nil)
	mAPI.EXPECT().JobScheduleExists(ctx, "sch1").Return(false, fmt.Errorf("failed to get schedule"))
	mAPI.EXPECT().SVMName().Return("svm")

	err = d.EstablishMirror(ctx, "vol1", "vol1", "", "sch1")

	assert.ErrorContains(t, err, "could not parse remoteVolumeHandle")
}

func TestNVMeReestablishMirror_Errors(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)

	// Empty replication policy and schedule trying out synchronously.
	mAPI.EXPECT().SVMName().Return("svm")

	err := d.ReestablishMirror(ctx, "vol1", "vol1", "", "")

	assert.ErrorContains(t, err, "could not parse remoteVolumeHandle")

	// Empty replication policy and schedule trying out synchronously.
	d.Config.ReplicationPolicy = "pol1"
	snapmirrorPolicy := &api.SnapmirrorPolicy{
		Type:             "async",
		CopyAllSnapshots: true,
	}
	mAPI.EXPECT().SnapmirrorPolicyGet(ctx, "pol1").Return(snapmirrorPolicy,
		fmt.Errorf("failed to get policy")).Times(2)
	mAPI.EXPECT().SVMName().Return("svm")

	err = d.ReestablishMirror(ctx, "vol1", "vol1", "", "")

	assert.ErrorContains(t, err, "could not parse remoteVolumeHandle")

	// Empty replication policy and schedule trying out asynchronously.
	mAPI.EXPECT().SnapmirrorPolicyGet(ctx, "pol1").Return(snapmirrorPolicy, nil)
	mAPI.EXPECT().SVMName().Return("svm")

	err = d.ReestablishMirror(ctx, "vol1", "vol1", "", "")

	assert.ErrorContains(t, err, "could not parse remoteVolumeHandle")

	// Non-empty replication schedule trying out asynchronously.
	mAPI.EXPECT().SnapmirrorPolicyGet(ctx, "pol1").Return(snapmirrorPolicy, nil)
	mAPI.EXPECT().JobScheduleExists(ctx, "sch1").Return(false, fmt.Errorf("failed to get schedule"))
	mAPI.EXPECT().SVMName().Return("svm")

	err = d.ReestablishMirror(ctx, "vol1", "vol1", "", "sch1")

	assert.ErrorContains(t, err, "could not parse remoteVolumeHandle")
}

func TestNVMePromoteMirror_Error(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)

	promote, err := d.PromoteMirror(ctx, "", "remoteHandle", "")

	assert.False(t, promote, "Mirror promotion succeeded.")
	assert.ErrorContains(t, err, "invalid volume name")
}

func TestNVMeGetMirrorStatus_Error(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)

	status, err := d.GetMirrorStatus(ctx, "", "remoteHandle")

	assert.Empty(t, status, "Mirror status non-empty.")
	assert.ErrorContains(t, err, "invalid volume name")
}

func TestNVMeReleaseMirror_Error(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)

	err := d.ReleaseMirror(ctx, "")

	assert.ErrorContains(t, err, "invalid volume name")
}

func TestNVMeGetReplicationDetails_Error(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	mAPI.EXPECT().SVMName().Return("svm")

	_, _, _, err := d.GetReplicationDetails(ctx, "", "remoteHandle")

	assert.ErrorContains(t, err, "invalid volume name")
}

func getNVMeCreateArgs(d *NVMeStorageDriver) (storage.Pool, *storage.VolumeConfig, map[string]sa.Request) {
	pool1 := d.virtualPools["pool1"]
	volConfig := &storage.VolumeConfig{InternalName: "vol1", Size: "200000000"}
	volAttrs := map[string]sa.Request{}

	return pool1, volConfig, volAttrs
}

func TestNVMeCreate_VolumeExists(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	pool1, volConfig, volAttrs := getNVMeCreateArgs(d)

	// Volume exists API error test case.
	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, fmt.Errorf("api invocation error"))

	err := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, err, "api invocation error")

	// Volume exists test case.
	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(true, nil)

	err = d.Create(ctx, volConfig, pool1, volAttrs)

	assert.True(t, drivers.IsVolumeExistsError(err), "Volume doesn't exist.")
}

func TestNVMeCreate_InvalidVolHandle(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	pool1, volConfig, volAttrs := getNVMeCreateArgs(d)
	volConfig.PeerVolumeHandle = "volHandle"

	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, nil)
	mAPI.EXPECT().SVMName().Return("svm")

	err := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, err, "invalid volume handle")
}

func TestNVMeCreate_GetPoolsError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, volAttrs := getNVMeCreateArgs(d)
	pool1 := storage.NewStoragePool(nil, "invalid-pool")

	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, nil)

	err := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, err, "could not find pool")
}

func TestNVMeCreate_InvalidSnapshotReserve(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	pool1, volConfig, volAttrs := getNVMeCreateArgs(d)
	pool1.InternalAttributes()[SnapshotReserve] = "snapReserve"

	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, nil)

	err := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, err, "invalid value for snapshotReserve")
}

func TestNVMeCreate_VolSizeErrors(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	pool1, volConfig, volAttrs := getNVMeCreateArgs(d)

	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, nil).Times(4)

	// Convert volume size error.
	volConfig.Size = "convert-size"

	err := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, err, "could not convert volume size")

	// Invalid volume size error.
	volConfig.Size = "-100"

	err = d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, err, "invalid size")

	// Unsupported capacity error; size < 20 MiB.
	volConfig.Size = "10"

	err = d.Create(ctx, volConfig, pool1, volAttrs)
	isUnsupportedErr, _ := errors.HasUnsupportedCapacityRangeError(err)

	assert.True(t, isUnsupportedErr, "Volume size supported.")

	// Required volume size more than backend config volume size.
	volConfig.Size = "200000000"
	d.Config.LimitVolumeSize = "2000"

	err = d.Create(ctx, volConfig, pool1, volAttrs)
	isUnsupportedErr, _ = errors.HasUnsupportedCapacityRangeError(err)

	assert.True(t, isUnsupportedErr, "Volume size as per backend config.")
}

func TestNVMeCreate_InvalidEncryptionValue(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	pool1, volConfig, volAttrs := getNVMeCreateArgs(d)
	pool1.InternalAttributes()[Encryption] = "encrypt"

	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, nil)

	err := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, err, "invalid boolean value for encryption")
}

func TestNVMeCreate_InvalidFSType(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	pool1, volConfig, volAttrs := getNVMeCreateArgs(d)
	pool1.InternalAttributes()[FileSystemType] = "fake-fs"

	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, nil)

	err := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, err, "unsupported fileSystemType option")
}

func TestNVMeCreate_BothQoSPolicyError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	pool1, volConfig, volAttrs := getNVMeCreateArgs(d)
	pool1.InternalAttributes()[QosPolicy] = "QoSPol1"
	pool1.InternalAttributes()[AdaptiveQosPolicy] = "AQoSPol1"

	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, nil)
	mAPI.EXPECT().TieringPolicyValue(ctx).Return("TPolicy")

	err := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, err, "only one kind of QoS policy group may be defined")
}

func TestNVMeCreate_AggSpaceError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	pool1, volConfig, volAttrs := getNVMeCreateArgs(d)
	d.Config.LimitAggregateUsage = "10000000"

	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, nil)
	mAPI.EXPECT().TieringPolicyValue(ctx).Return("TPolicy")
	mAPI.EXPECT().GetSVMAggregateSpace(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to get aggr space"))

	err := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, err, "failed to get aggr space")
}

func TestNVMeCreate_LongLabelError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	pool1, volConfig, volAttrs := getNVMeCreateArgs(d)

	longLabelVal := "thisIsATestLabelWhoseLengthShouldExceed1023Characters_AddingSomeRandomCharacters_" +
		"V88bESTQlRIWRSS40sx9ND8P9yPf0LV8jPofiqtTp2iIXgotGh83zZ1HEeFlMGxZlIcOiPdoi07cJ" +
		"bQBuHvTRNX6pHRKUXaIrjEpygM4SpaqHYdZ8O1k2meeugg7eXu4dPhqetI3Sip3W4v9QuFkh1YBaI" +
		"9sHE9w5eRxpmTv0POpCB5xAqzmN6XCkxuXKc4yfNS9PRwcTSpvkA3PcKCF3TD1TJU3NYzcChsFQgm" +
		"bAsR32cbJRdsOwx6BkHNfRCji0xSnBFUFUu1sGHfYCmzzd3OmChADIP6RwRtpnqNzvt0CU6uumBnl" +
		"Lc5U7mBI1Ndmqhn0BBSh588thKOQcpD4bvnSBYU788tBeVxQtE8KkdUgKl8574eWldqWDiALwoiCS" +
		"Ae2GuZzwG4ACw2uHdIkjb6FEwapSKCEogr4yWFAVCYPp2pA37Mj88QWN82BEpyoTV6BRAOsubNPfT" +
		"N94X0qCcVaQp4L5bA4SPTQu0ag20a2k9LmVsocy5y11U3ewpzVGtENJmxyuyyAbxOFOkDxKLRMhgs" +
		"uJMhhplD894tkEcPoiFhdsYZbBZ4MOBF6KkuBF5aqMrQbOCFt2vvTN843nRhomVMpY01SNuUeb5mh" +
		"UN53wsqqHSGoYb1eUBDlTUDLFcCcNacxfsILqmthnrD1B5u85jRm1SfkFfuIDOgaaTM9UhxNQ1U6M" +
		"mBaRYBkuGtTScoVTXyF4lij2sj1WWrKb7qWlaUUjxHiaxgLovPWErldCXXkNFsHgc7UYLQLF4j6lO" +
		"I1QdTAyrtCcSxRwdkjBxj8mQy1HblHnaaBwP7Nax9FvIvxpeqyD6s3X1vfFNGAMuRsc9DKmPDfxjh" +
		"qGzRQawFEbbURWij9xleKsUr0yCjukyKsxuaOlwbXnoFh4V3wtidrwrNXieFD608EANwvCp7u2S8Q" +
		"px99T4O87AdQGa5cAX8Ccojd9tENOmQRmOAwVEuFtuogos96TFlq0YHyfESDTB2TWayIuGJvgTIpX" +
		"lthQFQfHVgPpUZdzZMjXry"
	labelMap := map[string]string{"key": longLabelVal}
	pool1.Attributes()[sa.Labels] = sa.NewLabelOffer(labelMap)

	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, nil)
	mAPI.EXPECT().TieringPolicyValue(ctx).Return("TPolicy")

	err := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, err, "exceeds the character limit")
}

func TestNVMeCreate_VolumeCreateAPIError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	pool1, volConfig, volAttrs := getNVMeCreateArgs(d)

	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, nil)
	mAPI.EXPECT().TieringPolicyValue(ctx).Return("TPolicy")
	mAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(fmt.Errorf("volume create failed"))

	err := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, err, "volume create failed")
}

func TestNVMeCreate_NamespaceCreateAPIError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	pool1, volConfig, volAttrs := getNVMeCreateArgs(d)

	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, nil).Times(2)
	mAPI.EXPECT().TieringPolicyValue(ctx).Return("TPolicy").Times(2)
	mAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(nil).Times(2)
	mAPI.EXPECT().NVMeNamespaceCreate(ctx, gomock.Any()).
		Return("", fmt.Errorf("failed to create namespace")).
		Times(2)
	// Volume destroy error test case.
	mAPI.EXPECT().VolumeDestroy(ctx, gomock.Any(), true).Return(fmt.Errorf("failed to delete volume"))

	err := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, err, "failed to create namespace")

	// Volume destroy success test case.
	mAPI.EXPECT().VolumeDestroy(ctx, gomock.Any(), true).Return(nil)

	err = d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, err, "failed to create namespace")
}

func TestNVMeCreate_Success(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	pool1, volConfig, volAttrs := getNVMeCreateArgs(d)

	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, nil)
	mAPI.EXPECT().TieringPolicyValue(ctx).Return("TPolicy")
	mAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(nil)
	mAPI.EXPECT().NVMeNamespaceCreate(ctx, gomock.Any()).Return("nsUUID", nil)

	err := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.NoError(t, err, "Failed to create NVMe volume.")
}

func TestNVMeDestroy_VolumeExists(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, _ := getNVMeCreateArgs(d)

	// Volume exist api call error.
	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, fmt.Errorf("api call failed"))

	err := d.Destroy(ctx, volConfig)

	assert.ErrorContains(t, err, "api call failed")

	// Volume doesn't exist to delete.
	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(false, nil)

	err = d.Destroy(ctx, volConfig)

	assert.NoError(t, err, "Volume exists.")
}

func TestNVMeDestroy_SnapMirrorAPIError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	d.Config.DriverContext = tridentconfig.ContextDocker
	_, volConfig, _ := getNVMeCreateArgs(d)

	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(true, nil)
	mAPI.EXPECT().SVMName().Return("svm")
	mAPI.EXPECT().SnapmirrorDeleteViaDestination(ctx, volConfig.InternalName, gomock.Any()).
		Return(fmt.Errorf("snap mirror api call failed"))

	err := d.Destroy(ctx, volConfig)

	assert.ErrorContains(t, err, "snap mirror api call failed")
}

func TestNVMeDestroy_VolumeDestroy_FSx(t *testing.T) {
	svmName := "SVM1"
	d, mAPI, mAWSAPI := newNVMeDriverAndMockApiAndAwsApi(t)
	d.Config.DriverContext = tridentconfig.ContextDocker
	_, volConfig, _ := getNVMeCreateArgs(d)

	tests := []struct {
		message  string
		nasType  string
		smbShare string
		state    string
	}{
		{"Test volume in FSx in available state", "nfs", "", "AVAILABLE"},
		{"Test volume in FSx in deleting state", "nfs", "", "DELETING"},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			vol := awsapi.Volume{
				State: test.state,
			}
			isVolumeExists := vol.State != ""
			mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(true, nil)
			mAWSAPI.EXPECT().VolumeExists(ctx, volConfig).Return(isVolumeExists, &vol, nil)
			if isVolumeExists {
				mAWSAPI.EXPECT().WaitForVolumeStates(
					ctx, &vol, []string{awsapi.StateDeleted}, []string{awsapi.StateFailed}, awsapi.RetryDeleteTimeout).Return("", nil)
				if vol.State == awsapi.StateAvailable {
					mAWSAPI.EXPECT().DeleteVolume(ctx, &vol).Return(nil)
				}
			} else {
				mAPI.EXPECT().SVMName().AnyTimes().Return(svmName)
				mAPI.EXPECT().SnapmirrorDeleteViaDestination(ctx, volConfig.InternalName, svmName).Return(nil)
				// mockAPI.EXPECT().SnapmirrorRelease(ctx, volConfig.InternalName, svmName).Return(nil)
				mAPI.EXPECT().VolumeDestroy(ctx, volConfig.InternalName, true).Return(nil)
			}
			result := d.Destroy(ctx, volConfig)

			assert.NoError(t, result)
		})
	}
}

func TestNVMeDestroy_VolumeDestroy(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	d.Config.DriverContext = tridentconfig.ContextDocker
	_, volConfig, _ := getNVMeCreateArgs(d)

	mAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(true, nil).Times(2)
	mAPI.EXPECT().SVMName().Return("svm").Times(2)
	mAPI.EXPECT().SnapmirrorDeleteViaDestination(ctx, volConfig.InternalName, gomock.Any()).Return(nil).Times(2)

	// Volume destroy API call failed.
	mAPI.EXPECT().VolumeDestroy(ctx, volConfig.InternalName, true).Return(fmt.Errorf("destroy volume failed"))

	err := d.Destroy(ctx, volConfig)

	assert.ErrorContains(t, err, "destroy volume failed")

	// Volume destroy success.
	mAPI.EXPECT().VolumeDestroy(ctx, volConfig.InternalName, true).Return(nil)

	err = d.Destroy(ctx, volConfig)

	assert.NoError(t, err, "Failed to destroy volume.")
}

func TestNVMeGetVolumeExternal_GetVolumeError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to get volume"))

	_, err := d.GetVolumeExternal(ctx, "vol")

	assert.ErrorContains(t, err, "failed to get volume")
}

func TestNVMeGetVolumeExternal_GetNamespaceError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)

	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(&api.Volume{}, nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to get ns"))

	_, err := d.GetVolumeExternal(ctx, "vol")

	assert.ErrorContains(t, err, "failed to get ns")
}

func TestNVMeGetVolumeExternal_Success(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	v := &api.Volume{Name: "test_vol", SnapshotPolicy: "pol", Aggregates: []string{"data"}}

	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(v, nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).Return(&api.NVMeNamespace{Size: "10"}, nil)

	vol, err := d.GetVolumeExternal(ctx, "vol")

	assert.NoError(t, err, "Failed to get Volume External struct.")
	assert.Equal(t, "vol", vol.Config.Name, "Found different volume.")
	assert.Equal(t, "test_vol", vol.Config.InternalName, "Found different volume.")
	assert.Equal(t, "10", vol.Config.Size, "Wrong volume size.")
	assert.Equal(t, "data", vol.Pool, "Found wrong pool.")
}

func TestNVMeGetVolumeExternalWrappers_VolumeListError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	sChannel := make(chan *storage.VolumeExternalWrapper)

	mAPI.EXPECT().VolumeListByPrefix(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to get volumes"))

	go d.GetVolumeExternalWrappers(ctx, sChannel)

	v := <-sChannel

	assert.ErrorContains(t, v.Error, "failed to get volumes")
	assert.Nil(t, v.Volume, "Found a volume.")
}

func TestNVMeGetVolumeExternalWrappers_NamespaceListError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	sChannel := make(chan *storage.VolumeExternalWrapper)

	mAPI.EXPECT().VolumeListByPrefix(ctx, gomock.Any()).Return(nil, nil)
	mAPI.EXPECT().NVMeNamespaceList(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to get namespaces"))

	go d.GetVolumeExternalWrappers(ctx, sChannel)

	v := <-sChannel

	assert.ErrorContains(t, v.Error, "failed to get namespaces")
	assert.Nil(t, v.Volume, "Found a volume.")
}

func TestNVMeGetVolumeExternalWrappers_Success(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	sChannel := make(chan *storage.VolumeExternalWrapper)
	vols := api.Volumes{&api.Volume{Name: "vol1", SnapshotPolicy: "pol1"}}
	namespaces := api.NVMeNamespaces{
		&api.NVMeNamespace{Name: "ns1", VolumeName: "vol1", Size: "10"},
		&api.NVMeNamespace{Name: "ns2", VolumeName: "vol2", Size: "20"},
	}

	mAPI.EXPECT().VolumeListByPrefix(ctx, gomock.Any()).Return(vols, nil)
	mAPI.EXPECT().NVMeNamespaceList(ctx, gomock.Any()).Return(namespaces, nil)

	go d.GetVolumeExternalWrappers(ctx, sChannel)

	v := <-sChannel

	assert.NoError(t, v.Error, "failed to get namespaces")
	assert.Equal(t, "vol1", v.Volume.Config.Name, "Found different volume.")
	assert.Equal(t, "vol1", v.Volume.Config.InternalName, "Found different volume.")
	assert.Equal(t, "10", v.Volume.Config.Size, "Wrong volume size.")
	assert.Equal(t, drivers.UnsetPool, v.Volume.Pool, "Found wrong pool.")
}

func TestNVMeCreateNVMeNamespaceCommentString(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)
	nsAttr := map[string]string{
		nsAttributeFSType:    "ext4",
		nsAttributeLUKS:      "luks",
		nsAttributeDriverCtx: "docker",
	}

	nsCommentString := `{"nsAttribute":{"LUKS":"luks","com.netapp.ndvp.fstype":"ext4","driverContext":"docker"}}`

	// Comment string exceeds max length
	nsComment, err := d.createNVMeNamespaceCommentString(ctx, nsAttr, 10)

	assert.ErrorContains(t, err, "exceeds the character limit")

	assert.Equal(t, "", nsComment, "Comment has garbage string.")

	// Success case
	nsComment, err = d.createNVMeNamespaceCommentString(ctx, nsAttr, nsMaxCommentLength)

	assert.NoError(t, err, "Failed to get namespace comment.")
	assert.Equal(t, nsCommentString, nsComment, "Incorrect namespace comment.")
}

func TestNVMeParseNVMeNamespaceCommentString(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)

	nsCommentString := `{"nsAttribute":{"LUKS":"luks","com.netapp.ndvp.fstype":"ext4","driverContext":"docker"}}`

	nsComment, err := d.ParseNVMeNamespaceCommentString(ctx, nsCommentString)

	assert.NoError(t, err)
	assert.Equal(t, "ext4", nsComment[nsAttributeFSType])
	assert.Equal(t, "luks", nsComment[nsAttributeLUKS])
	assert.Equal(t, "docker", nsComment[nsAttributeDriverCtx])
}

func TestGetNamespaceSpecificSubsystemName(t *testing.T) {
	name := "fakeName"
	expected_name := "s_fakeName"

	got_name := getNamespaceSpecificSubsystemName(name)

	assert.Equal(t, got_name, expected_name)
}

func TestGetNodeSpecificSubsystemName(t *testing.T) {
	// case 1: subsystem, name is shorter than 96 char
	nodeName := "fakeNodeName"
	tridentUUID := "fakeUUID"
	expected := "fakeNodeName-fakeUUID"

	got := getNodeSpecificSubsystemName(nodeName, tridentUUID)

	assert.Equal(t, got, expected)

	// case 2: subsystem name is longer than 96 char
	nodeName = "fakeNodeNamefakeNodeNamefakeNodeNamefakeNodeNamefakeNodeNamefakeNodeNamefakeNodeNamefakeNodeName"

	expected = "fakeNodeNamefakeNodeNamefakeNodeNamefakeNodeNamefakeNodeNamefakeNodeNamefakeNodeNamefak-fakeUUID"

	got = getNodeSpecificSubsystemName(nodeName, tridentUUID)
	assert.Equal(t, got, expected)
}

func TestPublish(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mock := mockapi.NewMockOntapAPI(mockCtrl)
	d := newNVMeDriver(mock, nil, nil)

	volConfig := &storage.VolumeConfig{
		Name:         "fakeVolName",
		InternalName: "fakeInternalName",
	}

	flexVol := &api.Volume{
		AccessType: VolTypeRW,
	}

	publishInfo := &utils.VolumePublishInfo{
		HostName:    "fakeHostName",
		TridentUUID: "fakeUUID",
	}

	subsystem := &api.NVMeSubsystem{
		Name: "fakeSubsysName",
		NQN:  "fakeNQN",
		UUID: "fakeUUID",
	}

	namespace := &api.NVMeNamespace{
		UUID:    "fakeNsUUID",
		Comment: "fakeComment",
	}
	// case 1: error getting volume Info
	mock.EXPECT().VolumeInfo(ctx, volConfig.InternalName).Return(flexVol, fmt.Errorf("Error Getting Volume Info")).Times(1)

	err := d.Publish(ctx, volConfig, publishInfo)

	assert.Error(t, err)

	// case 2: success getting volume Info
	flexVol.AccessType = VolTypeDP
	mock.EXPECT().VolumeInfo(ctx, volConfig.InternalName).Return(flexVol, nil).Times(1)

	err = d.Publish(ctx, volConfig, publishInfo)

	assert.Error(t, err)

	// case 3: Error getting namespace in Docker context
	flexVol.AccessType = VolTypeRW
	d.Config.DriverContext = tridentconfig.ContextDocker
	publishInfo.HostNQN = "fakeHostNQN"
	mock.EXPECT().VolumeInfo(ctx, volConfig.InternalName).Return(flexVol, nil).Times(1)
	mock.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).Return(nil, fmt.Errorf("Error getting namespace by name")).Times(1)

	err = d.Publish(ctx, volConfig, publishInfo)

	assert.Error(t, err)

	// case 4: Error getting namespace in Docker context
	d.Config.DriverContext = tridentconfig.ContextDocker
	publishInfo.HostNQN = "fakeHostNQN"
	mock.EXPECT().VolumeInfo(ctx, volConfig.InternalName).Return(flexVol, nil).Times(1)
	mock.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).Return(nil, nil).Times(1)

	err = d.Publish(ctx, volConfig, publishInfo)

	assert.Error(t, err)

	// case 5: Error creating subsystem in Docker context
	d.Config.DriverContext = tridentconfig.ContextDocker
	publishInfo.HostNQN = "fakeHostNQN"
	mock.EXPECT().VolumeInfo(ctx, volConfig.InternalName).Return(flexVol, nil).Times(1)
	mock.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).Return(namespace, nil).Times(1)
	mock.EXPECT().NVMeSubsystemCreate(ctx, "fakeHostName-fakeUUID").Return(subsystem, fmt.Errorf("Error creating subsystem")).Times(1)

	err = d.Publish(ctx, volConfig, publishInfo)

	assert.Error(t, err)

	// case 6: Error creating subsystem in CSI Context
	d.Config.DriverContext = tridentconfig.ContextCSI
	mock.EXPECT().VolumeInfo(ctx, volConfig.InternalName).Return(flexVol, nil).Times(1)
	mock.EXPECT().NVMeSubsystemCreate(ctx, "fakeHostName-fakeUUID").Return(subsystem, fmt.Errorf("Error creating subsystem")).Times(1)

	err = d.Publish(ctx, volConfig, publishInfo)

	assert.Error(t, err)

	// case 7: Host NQN not found in publish Info
	flexVol.AccessType = VolTypeRW
	publishInfo.HostNQN = ""
	mock.EXPECT().VolumeInfo(ctx, volConfig.InternalName).Return(flexVol, nil).Times(1)

	err = d.Publish(ctx, volConfig, publishInfo)

	assert.Error(t, err)

	// case 8: Error while adding host nqn to subsystem
	publishInfo.HostNQN = "fakeHostNQN"
	mock.EXPECT().VolumeInfo(ctx, volConfig.InternalName).Return(flexVol, nil).Times(1)
	mock.EXPECT().NVMeSubsystemCreate(ctx, "fakeHostName-fakeUUID").Return(subsystem, nil).Times(1)
	mock.EXPECT().NVMeAddHostToSubsystem(ctx, publishInfo.HostNQN, subsystem.UUID).Return(fmt.Errorf("Error adding host nqnq to subsystem")).Times(1)

	err = d.Publish(ctx, volConfig, publishInfo)

	assert.Error(t, err)

	// case 9: Error returned by NVMeEnsureNamespaceMapped
	mock.EXPECT().VolumeInfo(ctx, volConfig.InternalName).Return(flexVol, nil).Times(1)
	mock.EXPECT().NVMeSubsystemCreate(ctx, "fakeHostName-fakeUUID").Return(subsystem, nil).Times(1)
	mock.EXPECT().NVMeAddHostToSubsystem(ctx, publishInfo.HostNQN, subsystem.UUID).Return(nil).Times(1)
	mock.EXPECT().NVMeEnsureNamespaceMapped(ctx, gomock.Any(), gomock.Any()).Return(fmt.Errorf("Error returned by NVMeEnsureNamespaceMapped")).Times(1)

	err = d.Publish(ctx, volConfig, publishInfo)

	assert.Error(t, err)

	// case 10: Success
	mock.EXPECT().VolumeInfo(ctx, volConfig.InternalName).Return(flexVol, nil).Times(1)
	mock.EXPECT().NVMeSubsystemCreate(ctx, "fakeHostName-fakeUUID").Return(subsystem, nil).Times(1)
	mock.EXPECT().NVMeAddHostToSubsystem(ctx, publishInfo.HostNQN, subsystem.UUID).Return(nil).Times(1)
	mock.EXPECT().NVMeEnsureNamespaceMapped(ctx, gomock.Any(), gomock.Any()).Return(nil).Times(1)

	err = d.Publish(ctx, volConfig, publishInfo)

	assert.NoError(t, err)
}

func TestUnpublish(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mock := mockapi.NewMockOntapAPI(mockCtrl)
	d := newNVMeDriver(mock, nil, nil)

	volConfig := &storage.VolumeConfig{
		Name:         "fakeVolName",
		InternalName: "fakeInternalName",
	}

	publishInfo := &utils.VolumePublishInfo{
		HostName:    "fakeHostName",
		TridentUUID: "fakeUUID",
	}

	// case 1: NVMeEnsureNamespaceUnmapped returned error
	volConfig.AccessInfo.NVMeNamespaceUUID = "fakeUUID"
	tridentconfig.CurrentDriverContext = tridentconfig.ContextCSI
	mock.EXPECT().NVMeEnsureNamespaceUnmapped(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(false, fmt.Errorf("NVMeEnsureNamespaceUnmapped returned error"))

	err := d.Unpublish(ctx, volConfig, publishInfo)

	assert.Error(t, err)

	// case 2: Success
	volConfig.AccessInfo.PublishEnforcement = true
	volConfig.AccessInfo.NVMeNamespaceUUID = "fakeUUID"
	tridentconfig.CurrentDriverContext = tridentconfig.ContextCSI
	mock.EXPECT().NVMeEnsureNamespaceUnmapped(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(true, nil)

	err = d.Unpublish(ctx, volConfig, publishInfo)

	assert.NoError(t, err)
}

func TestCreatePrepare(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mock := mockapi.NewMockOntapAPI(mockCtrl)
	d := newNVMeDriver(mock, nil, nil)
	volConfig := &storage.VolumeConfig{
		Name:         "fakeVolName",
		InternalName: "fakeInternalName",
	}
	volConfig.ImportNotManaged = false
	tridentconfig.CurrentDriverContext = tridentconfig.ContextCSI

	d.CreatePrepare(ctx, volConfig)
	tridentconfig.CurrentDriverContext = ""

	assert.Equal(t, "test_fakeVolName", volConfig.InternalName, "Incorrect volume internal name.")
	assert.True(t, volConfig.AccessInfo.PublishEnforcement, "Publish enforcement not enabled.")
}

func TestNVMeResize_VolumeExistsErrors(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, _ := getNVMeCreateArgs(d)

	// API call fails use case.
	mAPI.EXPECT().VolumeExists(ctx, gomock.Any()).Return(false, fmt.Errorf("api call failed"))

	err := d.Resize(ctx, volConfig, 100)

	assert.ErrorContains(t, err, "error occurred checking for existing volume")

	// Volume doesn't exist.
	mAPI.EXPECT().VolumeExists(ctx, gomock.Any()).Return(false, nil)

	err = d.Resize(ctx, volConfig, 100)

	assert.ErrorContains(t, err, "does not exist")
}

func TestNVMeResize_GetVolumeSizeError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, _ := getNVMeCreateArgs(d)

	mAPI.EXPECT().VolumeExists(ctx, gomock.Any()).Return(true, nil)
	mAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(0), fmt.Errorf("api call failed"))

	err := d.Resize(ctx, volConfig, 100)

	assert.ErrorContains(t, err, "error occurred when checking volume size")
}

func TestNVMeResize_GetNamespaceError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, _ := getNVMeCreateArgs(d)

	mAPI.EXPECT().VolumeExists(ctx, gomock.Any()).Return(true, nil)
	mAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(200), nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).Return(nil, fmt.Errorf("api call failed"))

	err := d.Resize(ctx, volConfig, 100)

	assert.ErrorContains(t, err, "api call failed")
}

func TestNVMeResize_ParseNamespaceSizeError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, _ := getNVMeCreateArgs(d)
	ns := &api.NVMeNamespace{Name: "/vol/vol1/namespace0", Size: "-100"}

	mAPI.EXPECT().VolumeExists(ctx, gomock.Any()).Return(true, nil)
	mAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(200), nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).Return(ns, nil)

	err := d.Resize(ctx, volConfig, 100)

	assert.ErrorContains(t, err, "error while parsing namespace size")
}

func TestNVMeResize_LessRequestedSizeError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, _ := getNVMeCreateArgs(d)
	ns := &api.NVMeNamespace{Name: "/vol/vol1/namespace0", Size: "100"}

	mAPI.EXPECT().VolumeExists(ctx, gomock.Any()).Return(true, nil)
	mAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(200), nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).Return(ns, nil)

	err := d.Resize(ctx, volConfig, 99)

	assert.ErrorContains(t, err, "less than existing volume size")
}

func TestNVMeResize_SameNamespaceSize(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, _ := getNVMeCreateArgs(d)
	ns := &api.NVMeNamespace{Name: "/vol/vol1/namespace0", Size: "100"}

	mAPI.EXPECT().VolumeExists(ctx, gomock.Any()).Return(true, nil)
	mAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(200), nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).Return(ns, nil)
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(nil, fmt.Errorf("api call failed"))

	err := d.Resize(ctx, volConfig, 120)

	assert.NoError(t, err, "Failed to resize.")
	assert.Equal(t, "100", volConfig.Size, "Resize failed.")
}

func TestNVMeResize_AggrLimitsError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, _ := getNVMeCreateArgs(d)
	ns := &api.NVMeNamespace{Name: "/vol/vol1/namespace0", Size: "100"}

	mAPI.EXPECT().VolumeExists(ctx, gomock.Any()).Return(true, nil)
	mAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(200), nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).Return(ns, nil)
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(nil, fmt.Errorf("api call failed")).Times(2)

	err := d.Resize(ctx, volConfig, tridentconfig.SANResizeDelta+200)

	assert.ErrorContains(t, err, "api call failed")
}

func TestNVMeResize_DriverVolumeLimitError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, _ := getNVMeCreateArgs(d)
	ns := &api.NVMeNamespace{Name: "/vol/vol1/namespace0", Size: "100"}
	vol := &api.Volume{Aggregates: []string{"data"}}
	d.Config.LimitVolumeSize = "1000.1000"

	mAPI.EXPECT().VolumeExists(ctx, gomock.Any()).Return(true, nil)
	mAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(200), nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).Return(ns, nil)
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed"))
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(vol, nil)

	err := d.Resize(ctx, volConfig, tridentconfig.SANResizeDelta+200)

	assert.ErrorContains(t, err, "error parsing limitVolumeSize")
}

func TestNVMeResize_VolumeSetSizeError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, _ := getNVMeCreateArgs(d)
	ns := &api.NVMeNamespace{Name: "/vol/vol1/namespace0", Size: "100"}
	vol := &api.Volume{Aggregates: []string{"data"}}

	mAPI.EXPECT().VolumeExists(ctx, gomock.Any()).Return(true, nil)
	mAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(200), nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).Return(ns, nil)
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed"))
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(vol, nil)
	mAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(fmt.Errorf("api call failed"))

	err := d.Resize(ctx, volConfig, tridentconfig.SANResizeDelta+200)

	assert.ErrorContains(t, err, "volume resize failed")
}

func TestNVMeResize_NamespaceSetSizeError(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, _ := getNVMeCreateArgs(d)
	ns := &api.NVMeNamespace{Name: "/vol/vol1/namespace0", Size: "100"}
	vol := &api.Volume{Aggregates: []string{"data"}}

	mAPI.EXPECT().VolumeExists(ctx, gomock.Any()).Return(true, nil)
	mAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(200), nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).Return(ns, nil)
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed"))
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(vol, nil)
	mAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(nil)
	mAPI.EXPECT().NVMeNamespaceSetSize(ctx, gomock.Any(), gomock.Any()).Return(fmt.Errorf("api call failed"))

	err := d.Resize(ctx, volConfig, tridentconfig.SANResizeDelta+200)

	assert.ErrorContains(t, err, "volume resize failed")
}

func TestNVMeResize_Success(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, _ := getNVMeCreateArgs(d)
	ns := &api.NVMeNamespace{Name: "/vol/vol1/namespace0", Size: "100"}
	vol := &api.Volume{Aggregates: []string{"data"}}

	mAPI.EXPECT().VolumeExists(ctx, gomock.Any()).Return(true, nil)
	mAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(200), nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).Return(ns, nil)
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed"))
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(vol, nil)
	mAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(nil)
	mAPI.EXPECT().NVMeNamespaceSetSize(ctx, gomock.Any(), gomock.Any()).Return(nil)

	err := d.Resize(ctx, volConfig, tridentconfig.SANResizeDelta+200)

	assert.NoError(t, err, "Volume resize failed.")
	assert.Equal(t, "50000200", volConfig.Size, "Incorrect namespace size after resize.")
}

func TestCreateClone(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, _ := getNVMeCreateArgs(d)
	cloneConfig := &storage.VolumeConfig{
		InternalName:                "cloneVol1",
		Size:                        "200000000",
		CloneSourceVolumeInternal:   "fakeSource",
		CloneSourceSnapshotInternal: "fakeSourceSnapshot",
	}
	vol := &api.Volume{Aggregates: []string{"data"}}
	pool := storage.NewStoragePool(nil, "fakepool")
	pool.InternalAttributes()[FileSystemType] = tridentconfig.FsExt4
	pool.InternalAttributes()[SplitOnClone] = "true"
	ns := &api.NVMeNamespace{Name: "/vol/cloneVol1/namespace0", Size: "100"}
	volConfig.InternalID = "/vol/cloneVol1/namespace0"

	// Test1: Success - creating clone
	mAPI.EXPECT().VolumeExists(ctx, gomock.Any()).AnyTimes().Return(false, nil)
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(vol, nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, gomock.Any()).AnyTimes().Return(ns, nil)
	mAPI.EXPECT().VolumeCloneCreate(ctx, gomock.Any(), gomock.Any(), gomock.Any(), false).AnyTimes().Return(nil)
	mAPI.EXPECT().VolumeWaitForStates(ctx, "cloneVol1", []string{"online"}, []string{"error"}, maxFlexvolCloneWait).AnyTimes().Return("online", nil)
	mAPI.EXPECT().VolumeSetComment(ctx, gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mAPI.EXPECT().VolumeCloneSplitStart(ctx, gomock.Any()).AnyTimes().Return(nil)

	err := d.CreateClone(ctx, volConfig, cloneConfig, pool)

	assert.NoError(t, err)

	// Test2: Error - Not able to check the volume
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(nil, fmt.Errorf("unable to get volume Info"))

	err = d.CreateClone(ctx, volConfig, cloneConfig, pool)

	assert.Error(t, err)

	// Test3: Error - Flexvol not found
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(nil, nil)

	err = d.CreateClone(ctx, volConfig, cloneConfig, pool)

	assert.Error(t, err)

	// Test4: Success - Setting comment
	vol.Comment = "fakeComment"
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(vol, nil)

	err = d.CreateClone(ctx, volConfig, cloneConfig, pool)

	assert.NoError(t, err)

	// Test5: Error - Storage pool name not set
	pool.SetName("")

	err = d.CreateClone(ctx, volConfig, cloneConfig, pool)

	assert.Error(t, err)
}

func TestImport(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)
	_, volConfig, _ := getNVMeCreateArgs(d)
	originalName := "fakeOriginalName"
	vol := &api.Volume{Aggregates: []string{"data"}}
	ns := &api.NVMeNamespace{Name: "/vol/cloneVol1/namespace0", Size: "100", UUID: "fakeUUID"}

	// Test1: Error - Error getting volume info
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(nil, fmt.Errorf("Error getting volume info"))

	err := d.Import(ctx, volConfig, originalName)

	assert.Error(t, err)

	// Test2: Error - Failed to get the volume info
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(nil, nil)

	err = d.Import(ctx, volConfig, originalName)

	assert.Error(t, err)

	// Test3: Error - volume is not read-write
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(vol, nil)
	vol.AccessType = "dp"

	err = d.Import(ctx, volConfig, originalName)

	assert.Error(t, err)

	// Test4: Error - Error getting namespace
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(vol, nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, "/vol/"+originalName+"/*").Return(nil,
		fmt.Errorf("error getting namespace info"))
	vol.AccessType = "rw"

	err = d.Import(ctx, volConfig, originalName)

	assert.Error(t, err)

	// Test5: Error - Failed to get namespace
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(vol, nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, "/vol/"+originalName+"/*").Return(nil, nil)

	err = d.Import(ctx, volConfig, originalName)

	assert.Error(t, err)

	// Test6: Error - Namespace not online
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(vol, nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, "/vol/"+originalName+"/*").Return(ns, nil)
	ns.State = "offline"

	err = d.Import(ctx, volConfig, originalName)

	assert.Error(t, err)

	// Test7: Error - Namespace mapped to subsystem
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(vol, nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, "/vol/"+originalName+"/*").Return(ns, nil)
	mAPI.EXPECT().NVMeIsNamespaceMapped(ctx, "", ns.UUID).Return(true, nil)
	ns.State = "online"

	err = d.Import(ctx, volConfig, originalName)

	assert.Error(t, err)

	// Test8: Error - Checking if namespace is mapped returns error
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(vol, nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, "/vol/"+originalName+"/*").Return(ns, nil)
	mAPI.EXPECT().NVMeIsNamespaceMapped(ctx, "", ns.UUID).Return(false,
		fmt.Errorf("error while checking namespace mapped"))

	err = d.Import(ctx, volConfig, originalName)

	assert.Error(t, err)

	// Test9: Error - while renaming the volume
	volConfig.ImportNotManaged = false
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(vol, nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, "/vol/"+originalName+"/*").Return(ns, nil)
	mAPI.EXPECT().NVMeIsNamespaceMapped(ctx, "", ns.UUID).Return(false, nil)
	mAPI.EXPECT().VolumeRename(ctx, originalName,
		volConfig.InternalName).Return(fmt.Errorf("error renaming volume"))

	err = d.Import(ctx, volConfig, originalName)

	assert.Error(t, err)

	// Test10: Success
	vol.Comment = "fakeComment"
	mAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(vol, nil)
	mAPI.EXPECT().NVMeNamespaceGetByName(ctx, "/vol/"+originalName+"/*").Return(ns, nil)
	mAPI.EXPECT().NVMeIsNamespaceMapped(ctx, "", ns.UUID).Return(false, nil)
	mAPI.EXPECT().VolumeRename(ctx, originalName, volConfig.InternalName).Return(nil)

	err = d.Import(ctx, volConfig, originalName)

	assert.NoError(t, err)
}

func TestExtractNamespaceName(t *testing.T) {
	var nsStr, nsNameGot, nsNameExpected string
	// Test 1 : Empty String
	nsStr = ""
	nsNameExpected = "namespace0"

	nsNameGot = extractNamespaceName(nsStr)

	assert.Equal(t, nsNameExpected, nsNameGot)

	// Test 2 : Success - Namespace Name in proper format
	nsStr = "/vol/fakeFlexVolName/fakeNamespaceName"
	nsNameExpected = "fakeNamespaceName"

	nsNameGot = extractNamespaceName(nsStr)

	assert.Equal(t, nsNameExpected, nsNameGot)

	// Test 3 : Namespace Name has more length
	nsStr = "/vol/fakeFlexVolName/hasMoreLength/fakeNamespaceName"
	nsNameExpected = "MalformedNamespace"

	nsNameGot = extractNamespaceName(nsStr)

	assert.Equal(t, nsNameExpected, nsNameGot)

	// Test 4 : Namespace Name doesn't match the format
	nsStr = "/vol/fakeFlexVolName"
	nsNameExpected = "MalformedNamespace"

	nsNameGot = extractNamespaceName(nsStr)

	assert.Equal(t, nsNameExpected, nsNameGot)
}

func TestCreateNamespacePath(t *testing.T) {
	nsNameExpected := "/vol/fakeFlexVolName/fakeNamespaceName"
	flexVolName := "fakeFlexVolName"
	nsName := "fakeNamespaceName"

	nsNameGot := createNamespacePath(flexVolName, nsName)

	assert.Equal(t, nsNameExpected, nsNameGot)
}

func TestGetBackendState(t *testing.T) {
	d, mAPI := newNVMeDriverAndMockApi(t)

	mAPI.EXPECT().GetSVMState(ctx).Return("", fmt.Errorf("returning test error"))

	reason, changeMap := d.GetBackendState(ctx)

	assert.Equal(t, reason, StateReasonSVMUnreachable, "should be 'SVM is not reachable'")
	assert.NotNil(t, changeMap, "should not be nil")
}

func TestEnablePublishEnforcement(t *testing.T) {
	d := newNVMeDriver(nil, nil, nil)
	vol := storage.Volume{Config: getVolumeConfig()}

	assert.True(t, d.CanEnablePublishEnforcement(), "Cannot enable publish enforcement.")

	d.EnablePublishEnforcement(ctx, &vol)

	assert.True(t, vol.Config.AccessInfo.PublishEnforcement, "Incorrect publish enforcement value.")
}
