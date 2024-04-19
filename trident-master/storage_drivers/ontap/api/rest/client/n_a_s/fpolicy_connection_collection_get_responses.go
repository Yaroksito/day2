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

// FpolicyConnectionCollectionGetReader is a Reader for the FpolicyConnectionCollectionGet structure.
type FpolicyConnectionCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyConnectionCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyConnectionCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyConnectionCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyConnectionCollectionGetOK creates a FpolicyConnectionCollectionGetOK with default headers values
func NewFpolicyConnectionCollectionGetOK() *FpolicyConnectionCollectionGetOK {
	return &FpolicyConnectionCollectionGetOK{}
}

/*
FpolicyConnectionCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyConnectionCollectionGetOK struct {
	Payload *models.FpolicyConnectionResponse
}

// IsSuccess returns true when this fpolicy connection collection get o k response has a 2xx status code
func (o *FpolicyConnectionCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy connection collection get o k response has a 3xx status code
func (o *FpolicyConnectionCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy connection collection get o k response has a 4xx status code
func (o *FpolicyConnectionCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy connection collection get o k response has a 5xx status code
func (o *FpolicyConnectionCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy connection collection get o k response a status code equal to that given
func (o *FpolicyConnectionCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *FpolicyConnectionCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/connections][%d] fpolicyConnectionCollectionGetOK  %+v", 200, o.Payload)
}

func (o *FpolicyConnectionCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/connections][%d] fpolicyConnectionCollectionGetOK  %+v", 200, o.Payload)
}

func (o *FpolicyConnectionCollectionGetOK) GetPayload() *models.FpolicyConnectionResponse {
	return o.Payload
}

func (o *FpolicyConnectionCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FpolicyConnectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFpolicyConnectionCollectionGetDefault creates a FpolicyConnectionCollectionGetDefault with default headers values
func NewFpolicyConnectionCollectionGetDefault(code int) *FpolicyConnectionCollectionGetDefault {
	return &FpolicyConnectionCollectionGetDefault{
		_statusCode: code,
	}
}

/*
FpolicyConnectionCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type FpolicyConnectionCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fpolicy connection collection get default response
func (o *FpolicyConnectionCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this fpolicy connection collection get default response has a 2xx status code
func (o *FpolicyConnectionCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy connection collection get default response has a 3xx status code
func (o *FpolicyConnectionCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy connection collection get default response has a 4xx status code
func (o *FpolicyConnectionCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy connection collection get default response has a 5xx status code
func (o *FpolicyConnectionCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy connection collection get default response a status code equal to that given
func (o *FpolicyConnectionCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FpolicyConnectionCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/connections][%d] fpolicy_connection_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *FpolicyConnectionCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/fpolicy/{svm.uuid}/connections][%d] fpolicy_connection_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *FpolicyConnectionCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyConnectionCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}