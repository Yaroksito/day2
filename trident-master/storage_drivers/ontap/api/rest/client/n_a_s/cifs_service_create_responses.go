// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// CifsServiceCreateReader is a Reader for the CifsServiceCreate structure.
type CifsServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCifsServiceCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsServiceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsServiceCreateAccepted creates a CifsServiceCreateAccepted with default headers values
func NewCifsServiceCreateAccepted() *CifsServiceCreateAccepted {
	return &CifsServiceCreateAccepted{}
}

/*
CifsServiceCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type CifsServiceCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cifs service create accepted response has a 2xx status code
func (o *CifsServiceCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs service create accepted response has a 3xx status code
func (o *CifsServiceCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs service create accepted response has a 4xx status code
func (o *CifsServiceCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs service create accepted response has a 5xx status code
func (o *CifsServiceCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs service create accepted response a status code equal to that given
func (o *CifsServiceCreateAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CifsServiceCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /protocols/cifs/services][%d] cifsServiceCreateAccepted  %+v", 202, o.Payload)
}

func (o *CifsServiceCreateAccepted) String() string {
	return fmt.Sprintf("[POST /protocols/cifs/services][%d] cifsServiceCreateAccepted  %+v", 202, o.Payload)
}

func (o *CifsServiceCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *CifsServiceCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsServiceCreateDefault creates a CifsServiceCreateDefault with default headers values
func NewCifsServiceCreateDefault(code int) *CifsServiceCreateDefault {
	return &CifsServiceCreateDefault{
		_statusCode: code,
	}
}

/*
	CifsServiceCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 4915251    | STARTTLS and LDAPS cannot be used together.|
*/
type CifsServiceCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs service create default response
func (o *CifsServiceCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cifs service create default response has a 2xx status code
func (o *CifsServiceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs service create default response has a 3xx status code
func (o *CifsServiceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs service create default response has a 4xx status code
func (o *CifsServiceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs service create default response has a 5xx status code
func (o *CifsServiceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs service create default response a status code equal to that given
func (o *CifsServiceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CifsServiceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/cifs/services][%d] cifs_service_create default  %+v", o._statusCode, o.Payload)
}

func (o *CifsServiceCreateDefault) String() string {
	return fmt.Sprintf("[POST /protocols/cifs/services][%d] cifs_service_create default  %+v", o._statusCode, o.Payload)
}

func (o *CifsServiceCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsServiceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}