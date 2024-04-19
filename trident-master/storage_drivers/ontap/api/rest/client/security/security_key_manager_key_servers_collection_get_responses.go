// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SecurityKeyManagerKeyServersCollectionGetReader is a Reader for the SecurityKeyManagerKeyServersCollectionGet structure.
type SecurityKeyManagerKeyServersCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerKeyServersCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeyManagerKeyServersCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerKeyServersCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerKeyServersCollectionGetOK creates a SecurityKeyManagerKeyServersCollectionGetOK with default headers values
func NewSecurityKeyManagerKeyServersCollectionGetOK() *SecurityKeyManagerKeyServersCollectionGetOK {
	return &SecurityKeyManagerKeyServersCollectionGetOK{}
}

/*
SecurityKeyManagerKeyServersCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeyManagerKeyServersCollectionGetOK struct {
	Payload *models.KeyServerResponse
}

// IsSuccess returns true when this security key manager key servers collection get o k response has a 2xx status code
func (o *SecurityKeyManagerKeyServersCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security key manager key servers collection get o k response has a 3xx status code
func (o *SecurityKeyManagerKeyServersCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security key manager key servers collection get o k response has a 4xx status code
func (o *SecurityKeyManagerKeyServersCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security key manager key servers collection get o k response has a 5xx status code
func (o *SecurityKeyManagerKeyServersCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security key manager key servers collection get o k response a status code equal to that given
func (o *SecurityKeyManagerKeyServersCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *SecurityKeyManagerKeyServersCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /security/key-managers/{uuid}/key-servers][%d] securityKeyManagerKeyServersCollectionGetOK  %+v", 200, o.Payload)
}

func (o *SecurityKeyManagerKeyServersCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /security/key-managers/{uuid}/key-servers][%d] securityKeyManagerKeyServersCollectionGetOK  %+v", 200, o.Payload)
}

func (o *SecurityKeyManagerKeyServersCollectionGetOK) GetPayload() *models.KeyServerResponse {
	return o.Payload
}

func (o *SecurityKeyManagerKeyServersCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyServerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeyManagerKeyServersCollectionGetDefault creates a SecurityKeyManagerKeyServersCollectionGetDefault with default headers values
func NewSecurityKeyManagerKeyServersCollectionGetDefault(code int) *SecurityKeyManagerKeyServersCollectionGetDefault {
	return &SecurityKeyManagerKeyServersCollectionGetDefault{
		_statusCode: code,
	}
}

/*
SecurityKeyManagerKeyServersCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityKeyManagerKeyServersCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the security key manager key servers collection get default response
func (o *SecurityKeyManagerKeyServersCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this security key manager key servers collection get default response has a 2xx status code
func (o *SecurityKeyManagerKeyServersCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security key manager key servers collection get default response has a 3xx status code
func (o *SecurityKeyManagerKeyServersCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security key manager key servers collection get default response has a 4xx status code
func (o *SecurityKeyManagerKeyServersCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security key manager key servers collection get default response has a 5xx status code
func (o *SecurityKeyManagerKeyServersCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security key manager key servers collection get default response a status code equal to that given
func (o *SecurityKeyManagerKeyServersCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SecurityKeyManagerKeyServersCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /security/key-managers/{uuid}/key-servers][%d] security_key_manager_key_servers_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *SecurityKeyManagerKeyServersCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /security/key-managers/{uuid}/key-servers][%d] security_key_manager_key_servers_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *SecurityKeyManagerKeyServersCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerKeyServersCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
