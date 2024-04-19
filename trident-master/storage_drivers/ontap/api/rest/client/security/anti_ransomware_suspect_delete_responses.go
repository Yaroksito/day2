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

// AntiRansomwareSuspectDeleteReader is a Reader for the AntiRansomwareSuspectDelete structure.
type AntiRansomwareSuspectDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AntiRansomwareSuspectDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewAntiRansomwareSuspectDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAntiRansomwareSuspectDeleteAccepted creates a AntiRansomwareSuspectDeleteAccepted with default headers values
func NewAntiRansomwareSuspectDeleteAccepted() *AntiRansomwareSuspectDeleteAccepted {
	return &AntiRansomwareSuspectDeleteAccepted{}
}

/*
AntiRansomwareSuspectDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AntiRansomwareSuspectDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this anti ransomware suspect delete accepted response has a 2xx status code
func (o *AntiRansomwareSuspectDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this anti ransomware suspect delete accepted response has a 3xx status code
func (o *AntiRansomwareSuspectDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this anti ransomware suspect delete accepted response has a 4xx status code
func (o *AntiRansomwareSuspectDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this anti ransomware suspect delete accepted response has a 5xx status code
func (o *AntiRansomwareSuspectDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this anti ransomware suspect delete accepted response a status code equal to that given
func (o *AntiRansomwareSuspectDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *AntiRansomwareSuspectDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /security/anti-ransomware/suspects/{volume.uuid}][%d] antiRansomwareSuspectDeleteAccepted  %+v", 202, o.Payload)
}

func (o *AntiRansomwareSuspectDeleteAccepted) String() string {
	return fmt.Sprintf("[DELETE /security/anti-ransomware/suspects/{volume.uuid}][%d] antiRansomwareSuspectDeleteAccepted  %+v", 202, o.Payload)
}

func (o *AntiRansomwareSuspectDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *AntiRansomwareSuspectDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}