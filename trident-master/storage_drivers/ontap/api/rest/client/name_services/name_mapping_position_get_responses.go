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

// NameMappingPositionGetReader is a Reader for the NameMappingPositionGet structure.
type NameMappingPositionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NameMappingPositionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNameMappingPositionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNameMappingPositionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNameMappingPositionGetOK creates a NameMappingPositionGetOK with default headers values
func NewNameMappingPositionGetOK() *NameMappingPositionGetOK {
	return &NameMappingPositionGetOK{}
}

/*
NameMappingPositionGetOK describes a response with status code 200, with default header values.

OK
*/
type NameMappingPositionGetOK struct {
	Payload *models.NameMapping
}

// IsSuccess returns true when this name mapping position get o k response has a 2xx status code
func (o *NameMappingPositionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this name mapping position get o k response has a 3xx status code
func (o *NameMappingPositionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this name mapping position get o k response has a 4xx status code
func (o *NameMappingPositionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this name mapping position get o k response has a 5xx status code
func (o *NameMappingPositionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this name mapping position get o k response a status code equal to that given
func (o *NameMappingPositionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *NameMappingPositionGetOK) Error() string {
	return fmt.Sprintf("[GET /name-services/name-mappings/{svm.uuid}/{direction}/{index}][%d] nameMappingPositionGetOK  %+v", 200, o.Payload)
}

func (o *NameMappingPositionGetOK) String() string {
	return fmt.Sprintf("[GET /name-services/name-mappings/{svm.uuid}/{direction}/{index}][%d] nameMappingPositionGetOK  %+v", 200, o.Payload)
}

func (o *NameMappingPositionGetOK) GetPayload() *models.NameMapping {
	return o.Payload
}

func (o *NameMappingPositionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NameMapping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNameMappingPositionGetDefault creates a NameMappingPositionGetDefault with default headers values
func NewNameMappingPositionGetDefault(code int) *NameMappingPositionGetDefault {
	return &NameMappingPositionGetDefault{
		_statusCode: code,
	}
}

/*
NameMappingPositionGetDefault describes a response with status code -1, with default header values.

Error
*/
type NameMappingPositionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the name mapping position get default response
func (o *NameMappingPositionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this name mapping position get default response has a 2xx status code
func (o *NameMappingPositionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this name mapping position get default response has a 3xx status code
func (o *NameMappingPositionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this name mapping position get default response has a 4xx status code
func (o *NameMappingPositionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this name mapping position get default response has a 5xx status code
func (o *NameMappingPositionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this name mapping position get default response a status code equal to that given
func (o *NameMappingPositionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NameMappingPositionGetDefault) Error() string {
	return fmt.Sprintf("[GET /name-services/name-mappings/{svm.uuid}/{direction}/{index}][%d] name_mapping_position_get default  %+v", o._statusCode, o.Payload)
}

func (o *NameMappingPositionGetDefault) String() string {
	return fmt.Sprintf("[GET /name-services/name-mappings/{svm.uuid}/{direction}/{index}][%d] name_mapping_position_get default  %+v", o._statusCode, o.Payload)
}

func (o *NameMappingPositionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NameMappingPositionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}