// Code generated by go-swagger; DO NOT EDIT.

package snaplock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnaplockLogGetReader is a Reader for the SnaplockLogGet structure.
type SnaplockLogGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLogGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockLogGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLogGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLogGetOK creates a SnaplockLogGetOK with default headers values
func NewSnaplockLogGetOK() *SnaplockLogGetOK {
	return &SnaplockLogGetOK{}
}

/*
SnaplockLogGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockLogGetOK struct {
	Payload *models.SnaplockLog
}

// IsSuccess returns true when this snaplock log get o k response has a 2xx status code
func (o *SnaplockLogGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock log get o k response has a 3xx status code
func (o *SnaplockLogGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock log get o k response has a 4xx status code
func (o *SnaplockLogGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock log get o k response has a 5xx status code
func (o *SnaplockLogGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock log get o k response a status code equal to that given
func (o *SnaplockLogGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *SnaplockLogGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplockLogGetOK  %+v", 200, o.Payload)
}

func (o *SnaplockLogGetOK) String() string {
	return fmt.Sprintf("[GET /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplockLogGetOK  %+v", 200, o.Payload)
}

func (o *SnaplockLogGetOK) GetPayload() *models.SnaplockLog {
	return o.Payload
}

func (o *SnaplockLogGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockLog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLogGetDefault creates a SnaplockLogGetDefault with default headers values
func NewSnaplockLogGetDefault(code int) *SnaplockLogGetDefault {
	return &SnaplockLogGetDefault{
		_statusCode: code,
	}
}

/*
SnaplockLogGetDefault describes a response with status code -1, with default header values.

Error
*/
type SnaplockLogGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock log get default response
func (o *SnaplockLogGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this snaplock log get default response has a 2xx status code
func (o *SnaplockLogGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock log get default response has a 3xx status code
func (o *SnaplockLogGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock log get default response has a 4xx status code
func (o *SnaplockLogGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock log get default response has a 5xx status code
func (o *SnaplockLogGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock log get default response a status code equal to that given
func (o *SnaplockLogGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SnaplockLogGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplock_log_get default  %+v", o._statusCode, o.Payload)
}

func (o *SnaplockLogGetDefault) String() string {
	return fmt.Sprintf("[GET /storage/snaplock/audit-logs/{svm.uuid}][%d] snaplock_log_get default  %+v", o._statusCode, o.Payload)
}

func (o *SnaplockLogGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLogGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}