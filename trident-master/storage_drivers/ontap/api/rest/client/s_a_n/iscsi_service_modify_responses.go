// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// IscsiServiceModifyReader is a Reader for the IscsiServiceModify structure.
type IscsiServiceModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IscsiServiceModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIscsiServiceModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIscsiServiceModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIscsiServiceModifyOK creates a IscsiServiceModifyOK with default headers values
func NewIscsiServiceModifyOK() *IscsiServiceModifyOK {
	return &IscsiServiceModifyOK{}
}

/*
IscsiServiceModifyOK describes a response with status code 200, with default header values.

OK
*/
type IscsiServiceModifyOK struct {
}

// IsSuccess returns true when this iscsi service modify o k response has a 2xx status code
func (o *IscsiServiceModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iscsi service modify o k response has a 3xx status code
func (o *IscsiServiceModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iscsi service modify o k response has a 4xx status code
func (o *IscsiServiceModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iscsi service modify o k response has a 5xx status code
func (o *IscsiServiceModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iscsi service modify o k response a status code equal to that given
func (o *IscsiServiceModifyOK) IsCode(code int) bool {
	return code == 200
}

func (o *IscsiServiceModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/san/iscsi/services/{svm.uuid}][%d] iscsiServiceModifyOK ", 200)
}

func (o *IscsiServiceModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/san/iscsi/services/{svm.uuid}][%d] iscsiServiceModifyOK ", 200)
}

func (o *IscsiServiceModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIscsiServiceModifyDefault creates a IscsiServiceModifyDefault with default headers values
func NewIscsiServiceModifyDefault(code int) *IscsiServiceModifyDefault {
	return &IscsiServiceModifyDefault{
		_statusCode: code,
	}
}

/*
	IscsiServiceModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621462 | An SVM with the specified UUID does not exist. |
| 5374078 | The SVM does not have an iSCSI service. |
*/
type IscsiServiceModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the iscsi service modify default response
func (o *IscsiServiceModifyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iscsi service modify default response has a 2xx status code
func (o *IscsiServiceModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iscsi service modify default response has a 3xx status code
func (o *IscsiServiceModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iscsi service modify default response has a 4xx status code
func (o *IscsiServiceModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iscsi service modify default response has a 5xx status code
func (o *IscsiServiceModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iscsi service modify default response a status code equal to that given
func (o *IscsiServiceModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IscsiServiceModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/san/iscsi/services/{svm.uuid}][%d] iscsi_service_modify default  %+v", o._statusCode, o.Payload)
}

func (o *IscsiServiceModifyDefault) String() string {
	return fmt.Sprintf("[PATCH /protocols/san/iscsi/services/{svm.uuid}][%d] iscsi_service_modify default  %+v", o._statusCode, o.Payload)
}

func (o *IscsiServiceModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IscsiServiceModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}