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

// ApplicationComponentCollectionGetReader is a Reader for the ApplicationComponentCollectionGet structure.
type ApplicationComponentCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationComponentCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationComponentCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationComponentCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationComponentCollectionGetOK creates a ApplicationComponentCollectionGetOK with default headers values
func NewApplicationComponentCollectionGetOK() *ApplicationComponentCollectionGetOK {
	return &ApplicationComponentCollectionGetOK{}
}

/*
ApplicationComponentCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ApplicationComponentCollectionGetOK struct {
	Payload *models.ApplicationComponentResponse
}

// IsSuccess returns true when this application component collection get o k response has a 2xx status code
func (o *ApplicationComponentCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application component collection get o k response has a 3xx status code
func (o *ApplicationComponentCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application component collection get o k response has a 4xx status code
func (o *ApplicationComponentCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this application component collection get o k response has a 5xx status code
func (o *ApplicationComponentCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this application component collection get o k response a status code equal to that given
func (o *ApplicationComponentCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *ApplicationComponentCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/components][%d] applicationComponentCollectionGetOK  %+v", 200, o.Payload)
}

func (o *ApplicationComponentCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/components][%d] applicationComponentCollectionGetOK  %+v", 200, o.Payload)
}

func (o *ApplicationComponentCollectionGetOK) GetPayload() *models.ApplicationComponentResponse {
	return o.Payload
}

func (o *ApplicationComponentCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationComponentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationComponentCollectionGetDefault creates a ApplicationComponentCollectionGetDefault with default headers values
func NewApplicationComponentCollectionGetDefault(code int) *ApplicationComponentCollectionGetDefault {
	return &ApplicationComponentCollectionGetDefault{
		_statusCode: code,
	}
}

/*
ApplicationComponentCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationComponentCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the application component collection get default response
func (o *ApplicationComponentCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this application component collection get default response has a 2xx status code
func (o *ApplicationComponentCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application component collection get default response has a 3xx status code
func (o *ApplicationComponentCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application component collection get default response has a 4xx status code
func (o *ApplicationComponentCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application component collection get default response has a 5xx status code
func (o *ApplicationComponentCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application component collection get default response a status code equal to that given
func (o *ApplicationComponentCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ApplicationComponentCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/components][%d] application_component_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationComponentCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /application/applications/{application.uuid}/components][%d] application_component_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationComponentCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationComponentCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
