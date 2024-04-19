// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// CloudStoreCollectionGetReader is a Reader for the CloudStoreCollectionGet structure.
type CloudStoreCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudStoreCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudStoreCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCloudStoreCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudStoreCollectionGetOK creates a CloudStoreCollectionGetOK with default headers values
func NewCloudStoreCollectionGetOK() *CloudStoreCollectionGetOK {
	return &CloudStoreCollectionGetOK{}
}

/*
CloudStoreCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type CloudStoreCollectionGetOK struct {
	Payload *models.CloudStoreResponse
}

// IsSuccess returns true when this cloud store collection get o k response has a 2xx status code
func (o *CloudStoreCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cloud store collection get o k response has a 3xx status code
func (o *CloudStoreCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cloud store collection get o k response has a 4xx status code
func (o *CloudStoreCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cloud store collection get o k response has a 5xx status code
func (o *CloudStoreCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cloud store collection get o k response a status code equal to that given
func (o *CloudStoreCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *CloudStoreCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/aggregates/{aggregate.uuid}/cloud-stores][%d] cloudStoreCollectionGetOK  %+v", 200, o.Payload)
}

func (o *CloudStoreCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /storage/aggregates/{aggregate.uuid}/cloud-stores][%d] cloudStoreCollectionGetOK  %+v", 200, o.Payload)
}

func (o *CloudStoreCollectionGetOK) GetPayload() *models.CloudStoreResponse {
	return o.Payload
}

func (o *CloudStoreCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudStoreResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudStoreCollectionGetDefault creates a CloudStoreCollectionGetDefault with default headers values
func NewCloudStoreCollectionGetDefault(code int) *CloudStoreCollectionGetDefault {
	return &CloudStoreCollectionGetDefault{
		_statusCode: code,
	}
}

/*
CloudStoreCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type CloudStoreCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cloud store collection get default response
func (o *CloudStoreCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cloud store collection get default response has a 2xx status code
func (o *CloudStoreCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cloud store collection get default response has a 3xx status code
func (o *CloudStoreCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cloud store collection get default response has a 4xx status code
func (o *CloudStoreCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cloud store collection get default response has a 5xx status code
func (o *CloudStoreCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cloud store collection get default response a status code equal to that given
func (o *CloudStoreCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CloudStoreCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/aggregates/{aggregate.uuid}/cloud-stores][%d] cloud_store_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *CloudStoreCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /storage/aggregates/{aggregate.uuid}/cloud-stores][%d] cloud_store_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *CloudStoreCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CloudStoreCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}