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

// EmsFilterCollectionGetReader is a Reader for the EmsFilterCollectionGet structure.
type EmsFilterCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmsFilterCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmsFilterCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEmsFilterCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEmsFilterCollectionGetOK creates a EmsFilterCollectionGetOK with default headers values
func NewEmsFilterCollectionGetOK() *EmsFilterCollectionGetOK {
	return &EmsFilterCollectionGetOK{}
}

/*
EmsFilterCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type EmsFilterCollectionGetOK struct {
	Payload *models.EmsFilterResponse
}

// IsSuccess returns true when this ems filter collection get o k response has a 2xx status code
func (o *EmsFilterCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ems filter collection get o k response has a 3xx status code
func (o *EmsFilterCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ems filter collection get o k response has a 4xx status code
func (o *EmsFilterCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ems filter collection get o k response has a 5xx status code
func (o *EmsFilterCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ems filter collection get o k response a status code equal to that given
func (o *EmsFilterCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *EmsFilterCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /support/ems/filters][%d] emsFilterCollectionGetOK  %+v", 200, o.Payload)
}

func (o *EmsFilterCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /support/ems/filters][%d] emsFilterCollectionGetOK  %+v", 200, o.Payload)
}

func (o *EmsFilterCollectionGetOK) GetPayload() *models.EmsFilterResponse {
	return o.Payload
}

func (o *EmsFilterCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmsFilterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmsFilterCollectionGetDefault creates a EmsFilterCollectionGetDefault with default headers values
func NewEmsFilterCollectionGetDefault(code int) *EmsFilterCollectionGetDefault {
	return &EmsFilterCollectionGetDefault{
		_statusCode: code,
	}
}

/*
EmsFilterCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type EmsFilterCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ems filter collection get default response
func (o *EmsFilterCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ems filter collection get default response has a 2xx status code
func (o *EmsFilterCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ems filter collection get default response has a 3xx status code
func (o *EmsFilterCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ems filter collection get default response has a 4xx status code
func (o *EmsFilterCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ems filter collection get default response has a 5xx status code
func (o *EmsFilterCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ems filter collection get default response a status code equal to that given
func (o *EmsFilterCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EmsFilterCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /support/ems/filters][%d] ems_filter_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *EmsFilterCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /support/ems/filters][%d] ems_filter_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *EmsFilterCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *EmsFilterCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}