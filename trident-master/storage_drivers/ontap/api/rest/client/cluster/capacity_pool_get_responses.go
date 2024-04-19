// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// CapacityPoolGetReader is a Reader for the CapacityPoolGet structure.
type CapacityPoolGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CapacityPoolGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCapacityPoolGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCapacityPoolGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCapacityPoolGetOK creates a CapacityPoolGetOK with default headers values
func NewCapacityPoolGetOK() *CapacityPoolGetOK {
	return &CapacityPoolGetOK{}
}

/*
CapacityPoolGetOK describes a response with status code 200, with default header values.

OK
*/
type CapacityPoolGetOK struct {
	Payload *models.CapacityPool
}

// IsSuccess returns true when this capacity pool get o k response has a 2xx status code
func (o *CapacityPoolGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this capacity pool get o k response has a 3xx status code
func (o *CapacityPoolGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this capacity pool get o k response has a 4xx status code
func (o *CapacityPoolGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this capacity pool get o k response has a 5xx status code
func (o *CapacityPoolGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this capacity pool get o k response a status code equal to that given
func (o *CapacityPoolGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *CapacityPoolGetOK) Error() string {
	return fmt.Sprintf("[GET /cluster/licensing/capacity-pools/{serial_number}][%d] capacityPoolGetOK  %+v", 200, o.Payload)
}

func (o *CapacityPoolGetOK) String() string {
	return fmt.Sprintf("[GET /cluster/licensing/capacity-pools/{serial_number}][%d] capacityPoolGetOK  %+v", 200, o.Payload)
}

func (o *CapacityPoolGetOK) GetPayload() *models.CapacityPool {
	return o.Payload
}

func (o *CapacityPoolGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CapacityPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCapacityPoolGetDefault creates a CapacityPoolGetDefault with default headers values
func NewCapacityPoolGetDefault(code int) *CapacityPoolGetDefault {
	return &CapacityPoolGetDefault{
		_statusCode: code,
	}
}

/*
CapacityPoolGetDefault describes a response with status code -1, with default header values.

Error
*/
type CapacityPoolGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the capacity pool get default response
func (o *CapacityPoolGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this capacity pool get default response has a 2xx status code
func (o *CapacityPoolGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this capacity pool get default response has a 3xx status code
func (o *CapacityPoolGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this capacity pool get default response has a 4xx status code
func (o *CapacityPoolGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this capacity pool get default response has a 5xx status code
func (o *CapacityPoolGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this capacity pool get default response a status code equal to that given
func (o *CapacityPoolGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CapacityPoolGetDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/licensing/capacity-pools/{serial_number}][%d] capacity_pool_get default  %+v", o._statusCode, o.Payload)
}

func (o *CapacityPoolGetDefault) String() string {
	return fmt.Sprintf("[GET /cluster/licensing/capacity-pools/{serial_number}][%d] capacity_pool_get default  %+v", o._statusCode, o.Payload)
}

func (o *CapacityPoolGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CapacityPoolGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
