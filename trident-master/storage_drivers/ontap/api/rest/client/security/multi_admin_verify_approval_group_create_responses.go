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

// MultiAdminVerifyApprovalGroupCreateReader is a Reader for the MultiAdminVerifyApprovalGroupCreate structure.
type MultiAdminVerifyApprovalGroupCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MultiAdminVerifyApprovalGroupCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewMultiAdminVerifyApprovalGroupCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMultiAdminVerifyApprovalGroupCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMultiAdminVerifyApprovalGroupCreateCreated creates a MultiAdminVerifyApprovalGroupCreateCreated with default headers values
func NewMultiAdminVerifyApprovalGroupCreateCreated() *MultiAdminVerifyApprovalGroupCreateCreated {
	return &MultiAdminVerifyApprovalGroupCreateCreated{}
}

/*
MultiAdminVerifyApprovalGroupCreateCreated describes a response with status code 201, with default header values.

Created
*/
type MultiAdminVerifyApprovalGroupCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.MultiAdminVerifyApprovalGroupResponse
}

// IsSuccess returns true when this multi admin verify approval group create created response has a 2xx status code
func (o *MultiAdminVerifyApprovalGroupCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this multi admin verify approval group create created response has a 3xx status code
func (o *MultiAdminVerifyApprovalGroupCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this multi admin verify approval group create created response has a 4xx status code
func (o *MultiAdminVerifyApprovalGroupCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this multi admin verify approval group create created response has a 5xx status code
func (o *MultiAdminVerifyApprovalGroupCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this multi admin verify approval group create created response a status code equal to that given
func (o *MultiAdminVerifyApprovalGroupCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *MultiAdminVerifyApprovalGroupCreateCreated) Error() string {
	return fmt.Sprintf("[POST /security/multi-admin-verify/approval-groups][%d] multiAdminVerifyApprovalGroupCreateCreated  %+v", 201, o.Payload)
}

func (o *MultiAdminVerifyApprovalGroupCreateCreated) String() string {
	return fmt.Sprintf("[POST /security/multi-admin-verify/approval-groups][%d] multiAdminVerifyApprovalGroupCreateCreated  %+v", 201, o.Payload)
}

func (o *MultiAdminVerifyApprovalGroupCreateCreated) GetPayload() *models.MultiAdminVerifyApprovalGroupResponse {
	return o.Payload
}

func (o *MultiAdminVerifyApprovalGroupCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.MultiAdminVerifyApprovalGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiAdminVerifyApprovalGroupCreateDefault creates a MultiAdminVerifyApprovalGroupCreateDefault with default headers values
func NewMultiAdminVerifyApprovalGroupCreateDefault(code int) *MultiAdminVerifyApprovalGroupCreateDefault {
	return &MultiAdminVerifyApprovalGroupCreateDefault{
		_statusCode: code,
	}
}

/*
	MultiAdminVerifyApprovalGroupCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262309 | The feature must be enabled first. |
*/
type MultiAdminVerifyApprovalGroupCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the multi admin verify approval group create default response
func (o *MultiAdminVerifyApprovalGroupCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this multi admin verify approval group create default response has a 2xx status code
func (o *MultiAdminVerifyApprovalGroupCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this multi admin verify approval group create default response has a 3xx status code
func (o *MultiAdminVerifyApprovalGroupCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this multi admin verify approval group create default response has a 4xx status code
func (o *MultiAdminVerifyApprovalGroupCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this multi admin verify approval group create default response has a 5xx status code
func (o *MultiAdminVerifyApprovalGroupCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this multi admin verify approval group create default response a status code equal to that given
func (o *MultiAdminVerifyApprovalGroupCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *MultiAdminVerifyApprovalGroupCreateDefault) Error() string {
	return fmt.Sprintf("[POST /security/multi-admin-verify/approval-groups][%d] multi_admin_verify_approval_group_create default  %+v", o._statusCode, o.Payload)
}

func (o *MultiAdminVerifyApprovalGroupCreateDefault) String() string {
	return fmt.Sprintf("[POST /security/multi-admin-verify/approval-groups][%d] multi_admin_verify_approval_group_create default  %+v", o._statusCode, o.Payload)
}

func (o *MultiAdminVerifyApprovalGroupCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MultiAdminVerifyApprovalGroupCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
