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

// SoftwareModifyReader is a Reader for the SoftwareModify structure.
type SoftwareModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoftwareModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSoftwareModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSoftwareModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSoftwareModifyAccepted creates a SoftwareModifyAccepted with default headers values
func NewSoftwareModifyAccepted() *SoftwareModifyAccepted {
	return &SoftwareModifyAccepted{}
}

/*
SoftwareModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SoftwareModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this software modify accepted response has a 2xx status code
func (o *SoftwareModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this software modify accepted response has a 3xx status code
func (o *SoftwareModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this software modify accepted response has a 4xx status code
func (o *SoftwareModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this software modify accepted response has a 5xx status code
func (o *SoftwareModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this software modify accepted response a status code equal to that given
func (o *SoftwareModifyAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *SoftwareModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /cluster/software][%d] softwareModifyAccepted  %+v", 202, o.Payload)
}

func (o *SoftwareModifyAccepted) String() string {
	return fmt.Sprintf("[PATCH /cluster/software][%d] softwareModifyAccepted  %+v", 202, o.Payload)
}

func (o *SoftwareModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SoftwareModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwareModifyDefault creates a SoftwareModifyDefault with default headers values
func NewSoftwareModifyDefault(code int) *SoftwareModifyDefault {
	return &SoftwareModifyDefault{
		_statusCode: code,
	}
}

/*
SoftwareModifyDefault describes a response with status code -1, with default header values.

Error
*/
type SoftwareModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the software modify default response
func (o *SoftwareModifyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this software modify default response has a 2xx status code
func (o *SoftwareModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this software modify default response has a 3xx status code
func (o *SoftwareModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this software modify default response has a 4xx status code
func (o *SoftwareModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this software modify default response has a 5xx status code
func (o *SoftwareModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this software modify default response a status code equal to that given
func (o *SoftwareModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SoftwareModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /cluster/software][%d] software_modify default  %+v", o._statusCode, o.Payload)
}

func (o *SoftwareModifyDefault) String() string {
	return fmt.Sprintf("[PATCH /cluster/software][%d] software_modify default  %+v", o._statusCode, o.Payload)
}

func (o *SoftwareModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SoftwareModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
