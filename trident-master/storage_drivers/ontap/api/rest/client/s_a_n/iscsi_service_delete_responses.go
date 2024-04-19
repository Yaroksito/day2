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

// IscsiServiceDeleteReader is a Reader for the IscsiServiceDelete structure.
type IscsiServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IscsiServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIscsiServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIscsiServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIscsiServiceDeleteOK creates a IscsiServiceDeleteOK with default headers values
func NewIscsiServiceDeleteOK() *IscsiServiceDeleteOK {
	return &IscsiServiceDeleteOK{}
}

/*
IscsiServiceDeleteOK describes a response with status code 200, with default header values.

OK
*/
type IscsiServiceDeleteOK struct {
}

// IsSuccess returns true when this iscsi service delete o k response has a 2xx status code
func (o *IscsiServiceDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iscsi service delete o k response has a 3xx status code
func (o *IscsiServiceDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iscsi service delete o k response has a 4xx status code
func (o *IscsiServiceDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iscsi service delete o k response has a 5xx status code
func (o *IscsiServiceDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iscsi service delete o k response a status code equal to that given
func (o *IscsiServiceDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *IscsiServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/san/iscsi/services/{svm.uuid}][%d] iscsiServiceDeleteOK ", 200)
}

func (o *IscsiServiceDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/san/iscsi/services/{svm.uuid}][%d] iscsiServiceDeleteOK ", 200)
}

func (o *IscsiServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIscsiServiceDeleteDefault creates a IscsiServiceDeleteDefault with default headers values
func NewIscsiServiceDeleteDefault(code int) *IscsiServiceDeleteDefault {
	return &IscsiServiceDeleteDefault{
		_statusCode: code,
	}
}

/*
	IscsiServiceDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621462 | An SVM with the specified UUID does not exist. |
| 5373960 | The iSCSI service is enabled. The iSCSI service must be disabled before it can be deleted. |
| 5374078 | The SVM does not have an iSCSI service. |
*/
type IscsiServiceDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the iscsi service delete default response
func (o *IscsiServiceDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iscsi service delete default response has a 2xx status code
func (o *IscsiServiceDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iscsi service delete default response has a 3xx status code
func (o *IscsiServiceDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iscsi service delete default response has a 4xx status code
func (o *IscsiServiceDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iscsi service delete default response has a 5xx status code
func (o *IscsiServiceDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iscsi service delete default response a status code equal to that given
func (o *IscsiServiceDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IscsiServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/san/iscsi/services/{svm.uuid}][%d] iscsi_service_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IscsiServiceDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /protocols/san/iscsi/services/{svm.uuid}][%d] iscsi_service_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IscsiServiceDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IscsiServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}