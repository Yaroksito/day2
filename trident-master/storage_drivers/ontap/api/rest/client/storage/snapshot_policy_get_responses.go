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

// SnapshotPolicyGetReader is a Reader for the SnapshotPolicyGet structure.
type SnapshotPolicyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotPolicyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotPolicyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotPolicyGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotPolicyGetOK creates a SnapshotPolicyGetOK with default headers values
func NewSnapshotPolicyGetOK() *SnapshotPolicyGetOK {
	return &SnapshotPolicyGetOK{}
}

/*
SnapshotPolicyGetOK describes a response with status code 200, with default header values.

OK
*/
type SnapshotPolicyGetOK struct {
	Payload *models.SnapshotPolicy
}

// IsSuccess returns true when this snapshot policy get o k response has a 2xx status code
func (o *SnapshotPolicyGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshot policy get o k response has a 3xx status code
func (o *SnapshotPolicyGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot policy get o k response has a 4xx status code
func (o *SnapshotPolicyGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshot policy get o k response has a 5xx status code
func (o *SnapshotPolicyGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot policy get o k response a status code equal to that given
func (o *SnapshotPolicyGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *SnapshotPolicyGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/snapshot-policies/{uuid}][%d] snapshotPolicyGetOK  %+v", 200, o.Payload)
}

func (o *SnapshotPolicyGetOK) String() string {
	return fmt.Sprintf("[GET /storage/snapshot-policies/{uuid}][%d] snapshotPolicyGetOK  %+v", 200, o.Payload)
}

func (o *SnapshotPolicyGetOK) GetPayload() *models.SnapshotPolicy {
	return o.Payload
}

func (o *SnapshotPolicyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapshotPolicyGetDefault creates a SnapshotPolicyGetDefault with default headers values
func NewSnapshotPolicyGetDefault(code int) *SnapshotPolicyGetDefault {
	return &SnapshotPolicyGetDefault{
		_statusCode: code,
	}
}

/*
SnapshotPolicyGetDefault describes a response with status code -1, with default header values.

Error
*/
type SnapshotPolicyGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snapshot policy get default response
func (o *SnapshotPolicyGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this snapshot policy get default response has a 2xx status code
func (o *SnapshotPolicyGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapshot policy get default response has a 3xx status code
func (o *SnapshotPolicyGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapshot policy get default response has a 4xx status code
func (o *SnapshotPolicyGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapshot policy get default response has a 5xx status code
func (o *SnapshotPolicyGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapshot policy get default response a status code equal to that given
func (o *SnapshotPolicyGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SnapshotPolicyGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/snapshot-policies/{uuid}][%d] snapshot_policy_get default  %+v", o._statusCode, o.Payload)
}

func (o *SnapshotPolicyGetDefault) String() string {
	return fmt.Sprintf("[GET /storage/snapshot-policies/{uuid}][%d] snapshot_policy_get default  %+v", o._statusCode, o.Payload)
}

func (o *SnapshotPolicyGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotPolicyGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
