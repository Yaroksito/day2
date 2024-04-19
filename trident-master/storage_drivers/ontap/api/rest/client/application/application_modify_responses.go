// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ApplicationModifyReader is a Reader for the ApplicationModify structure.
type ApplicationModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewApplicationModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationModifyAccepted creates a ApplicationModifyAccepted with default headers values
func NewApplicationModifyAccepted() *ApplicationModifyAccepted {
	return &ApplicationModifyAccepted{}
}

/*
ApplicationModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ApplicationModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this application modify accepted response has a 2xx status code
func (o *ApplicationModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this application modify accepted response has a 3xx status code
func (o *ApplicationModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this application modify accepted response has a 4xx status code
func (o *ApplicationModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this application modify accepted response has a 5xx status code
func (o *ApplicationModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this application modify accepted response a status code equal to that given
func (o *ApplicationModifyAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *ApplicationModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /application/applications/{uuid}][%d] applicationModifyAccepted  %+v", 202, o.Payload)
}

func (o *ApplicationModifyAccepted) String() string {
	return fmt.Sprintf("[PATCH /application/applications/{uuid}][%d] applicationModifyAccepted  %+v", 202, o.Payload)
}

func (o *ApplicationModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *ApplicationModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationModifyDefault creates a ApplicationModifyDefault with default headers values
func NewApplicationModifyDefault(code int) *ApplicationModifyDefault {
	return &ApplicationModifyDefault{
		_statusCode: code,
	}
}

/*
ApplicationModifyDefault describes a response with status code -1, with default header values.

Error
*/
type ApplicationModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the application modify default response
func (o *ApplicationModifyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this application modify default response has a 2xx status code
func (o *ApplicationModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this application modify default response has a 3xx status code
func (o *ApplicationModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this application modify default response has a 4xx status code
func (o *ApplicationModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this application modify default response has a 5xx status code
func (o *ApplicationModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this application modify default response a status code equal to that given
func (o *ApplicationModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ApplicationModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /application/applications/{uuid}][%d] application_modify default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationModifyDefault) String() string {
	return fmt.Sprintf("[PATCH /application/applications/{uuid}][%d] application_modify default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ApplicationModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}