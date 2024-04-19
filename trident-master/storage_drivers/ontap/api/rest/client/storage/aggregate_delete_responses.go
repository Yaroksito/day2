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

// AggregateDeleteReader is a Reader for the AggregateDelete structure.
type AggregateDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewAggregateDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAggregateDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregateDeleteAccepted creates a AggregateDeleteAccepted with default headers values
func NewAggregateDeleteAccepted() *AggregateDeleteAccepted {
	return &AggregateDeleteAccepted{}
}

/*
AggregateDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AggregateDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this aggregate delete accepted response has a 2xx status code
func (o *AggregateDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate delete accepted response has a 3xx status code
func (o *AggregateDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate delete accepted response has a 4xx status code
func (o *AggregateDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate delete accepted response has a 5xx status code
func (o *AggregateDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate delete accepted response a status code equal to that given
func (o *AggregateDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *AggregateDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /storage/aggregates/{uuid}][%d] aggregateDeleteAccepted  %+v", 202, o.Payload)
}

func (o *AggregateDeleteAccepted) String() string {
	return fmt.Sprintf("[DELETE /storage/aggregates/{uuid}][%d] aggregateDeleteAccepted  %+v", 202, o.Payload)
}

func (o *AggregateDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *AggregateDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateDeleteDefault creates a AggregateDeleteDefault with default headers values
func NewAggregateDeleteDefault(code int) *AggregateDeleteDefault {
	return &AggregateDeleteDefault{
		_statusCode: code,
	}
}

/*
	AggregateDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 460770 | The aggregate delete job failed to delete the aggregate. |
| 460777 | Failed to get information on the delete job. |
| 786435 | Internal Error. Failed to create a communication handle. |
| 786451 | Failed to delete specified aggregate. |
| 786468 | VLDB is offline. |
| 786472 | Node that hosts the aggregate is offline. |
| 786497 | Cannot delete an aggregate that has volumes. |
| 786771 | Aggregate does not exist. |
| 786867 | Specified aggregate resides on the remote cluster. |
| 786897 | Specified aggregate cannot be deleted as it is a switched-over root aggregate. |
*/
type AggregateDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the aggregate delete default response
func (o *AggregateDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this aggregate delete default response has a 2xx status code
func (o *AggregateDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aggregate delete default response has a 3xx status code
func (o *AggregateDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aggregate delete default response has a 4xx status code
func (o *AggregateDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aggregate delete default response has a 5xx status code
func (o *AggregateDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aggregate delete default response a status code equal to that given
func (o *AggregateDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AggregateDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/aggregates/{uuid}][%d] aggregate_delete default  %+v", o._statusCode, o.Payload)
}

func (o *AggregateDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /storage/aggregates/{uuid}][%d] aggregate_delete default  %+v", o._statusCode, o.Payload)
}

func (o *AggregateDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AggregateDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
