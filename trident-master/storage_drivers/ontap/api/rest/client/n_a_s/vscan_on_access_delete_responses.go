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

// VscanOnAccessDeleteReader is a Reader for the VscanOnAccessDelete structure.
type VscanOnAccessDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanOnAccessDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanOnAccessDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanOnAccessDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanOnAccessDeleteOK creates a VscanOnAccessDeleteOK with default headers values
func NewVscanOnAccessDeleteOK() *VscanOnAccessDeleteOK {
	return &VscanOnAccessDeleteOK{}
}

/*
VscanOnAccessDeleteOK describes a response with status code 200, with default header values.

OK
*/
type VscanOnAccessDeleteOK struct {
}

// IsSuccess returns true when this vscan on access delete o k response has a 2xx status code
func (o *VscanOnAccessDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan on access delete o k response has a 3xx status code
func (o *VscanOnAccessDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan on access delete o k response has a 4xx status code
func (o *VscanOnAccessDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan on access delete o k response has a 5xx status code
func (o *VscanOnAccessDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan on access delete o k response a status code equal to that given
func (o *VscanOnAccessDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *VscanOnAccessDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}/on-access-policies/{name}][%d] vscanOnAccessDeleteOK ", 200)
}

func (o *VscanOnAccessDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}/on-access-policies/{name}][%d] vscanOnAccessDeleteOK ", 200)
}

func (o *VscanOnAccessDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVscanOnAccessDeleteDefault creates a VscanOnAccessDeleteDefault with default headers values
func NewVscanOnAccessDeleteDefault(code int) *VscanOnAccessDeleteDefault {
	return &VscanOnAccessDeleteDefault{
		_statusCode: code,
	}
}

/*
	VscanOnAccessDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 10027034   | An On-Access policy associated with an administrative SVM cannot be deleted. |
| 10027040   | An On-Access policy with a status enabled cannot be deleted. Disable the policy and then delete the policy. |
*/
type VscanOnAccessDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the vscan on access delete default response
func (o *VscanOnAccessDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this vscan on access delete default response has a 2xx status code
func (o *VscanOnAccessDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan on access delete default response has a 3xx status code
func (o *VscanOnAccessDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan on access delete default response has a 4xx status code
func (o *VscanOnAccessDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan on access delete default response has a 5xx status code
func (o *VscanOnAccessDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan on access delete default response a status code equal to that given
func (o *VscanOnAccessDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VscanOnAccessDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}/on-access-policies/{name}][%d] vscan_on_access_delete default  %+v", o._statusCode, o.Payload)
}

func (o *VscanOnAccessDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /protocols/vscan/{svm.uuid}/on-access-policies/{name}][%d] vscan_on_access_delete default  %+v", o._statusCode, o.Payload)
}

func (o *VscanOnAccessDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanOnAccessDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
