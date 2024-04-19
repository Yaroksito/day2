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

// PortsetCollectionGetReader is a Reader for the PortsetCollectionGet structure.
type PortsetCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PortsetCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPortsetCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPortsetCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPortsetCollectionGetOK creates a PortsetCollectionGetOK with default headers values
func NewPortsetCollectionGetOK() *PortsetCollectionGetOK {
	return &PortsetCollectionGetOK{}
}

/*
PortsetCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type PortsetCollectionGetOK struct {
	Payload *models.PortsetResponse
}

// IsSuccess returns true when this portset collection get o k response has a 2xx status code
func (o *PortsetCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this portset collection get o k response has a 3xx status code
func (o *PortsetCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this portset collection get o k response has a 4xx status code
func (o *PortsetCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this portset collection get o k response has a 5xx status code
func (o *PortsetCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this portset collection get o k response a status code equal to that given
func (o *PortsetCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *PortsetCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/san/portsets][%d] portsetCollectionGetOK  %+v", 200, o.Payload)
}

func (o *PortsetCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/san/portsets][%d] portsetCollectionGetOK  %+v", 200, o.Payload)
}

func (o *PortsetCollectionGetOK) GetPayload() *models.PortsetResponse {
	return o.Payload
}

func (o *PortsetCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortsetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPortsetCollectionGetDefault creates a PortsetCollectionGetDefault with default headers values
func NewPortsetCollectionGetDefault(code int) *PortsetCollectionGetDefault {
	return &PortsetCollectionGetDefault{
		_statusCode: code,
	}
}

/*
PortsetCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type PortsetCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the portset collection get default response
func (o *PortsetCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this portset collection get default response has a 2xx status code
func (o *PortsetCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this portset collection get default response has a 3xx status code
func (o *PortsetCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this portset collection get default response has a 4xx status code
func (o *PortsetCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this portset collection get default response has a 5xx status code
func (o *PortsetCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this portset collection get default response a status code equal to that given
func (o *PortsetCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PortsetCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/san/portsets][%d] portset_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *PortsetCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/san/portsets][%d] portset_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *PortsetCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PortsetCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
