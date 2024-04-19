// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ApplicationGetReader is a Reader for the ApplicationGet structure.
type ApplicationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationGetOK creates a ApplicationGetOK with default headers values
func NewApplicationGetOK() *ApplicationGetOK {
	return &ApplicationGetOK{}
}

/*
ApplicationGetOK describes a response with status code 200, with default header values.

OK
*/
type ApplicationGetOK struct {
	Payload *models.Application
}

// IsSuccess returns true when this application get o k response has a 2xx status code
func (o *ApplicationGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application get o k response has a 3xx status code
func (o *ApplicationGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application get o k response has a 4xx status code
func (o *ApplicationGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this application get o k response has a 5xx status code
func (o *ApplicationGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this application get o k response a status code equal to that given
func (o *ApplicationGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *ApplicationGetOK) Error() string {
	return fmt.Sprintf("[GET /application/applications/{uuid}][%d] applicationGetOK  %+v", 200, o.Payload)
}

func (o *ApplicationGetOK) String() string {
	return fmt.Sprintf("[GET /application/applications/{uuid}][%d] applicationGetOK  %+v", 200, o.Payload)
}

func (o *ApplicationGetOK) GetPayload() *models.Application {
	return o.Payload
}

func (o *ApplicationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationGetDefault creates a ApplicationGetDefault with default headers values
func NewApplicationGetDefault(code int) *ApplicationGetDefault {
	return &ApplicationGetDefault{
		_statusCode: code,
	}
}

/*
ApplicationGetDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the application get default response
func (o *ApplicationGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this application get default response has a 2xx status code
func (o *ApplicationGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application get default response has a 3xx status code
func (o *ApplicationGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application get default response has a 4xx status code
func (o *ApplicationGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application get default response has a 5xx status code
func (o *ApplicationGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application get default response a status code equal to that given
func (o *ApplicationGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ApplicationGetDefault) Error() string {
	return fmt.Sprintf("[GET /application/applications/{uuid}][%d] application_get default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationGetDefault) String() string {
	return fmt.Sprintf("[GET /application/applications/{uuid}][%d] application_get default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}