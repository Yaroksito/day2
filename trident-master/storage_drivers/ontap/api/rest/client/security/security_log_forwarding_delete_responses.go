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

// SecurityLogForwardingDeleteReader is a Reader for the SecurityLogForwardingDelete structure.
type SecurityLogForwardingDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityLogForwardingDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityLogForwardingDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityLogForwardingDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityLogForwardingDeleteOK creates a SecurityLogForwardingDeleteOK with default headers values
func NewSecurityLogForwardingDeleteOK() *SecurityLogForwardingDeleteOK {
	return &SecurityLogForwardingDeleteOK{}
}

/*
SecurityLogForwardingDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SecurityLogForwardingDeleteOK struct {
}

// IsSuccess returns true when this security log forwarding delete o k response has a 2xx status code
func (o *SecurityLogForwardingDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security log forwarding delete o k response has a 3xx status code
func (o *SecurityLogForwardingDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security log forwarding delete o k response has a 4xx status code
func (o *SecurityLogForwardingDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security log forwarding delete o k response has a 5xx status code
func (o *SecurityLogForwardingDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security log forwarding delete o k response a status code equal to that given
func (o *SecurityLogForwardingDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *SecurityLogForwardingDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/audit/destinations/{address}/{port}][%d] securityLogForwardingDeleteOK ", 200)
}

func (o *SecurityLogForwardingDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/audit/destinations/{address}/{port}][%d] securityLogForwardingDeleteOK ", 200)
}

func (o *SecurityLogForwardingDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecurityLogForwardingDeleteDefault creates a SecurityLogForwardingDeleteDefault with default headers values
func NewSecurityLogForwardingDeleteDefault(code int) *SecurityLogForwardingDeleteDefault {
	return &SecurityLogForwardingDeleteDefault{
		_statusCode: code,
	}
}

/*
SecurityLogForwardingDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type SecurityLogForwardingDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the security log forwarding delete default response
func (o *SecurityLogForwardingDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this security log forwarding delete default response has a 2xx status code
func (o *SecurityLogForwardingDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security log forwarding delete default response has a 3xx status code
func (o *SecurityLogForwardingDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security log forwarding delete default response has a 4xx status code
func (o *SecurityLogForwardingDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security log forwarding delete default response has a 5xx status code
func (o *SecurityLogForwardingDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security log forwarding delete default response a status code equal to that given
func (o *SecurityLogForwardingDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SecurityLogForwardingDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /security/audit/destinations/{address}/{port}][%d] security_log_forwarding_delete default  %+v", o._statusCode, o.Payload)
}

func (o *SecurityLogForwardingDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /security/audit/destinations/{address}/{port}][%d] security_log_forwarding_delete default  %+v", o._statusCode, o.Payload)
}

func (o *SecurityLogForwardingDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityLogForwardingDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}