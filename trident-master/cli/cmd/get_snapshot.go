// Copyright 2019 NetApp, Inc. All Rights Reserved.

package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/netapp/trident/cli/api"
	"github.com/netapp/trident/frontend/rest"
	"github.com/netapp/trident/storage"
	"github.com/netapp/trident/utils/errors"
)

const maskDisplayOfSnapshotStateOnline = storage.SnapshotState("") // Used for display in 'tridentctl' query

var getSnapshotVolume string

func init() {
	getCmd.AddCommand(getSnapshotCmd)
	getSnapshotCmd.Flags().StringVar(&getSnapshotVolume, "volume", "", "Limit query to volume "+
		"(unless additional arguments are provided)")
}

var getSnapshotCmd = &cobra.Command{
	Use:     "snapshot [<volume name>/<snapshot name>...]",
	Short:   "Get one or more snapshots from Trident",
	Aliases: []string{"s", "snap", "snapshots"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if OperatingMode == ModeTunnel {
			command := []string{"get", "snapshot"}
			if getSnapshotVolume != "" {
				command = append(command, "--volume", getSnapshotVolume)
			}
			out, err := TunnelCommand(append(command, args...))
			printOutput(cmd, out, err)
			return err
		} else {
			return snapshotList(args)
		}
	},
}

func snapshotList(snapshotIDs []string) error {
	var err error

	// If no snapshots were specified, we'll get all of them
	getAll := false
	if len(snapshotIDs) == 0 {
		getAll = true
		snapshotIDs, err = GetSnapshots(getSnapshotVolume)
		if err != nil {
			return err
		}
	}

	snapshots := make([]storage.SnapshotExternal, 0, 10)

	// Get the actual snapshot objects
	for _, snapshotID := range snapshotIDs {

		snapshot, err := GetSnapshot(snapshotID)
		if err != nil {
			if getAll && errors.IsNotFoundError(err) {
				continue
			}
			return err
		}
		snapshots = append(snapshots, snapshot)
	}

	WriteSnapshots(snapshots)

	return nil
}

func GetSnapshots(volume string) ([]string, error) {
	var url string
	if volume == "" {
		url = BaseURL() + "/snapshot"
	} else {
		url = BaseURL() + "/volume/" + volume + "/snapshot"
	}

	response, responseBody, err := api.InvokeRESTAPI("GET", url, nil)
	if err != nil {
		return nil, err
	} else if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get snapshots: %v", GetErrorFromHTTPResponse(response, responseBody))
	}

	var listSnapshotsResponse rest.ListSnapshotsResponse
	err = json.Unmarshal(responseBody, &listSnapshotsResponse)
	if err != nil {
		return nil, err
	}

	return listSnapshotsResponse.Snapshots, nil
}

func GetSnapshot(snapshotID string) (storage.SnapshotExternal, error) {
	if !strings.ContainsRune(snapshotID, '/') {
		return storage.SnapshotExternal{}, errors.InvalidInputError(fmt.Sprintf("invalid snapshot ID: %s; "+
			"Please use the format <volume name>/<snapshot name>", snapshotID))
	}

	url := BaseURL() + "/snapshot/" + snapshotID
	response, responseBody, err := api.InvokeRESTAPI("GET", url, nil)
	if err != nil {
		return storage.SnapshotExternal{}, err
	} else if response.StatusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("could not get snapshot %s: %v", snapshotID,
			GetErrorFromHTTPResponse(response, responseBody))
		switch response.StatusCode {
		case http.StatusNotFound:
			return storage.SnapshotExternal{}, errors.NotFoundError(errorMessage)
		default:
			return storage.SnapshotExternal{}, errors.New(errorMessage)
		}
	}

	var getSnapshotResponse rest.GetSnapshotResponse
	err = json.Unmarshal(responseBody, &getSnapshotResponse)
	if err != nil {
		return storage.SnapshotExternal{}, err
	}
	if getSnapshotResponse.Snapshot == nil {
		return storage.SnapshotExternal{}, fmt.Errorf("could not get snapshot %s: no snapshot returned",
			snapshotID)
	}

	if getSnapshotResponse.Snapshot.State == storage.SnapshotStateOnline {
		// Currently, this is used only for display, mask 'online' state as "".
		// If in future any callers use this attribute, need to take care of it.
		getSnapshotResponse.Snapshot.State = maskDisplayOfSnapshotStateOnline
	}

	return *getSnapshotResponse.Snapshot, nil
}

func WriteSnapshots(snapshots []storage.SnapshotExternal) {
	switch OutputFormat {
	case FormatJSON:
		WriteJSON(api.MultipleSnapshotResponse{Items: snapshots})
	case FormatYAML:
		WriteYAML(api.MultipleSnapshotResponse{Items: snapshots})
	case FormatName:
		writeSnapshotIDs(snapshots)
	case FormatWide:
		writeWideSnapshotTable(snapshots)
	default:
		writeSnapshotTable(snapshots)
	}
}

func writeSnapshotTable(snapshots []storage.SnapshotExternal) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Volume", "Managed"})

	for _, snapshot := range snapshots {
		table.Append([]string{
			snapshot.Config.Name,
			snapshot.Config.VolumeName,
			strconv.FormatBool(!snapshot.Config.ImportNotManaged),
		})
	}

	table.Render()
}

func writeWideSnapshotTable(snapshots []storage.SnapshotExternal) {
	table := tablewriter.NewWriter(os.Stdout)
	header := []string{
		"Name",
		"Volume",
		"Created",
		"Size",
		"State",
		"Managed",
	}
	table.SetHeader(header)

	for _, snapshot := range snapshots {
		table.Append([]string{
			snapshot.Config.Name,
			snapshot.Config.VolumeName,
			snapshot.Created,
			humanize.IBytes(uint64(snapshot.SizeBytes)),
			string(snapshot.State),
			strconv.FormatBool(!snapshot.Config.ImportNotManaged),
		})
	}

	table.Render()
}

func writeSnapshotIDs(snapshots []storage.SnapshotExternal) {
	for _, s := range snapshots {
		fmt.Println(storage.MakeSnapshotID(s.Config.VolumeName, s.Config.Name))
	}
}
