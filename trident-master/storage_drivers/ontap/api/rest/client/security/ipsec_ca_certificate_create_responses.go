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

// IpsecCaCertificateCreateReader is a Reader for the IpsecCaCertificateCreate structure.
type IpsecCaCertificateCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpsecCaCertificateCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpsecCaCertificateCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpsecCaCertificateCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpsecCaCertificateCreateCreated creates a IpsecCaCertificateCreateCreated with default headers values
func NewIpsecCaCertificateCreateCreated() *IpsecCaCertificateCreateCreated {
	return &IpsecCaCertificateCreateCreated{}
}

/*
IpsecCaCertificateCreateCreated describes a response with status code 201, with default header values.

Created
*/
type IpsecCaCertificateCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.IpsecCaCertificateResponse
}

// IsSuccess returns true when this ipsec ca certificate create created response has a 2xx status code
func (o *IpsecCaCertificateCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipsec ca certificate create created response has a 3xx status code
func (o *IpsecCaCertificateCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipsec ca certificate create created response has a 4xx status code
func (o *IpsecCaCertificateCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipsec ca certificate create created response has a 5xx status code
func (o *IpsecCaCertificateCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipsec ca certificate create created response a status code equal to that given
func (o *IpsecCaCertificateCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *IpsecCaCertificateCreateCreated) Error() string {
	return fmt.Sprintf("[POST /security/ipsec/ca-certificates][%d] ipsecCaCertificateCreateCreated  %+v", 201, o.Payload)
}

func (o *IpsecCaCertificateCreateCreated) String() string {
	return fmt.Sprintf("[POST /security/ipsec/ca-certificates][%d] ipsecCaCertificateCreateCreated  %+v", 201, o.Payload)
}

func (o *IpsecCaCertificateCreateCreated) GetPayload() *models.IpsecCaCertificateResponse {
	return o.Payload
}

func (o *IpsecCaCertificateCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.IpsecCaCertificateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpsecCaCertificateCreateDefault creates a IpsecCaCertificateCreateDefault with default headers values
func NewIpsecCaCertificateCreateDefault(code int) *IpsecCaCertificateCreateDefault {
	return &IpsecCaCertificateCreateDefault{
		_statusCode: code,
	}
}

/*
	IpsecCaCertificateCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 66257296 | CA certificate is not installed. |
*/
type IpsecCaCertificateCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ipsec ca certificate create default response
func (o *IpsecCaCertificateCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipsec ca certificate create default response has a 2xx status code
func (o *IpsecCaCertificateCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipsec ca certificate create default response has a 3xx status code
func (o *IpsecCaCertificateCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipsec ca certificate create default response has a 4xx status code
func (o *IpsecCaCertificateCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipsec ca certificate create default response has a 5xx status code
func (o *IpsecCaCertificateCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipsec ca certificate create default response a status code equal to that given
func (o *IpsecCaCertificateCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpsecCaCertificateCreateDefault) Error() string {
	return fmt.Sprintf("[POST /security/ipsec/ca-certificates][%d] ipsec_ca_certificate_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpsecCaCertificateCreateDefault) String() string {
	return fmt.Sprintf("[POST /security/ipsec/ca-certificates][%d] ipsec_ca_certificate_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpsecCaCertificateCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IpsecCaCertificateCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}