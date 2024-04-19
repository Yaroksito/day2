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

// EffectivePermissionGetReader is a Reader for the EffectivePermissionGet structure.
type EffectivePermissionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EffectivePermissionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEffectivePermissionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEffectivePermissionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEffectivePermissionGetOK creates a EffectivePermissionGetOK with default headers values
func NewEffectivePermissionGetOK() *EffectivePermissionGetOK {
	return &EffectivePermissionGetOK{}
}

/*
EffectivePermissionGetOK describes a response with status code 200, with default header values.

OK
*/
type EffectivePermissionGetOK struct {
	Payload *models.EffectivePermission
}

// IsSuccess returns true when this effective permission get o k response has a 2xx status code
func (o *EffectivePermissionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this effective permission get o k response has a 3xx status code
func (o *EffectivePermissionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this effective permission get o k response has a 4xx status code
func (o *EffectivePermissionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this effective permission get o k response has a 5xx status code
func (o *EffectivePermissionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this effective permission get o k response a status code equal to that given
func (o *EffectivePermissionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *EffectivePermissionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/file-security/effective-permissions/{svm.uuid}/{path}][%d] effectivePermissionGetOK  %+v", 200, o.Payload)
}

func (o *EffectivePermissionGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/file-security/effective-permissions/{svm.uuid}/{path}][%d] effectivePermissionGetOK  %+v", 200, o.Payload)
}

func (o *EffectivePermissionGetOK) GetPayload() *models.EffectivePermission {
	return o.Payload
}

func (o *EffectivePermissionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EffectivePermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEffectivePermissionGetDefault creates a EffectivePermissionGetDefault with default headers values
func NewEffectivePermissionGetDefault(code int) *EffectivePermissionGetDefault {
	return &EffectivePermissionGetDefault{
		_statusCode: code,
	}
}

/*
EffectivePermissionGetDefault describes a response with status code -1, with default header values.

ONTAP Error Response Codes | Error Code | Description | | ---------- | ----------- | | 655865     | The specified path cannot be used, if the file does not exist.|
*/
type EffectivePermissionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the effective permission get default response
func (o *EffectivePermissionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this effective permission get default response has a 2xx status code
func (o *EffectivePermissionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this effective permission get default response has a 3xx status code
func (o *EffectivePermissionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this effective permission get default response has a 4xx status code
func (o *EffectivePermissionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this effective permission get default response has a 5xx status code
func (o *EffectivePermissionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this effective permission get default response a status code equal to that given
func (o *EffectivePermissionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EffectivePermissionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/file-security/effective-permissions/{svm.uuid}/{path}][%d] effective_permission_get default  %+v", o._statusCode, o.Payload)
}

func (o *EffectivePermissionGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/file-security/effective-permissions/{svm.uuid}/{path}][%d] effective_permission_get default  %+v", o._statusCode, o.Payload)
}

func (o *EffectivePermissionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EffectivePermissionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
