// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnapshotCreateReader is a Reader for the SnapshotCreate structure.
type SnapshotCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSnapshotCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotCreateAccepted creates a SnapshotCreateAccepted with default headers values
func NewSnapshotCreateAccepted() *SnapshotCreateAccepted {
	return &SnapshotCreateAccepted{}
}

/*
SnapshotCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnapshotCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snapshot create accepted response has a 2xx status code
func (o *SnapshotCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshot create accepted response has a 3xx status code
func (o *SnapshotCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot create accepted response has a 4xx status code
func (o *SnapshotCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshot create accepted response has a 5xx status code
func (o *SnapshotCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot create accepted response a status code equal to that given
func (o *SnapshotCreateAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *SnapshotCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /storage/volumes/{volume.uuid}/snapshots][%d] snapshotCreateAccepted  %+v", 202, o.Payload)
}

func (o *SnapshotCreateAccepted) String() string {
	return fmt.Sprintf("[POST /storage/volumes/{volume.uuid}/snapshots][%d] snapshotCreateAccepted  %+v", 202, o.Payload)
}

func (o *SnapshotCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapshotCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotCreateDefault creates a SnapshotCreateDefault with default headers values
func NewSnapshotCreateDefault(code int) *SnapshotCreateDefault {
	return &SnapshotCreateDefault{
		_statusCode: code,
	}
}

/*
	SnapshotCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Code

| Error Code | Description |
| ---------- | ----------- |
| 524479     | The specified volume is not online or does not have enough space to create a Snapshot copy. |
| 2621462    | The specified SVM name does not exist. |
| 1638433    | A Snapshot copy with the specified name already exists. |
| 1638461    | Snapshot copies can only be created on read/write (RW) volumes. |
| 1638477    | User-created Snapshot copy names cannot begin with the specified prefix. |
| 1638518    | The specified Snapshot copy name is invalid. |
| 1638532    | Failed to create the Snapshot copy on the specified volume because a revert operation is in progress. |
| 1638537    | Cannot determine the status of the Snapshot copy create operation for the specified volume. |
| 1638616    | Bulk Snapshot copy create is not supported with multiple Snapshot copy names. |
| 1638617    | Bulk Snapshot copy create is not supported with volume names in a mixed-version cluster. |
| 1638618    | The property cannot be specified for Snapshot copy create. |
*/
type SnapshotCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snapshot create default response
func (o *SnapshotCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this snapshot create default response has a 2xx status code
func (o *SnapshotCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapshot create default response has a 3xx status code
func (o *SnapshotCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapshot create default response has a 4xx status code
func (o *SnapshotCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapshot create default response has a 5xx status code
func (o *SnapshotCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapshot create default response a status code equal to that given
func (o *SnapshotCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SnapshotCreateDefault) Error() string {
	return fmt.Sprintf("[POST /storage/volumes/{volume.uuid}/snapshots][%d] snapshot_create default  %+v", o._statusCode, o.Payload)
}

func (o *SnapshotCreateDefault) String() string {
	return fmt.Sprintf("[POST /storage/volumes/{volume.uuid}/snapshots][%d] snapshot_create default  %+v", o._statusCode, o.Payload)
}

func (o *SnapshotCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
