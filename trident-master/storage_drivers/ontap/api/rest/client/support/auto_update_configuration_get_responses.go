// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// AutoUpdateConfigurationGetReader is a Reader for the AutoUpdateConfigurationGet structure.
type AutoUpdateConfigurationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutoUpdateConfigurationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutoUpdateConfigurationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAutoUpdateConfigurationGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAutoUpdateConfigurationGetOK creates a AutoUpdateConfigurationGetOK with default headers values
func NewAutoUpdateConfigurationGetOK() *AutoUpdateConfigurationGetOK {
	return &AutoUpdateConfigurationGetOK{}
}

/*
AutoUpdateConfigurationGetOK describes a response with status code 200, with default header values.

OK
*/
type AutoUpdateConfigurationGetOK struct {
	Payload *models.AutoUpdateConfiguration
}

// IsSuccess returns true when this auto update configuration get o k response has a 2xx status code
func (o *AutoUpdateConfigurationGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auto update configuration get o k response has a 3xx status code
func (o *AutoUpdateConfigurationGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auto update configuration get o k response has a 4xx status code
func (o *AutoUpdateConfigurationGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auto update configuration get o k response has a 5xx status code
func (o *AutoUpdateConfigurationGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auto update configuration get o k response a status code equal to that given
func (o *AutoUpdateConfigurationGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *AutoUpdateConfigurationGetOK) Error() string {
	return fmt.Sprintf("[GET /support/auto-update/configurations/{uuid}][%d] autoUpdateConfigurationGetOK  %+v", 200, o.Payload)
}

func (o *AutoUpdateConfigurationGetOK) String() string {
	return fmt.Sprintf("[GET /support/auto-update/configurations/{uuid}][%d] autoUpdateConfigurationGetOK  %+v", 200, o.Payload)
}

func (o *AutoUpdateConfigurationGetOK) GetPayload() *models.AutoUpdateConfiguration {
	return o.Payload
}

func (o *AutoUpdateConfigurationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AutoUpdateConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAutoUpdateConfigurationGetDefault creates a AutoUpdateConfigurationGetDefault with default headers values
func NewAutoUpdateConfigurationGetDefault(code int) *AutoUpdateConfigurationGetDefault {
	return &AutoUpdateConfigurationGetDefault{
		_statusCode: code,
	}
}

/*
AutoUpdateConfigurationGetDefault describes a response with status code -1, with default header values.

Error
*/
type AutoUpdateConfigurationGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the auto update configuration get default response
func (o *AutoUpdateConfigurationGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this auto update configuration get default response has a 2xx status code
func (o *AutoUpdateConfigurationGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this auto update configuration get default response has a 3xx status code
func (o *AutoUpdateConfigurationGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this auto update configuration get default response has a 4xx status code
func (o *AutoUpdateConfigurationGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this auto update configuration get default response has a 5xx status code
func (o *AutoUpdateConfigurationGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this auto update configuration get default response a status code equal to that given
func (o *AutoUpdateConfigurationGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AutoUpdateConfigurationGetDefault) Error() string {
	return fmt.Sprintf("[GET /support/auto-update/configurations/{uuid}][%d] auto_update_configuration_get default  %+v", o._statusCode, o.Payload)
}

func (o *AutoUpdateConfigurationGetDefault) String() string {
	return fmt.Sprintf("[GET /support/auto-update/configurations/{uuid}][%d] auto_update_configuration_get default  %+v", o._statusCode, o.Payload)
}

func (o *AutoUpdateConfigurationGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutoUpdateConfigurationGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
