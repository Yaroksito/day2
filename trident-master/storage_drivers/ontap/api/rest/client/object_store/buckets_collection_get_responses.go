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

// BucketsCollectionGetReader is a Reader for the BucketsCollectionGet structure.
type BucketsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BucketsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBucketsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBucketsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBucketsCollectionGetOK creates a BucketsCollectionGetOK with default headers values
func NewBucketsCollectionGetOK() *BucketsCollectionGetOK {
	return &BucketsCollectionGetOK{}
}

/*
BucketsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type BucketsCollectionGetOK struct {
	Payload *models.S3BucketSvmResponse
}

// IsSuccess returns true when this buckets collection get o k response has a 2xx status code
func (o *BucketsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this buckets collection get o k response has a 3xx status code
func (o *BucketsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this buckets collection get o k response has a 4xx status code
func (o *BucketsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this buckets collection get o k response has a 5xx status code
func (o *BucketsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this buckets collection get o k response a status code equal to that given
func (o *BucketsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *BucketsCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/buckets][%d] bucketsCollectionGetOK  %+v", 200, o.Payload)
}

func (o *BucketsCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/buckets][%d] bucketsCollectionGetOK  %+v", 200, o.Payload)
}

func (o *BucketsCollectionGetOK) GetPayload() *models.S3BucketSvmResponse {
	return o.Payload
}

func (o *BucketsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3BucketSvmResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBucketsCollectionGetDefault creates a BucketsCollectionGetDefault with default headers values
func NewBucketsCollectionGetDefault(code int) *BucketsCollectionGetDefault {
	return &BucketsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
BucketsCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type BucketsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the buckets collection get default response
func (o *BucketsCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this buckets collection get default response has a 2xx status code
func (o *BucketsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this buckets collection get default response has a 3xx status code
func (o *BucketsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this buckets collection get default response has a 4xx status code
func (o *BucketsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this buckets collection get default response has a 5xx status code
func (o *BucketsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this buckets collection get default response a status code equal to that given
func (o *BucketsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *BucketsCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/buckets][%d] buckets_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *BucketsCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/buckets][%d] buckets_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *BucketsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BucketsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}