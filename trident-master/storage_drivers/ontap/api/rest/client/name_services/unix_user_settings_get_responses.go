// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// UnixUserSettingsGetReader is a Reader for the UnixUserSettingsGet structure.
type UnixUserSettingsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnixUserSettingsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnixUserSettingsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnixUserSettingsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnixUserSettingsGetOK creates a UnixUserSettingsGetOK with default headers values
func NewUnixUserSettingsGetOK() *UnixUserSettingsGetOK {
	return &UnixUserSettingsGetOK{}
}

/*
UnixUserSettingsGetOK describes a response with status code 200, with default header values.

OK
*/
type UnixUserSettingsGetOK struct {
	Payload *models.UnixUserSettings
}

// IsSuccess returns true when this unix user settings get o k response has a 2xx status code
func (o *UnixUserSettingsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unix user settings get o k response has a 3xx status code
func (o *UnixUserSettingsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unix user settings get o k response has a 4xx status code
func (o *UnixUserSettingsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unix user settings get o k response has a 5xx status code
func (o *UnixUserSettingsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unix user settings get o k response a status code equal to that given
func (o *UnixUserSettingsGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *UnixUserSettingsGetOK) Error() string {
	return fmt.Sprintf("[GET /name-services/cache/unix-user/settings/{svm.uuid}][%d] unixUserSettingsGetOK  %+v", 200, o.Payload)
}

func (o *UnixUserSettingsGetOK) String() string {
	return fmt.Sprintf("[GET /name-services/cache/unix-user/settings/{svm.uuid}][%d] unixUserSettingsGetOK  %+v", 200, o.Payload)
}

func (o *UnixUserSettingsGetOK) GetPayload() *models.UnixUserSettings {
	return o.Payload
}

func (o *UnixUserSettingsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnixUserSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnixUserSettingsGetDefault creates a UnixUserSettingsGetDefault with default headers values
func NewUnixUserSettingsGetDefault(code int) *UnixUserSettingsGetDefault {
	return &UnixUserSettingsGetDefault{
		_statusCode: code,
	}
}

/*
UnixUserSettingsGetDefault describes a response with status code -1, with default header values.

Error
*/
type UnixUserSettingsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the unix user settings get default response
func (o *UnixUserSettingsGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this unix user settings get default response has a 2xx status code
func (o *UnixUserSettingsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unix user settings get default response has a 3xx status code
func (o *UnixUserSettingsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unix user settings get default response has a 4xx status code
func (o *UnixUserSettingsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unix user settings get default response has a 5xx status code
func (o *UnixUserSettingsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unix user settings get default response a status code equal to that given
func (o *UnixUserSettingsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UnixUserSettingsGetDefault) Error() string {
	return fmt.Sprintf("[GET /name-services/cache/unix-user/settings/{svm.uuid}][%d] unix_user_settings_get default  %+v", o._statusCode, o.Payload)
}

func (o *UnixUserSettingsGetDefault) String() string {
	return fmt.Sprintf("[GET /name-services/cache/unix-user/settings/{svm.uuid}][%d] unix_user_settings_get default  %+v", o._statusCode, o.Payload)
}

func (o *UnixUserSettingsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnixUserSettingsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
