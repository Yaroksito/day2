// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FcInterfaceGetReader is a Reader for the FcInterfaceGet structure.
type FcInterfaceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FcInterfaceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFcInterfaceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFcInterfaceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFcInterfaceGetOK creates a FcInterfaceGetOK with default headers values
func NewFcInterfaceGetOK() *FcInterfaceGetOK {
	return &FcInterfaceGetOK{}
}

/*
FcInterfaceGetOK describes a response with status code 200, with default header values.

OK
*/
type FcInterfaceGetOK struct {
	Payload *models.FcInterface
}

// IsSuccess returns true when this fc interface get o k response has a 2xx status code
func (o *FcInterfaceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fc interface get o k response has a 3xx status code
func (o *FcInterfaceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fc interface get o k response has a 4xx status code
func (o *FcInterfaceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fc interface get o k response has a 5xx status code
func (o *FcInterfaceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fc interface get o k response a status code equal to that given
func (o *FcInterfaceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *FcInterfaceGetOK) Error() string {
	return fmt.Sprintf("[GET /network/fc/interfaces/{uuid}][%d] fcInterfaceGetOK  %+v", 200, o.Payload)
}

func (o *FcInterfaceGetOK) String() string {
	return fmt.Sprintf("[GET /network/fc/interfaces/{uuid}][%d] fcInterfaceGetOK  %+v", 200, o.Payload)
}

func (o *FcInterfaceGetOK) GetPayload() *models.FcInterface {
	return o.Payload
}

func (o *FcInterfaceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FcInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFcInterfaceGetDefault creates a FcInterfaceGetDefault with default headers values
func NewFcInterfaceGetDefault(code int) *FcInterfaceGetDefault {
	return &FcInterfaceGetDefault{
		_statusCode: code,
	}
}

/*
FcInterfaceGetDefault describes a response with status code -1, with default header values.

Error
*/
type FcInterfaceGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fc interface get default response
func (o *FcInterfaceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this fc interface get default response has a 2xx status code
func (o *FcInterfaceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fc interface get default response has a 3xx status code
func (o *FcInterfaceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fc interface get default response has a 4xx status code
func (o *FcInterfaceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fc interface get default response has a 5xx status code
func (o *FcInterfaceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fc interface get default response a status code equal to that given
func (o *FcInterfaceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FcInterfaceGetDefault) Error() string {
	return fmt.Sprintf("[GET /network/fc/interfaces/{uuid}][%d] fc_interface_get default  %+v", o._statusCode, o.Payload)
}

func (o *FcInterfaceGetDefault) String() string {
	return fmt.Sprintf("[GET /network/fc/interfaces/{uuid}][%d] fc_interface_get default  %+v", o._statusCode, o.Payload)
}

func (o *FcInterfaceGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FcInterfaceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}