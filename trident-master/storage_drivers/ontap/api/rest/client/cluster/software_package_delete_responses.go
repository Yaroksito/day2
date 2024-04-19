// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SoftwarePackageDeleteReader is a Reader for the SoftwarePackageDelete structure.
type SoftwarePackageDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoftwarePackageDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSoftwarePackageDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSoftwarePackageDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSoftwarePackageDeleteAccepted creates a SoftwarePackageDeleteAccepted with default headers values
func NewSoftwarePackageDeleteAccepted() *SoftwarePackageDeleteAccepted {
	return &SoftwarePackageDeleteAccepted{}
}

/*
SoftwarePackageDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SoftwarePackageDeleteAccepted struct {
}

// IsSuccess returns true when this software package delete accepted response has a 2xx status code
func (o *SoftwarePackageDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this software package delete accepted response has a 3xx status code
func (o *SoftwarePackageDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this software package delete accepted response has a 4xx status code
func (o *SoftwarePackageDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this software package delete accepted response has a 5xx status code
func (o *SoftwarePackageDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this software package delete accepted response a status code equal to that given
func (o *SoftwarePackageDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *SoftwarePackageDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /cluster/software/packages/{version}][%d] softwarePackageDeleteAccepted ", 202)
}

func (o *SoftwarePackageDeleteAccepted) String() string {
	return fmt.Sprintf("[DELETE /cluster/software/packages/{version}][%d] softwarePackageDeleteAccepted ", 202)
}

func (o *SoftwarePackageDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSoftwarePackageDeleteDefault creates a SoftwarePackageDeleteDefault with default headers values
func NewSoftwarePackageDeleteDefault(code int) *SoftwarePackageDeleteDefault {
	return &SoftwarePackageDeleteDefault{
		_statusCode: code,
	}
}

/*
	SoftwarePackageDeleteDefault describes a response with status code -1, with default header values.

	ONTAP error response codes

| Error codes | Description |
| ----------- | ----------- |
| 10551315    | Package store is empty |
| 10551322    | Error in retrieving package cleanup status |
| 10551323    | Error in cleaning up package information on a node |
| 10551324    | Error in cleaning up package information on multiple nodes |
| 10551325    | Package does not exist on the system |
| 10551326    | Error in deleting older package cleanup tasks. Clean up images from the store and retry |
| 10551346    | Package delete failed since a validation is in progress |
| 10551347    | Package delete failed since an update is in progress |
| 10551367    | A package synchronization is in progress |
| 10551388    | Package delete operation timed out |
*/
type SoftwarePackageDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the software package delete default response
func (o *SoftwarePackageDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this software package delete default response has a 2xx status code
func (o *SoftwarePackageDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this software package delete default response has a 3xx status code
func (o *SoftwarePackageDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this software package delete default response has a 4xx status code
func (o *SoftwarePackageDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this software package delete default response has a 5xx status code
func (o *SoftwarePackageDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this software package delete default response a status code equal to that given
func (o *SoftwarePackageDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SoftwarePackageDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /cluster/software/packages/{version}][%d] software_package_delete default  %+v", o._statusCode, o.Payload)
}

func (o *SoftwarePackageDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /cluster/software/packages/{version}][%d] software_package_delete default  %+v", o._statusCode, o.Payload)
}

func (o *SoftwarePackageDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SoftwarePackageDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}