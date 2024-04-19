// Copyright 2019 NetApp, Inc. All Rights Reserved.

package storage

import (
	"fmt"
	"testing"

	"github.com/netapp/trident/utils"

	"github.com/stretchr/testify/assert"
)

func TestVolumeState(t *testing.T) {
	tests := map[string]struct {
		input     VolumeState
		output    string
		predicate func(VolumeState) bool
	}{
		"Unknown state (bad)": {
			input:  "",
			output: "unknown",
			predicate: func(input VolumeState) bool {
				return input.IsUnknown()
			},
		},
		"Unknown state": {
			input:  VolumeStateUnknown,
			output: "unknown",
			predicate: func(input VolumeState) bool {
				return input.IsUnknown()
			},
		},
		"Online state": {
			input:  VolumeStateOnline,
			output: "online",
			predicate: func(input VolumeState) bool {
				return input.IsOnline()
			},
		},
		"Deleting state": {
			input:  VolumeStateDeleting,
			output: "deleting",
			predicate: func(input VolumeState) bool {
				return input.IsDeleting()
			},
		},
	}
	for testName, test := range tests {
		t.Logf("Running test case '%s'", testName)

		assert.Equal(t, test.input.String(), test.output, "Strings not equal")
		assert.True(t, test.predicate(test.input), "Predicate failed")
	}
}

func TestVolumeConfig_ConstructClone(t *testing.T) {
	tests := map[string]*VolumeConfig{
		"IscsiAccessInfoIsCloned":    constructIscsiVolumeConfig(),
		"NfsAccessInfoIsCloned":      constructNFSVolumeConfig(),
		"SMBAccessInfoIsCloned":      constructSMBVolumeConfig(),
		"NfsBlockAccessInfoIsCloned": constructNfsBlockVolumeConfig(),
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			copy := test.ConstructClone()
			testConfigAddress := fmt.Sprintf("%p", &test)
			copyConfigAddress := fmt.Sprintf("%p", &copy)
			testAccessInfoAddress := fmt.Sprintf("%p", &test.AccessInfo)
			copyAccessInfoAddress := fmt.Sprintf("%p", &copy.AccessInfo)

			assert.EqualValues(t, test, copy, "Expected equal values")
			assert.NotEqual(t, testConfigAddress, copyConfigAddress, "Expected different addresses")
			assert.NotEqual(t, testAccessInfoAddress, copyAccessInfoAddress, "Expected different addresses")
		})
	}
}

func constructIscsiVolumeConfig() *VolumeConfig {
	return &VolumeConfig{
		Version:      "1",
		Name:         "foo",
		InternalName: "internal_foo",
		Size:         "1Gi",
		StorageClass: "san-sc",
		AccessMode:   "0",
		AccessInfo: utils.VolumeAccessInfo{
			IscsiAccessInfo: utils.IscsiAccessInfo{
				IscsiTargetPortal: "10.0.0.0",
				IscsiPortals: []string{
					"10.0.0.0", "10.0.0.1",
				},
				IscsiTargetIQN: "",
				IscsiLunNumber: 0,
				IscsiInterface: "",
				IscsiIgroup:    "per-node-igroup",
				IscsiChapInfo: utils.IscsiChapInfo{
					UseCHAP:              false,
					IscsiUsername:        "user",
					IscsiInitiatorSecret: "shh!",
					IscsiTargetUsername:  "target",
					IscsiTargetSecret:    "ssh!",
				},
			},
			PublishEnforcement: true,
			ReadOnly:           false,
			AccessMode:         0,
		},
		BlockSize: "1",
	}
}

func constructNFSVolumeConfig() *VolumeConfig {
	return &VolumeConfig{
		Version:      "1",
		Name:         "foo",
		InternalName: "internal_foo",
		Size:         "1Gi",
		StorageClass: "nas-sc",
		AccessMode:   "1",
		AccessInfo: utils.VolumeAccessInfo{
			NfsAccessInfo: utils.NfsAccessInfo{
				NfsServerIP: "10.0.0.0",
				NfsPath:     "/nfsshare",
				NfsUniqueID: "uuid4",
			},
			PublishEnforcement: false,
			ReadOnly:           false,
			AccessMode:         1,
		},
		FileSystem: "ext4",
	}
}

func constructSMBVolumeConfig() *VolumeConfig {
	return &VolumeConfig{
		Version:      "1",
		Name:         "foo",
		InternalName: "internal_foo",
		Size:         "1Gi",
		StorageClass: "nas-sc",
		AccessMode:   "1",
		AccessInfo: utils.VolumeAccessInfo{
			SMBAccessInfo: utils.SMBAccessInfo{
				SMBServer: "server",
				SMBPath:   "/smbshare",
			},
			PublishEnforcement: false,
			ReadOnly:           false,
			AccessMode:         1,
		},
		FileSystem: "ext4",
	}
}

func constructNfsBlockVolumeConfig() *VolumeConfig {
	return &VolumeConfig{
		Version:      "1",
		Name:         "foo",
		InternalName: "internal_foo",
		Size:         "1Gi",
		StorageClass: "nas-sc",
		AccessMode:   "1",
		AccessInfo: utils.VolumeAccessInfo{
			NfsBlockAccessInfo: utils.NfsBlockAccessInfo{
				SubvolumeName:         "bar",
				SubvolumeMountOptions: "mountoptions",
				NFSMountpoint:         "/share",
			},
			PublishEnforcement: false,
			ReadOnly:           false,
			AccessMode:         1,
		},
		FileSystem: "ext4",
	}
}
