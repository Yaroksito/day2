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

// LocalCifsGroupGetReader is a Reader for the LocalCifsGroupGet structure.
type LocalCifsGroupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalCifsGroupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalCifsGroupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalCifsGroupGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalCifsGroupGetOK creates a LocalCifsGroupGetOK with default headers values
func NewLocalCifsGroupGetOK() *LocalCifsGroupGetOK {
	return &LocalCifsGroupGetOK{}
}

/*
LocalCifsGroupGetOK describes a response with status code 200, with default header values.

OK
*/
type LocalCifsGroupGetOK struct {
	Payload *models.LocalCifsGroup
}

// IsSuccess returns true when this local cifs group get o k response has a 2xx status code
func (o *LocalCifsGroupGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local cifs group get o k response has a 3xx status code
func (o *LocalCifsGroupGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local cifs group get o k response has a 4xx status code
func (o *LocalCifsGroupGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this local cifs group get o k response has a 5xx status code
func (o *LocalCifsGroupGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this local cifs group get o k response a status code equal to that given
func (o *LocalCifsGroupGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *LocalCifsGroupGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/cifs/local-groups/{svm.uuid}/{sid}][%d] localCifsGroupGetOK  %+v", 200, o.Payload)
}

func (o *LocalCifsGroupGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/cifs/local-groups/{svm.uuid}/{sid}][%d] localCifsGroupGetOK  %+v", 200, o.Payload)
}

func (o *LocalCifsGroupGetOK) GetPayload() *models.LocalCifsGroup {
	return o.Payload
}

func (o *LocalCifsGroupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocalCifsGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocalCifsGroupGetDefault creates a LocalCifsGroupGetDefault with default headers values
func NewLocalCifsGroupGetDefault(code int) *LocalCifsGroupGetDefault {
	return &LocalCifsGroupGetDefault{
		_statusCode: code,
	}
}

/*
LocalCifsGroupGetDefault describes a response with status code -1, with default header values.

Error
*/
type LocalCifsGroupGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the local cifs group get default response
func (o *LocalCifsGroupGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this local cifs group get default response has a 2xx status code
func (o *LocalCifsGroupGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local cifs group get default response has a 3xx status code
func (o *LocalCifsGroupGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local cifs group get default response has a 4xx status code
func (o *LocalCifsGroupGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local cifs group get default response has a 5xx status code
func (o *LocalCifsGroupGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local cifs group get default response a status code equal to that given
func (o *LocalCifsGroupGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *LocalCifsGroupGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/cifs/local-groups/{svm.uuid}/{sid}][%d] local_cifs_group_get default  %+v", o._statusCode, o.Payload)
}

func (o *LocalCifsGroupGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/cifs/local-groups/{svm.uuid}/{sid}][%d] local_cifs_group_get default  %+v", o._statusCode, o.Payload)
}

func (o *LocalCifsGroupGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalCifsGroupGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
