// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// VscanOnAccessCreateReader is a Reader for the VscanOnAccessCreate structure.
type VscanOnAccessCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanOnAccessCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVscanOnAccessCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanOnAccessCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanOnAccessCreateCreated creates a VscanOnAccessCreateCreated with default headers values
func NewVscanOnAccessCreateCreated() *VscanOnAccessCreateCreated {
	return &VscanOnAccessCreateCreated{}
}

/*
VscanOnAccessCreateCreated describes a response with status code 201, with default header values.

Created
*/
type VscanOnAccessCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.VscanOnAccessResponse
}

// IsSuccess returns true when this vscan on access create created response has a 2xx status code
func (o *VscanOnAccessCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan on access create created response has a 3xx status code
func (o *VscanOnAccessCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan on access create created response has a 4xx status code
func (o *VscanOnAccessCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan on access create created response has a 5xx status code
func (o *VscanOnAccessCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan on access create created response a status code equal to that given
func (o *VscanOnAccessCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *VscanOnAccessCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscanOnAccessCreateCreated  %+v", 201, o.Payload)
}

func (o *VscanOnAccessCreateCreated) String() string {
	return fmt.Sprintf("[POST /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscanOnAccessCreateCreated  %+v", 201, o.Payload)
}

func (o *VscanOnAccessCreateCreated) GetPayload() *models.VscanOnAccessResponse {
	return o.Payload
}

func (o *VscanOnAccessCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.VscanOnAccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVscanOnAccessCreateDefault creates a VscanOnAccessCreateDefault with default headers values
func NewVscanOnAccessCreateDefault(code int) *VscanOnAccessCreateDefault {
	return &VscanOnAccessCreateDefault{
		_statusCode: code,
	}
}

/*
	VscanOnAccessCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 10027043   | The new On-Access policy cannot be created as the SVM has reached the maximum number of On-Access policies allowed. Delete an existing policy in order to create a new policy |
| 10027101   | The file size must be in the range 1KB to 1TB |
| 10027107   | The include extensions list cannot be empty. Specify at least one extension for inclusion |
| 10027109   | The specified CIFS path is invalid. It must be in the form \"\\dir1\\dir2\" or \"\\dir1\\dir2\\\" |
| 10027249   | The On-Access policy created successfully but failed to enable the policy. The reason for enable policy operation failure might be that another policy is enabled. Disable the enabled policy and then enable the newly created policy using the PATCH operation |
| 10027253   | The number of paths specified exceeds the configured number of maximum paths. You cannot specify more than the maximum number of configured paths |
| 10027254   | The number of extensions specified exceeds the configured maximum number of extensions. You cannot specify more than the maximum number of configured extensions |
*/
type VscanOnAccessCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the vscan on access create default response
func (o *VscanOnAccessCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this vscan on access create default response has a 2xx status code
func (o *VscanOnAccessCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan on access create default response has a 3xx status code
func (o *VscanOnAccessCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan on access create default response has a 4xx status code
func (o *VscanOnAccessCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan on access create default response has a 5xx status code
func (o *VscanOnAccessCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan on access create default response a status code equal to that given
func (o *VscanOnAccessCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VscanOnAccessCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscan_on_access_create default  %+v", o._statusCode, o.Payload)
}

func (o *VscanOnAccessCreateDefault) String() string {
	return fmt.Sprintf("[POST /protocols/vscan/{svm.uuid}/on-access-policies][%d] vscan_on_access_create default  %+v", o._statusCode, o.Payload)
}

func (o *VscanOnAccessCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanOnAccessCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}