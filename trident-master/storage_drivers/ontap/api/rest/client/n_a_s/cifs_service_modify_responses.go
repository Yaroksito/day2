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

// CifsServiceModifyReader is a Reader for the CifsServiceModify structure.
type CifsServiceModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsServiceModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCifsServiceModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsServiceModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsServiceModifyAccepted creates a CifsServiceModifyAccepted with default headers values
func NewCifsServiceModifyAccepted() *CifsServiceModifyAccepted {
	return &CifsServiceModifyAccepted{}
}

/*
CifsServiceModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type CifsServiceModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this cifs service modify accepted response has a 2xx status code
func (o *CifsServiceModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs service modify accepted response has a 3xx status code
func (o *CifsServiceModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs service modify accepted response has a 4xx status code
func (o *CifsServiceModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs service modify accepted response has a 5xx status code
func (o *CifsServiceModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs service modify accepted response a status code equal to that given
func (o *CifsServiceModifyAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CifsServiceModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/services/{svm.uuid}][%d] cifsServiceModifyAccepted  %+v", 202, o.Payload)
}

func (o *CifsServiceModifyAccepted) String() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/services/{svm.uuid}][%d] cifsServiceModifyAccepted  %+v", 202, o.Payload)
}

func (o *CifsServiceModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *CifsServiceModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsServiceModifyDefault creates a CifsServiceModifyDefault with default headers values
func NewCifsServiceModifyDefault(code int) *CifsServiceModifyDefault {
	return &CifsServiceModifyDefault{
		_statusCode: code,
	}
}

/*
	CifsServiceModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 4915251    | STARTTLS and LDAPS cannot be used together.|
*/
type CifsServiceModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs service modify default response
func (o *CifsServiceModifyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cifs service modify default response has a 2xx status code
func (o *CifsServiceModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs service modify default response has a 3xx status code
func (o *CifsServiceModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs service modify default response has a 4xx status code
func (o *CifsServiceModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs service modify default response has a 5xx status code
func (o *CifsServiceModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs service modify default response a status code equal to that given
func (o *CifsServiceModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CifsServiceModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/services/{svm.uuid}][%d] cifs_service_modify default  %+v", o._statusCode, o.Payload)
}

func (o *CifsServiceModifyDefault) String() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/services/{svm.uuid}][%d] cifs_service_modify default  %+v", o._statusCode, o.Payload)
}

func (o *CifsServiceModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsServiceModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}