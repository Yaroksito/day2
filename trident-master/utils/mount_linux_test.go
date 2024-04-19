// Copyright 2022 NetApp, Inc. All Rights Reserved.

package utils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/netapp/trident/utils/errors"
)

func TestParseProcMountinfo(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		expected    []MountInfo
		expectedErr error
	}{
		{
			name: "10-13 fields",
			content: `40 35 0:34 / /sys/fs/cgroup/cpu,cpuacct rw,nosuid,nodev,noexec,relatime - cgroup cgroup rw,cpu,cpuacct
46 35 0:40 / /sys/fs/cgroup/blkio rw,nosuid,nodev,noexec,relatime shared:21 - cgroup cgroup rw,blkio
47 35 0:41 / /sys/fs/cgroup/rdma rw,nosuid,nodev,noexec,relatime shared:22 master:1 - cgroup cgroup rw,rdma
48 35 0:42 / /sys/fs/cgroup/devices rw,nosuid,nodev,noexec,relatime shared:23 shared:74 master:2 - cgroup cgroup rw,devices`,
			expected: []MountInfo{
				{
					MountId:      40,
					ParentId:     35,
					DeviceId:     "0:34",
					Root:         "/",
					MountPoint:   "/sys/fs/cgroup/cpu,cpuacct",
					MountOptions: []string{"rw", "nosuid", "nodev", "noexec", "relatime"},
					FsType:       "cgroup",
					MountSource:  "cgroup",
					SuperOptions: []string{"rw", "cpu", "cpuacct"},
				},
				{
					MountId:      46,
					ParentId:     35,
					DeviceId:     "0:40",
					Root:         "/",
					MountPoint:   "/sys/fs/cgroup/blkio",
					MountOptions: []string{"rw", "nosuid", "nodev", "noexec", "relatime"},
					FsType:       "cgroup",
					MountSource:  "cgroup",
					SuperOptions: []string{"rw", "blkio"},
				},
				{
					MountId:      47,
					ParentId:     35,
					DeviceId:     "0:41",
					Root:         "/",
					MountPoint:   "/sys/fs/cgroup/rdma",
					MountOptions: []string{"rw", "nosuid", "nodev", "noexec", "relatime"},
					FsType:       "cgroup",
					MountSource:  "cgroup",
					SuperOptions: []string{"rw", "rdma"},
				},
				{
					MountId:      48,
					ParentId:     35,
					DeviceId:     "0:42",
					Root:         "/",
					MountPoint:   "/sys/fs/cgroup/devices",
					MountOptions: []string{"rw", "nosuid", "nodev", "noexec", "relatime"},
					FsType:       "cgroup",
					MountSource:  "cgroup",
					SuperOptions: []string{"rw", "devices"},
				},
			},
		},
		{
			name: "one valid one invalid",
			content: `36 35 0:30 / /sys/fs/cgroup/unified rw,nosuid,nodev,noexec,relatime shared:10 - cgroup2 cgroup2 rw
47 35 0:41 / rw,nosuid,nodev,noexec,relatime - cgroup cgroup rw,rdma`,
			expectedErr: errors.New("wrong number of fields (expected at least 10, got 9): 47 35 0:41 / rw,nosuid,nodev,noexec,relatime - cgroup cgroup rw,rdma"),
		},
		{
			name:        "too few fields",
			content:     `36 35 0:30 / /sys/fs/cgroup/unified - cgroup2 cgroup2 rw`,
			expectedErr: errors.New("wrong number of fields (expected at least 10, got 9): 36 35 0:30 / /sys/fs/cgroup/unified - cgroup2 cgroup2 rw"),
		},
		{
			name:        "separator in 5th position",
			content:     `49 35 0:43 / /sys/fs/cgroup/pids rw,nosuid,nodev,noexec,relatime - shared:24 cgroup cgroup rw,pids`,
			expectedErr: errors.New("malformed mountinfo (could not find separator): 49 35 0:43 / /sys/fs/cgroup/pids rw,nosuid,nodev,noexec,relatime - shared:24 cgroup cgroup rw,pids"),
		},
		{
			name:        "no separator",
			content:     `52 26 0:12 / /sys/kernel/tracing rw,nosuid,nodev,noexec,relatime shared:27 tracefs tracefs rw`,
			expectedErr: errors.New("malformed mountinfo (could not find separator): 52 26 0:12 / /sys/kernel/tracing rw,nosuid,nodev,noexec,relatime shared:27 tracefs tracefs rw"),
		},
	}
	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			m, err := parseProcMountinfo([]byte(tests[i].content))
			assert.Equal(t, tests[i].expectedErr, err)
			assert.Equal(t, tests[i].expected, m)
		})
	}
}

func TestCheckMountOptionsPositive(t *testing.T) {
	ctx := context.Background()
	mountInfo := MountInfo{
		MountId:      0,
		ParentId:     0,
		DeviceId:     "",
		Root:         "",
		MountPoint:   "",
		MountOptions: []string{"m1", "m2"},
		FsType:       "",
		MountSource:  "",
		SuperOptions: []string{"sm1", "sm2"},
	}
	res := CheckMountOptions(ctx, mountInfo, "m1")
	assert.NoError(t, res, "mount option mismatch")

	res = CheckMountOptions(ctx, mountInfo, "sm2")
	assert.NoError(t, res, "mount option mismatch")

	// no mount option should not result in error
	res = CheckMountOptions(ctx, mountInfo, "")
	assert.NoError(t, res, "mount option mismatch")
}

func TestCheckMountOptionsNegative(t *testing.T) {
	ctx := context.Background()
	mountInfo := MountInfo{
		MountId:      0,
		ParentId:     0,
		DeviceId:     "",
		Root:         "",
		MountPoint:   "",
		MountOptions: []string{"m1", "m2"},
		FsType:       "",
		MountSource:  "",
		SuperOptions: []string{"sm1", "sm2"},
	}
	res := CheckMountOptions(ctx, mountInfo, "m9")
	assert.Error(t, res, "expecting mount option mismatch")

	res = CheckMountOptions(ctx, mountInfo, "sm9")
	assert.Error(t, res, "expecting mount option mismatch")
}

func TestMountSMBPath(t *testing.T) {
	ctx := context.Background()
	result := mountSMBPath(ctx, "\\export\\path", "\\mount\\path", "test-user", "password")
	assert.Error(t, result, "no error")
	assert.True(t, errors.IsUnsupportedError(result), "not UnsupportedError")
}

func TestUmountSMBPath(t *testing.T) {
	ctx := context.Background()
	result := UmountSMBPath(ctx, "", "test-target")
	assert.Error(t, result, "no error")
	assert.True(t, errors.IsUnsupportedError(result), "not UnsupportedError")
}

func TestWindowsBindMount(t *testing.T) {
	ctx := context.Background()
	result := WindowsBindMount(ctx, "test-source", "test-target", []string{"test-val1", "test-val2"})
	assert.Error(t, result, "no error")
	assert.True(t, errors.IsUnsupportedError(result), "not UnsupportedError")
}

func TestIsCompatible_NFSProtocol(t *testing.T) {
	ctx := context.Background()
	err := IsCompatible(ctx, "nfs")
	assert.NoError(t, err, "error")
}

func TestIsCompatible_SMBProtocol(t *testing.T) {
	ctx := context.Background()
	err := IsCompatible(ctx, "smb")
	assert.Error(t, err, "no error")
}

// TODO (arorar): Need to identify a way to override /proc/self/mountinfo
/*
func TestPVMountpointMappings(t *testing.T) {

	file, err := osFs.OpenFile(procSelfMountinfoPath, os.O_RDWR, 0777)
	assert.NoError(t, err, "error reading file")

	input := "3242 30 0:6 /abc /var/lib/kubelet/plugins/kubernetes." +
		"io/csi/volumeDevices/publish/pvc-abcc/0c587b40-e3db-4ff1-8c36-a47753e1ca2a" +
		" rw,nosuid,relatime shared:2 - devtmpfs udev rw,size=1971956k,nr_inodes=492989," +
		"mode=755\n3255 30 0:6 /abc /var/lib/kubelet/plugins/kubernetes." +
		"io/csi/volumeDevices/pvc-abc/dev/0c587b40-e3db-4ff1-8c36-a47753e1ca2a rw," +
		"relatime shared:2 - devtmpfs udev rw,size=1971956k,nr_inodes=492989," +
		"mode=755\n3511 30 253:1 / /var/lib/kubelet/pods/b2821391-5790-4191-be66-eac8e744a3dd/volumes/kubernetes." +
		"io~csi/pvc-xyz/mount rw,relatime shared:1304 - ext4 /xyz rw,stripe=16\n"

	_, err = file.Write([]byte(input))
	assert.NoError(t, err, "error writing")

	mapping, err := PVMountpointMappings(context.Background())
	assert.NoError(t, err, "error getting mapping")

	value, ok := mapping["/var/lib/kubelet/plugins/kubernetes."+
		"io/csi/volumeDevices/publish/pvc-abc/0c587b40-e3db-4ff1-8c36-a47753e1ca2a"]
	assert.True(t, ok, "expected mapping to be found in map!")
	assert.Equal(t, "/dev/abc", value)

	value, ok = mapping["/var/lib/kubelet/pods/b2821391-5790-4191-be66-eac8e744a3dd/volumes/kubernetes.io~csi/pvc-xyz/mount"]
	assert.True(t, ok, "expected mapping to be found in map!")
	assert.Equal(t, "/dev/xyz", value)
}*/
