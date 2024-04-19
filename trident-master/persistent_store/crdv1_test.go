// Copyright 2023 NetApp, Inc. All Rights Reserved.

package persistentstore

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	discoveryfake "k8s.io/client-go/discovery/fake"
	clientgotesting "k8s.io/client-go/testing"

	k8sclient "github.com/netapp/trident/cli/k8s_client"
	"github.com/netapp/trident/config"
	. "github.com/netapp/trident/logging"
	"github.com/netapp/trident/persistent_store/crd/client/clientset/versioned"
	"github.com/netapp/trident/persistent_store/crd/client/clientset/versioned/fake"
	v1 "github.com/netapp/trident/persistent_store/crd/client/clientset/versioned/typed/netapp/v1"
	v1fake "github.com/netapp/trident/persistent_store/crd/client/clientset/versioned/typed/netapp/v1/fake"
	"github.com/netapp/trident/storage"
	storageattribute "github.com/netapp/trident/storage_attribute"
	storageclass "github.com/netapp/trident/storage_class"
	drivers "github.com/netapp/trident/storage_drivers"
	"github.com/netapp/trident/storage_drivers/ontap"
	"github.com/netapp/trident/storage_drivers/solidfire"
	tridenttesting "github.com/netapp/trident/testing"
	"github.com/netapp/trident/utils"
)

var (
	storagePool = "aggr1"
	ctx         = context.Background

	crdScheme = runtime.NewScheme()
	codecs    = serializer.NewCodecFactory(crdScheme)
)

func init() {
	utilruntime.Must(fake.AddToScheme(crdScheme))
}

func GetTestKubernetesClient() (*CRDClientV1, k8sclient.KubernetesClient) {
	client := NewFakeClientset()
	k8sclientFake, _ := k8sclient.NewFakeKubeClient()

	return &CRDClientV1{
		crdClient: client,
		k8sClient: k8sclientFake,
		version: &config.PersistentStateVersion{
			PersistentStoreVersion: string(CRDV1Store),
			OrchestratorAPIVersion: config.OrchestratorAPIVersion,
		},
	}, k8sclientFake
}

// NewFakeClientset is a better replacement for the autogenerated fake.NewSimpleClientset that
// uses our own improved testing fixtures.
func NewFakeClientset(objects ...runtime.Object) *Clientset {
	o := tridenttesting.NewObjectTracker(crdScheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &discoveryfake.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", tridenttesting.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action clientgotesting.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watchIntf, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watchIntf, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.  Uses Trident's improved ObjectTracker.
type Clientset struct {
	clientgotesting.Fake
	discovery *discoveryfake.FakeDiscovery
	tracker   tridenttesting.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() tridenttesting.ObjectTracker {
	return c.tracker
}

var _ versioned.Interface = &Clientset{}

// TridentV1 retrieves the TridentV1Client
func (c *Clientset) TridentV1() v1.TridentV1Interface {
	return &v1fake.FakeTridentV1{Fake: &c.Fake}
}

func TestKubernetesBackend(t *testing.T) {
	p, _ := GetTestKubernetesClient()

	// Adding storage backend
	NFSServerConfig := drivers.OntapStorageDriverConfig{
		CommonStorageDriverConfig: &drivers.CommonStorageDriverConfig{
			StorageDriverName: config.OntapNASStorageDriverName,
		},
		ManagementLIF: "10.0.0.4",
		DataLIF:       "10.0.0.100",
		SVM:           "svm1",
		Username:      "admin",
		Password:      "netapp",
	}
	NFSDriver := ontap.NASStorageDriver{
		Config: NFSServerConfig,
	}
	NFSServer := &storage.StorageBackend{}
	NFSServer.SetDriver(&NFSDriver)
	NFSServer.SetName("nfs-server-1-" + NFSServerConfig.ManagementLIF)
	NFSServer.SetBackendUUID(uuid.New().String())

	err := p.AddBackend(ctx(), NFSServer)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	// Getting a storage backend
	var ontapConfig drivers.OntapStorageDriverConfig
	recoveredBackend, err := p.GetBackend(ctx(), NFSServer.Name())
	if err != nil {
		t.Fatal(err.Error())
	}
	configJSON, err := recoveredBackend.MarshalConfig()
	if err != nil {
		t.Fatal("Unable to marshal recovered backend config: ", err)
	}
	if err = json.Unmarshal([]byte(configJSON), &ontapConfig); err != nil {
		t.Error("Unable to unmarshal backend into ontap configuration: ", err)
	} else if ontapConfig.SVM != NFSServerConfig.SVM {
		t.Error("Backend state doesn't match!")
	}

	// Updating a storage backend
	NFSServerNewConfig := drivers.OntapStorageDriverConfig{
		CommonStorageDriverConfig: &drivers.CommonStorageDriverConfig{
			StorageDriverName: config.OntapNASStorageDriverName,
		},
		ManagementLIF: "10.0.0.4",
		DataLIF:       "10.0.0.100",
		SVM:           "svm1",
		Username:      "admin",
		Password:      "NETAPP",
	}
	NFSDriver.Config = NFSServerNewConfig
	err = p.UpdateBackend(ctx(), NFSServer)
	if err != nil {
		t.Error(err.Error())
	}
	recoveredBackend, err = p.GetBackend(ctx(), NFSServer.Name())
	if err != nil {
		t.Fatal(err.Error())
	}
	configJSON, err = recoveredBackend.MarshalConfig()
	if err != nil {
		t.Fatal("Unable to marshal recovered backend config: ", err)
	}
	if err = json.Unmarshal([]byte(configJSON), &ontapConfig); err != nil {
		t.Error("Unable to unmarshal backend into ontap configuration: ", err)
	} else if ontapConfig.SVM != NFSServerConfig.SVM {
		t.Error("Backend state doesn't match!")
	}

	// Deleting a storage backend
	err = p.DeleteBackend(ctx(), NFSServer)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestKubernetesBackends(t *testing.T) {
	p, _ := GetTestKubernetesClient()

	var persistentBackends []*storage.BackendPersistent
	var backends []*storage.StorageBackend
	var err error

	// Adding storage backends
	for i := 1; i <= 5; i++ {
		NFSServerConfig := drivers.OntapStorageDriverConfig{
			CommonStorageDriverConfig: &drivers.CommonStorageDriverConfig{
				StorageDriverName: config.OntapNASStorageDriverName,
			},
			ManagementLIF: "10.0.0." + strconv.Itoa(i),
			DataLIF:       "10.0.0.100",
			SVM:           "svm" + strconv.Itoa(i),
			Username:      "admin",
			Password:      "netapp",
		}
		NFSServer := &storage.StorageBackend{}
		NFSServer.SetDriver(&ontap.NASStorageDriver{
			Config: NFSServerConfig,
		})
		NFSServer.SetName("nfs-server-" + strconv.Itoa(i) + "-" + NFSServerConfig.ManagementLIF)
		NFSServer.SetBackendUUID(uuid.New().String())

		backends = append(backends, NFSServer)

		err = p.AddBackend(ctx(), NFSServer)
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		Log().Info(">>CHECKING BACKENDS.")
		persistentBackends, err = p.GetBackends(ctx())
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}
		for _, backend := range persistentBackends {
			secretName := p.backendSecretName(backend.BackendUUID)
			Log().WithFields(LogFields{
				"backend.Name":        backend.Name,
				"backend.BackendUUID": backend.BackendUUID,
				"backend":             backend,
				"secretName":          secretName,
			}).Info("Currently have.")

			secret, err := p.k8sClient.GetSecret(secretName)
			if err != nil {
				t.Error(err.Error())
				t.FailNow()
			}
			Log().WithFields(LogFields{
				"secret": secret,
			}).Info("Found secret.")

			// Decode secret data into map.  The fake client returns only StringData while the real
			// API returns only Data, so we must use both here to support the unit tests.
			secretMap := make(map[string]string)
			for key, value := range secret.Data {
				secretMap[key] = string(value)
			}
			for key, value := range secret.StringData {
				secretMap[key] = value
			}

			if secretMap["Username"] != NFSServerConfig.Username {
				t.FailNow()
			}
			if secretMap["Password"] != NFSServerConfig.Password {
				t.FailNow()
			}
		}
		Log().Info("<<CHECKING BACKENDS.")
	}

	// Retrieving all backends
	if persistentBackends, err = p.GetBackends(ctx()); err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if len(persistentBackends) != 5 {
		t.Error("Didn't retrieve all the backends!")
	}

	// Delete all backends
	for _, backend := range backends {
		if err = p.DeleteBackend(ctx(), backend); err != nil {
			t.Error(err.Error())
		}
	}

	if persistentBackends, err = p.GetBackends(ctx()); err != nil {
		t.Error(err.Error())
	}
	if len(persistentBackends) != 0 {
		t.Error("Deleting backends failed!")
	}
}

func TestKubernetesDuplicateBackend(t *testing.T) {
	p, _ := GetTestKubernetesClient()

	NFSServerConfig := drivers.OntapStorageDriverConfig{
		CommonStorageDriverConfig: &drivers.CommonStorageDriverConfig{
			StorageDriverName: config.OntapNASStorageDriverName,
		},
		ManagementLIF: "10.0.0.4",
		DataLIF:       "10.0.0.100",
		SVM:           "svm1",
		Username:      "admin",
		Password:      "netapp",
	}
	NFSServer := &storage.StorageBackend{}
	NFSServer.SetDriver(&ontap.NASStorageDriver{
		Config: NFSServerConfig,
	})
	NFSServer.SetName("nfs-server-1-" + NFSServerConfig.ManagementLIF)
	NFSServer.SetBackendUUID(uuid.New().String())

	err := p.AddBackend(ctx(), NFSServer)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	err = p.AddBackend(ctx(), NFSServer)
	if err == nil {
		t.Error("Second Create should have failed!")
	}

	err = p.DeleteBackend(ctx(), NFSServer)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestKubernetesVolume(t *testing.T) {
	p, _ := GetTestKubernetesClient()

	// Adding a volume
	NFSServerConfig := drivers.OntapStorageDriverConfig{
		CommonStorageDriverConfig: &drivers.CommonStorageDriverConfig{
			StorageDriverName: config.OntapNASStorageDriverName,
		},
		ManagementLIF: "10.0.0.4",
		DataLIF:       "10.0.0.100",
		SVM:           "svm1",
		Username:      "admin",
		Password:      "netapp",
	}
	NFSServer := &storage.StorageBackend{}
	NFSServer.SetDriver(&ontap.NASStorageDriver{
		Config: NFSServerConfig,
	})
	NFSServer.SetName("NFS_server-" + NFSServerConfig.ManagementLIF)
	NFSServer.SetBackendUUID(uuid.New().String())
	vol1Config := storage.VolumeConfig{
		Version:      config.OrchestratorAPIVersion,
		Name:         "vol1",
		Size:         "1GB",
		Protocol:     config.File,
		StorageClass: "gold",
	}
	vol1 := &storage.Volume{
		Config:      &vol1Config,
		BackendUUID: NFSServer.BackendUUID(),
		Pool:        storagePool,
	}
	err := p.AddVolume(ctx(), vol1)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	// Getting a volume
	var recoveredVolume *storage.VolumeExternal
	recoveredVolume, err = p.GetVolume(ctx(), vol1.Config.Name)
	if err != nil {
		t.Fatal(err.Error())
	}
	if recoveredVolume.BackendUUID != vol1.BackendUUID || recoveredVolume.Config.Size != vol1.Config.Size {
		t.Error("Recovered volume does not match!")
	}

	// Test idempotent replacement
	err = p.AddVolume(ctx(), vol1)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	// Getting a volume
	recoveredVolume, err = p.GetVolume(ctx(), vol1.Config.Name)
	if err != nil {
		t.Fatal(err.Error())
	}
	if recoveredVolume.BackendUUID != vol1.BackendUUID || recoveredVolume.Config.Size != vol1.Config.Size {
		t.Error("Recovered volume does not match!")
	}

	// Updating a volume
	vol1Config.Size = "2GB"
	err = p.UpdateVolume(ctx(), vol1)
	if err != nil {
		t.Error(err.Error())
	}
	recoveredVolume, err = p.GetVolume(ctx(), vol1.Config.Name)
	if err != nil {
		t.Fatal(err.Error())
	}
	if recoveredVolume.Config.Size != vol1Config.Size {
		t.Error("Volume update failed!")
	}

	// Deleting a volume
	err = p.DeleteVolume(ctx(), vol1)
	if err != nil {
		t.Error(err.Error())
	}

	tvol, err := p.crdClient.TridentV1().TridentVolumes(p.namespace).Get(ctx(), vol1.Config.Name, getOpts)
	if err != nil || tvol == nil || !tvol.HasTridentFinalizers() || tvol.DeletionTimestamp.IsZero() {
		t.Fatalf("Volume should have been updated; %v", err)
	}

	// Remove finalizers to ensure volume is deleted
	tvol.RemoveTridentFinalizers()
	_, _ = p.crdClient.TridentV1().TridentVolumes(p.namespace).Update(ctx(), tvol, updateOpts)

	_, getErr := p.GetVolume(ctx(), vol1.Config.Name)
	if !errors.IsNotFound(getErr) {
		t.Fatalf("Volume should have been deleted; %v", err)
	}

	// Deleting a non-existent volume
	err = p.DeleteVolume(ctx(), vol1)
	if err != nil {
		t.Error("DeleteVolume should have succeeded.")
	}
}

func TestKubernetesVolumes(t *testing.T) {
	var err error

	p, _ := GetTestKubernetesClient()

	// Adding volumes
	NFSServerConfig := drivers.OntapStorageDriverConfig{
		CommonStorageDriverConfig: &drivers.CommonStorageDriverConfig{
			StorageDriverName: config.OntapNASStorageDriverName,
		},
		ManagementLIF: "10.0.0.4",
		DataLIF:       "10.0.0.100",
		SVM:           "svm1",
		Username:      "admin",
		Password:      "netapp",
	}
	NFSServer := &storage.StorageBackend{}
	NFSServer.SetDriver(&ontap.NASStorageDriver{
		Config: NFSServerConfig,
	})
	NFSServer.SetName("NFS_server-" + NFSServerConfig.ManagementLIF)
	NFSServer.SetBackendUUID(uuid.New().String())

	for i := 1; i <= 5; i++ {
		volConfig := storage.VolumeConfig{
			Version:      config.OrchestratorAPIVersion,
			Name:         "vol" + strconv.Itoa(i),
			Size:         strconv.Itoa(i) + "GB",
			Protocol:     config.File,
			StorageClass: "gold",
		}
		vol := &storage.Volume{
			Config:      &volConfig,
			BackendUUID: NFSServer.BackendUUID(),
		}
		err = p.AddVolume(ctx(), vol)
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}
	}

	// Retrieving all volumes
	var volumes []*storage.VolumeExternal
	if volumes, err = p.GetVolumes(ctx()); err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if len(volumes) != 5 {
		t.Errorf("Expected %d volumes; retrieved %d", 5, len(volumes))
	}

	// Deleting all volumes
	if err = p.DeleteVolumes(ctx()); err != nil {
		t.Error(err.Error())
	}
	if volumes, err = p.GetVolumes(ctx()); err != nil {
		t.Error(err.Error())
	} else if len(volumes) != 0 {
		t.Error("Deleting volumes failed!")
	}
}

func TestKubernetesVolumeTransactions(t *testing.T) {
	var err error

	p, _ := GetTestKubernetesClient()

	// Adding volume transactions
	for i := 1; i <= 5; i++ {
		volConfig := storage.VolumeConfig{
			Version:      config.OrchestratorAPIVersion,
			Name:         "vol" + strconv.Itoa(i),
			Size:         strconv.Itoa(i) + "GB",
			Protocol:     config.File,
			StorageClass: "gold",
		}
		volTxn := &storage.VolumeTransaction{
			Config: &volConfig,
			Op:     storage.AddVolume,
		}
		if err = p.AddVolumeTransaction(ctx(), volTxn); err != nil {
			t.Error(err.Error())
			t.FailNow()
		}
	}

	// Retrieving volume transactions
	volTxns, err := p.GetVolumeTransactions(ctx())
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if len(volTxns) != 5 {
		t.Error("Did n't retrieve all volume transactions!")
	}

	// Getting and Deleting volume transactions
	for _, v := range volTxns {
		txn, err := p.GetVolumeTransaction(ctx(), v)
		if err != nil {
			t.Errorf("Unable to get existing volume transaction %s:  %v", v.Config.Name, err)
		}
		if !reflect.DeepEqual(txn, v) {
			t.Errorf("Got incorrect volume transaction for %s (got %s)", v.Config.Name, txn.Config.Name)
		}
		if err = p.DeleteVolumeTransaction(ctx(), v); err != nil {
			t.Errorf("Unable to delete volume transaction: %v", err)
		}

		deletedTxn, err := p.GetVolumeTransaction(ctx(), v)
		if deletedTxn != nil || err != nil && !errors.IsNotFound(err) {
			t.Fatalf("Transaction should have been deleted; %v", err)
		}

		// Deleting a non-existent transaction
		err = p.DeleteVolumeTransaction(ctx(), v)
		if err != nil {
			t.Error("DeleteVolumeTransaction should have succeeded.")
		}
	}
	volTxns, err = p.GetVolumeTransactions(ctx())
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if len(volTxns) != 0 {
		t.Error("Didn't delete all volume transactions!")
	}
}

func TestKubernetesDuplicateVolumeTransaction(t *testing.T) {
	var err error

	p, _ := GetTestKubernetesClient()

	firstTxn := &storage.VolumeTransaction{
		Config: &storage.VolumeConfig{
			Version:      config.OrchestratorAPIVersion,
			Name:         "testVol",
			Size:         "1 GB",
			Protocol:     config.File,
			StorageClass: "gold",
		},
		Op: storage.AddVolume,
	}
	secondTxn := &storage.VolumeTransaction{
		Config: &storage.VolumeConfig{
			Version:      config.OrchestratorAPIVersion,
			Name:         "testVol",
			Size:         "1 GB",
			Protocol:     config.File,
			StorageClass: "silver",
		},
		Op: storage.AddVolume,
	}

	if err = p.AddVolumeTransaction(ctx(), firstTxn); err != nil {
		t.Error("Unable to add first volume transaction: ", err)
	}
	if err = p.AddVolumeTransaction(ctx(), secondTxn); err == nil {
		t.Error("Should not have been able to add second volume transaction: ", err)
	}
	volTxns, err := p.GetVolumeTransactions(ctx())
	if err != nil {
		t.Fatal("Unable to retrieve volume transactions: ", err)
	}
	if len(volTxns) != 1 {
		t.Errorf("Expected one volume transaction; got %d", len(volTxns))
	} else if volTxns[0].Config.StorageClass != "gold" {
		t.Errorf("Vol transaction changed.  Expected storage class gold; got %s.", volTxns[0].Config.StorageClass)
	}
	for i, volTxn := range volTxns {
		if err = p.DeleteVolumeTransaction(ctx(), volTxn); err != nil {
			t.Errorf("Unable to delete volume transaction %d:  %v", i+1, err)
		}
	}
}

func TestKubernetesAddSolidFireBackend(t *testing.T) {
	var err error

	p, _ := GetTestKubernetesClient()

	sfConfig := drivers.SolidfireStorageDriverConfig{
		CommonStorageDriverConfig: &drivers.CommonStorageDriverConfig{
			StorageDriverName: config.SolidfireSANStorageDriverName,
		},
		TenantName: "docker",
	}
	sfBackend := &storage.StorageBackend{}
	sfBackend.SetDriver(&solidfire.SANStorageDriver{
		Config: sfConfig,
	})
	sfBackend.SetName("solidfire" + "_10.0.0.9")
	sfBackend.SetBackendUUID(uuid.New().String())
	if err = p.AddBackend(ctx(), sfBackend); err != nil {
		t.Fatal(err.Error())
	}

	var retrievedConfig drivers.SolidfireStorageDriverConfig
	recoveredBackend, err := p.GetBackend(ctx(), sfBackend.Name())
	if err != nil {
		t.Fatal(err.Error())
	}
	configJSON, err := recoveredBackend.MarshalConfig()
	if err != nil {
		t.Fatal("Unable to marshal recovered backend config: ", err)
	}
	if err = json.Unmarshal([]byte(configJSON), &retrievedConfig); err != nil {
		t.Error("Unable to unmarshal backend into ontap configuration: ", err)
	} else if retrievedConfig.Size != sfConfig.Size {
		t.Errorf("Backend state doesn't match: %v != %v", retrievedConfig.Size, sfConfig.Size)
	}

	if err = p.DeleteBackend(ctx(), sfBackend); err != nil {
		t.Fatal(err.Error())
	}
}

func TestKubernetesAddStorageClass(t *testing.T) {
	var err error

	p, _ := GetTestKubernetesClient()

	bronzeConfig := &storageclass.Config{
		Name:            "bronze",
		Attributes:      make(map[string]storageattribute.Request),
		AdditionalPools: make(map[string][]string),
	}
	bronzeConfig.Attributes["media"] = storageattribute.NewStringRequest("hdd")
	bronzeConfig.AdditionalPools["ontapnas_10.0.207.101"] = []string{"aggr1"}
	bronzeConfig.AdditionalPools["ontapsan_10.0.207.103"] = []string{"aggr1"}
	bronzeClass := storageclass.New(bronzeConfig)

	if err := p.AddStorageClass(ctx(), bronzeClass); err != nil {
		t.Fatal(err.Error())
	}

	retrievedSC, err := p.GetStorageClass(ctx(), bronzeConfig.Name)
	if err != nil {
		t.Error(err.Error())
	}

	sc := storageclass.NewFromPersistent(retrievedSC)
	// Validating correct retrieval of storage class attributes
	retrievedAttrs := sc.GetAttributes()
	if _, ok := retrievedAttrs["media"]; !ok {
		t.Error("Could not find storage class attribute!")
	}
	if retrievedAttrs["media"].Value().(string) != "hdd" || retrievedAttrs["media"].GetType() != "string" {
		t.Error("Could not retrieve storage class attribute!")
	}

	// Validating correct retrieval of storage pools in a storage class
	backendVCs := sc.GetAdditionalStoragePools()
	for k, v := range bronzeConfig.AdditionalPools {
		if vcs, ok := backendVCs[k]; !ok {
			t.Errorf("Could not find backend %s for the storage class!", k)
		} else {
			if len(vcs) != len(v) {
				t.Errorf("Could not retrieve the correct number of storage pools!")
			} else {
				for i, vc := range vcs {
					if vc != v[i] {
						t.Errorf("Could not find storage pools %s for the storage class!", vc)
					}
				}
			}
		}
	}

	if err := p.DeleteStorageClass(ctx(), bronzeClass); err != nil {
		t.Fatal(err.Error())
	}

	tsc, err := p.crdClient.TridentV1().TridentStorageClasses(p.namespace).Get(ctx(), bronzeClass.GetName(), getOpts)
	if err != nil || tsc == nil || !tsc.HasTridentFinalizers() || tsc.DeletionTimestamp.IsZero() {
		t.Fatalf("Storage class should have been updated; %v", err)
	}

	// Remove finalizers to ensure storage class is deleted
	tsc.RemoveTridentFinalizers()
	_, _ = p.crdClient.TridentV1().TridentStorageClasses(p.namespace).Update(ctx(), tsc, updateOpts)

	_, err = p.GetStorageClass(ctx(), bronzeClass.GetName())
	if !errors.IsNotFound(err) {
		t.Fatalf("Storage class should have been deleted; %v", err)
	}

	// Delete a non-existent storage class
	err = p.DeleteStorageClass(ctx(), bronzeClass)
	if err != nil {
		t.Error("DeleteStorageClass should have succeeded.")
	}
}

func TestKubernetesReplaceBackendAndUpdateVolumes(t *testing.T) {
	var err error

	p, _ := GetTestKubernetesClient()

	// Initialize the state by adding a storage backend and volumes
	NFSServerConfig := drivers.OntapStorageDriverConfig{
		CommonStorageDriverConfig: &drivers.CommonStorageDriverConfig{
			StorageDriverName: config.OntapNASStorageDriverName,
		},
		ManagementLIF: "10.0.0.4",
		DataLIF:       "10.0.0.100",
		SVM:           "svm1",
		Username:      "admin",
		Password:      "netapp",
	}
	NFSDriver := ontap.NASStorageDriver{
		Config: NFSServerConfig,
	}
	NFSServer := &storage.StorageBackend{}
	NFSServer.SetDriver(&NFSDriver)
	NFSServer.SetName("ontapnas_" + NFSServerConfig.DataLIF)
	NFSServer.SetBackendUUID(uuid.New().String())
	err = p.AddBackend(ctx(), NFSServer)
	if err != nil {
		t.Fatalf("Backend creation failed: %v\n", err)
	}

	recoveredBackend, err := p.GetBackend(ctx(), NFSServer.Name())
	if err != nil {
		t.Error(err.Error())
		t.Fatalf("Backend lookup failed for:%v err: %v\n", recoveredBackend, err)
	}
	Log().WithFields(LogFields{
		"recoveredBackend.Name":        recoveredBackend.Name,
		"recoveredBackend.BackendUUID": recoveredBackend.BackendUUID,
	}).Debug("recoveredBackend")

	for i := 0; i < 5; i++ {
		volConfig := storage.VolumeConfig{
			Version:      config.OrchestratorAPIVersion,
			Name:         fmt.Sprintf("vol%d", i),
			Size:         "1GB",
			Protocol:     config.File,
			StorageClass: "gold",
		}
		vol := &storage.Volume{
			Config:      &volConfig,
			BackendUUID: recoveredBackend.BackendUUID,
			Pool:        storagePool,
		}
		err = p.AddVolume(ctx(), vol)
	}

	backends, err := p.GetBackends(ctx())
	if err != nil || len(backends) != 1 {
		t.Fatalf("Backend retrieval failed; backends:%v err:%v\n", backends, err)
	}
	Log().Debugf("GetBackends: %v, %v\n", backends, err)
	backend, err := p.GetBackend(ctx(), backends[0].Name)
	if err != nil || backend.Name != NFSServer.Name() {
		t.Fatalf("Backend retrieval failed; backend: %v err: %v\n", backends[0].Name, err)
	}
	Log().Debugf("GetBackend(%v): %v, %v\n", backends[0].Name, backend, err)
	volumes, err := p.GetVolumes(ctx())
	if err != nil || len(volumes) != 5 {
		t.Fatalf("Volume retrieval failed; volumes:%v err:%v\n", volumes, err)
	}
	Log().Debugf("GetVolumes: %v, %v\n", volumes, err)
	for i := 0; i < 5; i++ {
		volume, err := p.GetVolume(ctx(), fmt.Sprintf("vol%d", i))
		Log().WithFields(LogFields{
			"volume": volume,
		}).Debug("GetVolumes.")
		if err != nil {
			t.Fatalf("Volume retrieval failed; volume:%v err:%v\n", volume, err)
		}
		Log().Debugf("GetVolume(vol%v): %v, %v\n", i, volume, err)
	}

	newNFSServer := &storage.StorageBackend{}
	newNFSServer.SetDriver(&NFSDriver)
	// Renaming the NFS server
	newNFSServer.SetName("AFF")
	newNFSServer.SetBackendUUID(NFSServer.BackendUUID())
	err = p.ReplaceBackendAndUpdateVolumes(ctx(), NFSServer, newNFSServer)
	if err != nil {
		t.Fatalf("ReplaceBackendAndUpdateVolumes failed: %v\n", err)
	}

	// Validate successful renaming of the backend
	backends, err = p.GetBackends(ctx())
	if err != nil || len(backends) != 1 {
		t.Fatalf("Backend retrieval failed; backends:%v err:%v\n", backends, err)
	}
	Log().Debugf("GetBackends: %v, %v\n", backends, err)
	backend, err = p.GetBackend(ctx(), backends[0].Name)
	if err != nil || backend.Name != newNFSServer.Name() {
		t.Fatalf("Backend retrieval failed; backend.Name: %v newNFSServer.Name: %v err:%v\n", backends[0].Name,
			newNFSServer.Name(), err)
	}
	Log().Debugf("GetBackend(%v): %v, %v\n", backends[0].Name, backend, err)

	// Validate successful renaming of the volumes
	volumes, err = p.GetVolumes(ctx())
	if err != nil || len(volumes) != 5 {
		t.Fatalf("Volume retrieval failed; volumes:%v err:%v\n", volumes, err)
	}
	Log().Debugf("GetVolumes: %v, %v\n", volumes, err)
	for i := 0; i < 5; i++ {
		volume, err := p.GetVolume(ctx(), fmt.Sprintf("vol%d", i))
		if err != nil || volume.BackendUUID != recoveredBackend.BackendUUID {
			t.Fatalf("Volume retrieval failed; volume:%v err:%v\n", volume, err)
		}
		Log().WithFields(LogFields{
			"volume":                   volume,
			"volume.BackendUUID":       volume.BackendUUID,
			"newNFSServer.Name":        newNFSServer.Name(),
			"newNFSServer.BackendUUID": newNFSServer.BackendUUID(),
			"NFSServer.Name":           newNFSServer.Name(),
			"NFSServer.BackendUUID":    NFSServer.BackendUUID(),
		}).Debug("GetVolumes.")

		Log().Debugf("GetVolume(vol%v): %v, %v\n", i, volume, err)
	}

	// Deleting the storage backend
	err = p.DeleteBackend(ctx(), newNFSServer)
	if err != nil {
		t.Error(err.Error())
	}

	// Deleting all volumes
	if err = p.DeleteVolumes(ctx()); err != nil {
		t.Error(err.Error())
	}
}

func TestKubernetesAddOrUpdateNode(t *testing.T) {
	p, _ := GetTestKubernetesClient()

	utilsNode := &utils.Node{
		Name: "test",
		IQN:  "iqn",
		IPs: []string{
			"192.168.0.1",
		},
		Deleted: false,
	}

	// should not exist
	_, err := p.GetNode(ctx(), utilsNode.Name)
	if !IsStatusNotFoundError(err) {
		t.Fatal(err.Error())
	}

	// should be added
	if err := p.AddOrUpdateNode(ctx(), utilsNode); err != nil {
		t.Fatal(err.Error())
	}

	// validate we can find what we added
	node, err := p.GetNode(ctx(), utilsNode.Name)
	if err != nil {
		t.Fatal(err.Error())
	}

	if node.Name != utilsNode.Name {
		t.Fatalf("%v differs:  '%v' != '%v'", "Name", node.Name, utilsNode.Name)
	}

	if node.IQN != utilsNode.IQN {
		t.Fatalf("%v differs:  '%v' != '%v'", "IQN", node.IQN, utilsNode.IQN)
	}

	if len(node.IPs) != len(utilsNode.IPs) {
		t.Fatalf("%v differs:  '%v' != '%v'", "IPs", node.IPs, utilsNode.IPs)
	}

	for i := range node.IPs {
		if node.IPs[i] != utilsNode.IPs[i] {
			t.Fatalf("%v differs:  '%v' != '%v'", "IPs", node.IPs, utilsNode.IPs)
		}
	}

	// update it
	utilsNode.IQN = "iqnUpdated"
	if err := p.AddOrUpdateNode(ctx(), utilsNode); err != nil {
		t.Fatal(err.Error())
	}
	node, err = p.GetNode(ctx(), utilsNode.Name)
	if err != nil {
		t.Fatal(err.Error())
	}
	if node.IQN != "iqnUpdated" {
		t.Fatalf("%v differs:  '%v' != '%v'", "IQN", node.IQN, "iqnUpdated")
	}

	// remove it
	if err := p.DeleteNode(ctx(), utilsNode); err != nil {
		t.Fatal(err.Error())
	}

	tnode, err := p.crdClient.TridentV1().TridentNodes(p.namespace).Get(ctx(), utilsNode.Name, getOpts)
	if err != nil || tnode == nil || !tnode.HasTridentFinalizers() || tnode.DeletionTimestamp.IsZero() {
		t.Fatalf("Node should have been updated; %v", err)
	}

	// Remove finalizers to ensure node is deleted
	tnode.RemoveTridentFinalizers()
	_, _ = p.crdClient.TridentV1().TridentNodes(p.namespace).Update(ctx(), tnode, updateOpts)

	// validate it's gone
	_, err = p.GetNode(ctx(), utilsNode.Name)
	if !errors.IsNotFound(err) {
		t.Fatalf("Node should have been deleted; %v", err)
	}

	// Deleting a non-existent node
	err = p.DeleteNode(ctx(), utilsNode)
	if err != nil {
		t.Error("DeleteNode should have succeeded.")
	}
}

func TestKubernetesSnapshot(t *testing.T) {
	p, _ := GetTestKubernetesClient()

	// Adding a snapshot
	NFSServerConfig := drivers.OntapStorageDriverConfig{
		CommonStorageDriverConfig: &drivers.CommonStorageDriverConfig{
			StorageDriverName: config.OntapNASStorageDriverName,
		},
		ManagementLIF: "10.0.0.4",
		DataLIF:       "10.0.0.100",
		SVM:           "svm1",
		Username:      "admin",
		Password:      "netapp",
	}
	NFSServer := &storage.StorageBackend{}
	NFSServer.SetDriver(&ontap.NASStorageDriver{
		Config: NFSServerConfig,
	})
	NFSServer.SetName("NFS_server-" + NFSServerConfig.ManagementLIF)
	NFSServer.SetBackendUUID(uuid.New().String())
	vol1Config := storage.VolumeConfig{
		Version:      config.OrchestratorAPIVersion,
		Name:         "vol1",
		Size:         "1GB",
		Protocol:     config.File,
		StorageClass: "gold",
	}
	vol1 := &storage.Volume{
		Config:      &vol1Config,
		BackendUUID: NFSServer.BackendUUID(),
		Pool:        storagePool,
	}
	err := p.AddVolume(ctx(), vol1)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	snap1Config := &storage.SnapshotConfig{
		Version:            "1",
		Name:               "testsnap1",
		InternalName:       "internal_testsnap1",
		VolumeName:         "vol1",
		VolumeInternalName: "internal_vol1",
	}
	now := time.Now().UTC().Format(storage.SnapshotNameFormat)
	size := int64(1000000000)
	snap1 := storage.NewSnapshot(snap1Config, now, size, storage.SnapshotStateOnline)
	err = p.AddSnapshot(ctx(), snap1)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	// Getting a snapshot
	recoveredSnapshot, err := p.GetSnapshot(ctx(), snap1.Config.VolumeName, snap1.Config.Name)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if recoveredSnapshot.SizeBytes != snap1.SizeBytes || recoveredSnapshot.Created != snap1.Created ||
		!reflect.DeepEqual(recoveredSnapshot.Config, snap1.Config) {
		t.Error("Recovered snapshot does not match!")
	}

	// Test idempotent replacement
	err = p.AddSnapshot(ctx(), snap1)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	// Getting a snapshot
	recoveredSnapshot, err = p.GetSnapshot(ctx(), snap1.Config.VolumeName, snap1.Config.Name)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if recoveredSnapshot.SizeBytes != snap1.SizeBytes || recoveredSnapshot.Created != snap1.Created ||
		!reflect.DeepEqual(recoveredSnapshot.Config, snap1.Config) {
		t.Error("Recovered snapshot does not match!")
	}

	// Deleting a snapshot
	err = p.DeleteSnapshot(ctx(), snap1)
	if err != nil {
		t.Error(err.Error())
	}

	tsnap, err := p.crdClient.TridentV1().TridentSnapshots(p.namespace).Get(ctx(), "vol1-testsnap1", getOpts)
	if err != nil || tsnap == nil || !tsnap.HasTridentFinalizers() || tsnap.DeletionTimestamp.IsZero() {
		t.Fatalf("Snapshot should have been updated; %v", err)
	}

	// Remove finalizers to ensure snapshot is deleted
	tsnap.RemoveTridentFinalizers()
	_, _ = p.crdClient.TridentV1().TridentSnapshots(p.namespace).Update(ctx(), tsnap, updateOpts)

	_, err = p.GetSnapshot(ctx(), snap1.Config.VolumeName, snap1.Config.Name)
	if !errors.IsNotFound(err) {
		t.Fatalf("Snapshot should have been deleted; %v", err)
	}

	// Deleting a non-existent snapshot
	err = p.DeleteSnapshot(ctx(), snap1)
	if err != nil {
		t.Error("DeleteSnapshot should have succeeded.")
	}
}

func TestKubernetesSnapshots(t *testing.T) {
	var err error

	p, _ := GetTestKubernetesClient()

	// Adding volumes
	NFSServerConfig := drivers.OntapStorageDriverConfig{
		CommonStorageDriverConfig: &drivers.CommonStorageDriverConfig{
			StorageDriverName: config.OntapNASStorageDriverName,
		},
		ManagementLIF: "10.0.0.4",
		DataLIF:       "10.0.0.100",
		SVM:           "svm1",
		Username:      "admin",
		Password:      "netapp",
	}
	NFSServer := &storage.StorageBackend{}
	NFSServer.SetDriver(&ontap.NASStorageDriver{
		Config: NFSServerConfig,
	})
	NFSServer.SetName("NFS_server-" + NFSServerConfig.ManagementLIF)
	NFSServer.SetBackendUUID(uuid.New().String())

	for i := 1; i <= 5; i++ {
		volConfig := storage.VolumeConfig{
			Version:      config.OrchestratorAPIVersion,
			Name:         "vol" + strconv.Itoa(i),
			Size:         strconv.Itoa(i) + "GB",
			Protocol:     config.File,
			StorageClass: "gold",
		}
		vol := &storage.Volume{
			Config:      &volConfig,
			BackendUUID: NFSServer.BackendUUID(),
		}
		err = p.AddVolume(ctx(), vol)
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}
		snapConfig := &storage.SnapshotConfig{
			Version:            "1",
			Name:               "snap" + strconv.Itoa(i),
			InternalName:       "internal_snap1" + strconv.Itoa(i),
			VolumeName:         "vol" + strconv.Itoa(i),
			VolumeInternalName: "internal_vol" + strconv.Itoa(i),
		}
		now := time.Now().UTC().Format(storage.SnapshotNameFormat)
		size := int64(1000000000)
		snap := storage.NewSnapshot(snapConfig, now, size, storage.SnapshotStateOnline)
		err = p.AddSnapshot(ctx(), snap)
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}
	}

	// Retrieving all snapshots
	var snapshots []*storage.SnapshotPersistent
	if snapshots, err = p.GetSnapshots(ctx()); err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if len(snapshots) != 5 {
		t.Errorf("Expected %d snapshots; retrieved %d", 5, len(snapshots))
	}

	// Deleting all snapshots
	if err = p.DeleteSnapshots(ctx()); err != nil {
		t.Error(err.Error())
	}
	if snapshots, err = p.GetSnapshots(ctx()); err != nil {
		t.Error(err.Error())
	} else if len(snapshots) != 0 {
		t.Error("Deleting snapshots failed!")
	}
}

/*
func TestBackend_RemoveFinalizers(t *testing.T) {

	var err error

	p := GetTestKubernetesClient()

	// Initialize the state by adding a storage backend and volumes
	NFSServerConfig := drivers.OntapStorageDriverConfig{
		CommonStorageDriverConfig: &drivers.CommonStorageDriverConfig{
			StorageDriverName: config.OntapNASStorageDriverName,
		},
		ManagementLIF: "10.0.0.4",
		DataLIF:       "10.0.0.100",
		SVM:           "svm1",
		Username:      "admin",
		Password:      "netapp",
	}
	NFSDriver := ontap.NASStorageDriver{
		Config: NFSServerConfig,
	}
	NFSServer := &storage.Backend{
		Driver: &NFSDriver,
		Name:   "ontapnas_" + NFSServerConfig.DataLIF,
	}
	err = p.AddBackend(NFSServer)
	if err != nil {
		t.Fatalf("Backend creation failed: %v\n", err)
	}

	backendList, err := p.client.TridentV1().TridentBackends(p.namespace).List(metav1.ListOptions{})
	if err != nil {
		t.Fatalf("Backend list failed: %v\n", err)
	}

	// Check the Finalizers exist, removing them as we go
	for _, item := range backendList.Items {
		Log().WithFields(LogFields{
			"Name":    item.Name,
			"ObjectMeta.Finalizers": item.ObjectMeta.Finalizers,
		}).Debug("Found Item")

		if !utils.SliceContainsString(item.ObjectMeta.Finalizers, v1.TridentFinalizerName) {
			t.Fatalf("Expected to find Trident finalizer %v in list: %v\n", v1.TridentFinalizerName, item.ObjectMeta.Finalizers)
		}

		// remove our finalizer and update
		item.ObjectMeta.Finalizers = utils.RemoveStringFromSlice(item.ObjectMeta.Finalizers, v1.TridentFinalizerName)
		_, err = p.client.TridentV1().TridentBackends(p.namespace).Update(item)
	}

	// Now, spin back through, validating the Finalizers were removed
	backendList, err = p.client.TridentV1().TridentBackends(p.namespace).List(metav1.ListOptions{})
	if err != nil {
		t.Fatalf("Backend list failed: %v\n", err)
	}
	for _, item := range backendList.Items {
		Log().WithFields(LogFields{
			"Name":    item.Name,
			"ObjectMeta.Finalizers": item.ObjectMeta.Finalizers,
		}).Debug("Found Item")

		if utils.SliceContainsString(item.ObjectMeta.Finalizers, v1.TridentFinalizerName) {
			t.Fatalf("Did NOT expect to find Trident finalizer %v in list: %v\n", v1.TridentFinalizerName, item.ObjectMeta.Finalizers)
		}
	}
}
*/