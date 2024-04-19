// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// AutoUpdateInfoModifyReader is a Reader for the AutoUpdateInfoModify structure.
type AutoUpdateInfoModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AutoUpdateInfoModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAutoUpdateInfoModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAutoUpdateInfoModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAutoUpdateInfoModifyOK creates a AutoUpdateInfoModifyOK with default headers values
func NewAutoUpdateInfoModifyOK() *AutoUpdateInfoModifyOK {
	return &AutoUpdateInfoModifyOK{}
}

/*
AutoUpdateInfoModifyOK describes a response with status code 200, with default header values.

OK
*/
type AutoUpdateInfoModifyOK struct {
}

// IsSuccess returns true when this auto update info modify o k response has a 2xx status code
func (o *AutoUpdateInfoModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auto update info modify o k response has a 3xx status code
func (o *AutoUpdateInfoModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auto update info modify o k response has a 4xx status code
func (o *AutoUpdateInfoModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auto update info modify o k response has a 5xx status code
func (o *AutoUpdateInfoModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auto update info modify o k response a status code equal to that given
func (o *AutoUpdateInfoModifyOK) IsCode(code int) bool {
	return code == 200
}

func (o *AutoUpdateInfoModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /support/auto-update][%d] autoUpdateInfoModifyOK ", 200)
}

func (o *AutoUpdateInfoModifyOK) String() string {
	return fmt.Sprintf("[PATCH /support/auto-update][%d] autoUpdateInfoModifyOK ", 200)
}

func (o *AutoUpdateInfoModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAutoUpdateInfoModifyDefault creates a AutoUpdateInfoModifyDefault with default headers values
func NewAutoUpdateInfoModifyDefault(code int) *AutoUpdateInfoModifyDefault {
	return &AutoUpdateInfoModifyDefault{
		_statusCode: code,
	}
}

/*
	AutoUpdateInfoModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262179 | Unexpected argument. |
| 131072216 | Automatic update requires AutoSupport to be enabled. |
| 131072217 | Automatic update requires AutoSupport OnDemand to be enabled. |
| 131072218 | Automatic update requires AutoSupport to use the HTTPS transport. |
*/
type AutoUpdateInfoModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the auto update info modify default response
func (o *AutoUpdateInfoModifyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this auto update info modify default response has a 2xx status code
func (o *AutoUpdateInfoModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this auto update info modify default response has a 3xx status code
func (o *AutoUpdateInfoModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this auto update info modify default response has a 4xx status code
func (o *AutoUpdateInfoModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this auto update info modify default response has a 5xx status code
func (o *AutoUpdateInfoModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this auto update info modify default response a status code equal to that given
func (o *AutoUpdateInfoModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AutoUpdateInfoModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /support/auto-update][%d] auto_update_info_modify default  %+v", o._statusCode, o.Payload)
}

func (o *AutoUpdateInfoModifyDefault) String() string {
	return fmt.Sprintf("[PATCH /support/auto-update][%d] auto_update_info_modify default  %+v", o._statusCode, o.Payload)
}

func (o *AutoUpdateInfoModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AutoUpdateInfoModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
