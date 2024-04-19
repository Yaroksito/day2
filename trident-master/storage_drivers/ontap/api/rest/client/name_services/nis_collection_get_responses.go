// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NisCollectionGetReader is a Reader for the NisCollectionGet structure.
type NisCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NisCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNisCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNisCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNisCollectionGetOK creates a NisCollectionGetOK with default headers values
func NewNisCollectionGetOK() *NisCollectionGetOK {
	return &NisCollectionGetOK{}
}

/*
NisCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type NisCollectionGetOK struct {
	Payload *models.NisServiceResponse
}

// IsSuccess returns true when this nis collection get o k response has a 2xx status code
func (o *NisCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nis collection get o k response has a 3xx status code
func (o *NisCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nis collection get o k response has a 4xx status code
func (o *NisCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nis collection get o k response has a 5xx status code
func (o *NisCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nis collection get o k response a status code equal to that given
func (o *NisCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *NisCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /name-services/nis][%d] nisCollectionGetOK  %+v", 200, o.Payload)
}

func (o *NisCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /name-services/nis][%d] nisCollectionGetOK  %+v", 200, o.Payload)
}

func (o *NisCollectionGetOK) GetPayload() *models.NisServiceResponse {
	return o.Payload
}

func (o *NisCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NisServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNisCollectionGetDefault creates a NisCollectionGetDefault with default headers values
func NewNisCollectionGetDefault(code int) *NisCollectionGetDefault {
	return &NisCollectionGetDefault{
		_statusCode: code,
	}
}

/*
NisCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type NisCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the nis collection get default response
func (o *NisCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this nis collection get default response has a 2xx status code
func (o *NisCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nis collection get default response has a 3xx status code
func (o *NisCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nis collection get default response has a 4xx status code
func (o *NisCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nis collection get default response has a 5xx status code
func (o *NisCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nis collection get default response a status code equal to that given
func (o *NisCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NisCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /name-services/nis][%d] nis_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *NisCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /name-services/nis][%d] nis_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *NisCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NisCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}