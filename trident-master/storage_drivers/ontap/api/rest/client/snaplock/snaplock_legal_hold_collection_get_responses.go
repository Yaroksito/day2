// Code generated by go-swagger; DO NOT EDIT.

package snaplock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnaplockLegalHoldCollectionGetReader is a Reader for the SnaplockLegalHoldCollectionGet structure.
type SnaplockLegalHoldCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLegalHoldCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockLegalHoldCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLegalHoldCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLegalHoldCollectionGetOK creates a SnaplockLegalHoldCollectionGetOK with default headers values
func NewSnaplockLegalHoldCollectionGetOK() *SnaplockLegalHoldCollectionGetOK {
	return &SnaplockLegalHoldCollectionGetOK{}
}

/*
SnaplockLegalHoldCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockLegalHoldCollectionGetOK struct {
	Payload *models.SnaplockLitigationResponse
}

// IsSuccess returns true when this snaplock legal hold collection get o k response has a 2xx status code
func (o *SnaplockLegalHoldCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock legal hold collection get o k response has a 3xx status code
func (o *SnaplockLegalHoldCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock legal hold collection get o k response has a 4xx status code
func (o *SnaplockLegalHoldCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock legal hold collection get o k response has a 5xx status code
func (o *SnaplockLegalHoldCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock legal hold collection get o k response a status code equal to that given
func (o *SnaplockLegalHoldCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *SnaplockLegalHoldCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/litigations][%d] snaplockLegalHoldCollectionGetOK  %+v", 200, o.Payload)
}

func (o *SnaplockLegalHoldCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /storage/snaplock/litigations][%d] snaplockLegalHoldCollectionGetOK  %+v", 200, o.Payload)
}

func (o *SnaplockLegalHoldCollectionGetOK) GetPayload() *models.SnaplockLitigationResponse {
	return o.Payload
}

func (o *SnaplockLegalHoldCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockLitigationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLegalHoldCollectionGetDefault creates a SnaplockLegalHoldCollectionGetDefault with default headers values
func NewSnaplockLegalHoldCollectionGetDefault(code int) *SnaplockLegalHoldCollectionGetDefault {
	return &SnaplockLegalHoldCollectionGetDefault{
		_statusCode: code,
	}
}

/*
	SnaplockLegalHoldCollectionGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 14090346    | Internal Error. Wait a few minutes, then try the command again  |
*/
type SnaplockLegalHoldCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock legal hold collection get default response
func (o *SnaplockLegalHoldCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this snaplock legal hold collection get default response has a 2xx status code
func (o *SnaplockLegalHoldCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock legal hold collection get default response has a 3xx status code
func (o *SnaplockLegalHoldCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock legal hold collection get default response has a 4xx status code
func (o *SnaplockLegalHoldCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock legal hold collection get default response has a 5xx status code
func (o *SnaplockLegalHoldCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock legal hold collection get default response a status code equal to that given
func (o *SnaplockLegalHoldCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SnaplockLegalHoldCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/litigations][%d] snaplock_legal_hold_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *SnaplockLegalHoldCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /storage/snaplock/litigations][%d] snaplock_legal_hold_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *SnaplockLegalHoldCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLegalHoldCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
