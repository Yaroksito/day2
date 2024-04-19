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

// ClusterGetReader is a Reader for the ClusterGet structure.
type ClusterGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterGetOK creates a ClusterGetOK with default headers values
func NewClusterGetOK() *ClusterGetOK {
	return &ClusterGetOK{}
}

/*
ClusterGetOK describes a response with status code 200, with default header values.

OK
*/
type ClusterGetOK struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this cluster get o k response has a 2xx status code
func (o *ClusterGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster get o k response has a 3xx status code
func (o *ClusterGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster get o k response has a 4xx status code
func (o *ClusterGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster get o k response has a 5xx status code
func (o *ClusterGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster get o k response a status code equal to that given
func (o *ClusterGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClusterGetOK) Error() string {
	return fmt.Sprintf("[GET /cluster][%d] clusterGetOK  %+v", 200, o.Payload)
}

func (o *ClusterGetOK) String() string {
	return fmt.Sprintf("[GET /cluster][%d] clusterGetOK  %+v", 200, o.Payload)
}

func (o *ClusterGetOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *ClusterGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterGetDefault creates a ClusterGetDefault with default headers values
func NewClusterGetDefault(code int) *ClusterGetDefault {
	return &ClusterGetDefault{
		_statusCode: code,
	}
}

/*
ClusterGetDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cluster get default response
func (o *ClusterGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster get default response has a 2xx status code
func (o *ClusterGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster get default response has a 3xx status code
func (o *ClusterGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster get default response has a 4xx status code
func (o *ClusterGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster get default response has a 5xx status code
func (o *ClusterGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster get default response a status code equal to that given
func (o *ClusterGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterGetDefault) Error() string {
	return fmt.Sprintf("[GET /cluster][%d] cluster_get default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterGetDefault) String() string {
	return fmt.Sprintf("[GET /cluster][%d] cluster_get default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
