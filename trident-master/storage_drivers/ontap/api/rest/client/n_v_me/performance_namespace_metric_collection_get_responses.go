// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// PerformanceNamespaceMetricCollectionGetReader is a Reader for the PerformanceNamespaceMetricCollectionGet structure.
type PerformanceNamespaceMetricCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformanceNamespaceMetricCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformanceNamespaceMetricCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformanceNamespaceMetricCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformanceNamespaceMetricCollectionGetOK creates a PerformanceNamespaceMetricCollectionGetOK with default headers values
func NewPerformanceNamespaceMetricCollectionGetOK() *PerformanceNamespaceMetricCollectionGetOK {
	return &PerformanceNamespaceMetricCollectionGetOK{}
}

/*
PerformanceNamespaceMetricCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type PerformanceNamespaceMetricCollectionGetOK struct {
	Payload *models.PerformanceNamespaceMetricResponse
}

// IsSuccess returns true when this performance namespace metric collection get o k response has a 2xx status code
func (o *PerformanceNamespaceMetricCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance namespace metric collection get o k response has a 3xx status code
func (o *PerformanceNamespaceMetricCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance namespace metric collection get o k response has a 4xx status code
func (o *PerformanceNamespaceMetricCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance namespace metric collection get o k response has a 5xx status code
func (o *PerformanceNamespaceMetricCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance namespace metric collection get o k response a status code equal to that given
func (o *PerformanceNamespaceMetricCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *PerformanceNamespaceMetricCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/namespaces/{uuid}/metrics][%d] performanceNamespaceMetricCollectionGetOK  %+v", 200, o.Payload)
}

func (o *PerformanceNamespaceMetricCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /storage/namespaces/{uuid}/metrics][%d] performanceNamespaceMetricCollectionGetOK  %+v", 200, o.Payload)
}

func (o *PerformanceNamespaceMetricCollectionGetOK) GetPayload() *models.PerformanceNamespaceMetricResponse {
	return o.Payload
}

func (o *PerformanceNamespaceMetricCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceNamespaceMetricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformanceNamespaceMetricCollectionGetDefault creates a PerformanceNamespaceMetricCollectionGetDefault with default headers values
func NewPerformanceNamespaceMetricCollectionGetDefault(code int) *PerformanceNamespaceMetricCollectionGetDefault {
	return &PerformanceNamespaceMetricCollectionGetDefault{
		_statusCode: code,
	}
}

/*
PerformanceNamespaceMetricCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type PerformanceNamespaceMetricCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the performance namespace metric collection get default response
func (o *PerformanceNamespaceMetricCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this performance namespace metric collection get default response has a 2xx status code
func (o *PerformanceNamespaceMetricCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performance namespace metric collection get default response has a 3xx status code
func (o *PerformanceNamespaceMetricCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performance namespace metric collection get default response has a 4xx status code
func (o *PerformanceNamespaceMetricCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performance namespace metric collection get default response has a 5xx status code
func (o *PerformanceNamespaceMetricCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performance namespace metric collection get default response a status code equal to that given
func (o *PerformanceNamespaceMetricCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PerformanceNamespaceMetricCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/namespaces/{uuid}/metrics][%d] performance_namespace_metric_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *PerformanceNamespaceMetricCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /storage/namespaces/{uuid}/metrics][%d] performance_namespace_metric_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *PerformanceNamespaceMetricCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PerformanceNamespaceMetricCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
