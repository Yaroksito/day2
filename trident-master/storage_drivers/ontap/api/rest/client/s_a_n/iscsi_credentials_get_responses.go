// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// IscsiCredentialsGetReader is a Reader for the IscsiCredentialsGet structure.
type IscsiCredentialsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IscsiCredentialsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIscsiCredentialsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIscsiCredentialsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIscsiCredentialsGetOK creates a IscsiCredentialsGetOK with default headers values
func NewIscsiCredentialsGetOK() *IscsiCredentialsGetOK {
	return &IscsiCredentialsGetOK{}
}

/*
IscsiCredentialsGetOK describes a response with status code 200, with default header values.

OK
*/
type IscsiCredentialsGetOK struct {
	Payload *models.IscsiCredentials
}

// IsSuccess returns true when this iscsi credentials get o k response has a 2xx status code
func (o *IscsiCredentialsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iscsi credentials get o k response has a 3xx status code
func (o *IscsiCredentialsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iscsi credentials get o k response has a 4xx status code
func (o *IscsiCredentialsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iscsi credentials get o k response has a 5xx status code
func (o *IscsiCredentialsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iscsi credentials get o k response a status code equal to that given
func (o *IscsiCredentialsGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *IscsiCredentialsGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/san/iscsi/credentials/{svm.uuid}/{initiator}][%d] iscsiCredentialsGetOK  %+v", 200, o.Payload)
}

func (o *IscsiCredentialsGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/san/iscsi/credentials/{svm.uuid}/{initiator}][%d] iscsiCredentialsGetOK  %+v", 200, o.Payload)
}

func (o *IscsiCredentialsGetOK) GetPayload() *models.IscsiCredentials {
	return o.Payload
}

func (o *IscsiCredentialsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IscsiCredentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIscsiCredentialsGetDefault creates a IscsiCredentialsGetDefault with default headers values
func NewIscsiCredentialsGetDefault(code int) *IscsiCredentialsGetDefault {
	return &IscsiCredentialsGetDefault{
		_statusCode: code,
	}
}

/*
IscsiCredentialsGetDefault describes a response with status code -1, with default header values.

Error
*/
type IscsiCredentialsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the iscsi credentials get default response
func (o *IscsiCredentialsGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iscsi credentials get default response has a 2xx status code
func (o *IscsiCredentialsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iscsi credentials get default response has a 3xx status code
func (o *IscsiCredentialsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iscsi credentials get default response has a 4xx status code
func (o *IscsiCredentialsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iscsi credentials get default response has a 5xx status code
func (o *IscsiCredentialsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iscsi credentials get default response a status code equal to that given
func (o *IscsiCredentialsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IscsiCredentialsGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/san/iscsi/credentials/{svm.uuid}/{initiator}][%d] iscsi_credentials_get default  %+v", o._statusCode, o.Payload)
}

func (o *IscsiCredentialsGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/san/iscsi/credentials/{svm.uuid}/{initiator}][%d] iscsi_credentials_get default  %+v", o._statusCode, o.Payload)
}

func (o *IscsiCredentialsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IscsiCredentialsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}