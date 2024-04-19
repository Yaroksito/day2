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

// QuotaReportCollectionGetReader is a Reader for the QuotaReportCollectionGet structure.
type QuotaReportCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuotaReportCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuotaReportCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQuotaReportCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuotaReportCollectionGetOK creates a QuotaReportCollectionGetOK with default headers values
func NewQuotaReportCollectionGetOK() *QuotaReportCollectionGetOK {
	return &QuotaReportCollectionGetOK{}
}

/*
QuotaReportCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type QuotaReportCollectionGetOK struct {
	Payload *models.QuotaReportResponse
}

// IsSuccess returns true when this quota report collection get o k response has a 2xx status code
func (o *QuotaReportCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this quota report collection get o k response has a 3xx status code
func (o *QuotaReportCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this quota report collection get o k response has a 4xx status code
func (o *QuotaReportCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this quota report collection get o k response has a 5xx status code
func (o *QuotaReportCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this quota report collection get o k response a status code equal to that given
func (o *QuotaReportCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *QuotaReportCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/quota/reports][%d] quotaReportCollectionGetOK  %+v", 200, o.Payload)
}

func (o *QuotaReportCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /storage/quota/reports][%d] quotaReportCollectionGetOK  %+v", 200, o.Payload)
}

func (o *QuotaReportCollectionGetOK) GetPayload() *models.QuotaReportResponse {
	return o.Payload
}

func (o *QuotaReportCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuotaReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotaReportCollectionGetDefault creates a QuotaReportCollectionGetDefault with default headers values
func NewQuotaReportCollectionGetDefault(code int) *QuotaReportCollectionGetDefault {
	return &QuotaReportCollectionGetDefault{
		_statusCode: code,
	}
}

/*
QuotaReportCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type QuotaReportCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quota report collection get default response
func (o *QuotaReportCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this quota report collection get default response has a 2xx status code
func (o *QuotaReportCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this quota report collection get default response has a 3xx status code
func (o *QuotaReportCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this quota report collection get default response has a 4xx status code
func (o *QuotaReportCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this quota report collection get default response has a 5xx status code
func (o *QuotaReportCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this quota report collection get default response a status code equal to that given
func (o *QuotaReportCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *QuotaReportCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/quota/reports][%d] quota_report_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *QuotaReportCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /storage/quota/reports][%d] quota_report_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *QuotaReportCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QuotaReportCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
