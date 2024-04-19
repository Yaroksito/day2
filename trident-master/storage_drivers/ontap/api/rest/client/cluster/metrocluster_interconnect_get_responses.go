// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// MetroclusterInterconnectGetReader is a Reader for the MetroclusterInterconnectGet structure.
type MetroclusterInterconnectGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroclusterInterconnectGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetroclusterInterconnectGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMetroclusterInterconnectGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMetroclusterInterconnectGetOK creates a MetroclusterInterconnectGetOK with default headers values
func NewMetroclusterInterconnectGetOK() *MetroclusterInterconnectGetOK {
	return &MetroclusterInterconnectGetOK{}
}

/*
MetroclusterInterconnectGetOK describes a response with status code 200, with default header values.

OK
*/
type MetroclusterInterconnectGetOK struct {
	Payload *models.MetroclusterInterconnect
}

// IsSuccess returns true when this metrocluster interconnect get o k response has a 2xx status code
func (o *MetroclusterInterconnectGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this metrocluster interconnect get o k response has a 3xx status code
func (o *MetroclusterInterconnectGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this metrocluster interconnect get o k response has a 4xx status code
func (o *MetroclusterInterconnectGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this metrocluster interconnect get o k response has a 5xx status code
func (o *MetroclusterInterconnectGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this metrocluster interconnect get o k response a status code equal to that given
func (o *MetroclusterInterconnectGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *MetroclusterInterconnectGetOK) Error() string {
	return fmt.Sprintf("[GET /cluster/metrocluster/interconnects/{node.uuid}/{partner_type}/{adapter}][%d] metroclusterInterconnectGetOK  %+v", 200, o.Payload)
}

func (o *MetroclusterInterconnectGetOK) String() string {
	return fmt.Sprintf("[GET /cluster/metrocluster/interconnects/{node.uuid}/{partner_type}/{adapter}][%d] metroclusterInterconnectGetOK  %+v", 200, o.Payload)
}

func (o *MetroclusterInterconnectGetOK) GetPayload() *models.MetroclusterInterconnect {
	return o.Payload
}

func (o *MetroclusterInterconnectGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MetroclusterInterconnect)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroclusterInterconnectGetDefault creates a MetroclusterInterconnectGetDefault with default headers values
func NewMetroclusterInterconnectGetDefault(code int) *MetroclusterInterconnectGetDefault {
	return &MetroclusterInterconnectGetDefault{
		_statusCode: code,
	}
}

/*
	MetroclusterInterconnectGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2425734 | An internal error occurred. Wait a few minutes, and try the operation again. For further assistance, contact technical support. |
| 2427132 | MetroCluster is not configured on this cluster. |
*/
type MetroclusterInterconnectGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the metrocluster interconnect get default response
func (o *MetroclusterInterconnectGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this metrocluster interconnect get default response has a 2xx status code
func (o *MetroclusterInterconnectGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this metrocluster interconnect get default response has a 3xx status code
func (o *MetroclusterInterconnectGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this metrocluster interconnect get default response has a 4xx status code
func (o *MetroclusterInterconnectGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this metrocluster interconnect get default response has a 5xx status code
func (o *MetroclusterInterconnectGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this metrocluster interconnect get default response a status code equal to that given
func (o *MetroclusterInterconnectGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *MetroclusterInterconnectGetDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/metrocluster/interconnects/{node.uuid}/{partner_type}/{adapter}][%d] metrocluster_interconnect_get default  %+v", o._statusCode, o.Payload)
}

func (o *MetroclusterInterconnectGetDefault) String() string {
	return fmt.Sprintf("[GET /cluster/metrocluster/interconnects/{node.uuid}/{partner_type}/{adapter}][%d] metrocluster_interconnect_get default  %+v", o._statusCode, o.Payload)
}

func (o *MetroclusterInterconnectGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MetroclusterInterconnectGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}