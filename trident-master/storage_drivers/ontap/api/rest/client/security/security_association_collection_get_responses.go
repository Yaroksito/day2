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

// SecurityAssociationCollectionGetReader is a Reader for the SecurityAssociationCollectionGet structure.
type SecurityAssociationCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityAssociationCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityAssociationCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityAssociationCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityAssociationCollectionGetOK creates a SecurityAssociationCollectionGetOK with default headers values
func NewSecurityAssociationCollectionGetOK() *SecurityAssociationCollectionGetOK {
	return &SecurityAssociationCollectionGetOK{}
}

/*
SecurityAssociationCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SecurityAssociationCollectionGetOK struct {
	Payload *models.SecurityAssociationResponse
}

// IsSuccess returns true when this security association collection get o k response has a 2xx status code
func (o *SecurityAssociationCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security association collection get o k response has a 3xx status code
func (o *SecurityAssociationCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security association collection get o k response has a 4xx status code
func (o *SecurityAssociationCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security association collection get o k response has a 5xx status code
func (o *SecurityAssociationCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security association collection get o k response a status code equal to that given
func (o *SecurityAssociationCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *SecurityAssociationCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /security/ipsec/security-associations][%d] securityAssociationCollectionGetOK  %+v", 200, o.Payload)
}

func (o *SecurityAssociationCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /security/ipsec/security-associations][%d] securityAssociationCollectionGetOK  %+v", 200, o.Payload)
}

func (o *SecurityAssociationCollectionGetOK) GetPayload() *models.SecurityAssociationResponse {
	return o.Payload
}

func (o *SecurityAssociationCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityAssociationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityAssociationCollectionGetDefault creates a SecurityAssociationCollectionGetDefault with default headers values
func NewSecurityAssociationCollectionGetDefault(code int) *SecurityAssociationCollectionGetDefault {
	return &SecurityAssociationCollectionGetDefault{
		_statusCode: code,
	}
}

/*
SecurityAssociationCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityAssociationCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the security association collection get default response
func (o *SecurityAssociationCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this security association collection get default response has a 2xx status code
func (o *SecurityAssociationCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security association collection get default response has a 3xx status code
func (o *SecurityAssociationCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security association collection get default response has a 4xx status code
func (o *SecurityAssociationCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security association collection get default response has a 5xx status code
func (o *SecurityAssociationCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security association collection get default response a status code equal to that given
func (o *SecurityAssociationCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SecurityAssociationCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /security/ipsec/security-associations][%d] security_association_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *SecurityAssociationCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /security/ipsec/security-associations][%d] security_association_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *SecurityAssociationCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityAssociationCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}