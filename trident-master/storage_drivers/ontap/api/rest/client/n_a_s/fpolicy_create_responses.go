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

// FpolicyCreateReader is a Reader for the FpolicyCreate structure.
type FpolicyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewFpolicyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyCreateCreated creates a FpolicyCreateCreated with default headers values
func NewFpolicyCreateCreated() *FpolicyCreateCreated {
	return &FpolicyCreateCreated{}
}

/*
FpolicyCreateCreated describes a response with status code 201, with default header values.

Created
*/
type FpolicyCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.FpolicyResponse
}

// IsSuccess returns true when this fpolicy create created response has a 2xx status code
func (o *FpolicyCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy create created response has a 3xx status code
func (o *FpolicyCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy create created response has a 4xx status code
func (o *FpolicyCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy create created response has a 5xx status code
func (o *FpolicyCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy create created response a status code equal to that given
func (o *FpolicyCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *FpolicyCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/fpolicy][%d] fpolicyCreateCreated  %+v", 201, o.Payload)
}

func (o *FpolicyCreateCreated) String() string {
	return fmt.Sprintf("[POST /protocols/fpolicy][%d] fpolicyCreateCreated  %+v", 201, o.Payload)
}

func (o *FpolicyCreateCreated) GetPayload() *models.FpolicyResponse {
	return o.Payload
}

func (o *FpolicyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.FpolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFpolicyCreateDefault creates a FpolicyCreateDefault with default headers values
func NewFpolicyCreateDefault(code int) *FpolicyCreateDefault {
	return &FpolicyCreateDefault{
		_statusCode: code,
	}
}

/*
	FpolicyCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 9765032    | The FPolicy engine, FPolicy event or FPolicy policy specified already exists |
| 9765031    | If any of the FPolicy engine, FPolicy event, or FPolicy policy creation fails due to a systematic error or hardware failure, the cause of the failure is detailed in the error message |
| 2621706    | The SVM UUID specified belongs to different SVM |
| 2621462    | The SVM name specified does not exist |
*/
type FpolicyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fpolicy create default response
func (o *FpolicyCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this fpolicy create default response has a 2xx status code
func (o *FpolicyCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy create default response has a 3xx status code
func (o *FpolicyCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy create default response has a 4xx status code
func (o *FpolicyCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy create default response has a 5xx status code
func (o *FpolicyCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy create default response a status code equal to that given
func (o *FpolicyCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FpolicyCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/fpolicy][%d] fpolicy_create default  %+v", o._statusCode, o.Payload)
}

func (o *FpolicyCreateDefault) String() string {
	return fmt.Sprintf("[POST /protocols/fpolicy][%d] fpolicy_create default  %+v", o._statusCode, o.Payload)
}

func (o *FpolicyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
