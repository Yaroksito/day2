// Code generated by go-swagger; DO NOT EDIT.

package object_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// S3GroupCollectionGetReader is a Reader for the S3GroupCollectionGet structure.
type S3GroupCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3GroupCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3GroupCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3GroupCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3GroupCollectionGetOK creates a S3GroupCollectionGetOK with default headers values
func NewS3GroupCollectionGetOK() *S3GroupCollectionGetOK {
	return &S3GroupCollectionGetOK{}
}

/*
S3GroupCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type S3GroupCollectionGetOK struct {
	Payload *models.S3GroupResponse
}

// IsSuccess returns true when this s3 group collection get o k response has a 2xx status code
func (o *S3GroupCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 group collection get o k response has a 3xx status code
func (o *S3GroupCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 group collection get o k response has a 4xx status code
func (o *S3GroupCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 group collection get o k response has a 5xx status code
func (o *S3GroupCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 group collection get o k response a status code equal to that given
func (o *S3GroupCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *S3GroupCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/groups][%d] s3GroupCollectionGetOK  %+v", 200, o.Payload)
}

func (o *S3GroupCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/groups][%d] s3GroupCollectionGetOK  %+v", 200, o.Payload)
}

func (o *S3GroupCollectionGetOK) GetPayload() *models.S3GroupResponse {
	return o.Payload
}

func (o *S3GroupCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3GroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3GroupCollectionGetDefault creates a S3GroupCollectionGetDefault with default headers values
func NewS3GroupCollectionGetDefault(code int) *S3GroupCollectionGetDefault {
	return &S3GroupCollectionGetDefault{
		_statusCode: code,
	}
}

/*
S3GroupCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type S3GroupCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the s3 group collection get default response
func (o *S3GroupCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this s3 group collection get default response has a 2xx status code
func (o *S3GroupCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 group collection get default response has a 3xx status code
func (o *S3GroupCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 group collection get default response has a 4xx status code
func (o *S3GroupCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 group collection get default response has a 5xx status code
func (o *S3GroupCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 group collection get default response a status code equal to that given
func (o *S3GroupCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *S3GroupCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/groups][%d] s3_group_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *S3GroupCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/groups][%d] s3_group_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *S3GroupCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3GroupCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
