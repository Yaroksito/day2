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

// FpolicyEventCollectionGetReader is a Reader for the FpolicyEventCollectionGet structure.
type FpolicyEventCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyEventCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyEventCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyEventCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyEventCollectionGetOK creates a FpolicyEventCollectionGetOK with default headers values
func NewFpolicyEventCollectionGetOK() *FpolicyEventCollectionGetOK {
	return &FpolicyEventCollectionGetOK{}
}

/*
FpolicyEventCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyEventCollectionGetOK struct {
	Payload *models.FpolicyEventResponse
}

// IsSuccess returns true when this fpolicy event collection get o k response has a 2xx status code
func (o *FpolicyEventCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy event collection get o k response has a 3xx status code
func (o *FpolicyEventCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy event collection get o k response has a 4xx status code
func (o *FpolicyEventCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy event collection get o k response has a 5xx status code
func (o *FpolicyEventCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy event collection get o k response a status code equal to that given
func (o *FpolicyEventCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *FpolicyEventCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/events][%d] fpolicyEventCollectionGetOK  %+v", 200, o.Payload)
}

func (o *FpolicyEventCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/events][%d] fpolicyEventCollectionGetOK  %+v", 200, o.Payload)
}

func (o *FpolicyEventCollectionGetOK) GetPayload() *models.FpolicyEventResponse {
	return o.Payload
}

func (o *FpolicyEventCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FpolicyEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFpolicyEventCollectionGetDefault creates a FpolicyEventCollectionGetDefault with default headers values
func NewFpolicyEventCollectionGetDefault(code int) *FpolicyEventCollectionGetDefault {
	return &FpolicyEventCollectionGetDefault{
		_statusCode: code,
	}
}

/*
FpolicyEventCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type FpolicyEventCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fpolicy event collection get default response
func (o *FpolicyEventCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this fpolicy event collection get default response has a 2xx status code
func (o *FpolicyEventCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy event collection get default response has a 3xx status code
func (o *FpolicyEventCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy event collection get default response has a 4xx status code
func (o *FpolicyEventCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy event collection get default response has a 5xx status code
func (o *FpolicyEventCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy event collection get default response a status code equal to that given
func (o *FpolicyEventCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FpolicyEventCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/events][%d] fpolicy_event_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *FpolicyEventCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/events][%d] fpolicy_event_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *FpolicyEventCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyEventCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}