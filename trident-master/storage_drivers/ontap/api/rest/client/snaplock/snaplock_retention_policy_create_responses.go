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

// SnaplockRetentionPolicyCreateReader is a Reader for the SnaplockRetentionPolicyCreate structure.
type SnaplockRetentionPolicyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockRetentionPolicyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSnaplockRetentionPolicyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockRetentionPolicyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockRetentionPolicyCreateCreated creates a SnaplockRetentionPolicyCreateCreated with default headers values
func NewSnaplockRetentionPolicyCreateCreated() *SnaplockRetentionPolicyCreateCreated {
	return &SnaplockRetentionPolicyCreateCreated{}
}

/*
SnaplockRetentionPolicyCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SnaplockRetentionPolicyCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.SnaplockRetentionPolicy
}

// IsSuccess returns true when this snaplock retention policy create created response has a 2xx status code
func (o *SnaplockRetentionPolicyCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock retention policy create created response has a 3xx status code
func (o *SnaplockRetentionPolicyCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock retention policy create created response has a 4xx status code
func (o *SnaplockRetentionPolicyCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock retention policy create created response has a 5xx status code
func (o *SnaplockRetentionPolicyCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock retention policy create created response a status code equal to that given
func (o *SnaplockRetentionPolicyCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *SnaplockRetentionPolicyCreateCreated) Error() string {
	return fmt.Sprintf("[POST /storage/snaplock/event-retention/policies][%d] snaplockRetentionPolicyCreateCreated  %+v", 201, o.Payload)
}

func (o *SnaplockRetentionPolicyCreateCreated) String() string {
	return fmt.Sprintf("[POST /storage/snaplock/event-retention/policies][%d] snaplockRetentionPolicyCreateCreated  %+v", 201, o.Payload)
}

func (o *SnaplockRetentionPolicyCreateCreated) GetPayload() *models.SnaplockRetentionPolicy {
	return o.Payload
}

func (o *SnaplockRetentionPolicyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.SnaplockRetentionPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockRetentionPolicyCreateDefault creates a SnaplockRetentionPolicyCreateDefault with default headers values
func NewSnaplockRetentionPolicyCreateDefault(code int) *SnaplockRetentionPolicyCreateDefault {
	return &SnaplockRetentionPolicyCreateDefault{
		_statusCode: code,
	}
}

/*
SnaplockRetentionPolicyCreateDefault describes a response with status code -1, with default header values.

Error
*/
type SnaplockRetentionPolicyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock retention policy create default response
func (o *SnaplockRetentionPolicyCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this snaplock retention policy create default response has a 2xx status code
func (o *SnaplockRetentionPolicyCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock retention policy create default response has a 3xx status code
func (o *SnaplockRetentionPolicyCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock retention policy create default response has a 4xx status code
func (o *SnaplockRetentionPolicyCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock retention policy create default response has a 5xx status code
func (o *SnaplockRetentionPolicyCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock retention policy create default response a status code equal to that given
func (o *SnaplockRetentionPolicyCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SnaplockRetentionPolicyCreateDefault) Error() string {
	return fmt.Sprintf("[POST /storage/snaplock/event-retention/policies][%d] snaplock_retention_policy_create default  %+v", o._statusCode, o.Payload)
}

func (o *SnaplockRetentionPolicyCreateDefault) String() string {
	return fmt.Sprintf("[POST /storage/snaplock/event-retention/policies][%d] snaplock_retention_policy_create default  %+v", o._statusCode, o.Payload)
}

func (o *SnaplockRetentionPolicyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockRetentionPolicyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
