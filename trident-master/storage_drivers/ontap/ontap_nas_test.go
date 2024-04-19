// Copyright 2021 NetApp, Inc. All Rights Reserved.

package ontap

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/RoaringBitmap/roaring"
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

// //////////////////////////////////////////////////////////////////////////////////////////
// /             _____________________
// /            |   <<Interface>>    |
// /            |       ONTAPI       |
// /            |____________________|
// /                ^             ^
// /     Implements |             | Implements
// /   ____________________    ____________________
// /  |  ONTAPAPIREST     |   |  ONTAPAPIZAPI     |
// /  |___________________|   |___________________|
// /  | +API: RestClient  |   | +API: *Client     |
// /  |___________________|   |___________________|
// /
// //////////////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////////////////////////////
// Drivers that offer dual support are to call ONTAP REST or ZAPI's
// via abstraction layer (ONTAPI interface)
// //////////////////////////////////////////////////////////////////////////////////////////

var (
	ctx             = context.Background()
	debugTraceFlags = map[string]bool{"method": true, "api": true, "discovery": true}
)

const (
	BackendUUID                 = "deadbeef-03af-4394-ace4-e177cdbcaf28"
	ONTAPTEST_LOCALHOST         = "127.0.0.1"
	ONTAPTEST_VSERVER_AGGR_NAME = "data"
	FSX_ID                      = "fsx-1234"
)

func TestOntapNasStorageDriverConfigString(t *testing.T) {
	vserverAdminHost := ONTAPTEST_LOCALHOST
	vserverAdminPort := "0"
	vserverAggrName := ONTAPTEST_VSERVER_AGGR_NAME

	ontapNasDrivers := []NASStorageDriver{
		*newTestOntapNASDriver(vserverAdminHost, vserverAdminPort, vserverAggrName,
			"CSI", true, nil),
		*newTestOntapNASDriver(vserverAdminHost, vserverAdminPort, vserverAggrName,
			"CSI", false, nil),
	}

	sensitiveIncludeList := map[string]string{
		"username":        "ontap-nas-user",
		"password":        "password1!",
		"client username": "client_username",
		"client password": "client_password",
	}

	externalIncludeList := map[string]string{
		"<REDACTED>":                   "<REDACTED>",
		"username":                     "Username:<REDACTED>",
		"password":                     "Password:<REDACTED>",
		"api":                          "API:<REDACTED>",
		"chap username":                "ChapUsername:<REDACTED>",
		"chap initiator secret":        "ChapInitiatorSecret:<REDACTED>",
		"chap target username":         "ChapTargetUsername:<REDACTED>",
		"chap target initiator secret": "ChapTargetInitiatorSecret:<REDACTED>",
		"client private key":           "ClientPrivateKey:<REDACTED>",
	}

	for _, ontapNasDriver := range ontapNasDrivers {
		for key, val := range externalIncludeList {
			assert.Contains(t, ontapNasDriver.String(), val, "ontap-nas driver does not contain %v", key)
			assert.Contains(t, ontapNasDriver.GoString(), val, "ontap-nas driver does not contain %v", key)
		}

		for key, val := range sensitiveIncludeList {
			assert.NotContains(t, ontapNasDriver.String(), val, "ontap-nas driver contains %v", key)
			assert.NotContains(t, ontapNasDriver.GoString(), val, "ontap-nas driver contains %v", key)
		}
	}
}

func newTestOntapNASDriver(
	vserverAdminHost, vserverAdminPort, vserverAggrName string, driverContext tridentconfig.DriverContext, useREST bool, fsxId *string,
) *NASStorageDriver {
	config := &drivers.OntapStorageDriverConfig{}
	sp := func(s string) *string { return &s }

	config.CommonStorageDriverConfig = &drivers.CommonStorageDriverConfig{}
	config.CommonStorageDriverConfig.DebugTraceFlags = make(map[string]bool)
	config.CommonStorageDriverConfig.DebugTraceFlags["method"] = true
	config.ManagementLIF = vserverAdminHost + ":" + vserverAdminPort
	config.SVM = "SVM1"
	config.Aggregate = vserverAggrName
	config.Username = "ontap-nas-user"
	config.Password = "password1!"
	config.StorageDriverName = "ontap-nas"
	config.StoragePrefix = sp("test_")
	config.DriverContext = driverContext
	config.UseREST = useREST

	if fsxId != nil {
		config.AWSConfig = &drivers.AWSConfig{}
		config.AWSConfig.FSxFilesystemID = *fsxId
	}

	nasDriver := &NASStorageDriver{}
	nasDriver.Config = *config

	var ontapAPI api.OntapAPI

	nasDriver.API = ontapAPI
	nasDriver.telemetry = &Telemetry{
		Plugin:        nasDriver.Name(),
		SVM:           config.SVM,
		StoragePrefix: *nasDriver.GetConfig().StoragePrefix,
		Driver:        nasDriver,
	}

	return nasDriver
}

func TestInitializeStoragePoolsLabels(t *testing.T) {
	vserverAdminHost := ONTAPTEST_LOCALHOST
	vserverAggrName := ONTAPTEST_VSERVER_AGGR_NAME

	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	d := newTestOntapNASDriver(vserverAdminHost, "443", vserverAggrName, "CSI", false, nil)
	d.API = mockAPI
	d.Config.Storage = []drivers.OntapStorageDriverPool{
		{
			Region: "us_east_1",
			Zone:   "us_east_1a",
			SupportedTopologies: []map[string]string{
				{
					"topology.kubernetes.io/region": "us_east_1",
					"topology.kubernetes.io/zone":   "us_east_1a",
				},
			},
		},
	}

	poolAttributes := map[string]sa.Offer{
		sa.BackendType:      sa.NewStringOffer(d.Name()),
		sa.Snapshots:        sa.NewBoolOffer(true),
		sa.Clones:           sa.NewBoolOffer(true),
		sa.Encryption:       sa.NewBoolOffer(true),
		sa.ProvisioningType: sa.NewStringOffer("thick", "thin"),
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().GetSVMAggregateNames(gomock.Any()).AnyTimes().Return([]string{ONTAPTEST_VSERVER_AGGR_NAME}, nil)
	mockAPI.EXPECT().GetSVMAggregateAttributes(gomock.Any()).AnyTimes().Return(
		map[string]string{ONTAPTEST_VSERVER_AGGR_NAME: "vmdisk"}, nil,
	)

	cases := []struct {
		physicalPoolLabels   map[string]string
		virtualPoolLabels    map[string]string
		physicalExpected     string
		virtualExpected      string
		backendName          string
		physicalErrorMessage string
		virtualErrorMessage  string
	}{
		{
			nil, nil, "", "", "nas-backend",
			"Label is not empty", "Label is not empty",
		}, // no labels
		{
			map[string]string{"base-key": "base-value"},
			nil,
			`{"provisioning":{"base-key":"base-value"}}`,
			`{"provisioning":{"base-key":"base-value"}}`, "nas-backend",
			"Base label is not set correctly", "Base label is not set correctly",
		}, // base label only
		{
			nil,
			map[string]string{"virtual-key": "virtual-value"},
			"",
			`{"provisioning":{"virtual-key":"virtual-value"}}`, "nas-backend",
			"Base label is not empty", "Virtual pool label is not set correctly",
		}, // virtual label only
		{
			map[string]string{"base-key": "base-value"},
			map[string]string{"virtual-key": "virtual-value"},
			`{"provisioning":{"base-key":"base-value"}}`,
			`{"provisioning":{"base-key":"base-value","virtual-key":"virtual-value"}}`,
			"nas-backend",
			"Base label is not set correctly", "Virtual pool label is not set correctly",
		}, // base and virtual labels
	}

	for _, c := range cases {
		d.Config.Labels = c.physicalPoolLabels
		d.Config.Storage[0].Labels = c.virtualPoolLabels
		physicalPools, virtualPools, err := InitializeStoragePoolsCommon(ctx, d, poolAttributes,
			c.backendName)
		assert.NoError(t, err, "Error is not nil")

		physicalPool := physicalPools["data"]
		label, err := physicalPool.GetLabelsJSON(ctx, "provisioning", 1023)
		assert.NoError(t, err, "Error is not nil")
		assert.Equal(t, c.physicalExpected, label, c.physicalErrorMessage)

		virtualPool := virtualPools["nas-backend_pool_0"]
		label, err = virtualPool.GetLabelsJSON(ctx, "provisioning", 1023)
		assert.NoError(t, err, "Error is not nil")
		assert.Equal(t, c.virtualExpected, label, c.virtualErrorMessage)
	}
}

func TestOntapNasStorageDriverInitialize_WithTwoAuthMethods(t *testing.T) {
	vserverAdminHost := ONTAPTEST_LOCALHOST
	vserverAdminPort := "0"
	vserverAggrName := ONTAPTEST_VSERVER_AGGR_NAME

	commonConfig := &drivers.CommonStorageDriverConfig{
		Version:           1,
		StorageDriverName: "ontap-nas",
		BackendName:       "myOntapNasBackend",
		DriverContext:     tridentconfig.ContextCSI,
		DebugTraceFlags:   debugTraceFlags,
	}

	configJSON := `
	{
		"version":           1,
		"storageDriverName": "ontap-nas",
		"managementLIF":     "1.1.1.1:10",
		"svm":               "SVM1",
		"aggregate":         "data",
		"username":          "dummyuser",
		"password":          "dummypassword",
		"clientcertificate": "dummy-certificate",
		"clientprivatekey":  "dummy-client-private-key"
	}`
	ontapNasDriver := newTestOntapNASDriver(vserverAdminHost, vserverAdminPort, vserverAggrName,
		"CSI", false, nil)

	result := ontapNasDriver.Initialize(ctx, "CSI", configJSON, commonConfig,
		map[string]string{}, BackendUUID)

	assert.Error(t, result, "driver initialization succeeded despite more than one authentication methods in config")
	assert.Contains(t, result.Error(), "more than one authentication method", "expected error string not found")
}

func TestOntapNasStorageDriverInitialize_WithTwoAuthMethodsWithSecrets(t *testing.T) {
	vserverAdminHost := ONTAPTEST_LOCALHOST
	vserverAdminPort := "0"
	vserverAggrName := ONTAPTEST_VSERVER_AGGR_NAME

	commonConfig := &drivers.CommonStorageDriverConfig{
		Version:           1,
		StorageDriverName: "ontap-nas",
		BackendName:       "myOntapNasBackend",
		DriverContext:     tridentconfig.ContextCSI,
		DebugTraceFlags:   debugTraceFlags,
	}

	configJSON := `
	{
		"version":           1,
		"storageDriverName": "ontap-nas",
		"managementLIF":     "1.1.1.1:10",
		"svm":               "SVM1",
		"aggregate":         "data"
	}`
	secrets := map[string]string{
		"username":          "dummyuser",
		"password":          "dummypassword",
		"clientprivatekey":  "dummy-client-private-key",
		"clientcertificate": "dummy-certificate",
	}
	ontapNasDriver := newTestOntapNASDriver(vserverAdminHost, vserverAdminPort, vserverAggrName,
		"CSI", false, nil)

	result := ontapNasDriver.Initialize(ctx, "CSI", configJSON, commonConfig, secrets,
		BackendUUID)

	assert.Error(t, result, "driver initialization succeeded despite more than one authentication methods in config")
	assert.Contains(t, result.Error(), "more than one authentication method", "expected error string not found")
}

func TestOntapNasStorageDriverInitialize_WithTwoAuthMethodsWithConfigAndSecrets(t *testing.T) {
	vserverAdminHost := ONTAPTEST_LOCALHOST
	vserverAdminPort := "0"
	vserverAggrName := ONTAPTEST_VSERVER_AGGR_NAME

	commonConfig := &drivers.CommonStorageDriverConfig{
		Version:           1,
		StorageDriverName: "ontap-nas",
		BackendName:       "myOntapNasBackend",
		DriverContext:     tridentconfig.ContextCSI,
		DebugTraceFlags:   debugTraceFlags,
	}

	configJSON := `
	{
		"version":           1,
		"storageDriverName": "ontap-nas",
		"managementLIF":     "1.1.1.1:10",
		"svm":               "SVM1",
		"aggregate":         "data",
		"username":          "dummyuser",
		"password":          "dummypassword"
	}`
	secrets := map[string]string{
		"clientprivatekey":  "dummy-client-private-key",
		"clientcertificate": "dummy-certificate",
	}
	ontapNasDriver := newTestOntapNASDriver(vserverAdminHost, vserverAdminPort, vserverAggrName,
		"CSI", false, nil)

	result := ontapNasDriver.Initialize(ctx, "CSI", configJSON, commonConfig, secrets,
		BackendUUID)

	assert.Error(t, result, "driver initialization succeeded despite more than one authentication methods in config")
	assert.Contains(t, result.Error(), "more than one authentication method", "expected error string not found")
}

func newMockAWSOntapNASDriver(t *testing.T) (*mockapi.MockOntapAPI, *mockapi.MockAWSAPI, *NASStorageDriver) {
	mockCtrl := gomock.NewController(t)
	mockAPI, driver := newMockOntapNASDriver(t)
	mockAWSAPI := mockapi.NewMockAWSAPI(mockCtrl)
	driver.AWSAPI = mockAWSAPI
	return mockAPI, mockAWSAPI, driver
}

func newMockOntapNASDriver(t *testing.T) (*mockapi.MockOntapAPI, *NASStorageDriver) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)

	vserverAdminHost := ONTAPTEST_LOCALHOST
	vserverAdminPort := "0"
	vserverAggrName := ONTAPTEST_VSERVER_AGGR_NAME
	fsxId := FSX_ID

	driver := newTestOntapNASDriver(vserverAdminHost, vserverAdminPort, vserverAggrName,
		"CSI", false, &fsxId)
	driver.API = mockAPI
	return mockAPI, driver
}

func TestOntapNasStorageDriverInitialize(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")

	commonConfig := &drivers.CommonStorageDriverConfig{
		Version:           1,
		StorageDriverName: "ontap-nas",
		BackendName:       "myOntapNasBackend",
		DriverContext:     tridentconfig.ContextCSI,
		DebugTraceFlags:   debugTraceFlags,
	}

	configJSON := `
	{
		"version":           1,
		"storageDriverName": "ontap-nas",
		"managementLIF":     "127.0.0.1:0",
		"svm":               "SVM1",
		"aggregate":         "data",
		"username":          "dummyuser",
		"password":          "dummypassword"
	}`
	secrets := map[string]string{
		"clientcertificate": "dummy-certificate",
	}
	driver.telemetry = &Telemetry{
		Plugin: driver.Name(),
		SVM:    "SVM1",
		Driver: driver,
		done:   make(chan struct{}),
	}
	driver.telemetry.TridentVersion = tridentconfig.OrchestratorVersion.String()
	driver.telemetry.TridentBackendUUID = BackendUUID
	hostname, _ := os.Hostname()
	message, _ := json.Marshal(driver.GetTelemetry())

	mockAPI.EXPECT().IsSVMDRCapable(ctx).Return(true, nil)
	mockAPI.EXPECT().GetSVMAggregateNames(ctx).AnyTimes().Return([]string{ONTAPTEST_VSERVER_AGGR_NAME}, nil)
	mockAPI.EXPECT().GetSVMAggregateAttributes(gomock.Any()).AnyTimes().Return(
		map[string]string{ONTAPTEST_VSERVER_AGGR_NAME: "vmdisk"}, nil,
	)
	mockAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, "nfs").Return([]string{"dataLIF"}, nil)
	mockAPI.EXPECT().EmsAutosupportLog(ctx, "ontap-nas", "1", false, "heartbeat", hostname, string(message), 1,
		"trident", 5).AnyTimes()
	mockAPI.EXPECT().GetSVMUUID().Return("SVM1-uuid")

	result := driver.Initialize(ctx, "CSI", configJSON, commonConfig, secrets, BackendUUID)

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverInitialize_Failure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	driver.API = nil
	commonConfig := &drivers.CommonStorageDriverConfig{
		Version:           1,
		StorageDriverName: "ontap-nas",
		BackendName:       "myOntapNasBackend",
		DriverContext:     tridentconfig.ContextCSI,
		DebugTraceFlags:   debugTraceFlags,
	}

	configJSON := `
	{
		"version":           1,
		"storageDriverName": "ontap-nas",
		"managementLIF":     "127.0.0.1:0",
		"svm":               "SVM1",
		"aggregate":         "data",
		"username":          "dummyuser",
		"password":          "dummypassword"
	}`
	secrets := map[string]string{
		"clientcertificate": "dummy-certificate",
	}

	result := driver.Initialize(ctx, "CSI", configJSON, commonConfig, secrets, BackendUUID)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverInitialize_StoragePoolFailed(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")

	commonConfig := &drivers.CommonStorageDriverConfig{
		Version:           1,
		StorageDriverName: "ontap-nas",
		BackendName:       "myOntapNasBackend",
		DriverContext:     tridentconfig.ContextCSI,
		DebugTraceFlags:   debugTraceFlags,
	}

	configJSON := `
	{
		"version":           1,
		"storageDriverName": "ontap-nas",
		"managementLIF":     "127.0.0.1:0",
		"svm":               "SVM1",
		"aggregate":         "data",
		"username":          "dummyuser",
		"password":          "dummypassword"
	}`
	secrets := map[string]string{
		"clientcertificate": "dummy-certificate",
	}

	mockAPI.EXPECT().IsSVMDRCapable(ctx).Return(true, nil)
	mockAPI.EXPECT().GetSVMAggregateNames(ctx).AnyTimes().Return(nil, fmt.Errorf("no aggregates found"))

	result := driver.Initialize(ctx, "CSI", configJSON, commonConfig, secrets, BackendUUID)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverInitialize_ValidationFailed(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	commonConfig := &drivers.CommonStorageDriverConfig{
		Version:           1,
		StorageDriverName: "ontap-nas",
		BackendName:       "myOntapNasBackend",
		DriverContext:     tridentconfig.ContextCSI,
		DebugTraceFlags:   debugTraceFlags,
	}

	configJSON := `
	{
		"version":           1,
		"storageDriverName": "ontap-nas",
		"managementLIF":     "127.0.0.1:0",
		"svm":               "SVM1",
		"aggregate":         "data",
		"username":          "dummyuser",
		"password":          "dummypassword"
	}`
	secrets := map[string]string{
		"clientcertificate": "dummy-certificate",
	}

	mockAPI.EXPECT().IsSVMDRCapable(ctx).Return(true, nil)
	mockAPI.EXPECT().GetSVMAggregateNames(ctx).AnyTimes().Return([]string{ONTAPTEST_VSERVER_AGGR_NAME}, nil)
	mockAPI.EXPECT().GetSVMAggregateAttributes(gomock.Any()).AnyTimes().Return(
		map[string]string{ONTAPTEST_VSERVER_AGGR_NAME: "vmdisk"}, nil,
	)
	mockAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, "nfs").Return(nil, fmt.Errorf("validation failed"))

	result := driver.Initialize(ctx, "CSI", configJSON, commonConfig, secrets, BackendUUID)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverInitialized(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")

	driver.initialized = true
	result := driver.Initialized()
	assert.Equal(t, true, result, "Incorrect initialization status")

	driver.initialized = false
	result = driver.Initialized()
	assert.Equal(t, false, result, "Incorrect initialization status")
}

func TestOntapNasStorageDriverTerminate(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	tests := []struct {
		name string
		err  error
	}{
		{"no error", nil},
		{"error", fmt.Errorf("policy not found")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			driver.Config.AutoExportPolicy = true
			driver.telemetry = nil
			driver.initialized = true

			mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
			mockAPI.EXPECT().ExportPolicyDestroy(ctx, "trident-dummy").Return(test.err)

			driver.Terminate(ctx, "dummy")

			assert.False(t, driver.initialized)
		})
	}
}

func TestOntapNasStorageDriverTerminate_TelemetryFailure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	driver.Config.AutoExportPolicy = true
	driver.telemetry = &Telemetry{
		Plugin:        driver.Name(),
		SVM:           "SVM1",
		StoragePrefix: *driver.GetConfig().StoragePrefix,
		Driver:        driver,
		done:          make(chan struct{}),
	}
	driver.initialized = true

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().ExportPolicyDestroy(ctx, "trident-dummy").Return(fmt.Errorf("policy not found"))

	driver.Terminate(ctx, "dummy")

	assert.False(t, driver.initialized)
}

func TestOntapNasStorageDriverValidate(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	driver.Config.LUKSEncryption = "false"
	dataLIF := make([]string, 0)
	dataLIF = append(dataLIF, "10.0.201.1")

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, "nfs").Return(dataLIF, nil)

	result := driver.validate(ctx)

	assert.NoError(t, result, "Ontap NAS driver validation failed")
}

func TestOntapNasStorageDriverValidate_InvalidReplicationPolicy(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	driver.Config.LUKSEncryption = "false"
	driver.Config.ReplicationPolicy = "testpolicy"
	dataLIF := make([]string, 0)
	dataLIF = append(dataLIF, "10.0.201.1")

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, "testpolicy").Return(nil, fmt.Errorf("replication policy not found"))

	result := driver.validate(ctx)

	assert.Error(t, result, "Ontap NAS driver validation succeeded when it should have failed")
}

func TestOntapNasStorageDriverValidate_InvalidDataLIF(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	driver.Config.LUKSEncryption = "false"
	driver.Config.DataLIF = "foo"
	dataLIF := make([]string, 0)
	dataLIF = append(dataLIF, "10.0.201.1")

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, "nfs").Return(dataLIF, nil)

	result := driver.validate(ctx)

	assert.Error(t, result, "Ontap NAS driver validation failed")
}

func TestOntapNasStorageDriverValidate_InvalidPrefix(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	driver.Config.LUKSEncryption = "false"
	driver.Config.StoragePrefix = utils.Ptr("B@D")
	dataLIF := make([]string, 0)
	dataLIF = append(dataLIF, "10.0.201.1")

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, "nfs").Return(dataLIF, nil)

	result := driver.validate(ctx)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverValidate_InvalidStoragePools(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	driver.Config.LUKSEncryption = "false"
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		SpaceReserve: "iaminvalid",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	dataLIF := make([]string, 0)
	dataLIF = append(dataLIF, "10.0.201.1")

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, "nfs").Return(dataLIF, nil)

	result := driver.validate(ctx)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeClone(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.SplitOnClone = "false"

	volConfig := &storage.VolumeConfig{
		Size:                        "1g",
		Encryption:                  "false",
		FileSystem:                  "nfs",
		CloneSourceSnapshotInternal: "flexvol",
	}

	flexVol := api.Volume{
		Name:    "flexvol",
		Comment: "flexvol",
	}

	tests := []struct {
		NasType string
	}{
		{"nfs"},
		{"smb"},
	}

	for _, test := range tests {
		t.Run(test.NasType, func(t *testing.T) {
			mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
			mockAPI.EXPECT().VolumeInfo(ctx, volConfig.CloneSourceVolumeInternal).Return(&flexVol, nil)
			mockAPI.EXPECT().VolumeExists(ctx, "").Return(false, nil)
			mockAPI.EXPECT().VolumeCloneCreate(ctx, volConfig.InternalName, volConfig.CloneSourceVolumeInternal,
				volConfig.CloneSourceSnapshotInternal, false).Return(nil)
			mockAPI.EXPECT().VolumeWaitForStates(ctx, volConfig.InternalName, gomock.Any(), gomock.Any(),
				maxFlexvolCloneWait).Return("online", nil)
			mockAPI.EXPECT().VolumeSetComment(ctx, volConfig.InternalName, volConfig.InternalName, "flexvol").
				Return(nil)
			mockAPI.EXPECT().VolumeMount(ctx, volConfig.InternalName, "/"+volConfig.InternalName).Return(nil)

			if test.NasType == sa.SMB {
				driver.Config.NASType = sa.SMB
				mockAPI.EXPECT().SMBShareExists(ctx, volConfig.InternalName).Return(false, nil)
				mockAPI.EXPECT().SMBShareCreate(ctx, volConfig.InternalName, "/"+volConfig.InternalName).Return(nil)
			}

			result := driver.CreateClone(ctx, nil, volConfig, pool1)

			assert.NoError(t, result)
		})
	}
}

func TestOntapNasStorageDriverVolumeClone_ROClone(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.SplitOnClone = "false"

	volConfig := &storage.VolumeConfig{
		Size:                        "1g",
		Encryption:                  "false",
		FileSystem:                  "nfs",
		CloneSourceSnapshotInternal: "flexvol",
		ReadOnlyClone:               true,
	}

	flexVol := api.Volume{
		Name:        "flexvol",
		Comment:     "flexvol",
		SnapshotDir: true,
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, volConfig.CloneSourceVolumeInternal).Return(&flexVol, nil)

	result := driver.CreateClone(ctx, nil, volConfig, pool1)
	fmt.Println(result)

	assert.NoError(t, result, "received error")
}

func TestOntapNasStorageDriverVolumeClone_ROClone_Failure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.SplitOnClone = "false"

	volConfig := &storage.VolumeConfig{
		Size:                        "1g",
		Encryption:                  "false",
		FileSystem:                  "nfs",
		CloneSourceSnapshotInternal: "flexvol",
		ReadOnlyClone:               true,
	}

	// Set snapshot directory visibility to false
	flexVol := api.Volume{
		Name:        "flexvol",
		Comment:     "flexvol",
		SnapshotDir: false,
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")

	// Creating a readonly clone only results in the driver looking up volume information and no other calls to ONTAP.
	mockAPI.EXPECT().VolumeInfo(ctx, volConfig.CloneSourceVolumeInternal).Return(&flexVol, nil)

	result := driver.CreateClone(ctx, nil, volConfig, pool1)

	assert.Error(t, result, "expected error")
}

func TestOntapNasStorageDriverVolumeClone_StoragePoolUnset(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "nfs",
	}

	flexVol := api.Volume{
		Name:    "flexvol",
		Comment: "flexvol",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, volConfig.CloneSourceVolumeInternal).Return(&flexVol, nil)

	result := driver.CreateClone(ctx, nil, volConfig, nil)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeClone_VolumeDoesNotExist(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}

	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "nfs",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, volConfig.CloneSourceVolumeInternal).Return(nil, fmt.Errorf("volume not found"))

	result := driver.CreateClone(ctx, nil, volConfig, pool1)

	assert.Error(t, result, "could not create clone as original volume is not found")
}

func TestOntapNasStorageDriverVolumeClone_BothQosPolicy(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.SplitOnClone = "false"

	volConfig := &storage.VolumeConfig{
		Size:                "1g",
		Encryption:          "false",
		FileSystem:          "nfs",
		CloneSourceSnapshot: "flexvol",
		QosPolicy:           "fake",
		AdaptiveQosPolicy:   "fake",
	}

	flexVol := api.Volume{
		Name:    "flexvol",
		Comment: "flexvol",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, volConfig.CloneSourceVolumeInternal).Return(&flexVol, nil)

	result := driver.CreateClone(ctx, nil, volConfig, pool1)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeClone_LabelLengthExceeding(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	pool1 := storage.NewStoragePool(nil, "")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.SplitOnClone = "false"

	longLabel := "thisIsATestLabelWhoseLengthShouldExceed1023Characters_AddingSomeRandomCharacters_" +
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
	driver.Config.Labels = map[string]string{
		"cloud":   "anf",
		longLabel: "dev-test-cluster-1",
	}

	volConfig := &storage.VolumeConfig{
		Size:                "1g",
		Encryption:          "false",
		FileSystem:          "nfs",
		CloneSourceSnapshot: "flexvol",
	}
	flexVol := api.Volume{
		Name:    "flexvol",
		Comment: "flexvol",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, volConfig.CloneSourceVolumeInternal).Return(&flexVol, nil)

	result := driver.CreateClone(ctx, nil, volConfig, pool1)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeClone_CreateFail(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.SplitOnClone = "false"
	driver.Config.NASType = sa.SMB

	volConfig := &storage.VolumeConfig{
		Size:                        "1g",
		Encryption:                  "false",
		FileSystem:                  "nfs",
		CloneSourceSnapshotInternal: "flexvol",
	}

	flexVol := api.Volume{
		Name:    "flexvol",
		Comment: "flexvol",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, volConfig.CloneSourceVolumeInternal).Return(&flexVol, nil)
	mockAPI.EXPECT().VolumeExists(ctx, "").Return(false, nil)
	mockAPI.EXPECT().VolumeCloneCreate(ctx, volConfig.InternalName, volConfig.CloneSourceVolumeInternal,
		volConfig.CloneSourceSnapshotInternal, false).Return(fmt.Errorf("create clone fail"))

	result := driver.CreateClone(ctx, nil, volConfig, pool1)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeClone_SMBShareCreateFail(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.SplitOnClone = "false"
	driver.Config.NASType = sa.SMB

	volConfig := &storage.VolumeConfig{
		Size:                        "1g",
		Encryption:                  "false",
		FileSystem:                  "nfs",
		CloneSourceSnapshotInternal: "flexvol",
	}

	flexVol := api.Volume{
		Name:    "flexvol",
		Comment: "flexvol",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, volConfig.CloneSourceVolumeInternal).Return(&flexVol, nil)
	mockAPI.EXPECT().VolumeExists(ctx, "").Return(false, nil)
	mockAPI.EXPECT().VolumeCloneCreate(ctx, volConfig.InternalName, volConfig.CloneSourceVolumeInternal,
		volConfig.CloneSourceSnapshotInternal, false).Return(nil)
	mockAPI.EXPECT().VolumeWaitForStates(ctx, volConfig.InternalName, gomock.Any(), gomock.Any(),
		maxFlexvolCloneWait).Return("online", nil)
	mockAPI.EXPECT().VolumeSetComment(ctx, volConfig.InternalName, volConfig.InternalName, "flexvol").Return(nil)
	mockAPI.EXPECT().VolumeMount(ctx, volConfig.InternalName, "/"+volConfig.InternalName).Return(nil)
	mockAPI.EXPECT().SMBShareExists(ctx, volConfig.InternalName).Return(false, nil)
	mockAPI.EXPECT().SMBShareCreate(ctx, volConfig.InternalName,
		"/"+volConfig.InternalName).Return(fmt.Errorf("cannot create volume"))

	result := driver.CreateClone(ctx, nil, volConfig, pool1)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeDestroy_FSx(t *testing.T) {
	svmName := "SVM1"
	volName := "testVol"
	volNameInternal := volName + "Internal"
	mockAPI, mockAWSAPI, driver := newMockAWSOntapNASDriver(t)

	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Name:         volName,
		InternalName: volNameInternal,
		Encryption:   "false",
		FileSystem:   "xfs",
	}

	assert.NotNil(t, mockAPI)

	tests := []struct {
		message  string
		nasType  string
		smbShare string
		state    string
	}{
		{"Test NFS volume in FSx in available state", "nfs", "", "AVAILABLE"},
		{"Test NFS volume in FSx in deleting state", "nfs", "", "DELETING"},
		{"Test NFS volume does not exist in FSx", "nfs", "", ""},
		{"Test SMB volume does not exist in FSx", "smb", volConfig.InternalName, ""},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			vol := awsapi.Volume{
				State: test.state,
			}
			isVolumeExists := vol.State != ""
			mockAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(true, nil)
			mockAWSAPI.EXPECT().VolumeExists(ctx, volConfig).Return(isVolumeExists, &vol, nil)
			if isVolumeExists {
				mockAWSAPI.EXPECT().WaitForVolumeStates(
					ctx, &vol, []string{awsapi.StateDeleted}, []string{awsapi.StateFailed}, awsapi.RetryDeleteTimeout).Return("", nil)
				if vol.State == awsapi.StateAvailable {
					mockAWSAPI.EXPECT().DeleteVolume(ctx, &vol).Return(nil)
				}
			} else {
				mockAPI.EXPECT().SVMName().AnyTimes().Return(svmName)
				mockAPI.EXPECT().SnapmirrorDeleteViaDestination(ctx, volConfig.InternalName, svmName).Return(nil)
				mockAPI.EXPECT().SnapmirrorRelease(ctx, volConfig.InternalName, svmName).Return(nil)
				mockAPI.EXPECT().VolumeDestroy(ctx, volConfig.InternalName, true).Return(nil)
				if test.nasType == sa.SMB {
					if test.smbShare == "" {
						mockAPI.EXPECT().SMBShareExists(ctx, volConfig.InternalName).Return(true, nil)
						mockAPI.EXPECT().SMBShareDestroy(ctx, volConfig.InternalName).Return(nil)
					}
					driver.Config.NASType = sa.SMB
					driver.Config.SMBShare = test.smbShare
				}
			}
			result := driver.Destroy(ctx, volConfig)

			assert.NoError(t, result)
		})
	}
}

func TestOntapNasStorageDriverVolumeDestroy(t *testing.T) {
	svmName := "SVM1"
	volName := "testVol"
	volNameInternal := volName + "Internal"
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Name:         volName,
		InternalName: volNameInternal,
		Encryption:   "false",
		FileSystem:   "xfs",
	}

	assert.NotNil(t, mockAPI)

	tests := []struct {
		message  string
		nasType  string
		smbShare string
	}{
		{"Test NFS volume", "nfs", ""},
		{"Test SMB volume", "smb", ""},
		{"Test SMB volume with share", "smb", volConfig.InternalName},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			mockAPI.EXPECT().SVMName().AnyTimes().Return(svmName)
			mockAPI.EXPECT().VolumeExists(ctx, volConfig.InternalName).Return(true, nil)
			mockAPI.EXPECT().SnapmirrorDeleteViaDestination(ctx, volConfig.InternalName, svmName).Return(nil)
			mockAPI.EXPECT().SnapmirrorRelease(ctx, volConfig.InternalName, svmName).Return(nil)
			mockAPI.EXPECT().VolumeDestroy(ctx, volConfig.InternalName, true).Return(nil)
			if test.nasType == sa.SMB {
				if test.smbShare == "" {
					mockAPI.EXPECT().SMBShareExists(ctx, volConfig.InternalName).Return(true, nil)
					mockAPI.EXPECT().SMBShareDestroy(ctx, volConfig.InternalName).Return(nil)
				}
				driver.Config.NASType = sa.SMB
				driver.Config.SMBShare = test.smbShare
			}
			result := driver.Destroy(ctx, volConfig)

			assert.NoError(t, result)
		})
	}
}

func TestOntapNasStorageDriverVolumeDestroy_VolumeNotFound(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}

	tests := []struct {
		message string
		valid   bool
	}{
		{"volume already deleted", false},
		{"volume not found", true},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
			if test.valid {
				mockAPI.EXPECT().VolumeExists(ctx, "").Return(false, fmt.Errorf(test.message))
				result := driver.Destroy(ctx, volConfig)
				assert.Error(t, result)
			} else { // case where volume is already deleted
				mockAPI.EXPECT().VolumeExists(ctx, "").Return(false, nil)
				result := driver.Destroy(ctx, volConfig)
				assert.NoError(t, result)
			}
		})
	}
}

func TestOntapNasStorageDriverVolumeDestroy_SnapmirrorDeleteFail(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	svmName := "SVM1"
	volName := "testVol"
	volNameInternal := volName + "Internal"
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Name:         volName,
		InternalName: volNameInternal,
		Encryption:   "false",
		FileSystem:   "xfs",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return(svmName)
	mockAPI.EXPECT().VolumeExists(ctx, volNameInternal).Return(true, nil)
	mockAPI.EXPECT().SnapmirrorDeleteViaDestination(ctx, volNameInternal,
		svmName).Return(fmt.Errorf("error deleting snapmirror info for volume"))

	result := driver.Destroy(ctx, volConfig)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeDestroy_SnapmirrorReleaseFail(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	svmName := "SVM1"
	volName := "testVol"
	volNameInternal := volName + "Internal"
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Name:         volName,
		InternalName: volNameInternal,
		Encryption:   "false",
		FileSystem:   "xfs",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return(svmName)
	mockAPI.EXPECT().VolumeExists(ctx, volNameInternal).Return(true, nil)
	mockAPI.EXPECT().SnapmirrorDeleteViaDestination(ctx, volNameInternal, svmName).Return(nil)
	mockAPI.EXPECT().SnapmirrorRelease(ctx, volNameInternal,
		svmName).Return(fmt.Errorf("error releaseing snapmirror"))

	result := driver.Destroy(ctx, volConfig)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeDestroy_Fail(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	svmName := "SVM1"
	volName := "testVol"
	volNameInternal := volName + "Internal"
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Name:         volName,
		InternalName: volNameInternal,
		Encryption:   "false",
		FileSystem:   "xfs",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return(svmName)
	mockAPI.EXPECT().VolumeExists(ctx, volNameInternal).Return(true, nil)
	mockAPI.EXPECT().SnapmirrorDeleteViaDestination(ctx, volNameInternal, svmName).Return(nil)
	mockAPI.EXPECT().SnapmirrorRelease(ctx, volNameInternal, svmName).Return(nil)
	mockAPI.EXPECT().VolumeDestroy(ctx, volNameInternal, true).Return(fmt.Errorf("cannot delete volume"))

	result := driver.Destroy(ctx, volConfig)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverSMBShareDestroy_VolumeNotFound(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	svmName := "SVM1"
	volName := "testVol"
	volNameInternal := volName + "Internal"
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Name:         volName,
		InternalName: volNameInternal,
		Encryption:   "false",
		FileSystem:   "xfs",
	}
	driver.Config.NASType = sa.SMB

	tests := []struct {
		message           string
		serverReturnError bool
	}{
		{"ServerError", true},
		{"share not found", false},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			mockAPI.EXPECT().SVMName().AnyTimes().Return(svmName)
			mockAPI.EXPECT().VolumeExists(ctx, volNameInternal).Return(true, nil)
			mockAPI.EXPECT().SnapmirrorDeleteViaDestination(ctx, volNameInternal, svmName).Return(nil)
			mockAPI.EXPECT().SnapmirrorRelease(ctx, volNameInternal, svmName).Return(nil)
			mockAPI.EXPECT().VolumeDestroy(ctx, volNameInternal, true).Return(nil)
			if test.serverReturnError {
				mockAPI.EXPECT().SMBShareExists(ctx, volNameInternal).Return(false,
					fmt.Errorf("Server does not respond"))
				result := driver.Destroy(ctx, volConfig)
				assert.Error(t, result)
			} else {
				mockAPI.EXPECT().SMBShareExists(ctx, volNameInternal).Return(false, nil)
				result := driver.Destroy(ctx, volConfig)
				assert.NoError(t, result)
			}
		})
	}
}

func TestOntapNasStorageDriverSMBDestroy_Fail(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	svmName := "SVM1"
	volName := "testVol"
	volNameInternal := volName + "Internal"
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Name:         volName,
		InternalName: volNameInternal,
		Encryption:   "false",
		FileSystem:   "xfs",
	}

	driver.Config.NASType = sa.SMB

	mockAPI.EXPECT().SVMName().AnyTimes().Return(svmName)
	mockAPI.EXPECT().VolumeExists(ctx, volNameInternal).Return(true, nil)
	mockAPI.EXPECT().SnapmirrorDeleteViaDestination(ctx, volNameInternal, svmName).Return(nil)
	mockAPI.EXPECT().SnapmirrorRelease(ctx, volNameInternal, svmName).Return(nil)
	mockAPI.EXPECT().VolumeDestroy(ctx, volNameInternal, true).Return(nil)
	mockAPI.EXPECT().SMBShareExists(ctx, volNameInternal).Return(true, nil)
	mockAPI.EXPECT().SMBShareDestroy(ctx, volNameInternal).Return(fmt.Errorf("cannot delete SMB share"))

	result := driver.Destroy(ctx, volConfig)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeRename(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeRename(ctx, "volInternal", "newVolInternal").Return(nil)

	result := driver.Rename(ctx, "volInternal", "newVolInternal")

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverVolumeCanSnapshot(t *testing.T) {
	_, driver := newMockOntapNASDriver(t)
	result := driver.CanSnapshot(ctx, nil, nil)
	assert.NoError(t, result)
}

func TestOntapNasStorageDriverVolumeGetSnapshot(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap1",
		VolumeInternalName: "vol1",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeUsedSize(ctx, "vol1").Return(1, nil)
	mockAPI.EXPECT().VolumeSnapshotInfo(ctx, snapConfig.InternalName, snapConfig.VolumeInternalName).Return(
		api.Snapshot{
			CreateTime: "time",
			Name:       "snap1",
		},
		nil)

	snap, err := driver.GetSnapshot(ctx, snapConfig, volConfig)

	assert.NotNil(t, snap)
	assert.NoError(t, err)
}

func TestOntapNasStorageDriverVolumeGetSnapshot_VolumeSizeFailure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap1",
		VolumeInternalName: "vol1",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeUsedSize(ctx, "vol1").Return(0, fmt.Errorf("error reading volume size"))

	snap, err := driver.GetSnapshot(ctx, snapConfig, volConfig)

	assert.Nil(t, snap)
	assert.Error(t, err)
}

func TestOntapNasStorageDriverVolumeGetSnapshot_NoSnapshot(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap1",
		VolumeInternalName: "vol1",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeUsedSize(ctx, "vol1").Return(1, nil)
	mockAPI.EXPECT().VolumeSnapshotInfo(ctx, snapConfig.InternalName, snapConfig.VolumeInternalName).Return(
		api.Snapshot{},
		mockError)

	snap, err := driver.GetSnapshot(ctx, snapConfig, volConfig)

	assert.Nil(t, snap)
	assert.Error(t, err)
}

func TestOntapNasStorageDriverVolumeGetSnapshot_NoError(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap1",
		VolumeInternalName: "vol1",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeUsedSize(ctx, "vol1").Return(0, nil)
	mockAPI.EXPECT().VolumeSnapshotInfo(ctx, snapConfig.InternalName, snapConfig.VolumeInternalName).Return(
		api.Snapshot{},
		errors.NotFoundError(fmt.Sprintf("snapshot %v not found for volume %v", snapConfig.InternalName,
			snapConfig.VolumeInternalName)))

	snap, err := driver.GetSnapshot(ctx, snapConfig, volConfig)

	assert.Nil(t, snap)
	assert.NoError(t, err)
}

func TestOntapNasStorageDriverVolumeGetSnapshots(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	snapshots := api.Snapshots{}
	snapshots = append(snapshots, api.Snapshot{
		CreateTime: "time",
		Name:       "snap1",
	})

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeUsedSize(ctx, "vol1").Return(1, nil)
	mockAPI.EXPECT().VolumeSnapshotList(ctx, "vol1").Return(snapshots, nil)

	snap, err := driver.GetSnapshots(ctx, volConfig)

	assert.NotNil(t, snap)
	assert.NoError(t, err)
}

func TestOntapNasStorageDriverVolumeGetSnapshots_VolumeSizeFailure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeUsedSize(ctx, "vol1").Return(0, fmt.Errorf("error reading volume size"))

	snap, err := driver.GetSnapshots(ctx, volConfig)

	assert.Nil(t, snap)
	assert.Error(t, err)
}

func TestOntapNasStorageDriverVolumeGetSnapshots_NoSnapshot(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeUsedSize(ctx, "vol1").Return(0, nil)
	mockAPI.EXPECT().VolumeSnapshotList(ctx, "vol1").Return(nil, fmt.Errorf("no snapshots found"))

	snap, err := driver.GetSnapshots(ctx, volConfig)

	assert.Nil(t, snap)
	assert.Error(t, err)
}

func TestOntapNasStorageDriverVolumeCreateSnapshot(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap1",
		VolumeInternalName: "vol1",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(true, nil)
	mockAPI.EXPECT().VolumeUsedSize(ctx, "vol1").Return(1, nil)
	mockAPI.EXPECT().VolumeSnapshotCreate(ctx, "snap1", "vol1").Return(nil)
	mockAPI.EXPECT().VolumeSnapshotInfo(ctx,
		snapConfig.InternalName, snapConfig.VolumeInternalName).Return(
		api.Snapshot{
			CreateTime: "time",
			Name:       "snap1",
		},
		nil)

	snap, err := driver.CreateSnapshot(ctx, snapConfig, volConfig)

	assert.NotNil(t, snap)
	assert.NoError(t, err)
}

func TestOntapNasStorageDriverVolumeRestoreSnapshot(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap1",
		VolumeInternalName: "vol1",
	}

	mockAPI.EXPECT().SnapshotRestoreVolume(ctx, "snap1", "vol1").Return(nil)

	result := driver.RestoreSnapshot(ctx, snapConfig, volConfig)

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverVolumeRestoreSnapshot_Failure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap1",
		VolumeInternalName: "vol1",
	}

	mockAPI.EXPECT().SnapshotRestoreVolume(ctx, "snap1", "vol1").Return(fmt.Errorf("failed to restore volume"))

	result := driver.RestoreSnapshot(ctx, snapConfig, volConfig)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeDeleteSnapshot(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap1",
		VolumeInternalName: "vol1",
	}

	mockAPI.EXPECT().VolumeSnapshotDelete(ctx, "snap1", "vol1").Return(nil)

	result := driver.DeleteSnapshot(ctx, snapConfig, volConfig)

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverVolumeDeleteSnapshot_Failure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap1",
		VolumeInternalName: "vol1",
	}

	childVols := make([]string, 0)
	childVols = append(childVols, "vol1")

	mockAPI.EXPECT().VolumeSnapshotDelete(ctx, "snap1", "vol1").Return(api.SnapshotBusyError("snapshot is busy"))
	mockAPI.EXPECT().VolumeListBySnapshotParent(ctx, "snap1", "vol1").Return(childVols, nil)
	mockAPI.EXPECT().VolumeCloneSplitStart(ctx, "vol1").Return(nil)

	result := driver.DeleteSnapshot(ctx, snapConfig, volConfig)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeGet(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(true, nil)

	result := driver.Get(ctx, "vol1")

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverVolumeGet_Error(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, fmt.Errorf("error checking for existing volume"))

	result := driver.Get(ctx, "vol1")

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeGet_DoesNotExist(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)

	result := driver.Get(ctx, "vol1")

	assert.Error(t, result)
}

func TestOntapNasStorageDriverGetStorageBackendSpecs(t *testing.T) {
	_, driver := newMockOntapNASDriver(t)
	backend := storage.StorageBackend{}

	result := driver.GetStorageBackendSpecs(ctx, &backend)

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverGetStorageBackendPhysicalPoolNames(t *testing.T) {
	_, driver := newMockOntapNASDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}

	poolNames := driver.GetStorageBackendPhysicalPoolNames(ctx)

	assert.Equal(t, "pool1", poolNames[0], "Pool names are not equal")
}

func TestOntapNasStorageDriverGetStorageBackendPools(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
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
	assert.Equal(t, svmUUID, pools[0].SvmUUID)

	pool = pools[1]
	assert.NotNil(t, driver.physicalPools[pool.Aggregate])
	assert.Equal(t, driver.physicalPools[pool.Aggregate].Name(), pool.Aggregate)
	assert.Equal(t, svmUUID, pools[1].SvmUUID)
}

func TestOntapNasStorageDriverGetInternalVolumeName(t *testing.T) {
	_, driver := newMockOntapNASDriver(t)
	driver.Config.StoragePrefix = utils.Ptr("storagePrefix_")

	volName := driver.GetInternalVolumeName(ctx, "vol1")

	assert.Equal(t, "storagePrefix_vol1", volName, "Strings not equal")
}

func TestOntapNasStorageDriverGetProtocol(t *testing.T) {
	_, driver := newMockOntapNASDriver(t)

	protocol := driver.GetProtocol(ctx)
	assert.Equal(t, protocol, tridentconfig.File, "Protocols not equal")
}

func TestOntapNasStorageDriverGetVolumeExternalWrappers(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	channel := make(chan *storage.VolumeExternalWrapper, 1)

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeListByPrefix(ctx, gomock.Any()).Return(api.Volumes{&api.Volume{}}, nil).Times(1)

	driver.GetVolumeExternalWrappers(ctx, channel)
}

func TestOntapNasStorageDriverGetVolumeExternalWrappers_Failure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	channel := make(chan *storage.VolumeExternalWrapper, 1)

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeListByPrefix(ctx, gomock.Any()).Return(nil, fmt.Errorf("no volume found"))

	driver.GetVolumeExternalWrappers(ctx, channel)

	assert.Equal(t, 1, len(channel))
}

func TestOntapNasStorageDriverCreateFollowup_NASType_None(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	flexVol := api.Volume{
		Name:         "flexvol",
		Comment:      "flexvol",
		JunctionPath: "",
		AccessType:   "rw",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(&flexVol, nil)
	mockAPI.EXPECT().VolumeMount(ctx, "vol1", "/vol1").Return(nil)

	result := driver.CreateFollowup(ctx, volConfig)

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverCreateFollowup_NASType_SMB(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	driver.Config.NASType = "smb"

	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "cifs",
		InternalName: "vol1",
	}

	flexVol := api.Volume{
		Name:         "flexvol",
		Comment:      "flexvol",
		JunctionPath: "",
		AccessType:   "rw",
	}

	tests := []struct {
		message   string
		isSuccess bool
	}{
		{"SMB share created successfully", false},
		{"SMB share creation fail", true},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
			mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(&flexVol, nil)
			mockAPI.EXPECT().VolumeMount(ctx, "vol1", "/vol1").Return(nil)
			if test.isSuccess {
				mockAPI.EXPECT().SMBShareExists(ctx, "vol1").Return(false, nil)
				mockAPI.EXPECT().SMBShareCreate(ctx, "vol1", "/vol1").Return(nil)

				result := driver.CreateFollowup(ctx, volConfig)
				assert.NoError(t, result)
			} else {
				mockAPI.EXPECT().SMBShareExists(ctx, "vol1").Return(false, nil)
				mockAPI.EXPECT().SMBShareCreate(ctx, "vol1", "/vol1").Return(fmt.Errorf("SMB share creation failed"))

				result := driver.CreateFollowup(ctx, volConfig)
				assert.Error(t, result)
			}
		})
	}
}

func TestOntapNasStorageDriverCreateFollowup_VolumeInfoFailure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(nil, fmt.Errorf("could not find volume"))

	result := driver.CreateFollowup(ctx, volConfig)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverCreateFollowup_VolumeMountFailure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")

	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	flexVol := api.Volume{
		Name:         "flexvol",
		Comment:      "flexvol",
		JunctionPath: "",
		AccessType:   "dp",
		DPVolume:     true,
	}

	tests := []struct {
		message   string
		errorType string
	}{
		{"error mounting volume", "error"},
		{"api error", "api error"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(&flexVol, nil)
			if test.errorType == "api error" {
				mockAPI.EXPECT().VolumeMount(ctx, "vol1", "/vol1").Return(api.ApiError(test.message))
				result := driver.CreateFollowup(ctx, volConfig)
				assert.Nil(t, result)
			} else {
				mockAPI.EXPECT().VolumeMount(ctx, "vol1", "/vol1").Return(fmt.Errorf(test.message))
				result := driver.CreateFollowup(ctx, volConfig)
				assert.Error(t, result)
			}
		})
	}
}

func TestOntapNasStorageDriverCreateFollowup_WithJunctionPath_NASType_None(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	flexVol := api.Volume{
		Name:         "flexvol",
		Comment:      "flexvol",
		JunctionPath: "/vol1",
		AccessType:   "rw",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(&flexVol, nil)

	result := driver.CreateFollowup(ctx, volConfig)

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverCreateFollowup_WithJunctionPath_ROClone_Success(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:                      "1g",
		Encryption:                "false",
		FileSystem:                "nfs",
		InternalName:              "vol1",
		ReadOnlyClone:             true,
		CloneSourceVolumeInternal: "flexvol",
	}

	flexVol := api.Volume{
		Name:         "flexvol",
		Comment:      "flexvol",
		JunctionPath: "/vol1",
		AccessType:   "rw",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, "flexvol").Return(&flexVol, nil)

	result := driver.CreateFollowup(ctx, volConfig)

	assert.NoError(t, result, "error occurred")
}

func TestOntapNasStorageDriverCreateFollowup_WithJunctionPath_ROClone_Failure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:                      "1g",
		Encryption:                "false",
		FileSystem:                "nfs",
		InternalName:              "vol1",
		ReadOnlyClone:             true,
		CloneSourceVolumeInternal: "flexvol",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, "flexvol").Return(nil, api.ApiError("api error"))

	result := driver.CreateFollowup(ctx, volConfig)

	assert.Error(t, result, "expected error")
}

func TestOntapNasStorageDriverCreateFollowup_WithJunctionPath_NASType_SMB(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	driver.Config.NASType = "smb"

	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "cifs",
		InternalName: "vol1",
	}

	flexVol := api.Volume{
		Name:         "flexvol",
		Comment:      "flexvol",
		JunctionPath: "\\vol1",
		AccessType:   "rw",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(&flexVol, nil)

	result := driver.CreateFollowup(ctx, volConfig)

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverCreateFollowup_GetStoragePoolAttributes(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	mockAPI.EXPECT().IsSVMDRCapable(ctx).Return(false, nil)

	poolAttr := driver.getStoragePoolAttributes(ctx)

	assert.NotNil(t, poolAttr)
	assert.Equal(t, driver.Name(), poolAttr[BackendType].ToString())
	assert.Equal(t, "true", poolAttr[Snapshots].ToString())
	assert.Equal(t, "true", poolAttr[Clones].ToString())
	assert.Equal(t, "true", poolAttr[Encryption].ToString())
	assert.Equal(t, "false", poolAttr[Replication].ToString())
	assert.Equal(t, "thick,thin", poolAttr[ProvisioningType].ToString())
}

func TestOntapNasStorageDriverCreatePrepare(t *testing.T) {
	_, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}
	driver.CreatePrepare(ctx, volConfig)
}

func TestOntapNasStorageDriverGetUpdateType(t *testing.T) {
	mockAPI, oldDriver := newMockOntapNASDriver(t)

	oldDriver.API = mockAPI
	prefix1 := "test_"
	oldDriver.Config.StoragePrefix = &prefix1
	oldDriver.Config.Username = "user1"
	oldDriver.Config.Password = "password1"
	oldDriver.Config.Credentials = map[string]string{
		drivers.KeyName: "secret1",
		drivers.KeyType: string(drivers.CredentialStoreK8sSecret),
	}

	newDriver := newTestOntapNASDriver(ONTAPTEST_LOCALHOST, "0", ONTAPTEST_VSERVER_AGGR_NAME,
		"CSI", false, nil)
	newDriver.API = mockAPI
	prefix2 := "storage_"
	newDriver.Config.StoragePrefix = &prefix2
	newDriver.Config.Username = "user2"
	newDriver.Config.Password = "password2"
	newDriver.Config.Credentials = map[string]string{
		drivers.KeyName: "secret2",
		drivers.KeyType: string(drivers.CredentialStoreK8sSecret),
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")

	result := newDriver.GetUpdateType(ctx, oldDriver)

	expectedBitmap := &roaring.Bitmap{}
	expectedBitmap.Add(storage.UsernameChange)
	expectedBitmap.Add(storage.PasswordChange)
	expectedBitmap.Add(storage.PrefixChange)
	expectedBitmap.Add(storage.CredentialsChange)

	assert.Equal(t, expectedBitmap, result, "bitmap mismatch")
}

func TestOntapNasStorageDriverGetUpdateType_Failure(t *testing.T) {
	mockAPI, _ := newMockOntapNASDriver(t)
	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")

	oldDriver := newTestOntapSanEcoDriver(ONTAPTEST_LOCALHOST, "0", ONTAPTEST_VSERVER_AGGR_NAME, false, nil, mockAPI)
	oldDriver.API = mockAPI
	prefix1 := "test_"
	oldDriver.Config.StoragePrefix = &prefix1

	// Created a SAN driver
	newDriver := newTestOntapNASDriver(ONTAPTEST_LOCALHOST, "0", ONTAPTEST_VSERVER_AGGR_NAME,
		"CSI", false, nil)
	newDriver.API = mockAPI
	prefix2 := "storage_"
	newDriver.Config.StoragePrefix = &prefix2

	expectedBitmap := &roaring.Bitmap{}
	expectedBitmap.Add(storage.InvalidUpdate)

	result := newDriver.GetUpdateType(ctx, oldDriver)

	assert.Equal(t, expectedBitmap, result, "bitmap mismatch")
}

func TestOntapNasStorageDriverEstablishMirror(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	flexVol := api.Volume{
		Name:     "flexvol",
		Comment:  "flexvol",
		DPVolume: true,
	}
	mockAPI.EXPECT().VolumeInfo(ctx, "fakevolume1").Return(&flexVol, nil)

	snapmirror := &api.Snapmirror{
		State:              "uninitialized",
		RelationshipStatus: "idle",
	}
	snapmirror2 := &api.Snapmirror{
		State:              "snapmirrored",
		RelationshipStatus: "idle",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm1")
	mockAPI.EXPECT().SnapmirrorGet(ctx, "fakevolume1", "fakesvm1", "fakevolume2", "fakesvm2").Return(snapmirror, nil)
	mockAPI.EXPECT().SnapmirrorInitialize(ctx, "fakevolume1", "fakesvm1", "fakevolume2", "fakesvm2").Return(nil)
	mockAPI.EXPECT().SnapmirrorGet(ctx, "fakevolume1", "fakesvm1", "fakevolume2", "fakesvm2").Return(snapmirror2, nil)

	result := driver.EstablishMirror(ctx, "fakevolume1", "fakesvm2:fakevolume2", "", "")

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverEstablishMirror_WithReplicationPolicy(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	driver.Config.ReplicationPolicy = "testpolicy"
	snapmirrorPolicy := &api.SnapmirrorPolicy{
		Type: "async_mirror",
	}
	flexVol := api.Volume{
		Name:     "flexvol",
		Comment:  "flexvol",
		DPVolume: true,
	}
	snapmirror := &api.Snapmirror{
		State:              "snapmirrored",
		RelationshipStatus: "idle",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm1")
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, "testpolicy").Return(snapmirrorPolicy, nil).Times(2)
	mockAPI.EXPECT().VolumeInfo(ctx, "fakevolume1").Return(&flexVol, nil)
	mockAPI.EXPECT().SnapmirrorGet(ctx, "fakevolume1", "fakesvm1", "fakevolume2", "fakesvm2").Return(snapmirror, nil)

	result := driver.EstablishMirror(ctx, "fakevolume1", "fakesvm2:fakevolume2", "testpolicy", "")

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverEstablishMirror_WithReplicationPolicyAndSchedule(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	volName := "fakevolume1"

	driver.Config.ReplicationPolicy = "testpolicy"
	snapmirrorPolicy := &api.SnapmirrorPolicy{
		Type: "async_mirror",
	}
	flexVol := api.Volume{
		Name:     volName,
		Comment:  "flexvol",
		DPVolume: true,
	}
	snapmirror := &api.Snapmirror{
		State:              "snapmirrored",
		RelationshipStatus: "idle",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm1")
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, "testpolicy").Return(snapmirrorPolicy, nil).Times(2)
	mockAPI.EXPECT().VolumeInfo(ctx, "fakevolume1").Return(&flexVol, nil)
	mockAPI.EXPECT().SnapmirrorGet(ctx, "fakevolume1", "fakesvm1", "fakevolume2", "fakesvm2").Return(snapmirror, nil)
	mockAPI.EXPECT().JobScheduleExists(ctx, "testschedule").Return(true, nil)

	result := driver.EstablishMirror(ctx, "fakevolume1", "fakesvm2:fakevolume2", "testpolicy", "testschedule")

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverEstablishMirror_InvalidReplicationSchedule(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	driver.Config.ReplicationPolicy = "testpolicy"
	snapmirrorPolicy := &api.SnapmirrorPolicy{
		Type: "async_mirror",
	}
	flexVol := api.Volume{
		Name:     "flexvol",
		Comment:  "flexvol",
		DPVolume: true,
	}
	snapmirror := &api.Snapmirror{
		State:              "snapmirrored",
		RelationshipStatus: "idle",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm1")
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, "testpolicy").Return(snapmirrorPolicy, nil).Times(2)
	mockAPI.EXPECT().VolumeInfo(ctx, "fakevolume1").Return(&flexVol, nil)
	mockAPI.EXPECT().SnapmirrorGet(ctx, "fakevolume1", "fakesvm1", "fakevolume2", "fakesvm2").Return(snapmirror, nil)
	mockAPI.EXPECT().JobScheduleExists(ctx,
		"testschedule").Return(false, fmt.Errorf("specified replicationSchedule does not exist"))

	result := driver.EstablishMirror(ctx, "fakevolume1", "fakesvm2:fakevolume2", "testpolicy", "testschedule")

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverReestablishMirror(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	snapmirror := &api.Snapmirror{
		State:              "uninitialized",
		RelationshipStatus: "idle",
		UnhealthyReason:    "unhealthy",
	}
	snapmirror2 := &api.Snapmirror{
		State:              "snapmirrored",
		RelationshipStatus: "idle",
		IsHealthy:          true,
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm1")
	mockAPI.EXPECT().SnapmirrorGet(ctx, "fakevolume1", "fakesvm1", "fakevolume2", "fakesvm2").Return(snapmirror, nil)
	mockAPI.EXPECT().SnapmirrorGet(ctx, "fakevolume1", "fakesvm1", "fakevolume2", "fakesvm2").Return(snapmirror2, nil)
	mockAPI.EXPECT().SnapmirrorResync(ctx, "fakevolume1", "fakesvm1", "fakevolume2", "fakesvm2").Return(nil)

	result := driver.ReestablishMirror(ctx, "fakevolume1", "fakesvm2:fakevolume2", "", "")

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverReestablishMirror_WithReplicationPolicy(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	driver.Config.ReplicationPolicy = "testpolicy"
	snapmirrorPolicy := &api.SnapmirrorPolicy{
		Type: "async_mirror",
	}
	snapmirror := &api.Snapmirror{
		State:              "snapmirrored",
		RelationshipStatus: "idle",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm1")
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, "testpolicy").Return(snapmirrorPolicy, nil).Times(2)
	mockAPI.EXPECT().SnapmirrorGet(ctx, "fakevolume1", "fakesvm1", "fakevolume2", "fakesvm2").Return(snapmirror, nil)

	result := driver.ReestablishMirror(ctx, "fakevolume1", "fakesvm2:fakevolume2", "testpolicy", "")

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverReestablishMirror_WithReplicationPolicyAndSchedule(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	driver.Config.ReplicationPolicy = "testpolicy"
	snapmirrorPolicy := &api.SnapmirrorPolicy{
		Type: "async_mirror",
	}
	snapmirror := &api.Snapmirror{
		State:              "snapmirrored",
		RelationshipStatus: "idle",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm1")
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, "testpolicy").Return(snapmirrorPolicy, nil).Times(2)
	mockAPI.EXPECT().SnapmirrorGet(ctx, "fakevolume1", "fakesvm1", "fakevolume2", "fakesvm2").Return(snapmirror, nil)
	mockAPI.EXPECT().JobScheduleExists(ctx,
		"testschedule").Return(false, fmt.Errorf("specified replicationSchedule does not exist"))

	result := driver.ReestablishMirror(ctx, "fakevolume1", "fakesvm2:fakevolume2", "testpolicy",
		"testschedule")

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverPromoteMirror(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	driver.Config.ReplicationPolicy = "testpolicy"
	snapmirror := &api.Snapmirror{
		State:              "snapmirrored",
		RelationshipStatus: "idle",
	}
	snapmirrorPolicy := &api.SnapmirrorPolicy{
		Type: "async_mirror",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm1")
	mockAPI.EXPECT().SnapmirrorGet(ctx, "fakevolume1", "fakesvm1", "fakevolume2", "fakesvm2").Return(snapmirror, nil)
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, "testpolicy").Return(snapmirrorPolicy, nil)

	waitingForSnap, err := driver.PromoteMirror(ctx, "fakevolume1", "fakesvm2:fakevolume2", "snap1")

	assert.False(t, waitingForSnap)
	assert.Error(t, err)
}

func TestOntapNasStorageDriverGetMirrorStatus(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	snapmirror := &api.Snapmirror{
		State:              "snapmirrored",
		RelationshipStatus: "idle",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm1")
	mockAPI.EXPECT().SnapmirrorGet(ctx, "fakevolume1", "fakesvm1", "fakevolume2", "fakesvm2").Return(snapmirror, nil)

	status, err := driver.GetMirrorStatus(ctx, "fakevolume1", "fakesvm2:fakevolume2")

	assert.Equal(t, "established", status)
	assert.NoError(t, err)
}

func TestOntapNasStorageDriverReleaseMirror(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm1")
	mockAPI.EXPECT().SnapmirrorRelease(ctx, "fakevolume1", "fakesvm1").Return(nil)

	result := driver.ReleaseMirror(ctx, "fakevolume1")

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverGetReplicationDetails(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	snapmirror := &api.Snapmirror{
		State:               "snapmirrored",
		RelationshipStatus:  "idle",
		ReplicationPolicy:   "testpolicy",
		ReplicationSchedule: "testschedule",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm1")
	mockAPI.EXPECT().SnapmirrorGet(ctx, "fakevolume1", "fakesvm1", "fakevolume2", "fakesvm2").Return(snapmirror, nil)

	policy, schedule, SVMName, err := driver.GetReplicationDetails(ctx, "fakevolume1",
		"fakesvm2:fakevolume2")

	assert.Equal(t, "testpolicy", policy)
	assert.Equal(t, "testschedule", schedule)
	assert.Equal(t, "fakesvm1", SVMName)
	assert.NoError(t, err)
}

func TestOntapNasStorageDriverGetCommonConfig(t *testing.T) {
	_, driver := newMockOntapNASDriver(t)
	driver.Config.CommonStorageDriverConfig = &drivers.CommonStorageDriverConfig{}

	result := driver.GetCommonConfig(ctx)

	assert.NotNil(t, result)
}

func TestOntapNasStorageDriverReconcileNodeAccess(t *testing.T) {
	_, driver := newMockOntapNASDriver(t)
	nodes := make([]*utils.Node, 0)
	nodes = append(nodes, &utils.Node{Name: "node1"})

	result := driver.ReconcileNodeAccess(ctx, nodes, "1234", "")

	assert.NoError(t, result)
}

func TestNASStorageDriverGetBackendState(t *testing.T) {
	mockApi, mockDriver := newMockOntapNASDriver(t)

	mockApi.EXPECT().GetSVMState(ctx).Return("", fmt.Errorf("returning test error"))

	reason, changeMap := mockDriver.GetBackendState(ctx)
	assert.Equal(t, reason, StateReasonSVMUnreachable, "should be 'SVM is not reachable'")
	assert.NotNil(t, changeMap, "should not be nil")
}

func TestOntapNasStorageDriverResize(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	aggr := make([]string, 0)
	aggr = append(aggr, "aggr1")
	flexVol := api.Volume{
		Name:       "flexvol",
		Comment:    "flexvol",
		Aggregates: aggr,
	}
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(true, nil)
	mockAPI.EXPECT().VolumeSize(ctx, "vol1").Return(uint64(1073741824), nil)
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(&flexVol, nil).Times(2)
	mockAPI.EXPECT().VolumeSetSize(ctx, "vol1", "10737418240").Return(nil)

	result := driver.Resize(ctx, volConfig, 10737418240) // 10GB

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverResize_VolumeDoesNotExist(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)

	result := driver.Resize(ctx, volConfig, 10737418240) // 10GB

	assert.Error(t, result)
}

func TestOntapNasStorageDriverResize_SameSize(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(true, nil)
	mockAPI.EXPECT().VolumeSize(ctx, "vol1").Return(uint64(1073741824), nil)

	result := driver.Resize(ctx, volConfig, 1073741824) // 1GB

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverResize_NoVolumeInfo(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(true, nil)
	mockAPI.EXPECT().VolumeSize(ctx, "vol1").Return(uint64(1073741824), nil)
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(nil, fmt.Errorf("error fetching volume info")).Times(2)

	result := driver.Resize(ctx, volConfig, 10737418240) // 10GB

	assert.Error(t, result)
}

func TestOntapNasStorageDriverResize_WithAggregate(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	aggr := make([]string, 0)
	aggr = append(aggr, "aggr1")
	flexVol := api.Volume{
		Name:       "flexvol",
		Comment:    "flexvol",
		Aggregates: aggr,
	}
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(true, nil)
	mockAPI.EXPECT().VolumeSize(ctx, "vol1").Return(uint64(1073741824), nil)
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(&flexVol, nil).Times(2)
	mockAPI.EXPECT().VolumeSetSize(ctx, "vol1", "10737418240").Return(nil)

	result := driver.Resize(ctx, volConfig, 10737418240) // 10GB

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverResize_FakeLimitVolumeSize(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	// Added fake LimitVolumeSize value
	driver.Config.CommonStorageDriverConfig.LimitVolumeSize = "fake"
	aggr := make([]string, 0)
	aggr = append(aggr, "aggr1")
	flexVol := api.Volume{
		Name:       "flexvol",
		Comment:    "flexvol",
		Aggregates: aggr,
	}
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(true, nil)
	mockAPI.EXPECT().VolumeSize(ctx, "vol1").Return(uint64(1073741824), nil)
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(&flexVol, nil).Times(2)

	result := driver.Resize(ctx, volConfig, 10737418240) // 10GB

	assert.Error(t, result)
}

func TestOntapNasStorageDriverResize_Failure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	aggr := make([]string, 0)
	aggr = append(aggr, "aggr1")
	flexVol := api.Volume{
		Name:       "flexvol",
		Comment:    "flexvol",
		Aggregates: aggr,
	}
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(true, nil)
	mockAPI.EXPECT().VolumeSize(ctx, "vol1").Return(uint64(1073741824), nil)
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(&flexVol, nil).Times(2)
	mockAPI.EXPECT().VolumeSetSize(ctx, "vol1", "10737418240").Return(fmt.Errorf("cannot resize to specified size"))

	result := driver.Resize(ctx, volConfig, 10737418240) // 10GB

	assert.Error(t, result)
}

func TestOntapNasStorageDriverStoreConfig(t *testing.T) {
	_, driver := newMockOntapNASDriver(t)
	ontapConf := newOntapStorageDriverConfig()
	ontapConf.StorageDriverName = "ontap-nas"
	backendConfig := &storage.PersistentStorageBackendConfig{
		OntapConfig: ontapConf,
	}

	driver.StoreConfig(ctx, backendConfig)

	assert.Equal(t, &driver.Config, backendConfig.OntapConfig)
}

func TestOntapNasStorageDriverGetVolumeExternal(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	flexVol := api.Volume{
		Name:    "flexvol",
		Comment: "flexvol",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(&flexVol, nil)

	volExt, err := driver.GetVolumeExternal(ctx, "vol1")

	assert.NotNil(t, volExt)
	assert.NoError(t, err)
}

func TestOntapNasStorageDriverGetVolumeExternal_Failure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(nil, fmt.Errorf("error fetching volume info"))

	volExt, err := driver.GetVolumeExternal(ctx, "vol1")

	assert.Nil(t, volExt)
	assert.Error(t, err)
}

func TestOntapNasStorageDriverVolumeCreate(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakesvm2:vol1",
	}

	sb := &storage.StorageBackend{}
	sb.SetBackendUUID(BackendUUID)
	pool1 := storage.NewStoragePool(sb, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		SpaceReserve:      "none",
		SnapshotPolicy:    "fake-snap-policy",
		SnapshotReserve:   "10",
		UnixPermissions:   "0755",
		SnapshotDir:       "true",
		ExportPolicy:      "fake-export-policy",
		SecurityStyle:     "mixed",
		Encryption:        "false",
		TieringPolicy:     "",
		QosPolicy:         "fake-qos-policy",
		AdaptiveQosPolicy: "",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.AutoExportPolicy = true
	driver.Config.NASType = sa.SMB
	volAttrs := map[string]sa.Request{}

	tests := []struct {
		smbShare string
	}{
		{"vol1"},
		{""},
	}

	for _, test := range tests {
		t.Run(test.smbShare, func(t *testing.T) {
			driver.Config.SMBShare = test.smbShare
			mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
			mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
			mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)
			mockAPI.EXPECT().TieringPolicyValue(ctx).Return("none")
			mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(nil)
			mockAPI.EXPECT().VolumeMount(ctx, "vol1", "/vol1").Return(nil)
			if test.smbShare != "" {
				mockAPI.EXPECT().SMBShareExists(ctx, "vol1").Return(false, nil)
				mockAPI.EXPECT().SMBShareCreate(ctx, "vol1", "/").Return(nil)
			} else {
				mockAPI.EXPECT().SMBShareExists(ctx, "vol1").Return(false, nil)
				mockAPI.EXPECT().SMBShareCreate(ctx, "vol1", "/vol1").Return(nil)
			}

			result := driver.Create(ctx, volConfig, pool1, volAttrs)

			assert.NoError(t, result)
		})
	}

	assert.Equal(t, "none", volConfig.SpaceReserve)
	assert.Equal(t, "fake-snap-policy", volConfig.SnapshotPolicy)
	assert.Equal(t, "10", volConfig.SnapshotReserve)
	assert.Equal(t, "0755", volConfig.UnixPermissions)
	assert.Equal(t, "true", volConfig.SnapshotDir)
	assert.Equal(t, "trident-"+BackendUUID, volConfig.ExportPolicy)
	assert.Equal(t, "mixed", volConfig.SecurityStyle)
	assert.Equal(t, "false", volConfig.Encryption)
	assert.Equal(t, "fake-qos-policy", volConfig.QosPolicy)
	assert.Equal(t, "", volConfig.AdaptiveQosPolicy)
}

func TestOntapNasStorageDriverVolumeCreate_VolumeExists(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")

	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
		SnapshotDir:     "true",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}

	volAttrs := map[string]sa.Request{}

	tests := []struct {
		volumeExists bool
		message      string
	}{
		{true, ""},
		{true, "error checking for existing volume"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			if test.message == "" {
				mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(test.volumeExists, nil)
			} else {
				mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(test.volumeExists, fmt.Errorf(test.message))
			}
			result := driver.Create(ctx, volConfig, pool1, volAttrs)

			assert.Error(t, result)
		})
	}
}

func TestOntapNasStorageDriverVolumeCreate_PeerVolumeHandleFailure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakevol",
	}
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
		SnapshotDir:     "true",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)

	result := driver.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeCreate_NoPhysicalPool(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakesvm2:vol1",
	}

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
		SnapshotDir:     "true",
	})
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
	mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)

	result := driver.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeCreate_InvalidSnapshotReserve(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakesvm2:vol1",
	}

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
		SnapshotDir:     "true",
		SnapshotReserve: "fake",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}

	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
	mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)

	result := driver.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeCreate_VolumeSize(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")

	tests := []struct {
		volumeSize string
		valid      bool
	}{
		{"invalid", false},
		{"19m", false},
		{"-1002947563b", false},
		{"10g", true},
	}
	for _, test := range tests {
		t.Run(test.volumeSize, func(t *testing.T) {
			volConfig := &storage.VolumeConfig{
				Size:             test.volumeSize,
				Encryption:       "false",
				FileSystem:       "nfs",
				InternalName:     "vol1",
				PeerVolumeHandle: "fakesvm2:vol1",
			}

			pool1 := storage.NewStoragePool(nil, "pool1")
			pool1.SetInternalAttributes(map[string]string{
				"tieringPolicy": "none",
				SnapshotDir:     "true",
			})
			driver.physicalPools = map[string]storage.Pool{"pool1": pool1}

			volAttrs := map[string]sa.Request{}

			mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
			mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)

			if test.valid {
				mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(nil)
				mockAPI.EXPECT().VolumeMount(ctx, "vol1", "/vol1").Return(nil)
			}

			result := driver.Create(ctx, volConfig, pool1, volAttrs)
			if test.valid {
				assert.NoError(t, result)
			} else {
				assert.Error(t, result)
			}
		})
	}
}

func TestOntapNasStorageDriverVolumeCreate_LimitVolumeSize(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakesvm2:vol1",
	}

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
		SnapshotDir:     "true",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.CommonStorageDriverConfig.LimitVolumeSize = "invalid" // invalid int value
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
	mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)

	result := driver.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeCreate_InvalidSnapshotDir(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakesvm2:vol1",
	}

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
		SnapshotDir:     "invalid", // invalid bool value
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
	mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)

	result := driver.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeCreate_InvalidEncryptionValue(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "invalid", // invalid bool value
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakesvm2:vol1",
	}

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
		SnapshotDir:     "true",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
	mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)

	result := driver.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeCreate_BothQosPolicies(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakesvm2:vol1",
	}

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		TieringPolicy:     "none",
		SnapshotDir:       "true",
		QosPolicy:         "fake",
		AdaptiveQosPolicy: "fake",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
	mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)

	result := driver.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeCreate_NoAggregate(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakesvm2:vol1",
	}

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		TieringPolicy: "none",
		SnapshotDir:   "true",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.LimitAggregateUsage = "invalid"
	volAttrs := map[string]sa.Request{}

	var svmAggregateSpaceList []api.SVMAggregateSpace
	svmAggregateSpace := api.SVMAggregateSpace{}
	svmAggregateSpaceList = append(svmAggregateSpaceList, svmAggregateSpace)

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
	mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)
	mockAPI.EXPECT().GetSVMAggregateSpace(ctx, "pool1").Return(svmAggregateSpaceList, nil)

	result := driver.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeCreate_CreateFailed(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")

	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakesvm2:vol1",
	}

	sb := &storage.StorageBackend{}
	sb.SetBackendUUID(BackendUUID)
	pool1 := storage.NewStoragePool(sb, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		TieringPolicy: "",
		SnapshotDir:   "true",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.AutoExportPolicy = true
	volAttrs := map[string]sa.Request{}

	tests := []struct {
		message     string
		expectError bool
	}{
		{"volume creation failed", true},
		{"volume create job exists error", false},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
			mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)
			mockAPI.EXPECT().TieringPolicyValue(ctx).Return("none")
			if test.expectError {
				mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(fmt.Errorf(test.message))
				result := driver.Create(ctx, volConfig, pool1, volAttrs)
				assert.Error(t, result)
			} else {
				mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(api.VolumeCreateJobExistsError(test.message))
				result := driver.Create(ctx, volConfig, pool1, volAttrs)
				assert.NoError(t, result)
			}
		})
	}
}

func TestOntapNasStorageDriverVolumeCreate_SnapshotDisabled(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakesvm2:vol1",
	}

	sb := &storage.StorageBackend{}
	sb.SetBackendUUID(BackendUUID)
	pool1 := storage.NewStoragePool(sb, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		TieringPolicy: "",
		SnapshotDir:   "false",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.AutoExportPolicy = true
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
	mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)
	mockAPI.EXPECT().TieringPolicyValue(ctx).Return("none")
	mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(nil)
	mockAPI.EXPECT().VolumeModifySnapshotDirectoryAccess(ctx,
		"vol1", false).Return(fmt.Errorf("failed to disable snapshot directory access"))

	result := driver.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeCreate_IsMirrorDestination(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:                "1g",
		Encryption:          "false",
		FileSystem:          "nfs",
		InternalName:        "vol1",
		PeerVolumeHandle:    "fakesvm2:vol1",
		IsMirrorDestination: true,
	}

	sb := &storage.StorageBackend{}
	sb.SetBackendUUID(BackendUUID)
	pool1 := storage.NewStoragePool(sb, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		TieringPolicy: "",
		SnapshotDir:   "true",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.AutoExportPolicy = true
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
	mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)
	mockAPI.EXPECT().TieringPolicyValue(ctx).Return("none")
	mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(nil)

	result := driver.Create(ctx, volConfig, pool1, volAttrs)

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverVolumeCreate_MountFailed(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakesvm2:vol1",
	}

	sb := &storage.StorageBackend{}
	sb.SetBackendUUID(BackendUUID)
	pool1 := storage.NewStoragePool(sb, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		TieringPolicy: "",
		SnapshotDir:   "true",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.AutoExportPolicy = true
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
	mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)
	mockAPI.EXPECT().TieringPolicyValue(ctx).Return("none")
	mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(nil)
	mockAPI.EXPECT().VolumeMount(ctx, "vol1", "/vol1").Return(fmt.Errorf("failed to mount volume"))

	result := driver.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeCreate_LabelLengthExceeding(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakesvm2:vol1",
	}

	sb := &storage.StorageBackend{}
	sb.SetBackendUUID(BackendUUID)
	pool1 := storage.NewStoragePool(sb, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		TieringPolicy: "",
		SnapshotDir:   "true",
	})
	pool1.SetAttributes(make(map[string]sa.Offer))
	pool1.Attributes()["labels"] = sa.NewLabelOffer(map[string]string{
		"cloud": "anf",
		"thisIsATestLabelWhoseLengthShouldExceed1023Characters_AddingSomeRandomCharacters_" +
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
			"lthQFQfHVgPpUZdzZMjXry": "dev-test-cluster-1",
	})

	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.AutoExportPolicy = true
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
	mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
	mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)
	mockAPI.EXPECT().TieringPolicyValue(ctx).Return("none")

	result := driver.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeCreate_SMBShareCreatefail(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakesvm2:vol1",
	}

	sb := &storage.StorageBackend{}
	sb.SetBackendUUID(BackendUUID)
	pool1 := storage.NewStoragePool(sb, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		TieringPolicy: "",
		SnapshotDir:   "true",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.AutoExportPolicy = true
	driver.Config.NASType = sa.SMB
	volAttrs := map[string]sa.Request{}

	tests := []struct {
		smbShare string
	}{
		{"vol1"},
		{""},
	}

	for _, test := range tests {
		t.Run(test.smbShare, func(t *testing.T) {
			driver.Config.SMBShare = test.smbShare
			mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
			mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
			mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)
			mockAPI.EXPECT().TieringPolicyValue(ctx).Return("none")
			mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(nil)
			mockAPI.EXPECT().VolumeMount(ctx, "vol1", "/vol1").Return(nil)
			if test.smbShare != "" {
				mockAPI.EXPECT().SMBShareExists(ctx, "vol1").Return(false, nil)
				mockAPI.EXPECT().SMBShareCreate(ctx, "vol1", "/").Return(fmt.Errorf("cannot create volume"))
			} else {
				mockAPI.EXPECT().SMBShareExists(ctx, "vol1").Return(false, nil)
				mockAPI.EXPECT().SMBShareCreate(ctx, "vol1", "/vol1").Return(fmt.Errorf("cannot create volume"))
			}

			result := driver.Create(ctx, volConfig, pool1, volAttrs)

			assert.Error(t, result)
		})
	}
}

func TestOntapNasStorageDriverVolumeCreate_SMBShareExistsfail(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "fakesvm2:vol1",
	}

	sb := &storage.StorageBackend{}
	sb.SetBackendUUID(BackendUUID)
	pool1 := storage.NewStoragePool(sb, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		TieringPolicy: "",
		SnapshotDir:   "true",
	})
	driver.physicalPools = map[string]storage.Pool{"pool1": pool1}
	driver.Config.AutoExportPolicy = true
	driver.Config.NASType = sa.SMB
	volAttrs := map[string]sa.Request{}

	tests := []struct {
		message  string
		smbShare string
	}{
		{"User define SMB server validation fail", "vol1"},
		{"System define SMB share validation fail", ""},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			driver.Config.SMBShare = test.smbShare
			mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm")
			mockAPI.EXPECT().VolumeExists(ctx, "vol1").Return(false, nil)
			mockAPI.EXPECT().GetSVMPeers(ctx).Return([]string{"fakesvm2"}, nil)
			mockAPI.EXPECT().TieringPolicyValue(ctx).Return("none")
			mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(nil)
			mockAPI.EXPECT().VolumeMount(ctx, "vol1", "/vol1").Return(nil)
			mockAPI.EXPECT().SMBShareExists(ctx, "vol1").Return(false, fmt.Errorf("server error"))

			result := driver.Create(ctx, volConfig, pool1, volAttrs)
			assert.Error(t, result)
		})
	}
}

func TestOntapNasStorageDriverVolumeImport(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol2",
		PeerVolumeHandle: "fakesvm:vol1",
		ImportNotManaged: false,
		UnixPermissions:  DefaultUnixPermissions,
	}
	flexVol := &api.Volume{
		Name:    "flexvol",
		Comment: "flexvol",
		Size:    "1",
	}

	tests := []struct {
		NasType string
	}{
		{"nfs"},
		{"smb"},
	}

	for _, test := range tests {
		t.Run(test.NasType, func(t *testing.T) {
			mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
			mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(flexVol, nil)
			mockAPI.EXPECT().VolumeRename(ctx, "vol1", "vol2").Return(nil)
			mockAPI.EXPECT().VolumeModifyUnixPermissions(ctx, "vol2", "vol1", DefaultUnixPermissions).Return(nil)

			if test.NasType == sa.SMB {
				driver.Config.NASType = sa.SMB
			}
			result := driver.Import(ctx, volConfig, "vol1")
			assert.NoError(t, result)
		})
	}
}

func TestOntapNasStorageDriverVolumeImport_Failure(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")

	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol2",
		PeerVolumeHandle: "SVM1:vol1",
		ImportNotManaged: true,
		UnixPermissions:  DefaultUnixPermissions,
	}

	tests := []struct {
		name               string
		mockFlexvol        *api.Volume
		mockError          error
		expectedErrMessage string
	}{
		{
			"VolumeReadError",
			nil,
			api.VolumeReadError("error reading volume"),
			"error reading volume",
		},
		{
			"VolumeIdAttributesReadError",
			nil,
			api.VolumeIdAttributesReadError("error reading volume id attributes"),
			"error reading volume id attributes",
		},
		{
			"Invalid Access type",
			&api.Volume{Name: "flexvol", AccessType: "non-rw"},
			fmt.Errorf("volume vol1 type is non-rw, not rw"),
			"volume vol1 type is non-rw, not rw",
		},
		{
			"Empty Junction path of volume",
			&api.Volume{Name: "flexvol", AccessType: "rw", JunctionPath: ""},
			fmt.Errorf("junction path is not set for volume vol1"),
			"junction path is not set for volume vol1",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(test.mockFlexvol, test.mockError)
			result := driver.Import(ctx, volConfig, "vol1")
			assert.Error(t, result)
			assert.Contains(t, result.Error(), test.expectedErrMessage, "Error       message mismatch")
		})
	}
}

func TestOntapNasStorageDriverVolumeImport_RenameFailed(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol2",
		PeerVolumeHandle: "SVM1:vol1",
		ImportNotManaged: false,
		UnixPermissions:  DefaultUnixPermissions,
	}
	flexVol := &api.Volume{
		Name:    "flexvol",
		Comment: "flexvol",
		Size:    "1",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(flexVol, nil)
	mockAPI.EXPECT().VolumeRename(ctx, "vol1", "vol2").Return(fmt.Errorf("failed to rename volume"))

	result := driver.Import(ctx, volConfig, "vol1")

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeImport_ModifyComment(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol2",
		PeerVolumeHandle: "SVM1:vol1",
		ImportNotManaged: false,
		UnixPermissions:  DefaultUnixPermissions,
	}
	flexVol := &api.Volume{
		Name:    "flexvol",
		Comment: "{\"provisioning\": {\"storageDriverName\": \"ontap-nas\", \"backendName\": \"customBackendName\"}}",
		Size:    "1",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(flexVol, nil)
	mockAPI.EXPECT().VolumeRename(ctx, "vol1", "vol2").Return(nil)
	mockAPI.EXPECT().VolumeSetComment(ctx, "vol2", "vol1", "").Return(fmt.Errorf("error modifying comment"))

	result := driver.Import(ctx, volConfig, "vol1")

	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumeImport_UnixPermissions(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol2",
		PeerVolumeHandle: "SVM1:vol1",
		ImportNotManaged: false,
		UnixPermissions:  "",
	}
	flexVol := &api.Volume{
		Name:    "flexvol",
		Comment: "flexvol",
		Size:    "1",
	}

	tests := []struct {
		nasYype         string
		securityStyle   string
		unixPermissions string
	}{
		{nasYype: sa.SMB, securityStyle: "mixed", unixPermissions: DefaultUnixPermissions},
		{nasYype: sa.SMB, securityStyle: "ntfs", unixPermissions: ""},
		{nasYype: sa.NFS, securityStyle: "mixed", unixPermissions: DefaultUnixPermissions},
		{nasYype: sa.NFS, securityStyle: "unix", unixPermissions: DefaultUnixPermissions},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Unix Permissions: %d", i), func(t *testing.T) {
			driver.Config.NASType = test.nasYype
			driver.Config.SecurityStyle = test.securityStyle

			mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
			mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(flexVol, nil)
			mockAPI.EXPECT().VolumeRename(ctx, "vol1", "vol2").Return(nil)
			mockAPI.EXPECT().VolumeModifyUnixPermissions(ctx, "vol2",
				"vol1", "").Return(fmt.Errorf("error modifying unix permissions"))

			result := driver.Import(ctx, volConfig, "vol1")
			assert.Error(t, result)
		})
	}
}

func TestOntapNasStorageDriverVolumeImport_ModifySnapshotAccess(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		SnapshotDir:      "true",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol2",
		PeerVolumeHandle: "SVM1:vol1",
		ImportNotManaged: false,
		UnixPermissions:  DefaultUnixPermissions,
	}
	flexVol := &api.Volume{
		Name:    "flexvol",
		Comment: "{\"provisioning\": {\"storageDriverName\": \"ontap-nas\", \"backendName\": \"customBackendName\"}}",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(flexVol, nil)
	mockAPI.EXPECT().VolumeRename(ctx, "vol1", "vol2").Return(nil)
	mockAPI.EXPECT().VolumeSetComment(ctx, "vol2", "vol1", "").Return(nil)
	mockAPI.EXPECT().VolumeModifyUnixPermissions(ctx, "vol2", "vol1", DefaultUnixPermissions).Return(nil)
	mockAPI.EXPECT().VolumeModifySnapshotDirectoryAccess(ctx, "vol2", true).Return(nil)

	result := driver.Import(ctx, volConfig, "vol1")

	assert.NoError(t, result, "An error occurred")
}

func TestOntapNasStorageDriverVolumeImport_FailedModifySnapshotAccess(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		SnapshotDir:      "true",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol2",
		PeerVolumeHandle: "SVM1:vol1",
		ImportNotManaged: false,
		UnixPermissions:  DefaultUnixPermissions,
	}
	flexVol := &api.Volume{
		Name:    "flexvol",
		Comment: "{\"provisioning\": {\"storageDriverName\": \"ontap-nas\", \"backendName\": \"customBackendName\"}}",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(flexVol, nil)
	mockAPI.EXPECT().VolumeRename(ctx, "vol1", "vol2").Return(nil)
	mockAPI.EXPECT().VolumeSetComment(ctx, "vol2", "vol1", "").Return(nil)
	mockAPI.EXPECT().VolumeModifyUnixPermissions(ctx, "vol2", "vol1", DefaultUnixPermissions).Return(nil)
	mockAPI.EXPECT().VolumeModifySnapshotDirectoryAccess(ctx, "vol2", true).Return(mockError)

	result := driver.Import(ctx, volConfig, "vol1")

	assert.Error(t, result, "An error is expected")
}

func TestOntapNasStorageDriverVolumeImport_SMBShareCreateFail(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol2",
		PeerVolumeHandle: "fakesvm:vol1",
		ImportNotManaged: false,
		UnixPermissions:  DefaultUnixPermissions,
	}
	driver.Config.NASType = sa.SMB
	flexVol := &api.Volume{
		Name:         "flexvol",
		Comment:      "flexvol",
		JunctionPath: "/flexvol",
		Size:         "1",
	}
	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().VolumeInfo(ctx, "vol1").Return(flexVol, nil)
	mockAPI.EXPECT().VolumeRename(ctx, "vol1", "vol2").Return(nil)
	mockAPI.EXPECT().VolumeModifyUnixPermissions(ctx, "vol2", "vol1", DefaultUnixPermissions).Return(nil)
	mockAPI.EXPECT().SMBShareExists(ctx, "vol1").Return(false, nil)
	mockAPI.EXPECT().SMBShareCreate(ctx, "vol1", "/vol1").Return(fmt.Errorf("error creating SMB share"))

	result := driver.Import(ctx, volConfig, "vol1")
	assert.Error(t, result)
}

func TestOntapNasStorageDriverVolumePublish_NASType_None(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "nfs",
		InternalName:     "vol1",
		PeerVolumeHandle: "SVM1:vol1",
		ImportNotManaged: false,
		UnixPermissions:  "",
		MountOptions:     "-o nfsvers=3",
	}
	volConfig.AccessInfo.NfsPath = "/nfs"

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")

	result := driver.Publish(ctx, volConfig, &utils.VolumePublishInfo{})

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverVolumePublish_NASType_SMB(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)
	driver.Config.NASType = "smb"

	volConfig := &storage.VolumeConfig{
		Size:             "1g",
		Encryption:       "false",
		FileSystem:       "cifs",
		InternalName:     "vol1",
		PeerVolumeHandle: "SVM1:vol1",
		ImportNotManaged: false,
	}
	volConfig.AccessInfo.SMBPath = "/test_cifs_path"

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")

	result := driver.Publish(ctx, volConfig, &utils.VolumePublishInfo{})

	assert.NoError(t, result)
}

func TestOntapNasStorageDriverGetTelemetry(t *testing.T) {
	_, driver := newMockOntapNASDriver(t)
	driver.telemetry = &Telemetry{
		Plugin:        driver.Name(),
		SVM:           "SVM1",
		StoragePrefix: *driver.GetConfig().StoragePrefix,
		Driver:        driver,
		done:          make(chan struct{}),
	}

	assert.True(t, reflect.DeepEqual(driver.telemetry, driver.GetTelemetry()))
}

func TestOntapNasStorageDriverBackendName(t *testing.T) {
	_, driver := newMockOntapNASDriver(t)
	driver.Config.BackendName = "myBackend"

	result := driver.BackendName()

	assert.Equal(t, result, "myBackend")
}

func TestOntapNasStorageDriverUpdateMirror(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	mockAPI.EXPECT().SnapmirrorUpdate(ctx, "testVol", "testSnap")

	err := driver.UpdateMirror(ctx, "testVol", "testSnap")
	assert.Error(t, err, "expected error")
}

func TestOntapNasStorageDriverCheckMirrorTransferState(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	snapmirror := &api.Snapmirror{
		State:              "snapmirrored",
		RelationshipStatus: "idle",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm1")
	mockAPI.EXPECT().SnapmirrorGet(ctx, "fakevolume1", "fakesvm1", "", "").Return(snapmirror, nil)

	result, err := driver.CheckMirrorTransferState(ctx, "fakevolume1")

	assert.Nil(t, result, "expected nil")
	assert.Error(t, err, "expected error")
}

func TestOntapStorageDriverGetMirrorTransferTime(t *testing.T) {
	mockAPI, driver := newMockOntapNASDriver(t)

	timeNow := time.Now()
	snapmirror := &api.Snapmirror{
		State:              "snapmirrored",
		RelationshipStatus: "idle",
		EndTransferTime:    &timeNow,
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("fakesvm1")
	mockAPI.EXPECT().SnapmirrorGet(ctx, "fakevolume1", "fakesvm1", "", "").Return(snapmirror, nil)

	result, err := driver.GetMirrorTransferTime(ctx, "fakevolume1")
	assert.NotNil(t, result, "received nil")
	assert.NoError(t, err, "received error")
}
