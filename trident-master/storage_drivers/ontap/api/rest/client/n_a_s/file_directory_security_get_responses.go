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

// FileDirectorySecurityGetReader is a Reader for the FileDirectorySecurityGet structure.
type FileDirectorySecurityGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileDirectorySecurityGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFileDirectorySecurityGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileDirectorySecurityGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileDirectorySecurityGetOK creates a FileDirectorySecurityGetOK with default headers values
func NewFileDirectorySecurityGetOK() *FileDirectorySecurityGetOK {
	return &FileDirectorySecurityGetOK{}
}

/*
FileDirectorySecurityGetOK describes a response with status code 200, with default header values.

OK
*/
type FileDirectorySecurityGetOK struct {
	Payload *models.FileDirectorySecurity
}

// IsSuccess returns true when this file directory security get o k response has a 2xx status code
func (o *FileDirectorySecurityGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file directory security get o k response has a 3xx status code
func (o *FileDirectorySecurityGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file directory security get o k response has a 4xx status code
func (o *FileDirectorySecurityGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this file directory security get o k response has a 5xx status code
func (o *FileDirectorySecurityGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this file directory security get o k response a status code equal to that given
func (o *FileDirectorySecurityGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *FileDirectorySecurityGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/file-security/permissions/{svm.uuid}/{path}][%d] fileDirectorySecurityGetOK  %+v", 200, o.Payload)
}

func (o *FileDirectorySecurityGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/file-security/permissions/{svm.uuid}/{path}][%d] fileDirectorySecurityGetOK  %+v", 200, o.Payload)
}

func (o *FileDirectorySecurityGetOK) GetPayload() *models.FileDirectorySecurity {
	return o.Payload
}

func (o *FileDirectorySecurityGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileDirectorySecurity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileDirectorySecurityGetDefault creates a FileDirectorySecurityGetDefault with default headers values
func NewFileDirectorySecurityGetDefault(code int) *FileDirectorySecurityGetDefault {
	return &FileDirectorySecurityGetDefault{
		_statusCode: code,
	}
}

/*
FileDirectorySecurityGetDefault describes a response with status code -1, with default header values.

Error
*/
type FileDirectorySecurityGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the file directory security get default response
func (o *FileDirectorySecurityGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this file directory security get default response has a 2xx status code
func (o *FileDirectorySecurityGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this file directory security get default response has a 3xx status code
func (o *FileDirectorySecurityGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this file directory security get default response has a 4xx status code
func (o *FileDirectorySecurityGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this file directory security get default response has a 5xx status code
func (o *FileDirectorySecurityGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this file directory security get default response a status code equal to that given
func (o *FileDirectorySecurityGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FileDirectorySecurityGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/file-security/permissions/{svm.uuid}/{path}][%d] file_directory_security_get default  %+v", o._statusCode, o.Payload)
}

func (o *FileDirectorySecurityGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/file-security/permissions/{svm.uuid}/{path}][%d] file_directory_security_get default  %+v", o._statusCode, o.Payload)
}

func (o *FileDirectorySecurityGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileDirectorySecurityGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
