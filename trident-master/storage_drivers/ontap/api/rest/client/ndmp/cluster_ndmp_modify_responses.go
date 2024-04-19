// Code generated by go-swagger; DO NOT EDIT.

package ndmp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ClusterNdmpModifyReader is a Reader for the ClusterNdmpModify structure.
type ClusterNdmpModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNdmpModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterNdmpModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNdmpModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNdmpModifyOK creates a ClusterNdmpModifyOK with default headers values
func NewClusterNdmpModifyOK() *ClusterNdmpModifyOK {
	return &ClusterNdmpModifyOK{}
}

/*
ClusterNdmpModifyOK describes a response with status code 200, with default header values.

OK
*/
type ClusterNdmpModifyOK struct {
	Payload *models.ClusterNdmpProperties
}

// IsSuccess returns true when this cluster ndmp modify o k response has a 2xx status code
func (o *ClusterNdmpModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster ndmp modify o k response has a 3xx status code
func (o *ClusterNdmpModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster ndmp modify o k response has a 4xx status code
func (o *ClusterNdmpModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster ndmp modify o k response has a 5xx status code
func (o *ClusterNdmpModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster ndmp modify o k response a status code equal to that given
func (o *ClusterNdmpModifyOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClusterNdmpModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/ndmp][%d] clusterNdmpModifyOK  %+v", 200, o.Payload)
}

func (o *ClusterNdmpModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/ndmp][%d] clusterNdmpModifyOK  %+v", 200, o.Payload)
}

func (o *ClusterNdmpModifyOK) GetPayload() *models.ClusterNdmpProperties {
	return o.Payload
}

func (o *ClusterNdmpModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNdmpProperties)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNdmpModifyDefault creates a ClusterNdmpModifyDefault with default headers values
func NewClusterNdmpModifyDefault(code int) *ClusterNdmpModifyDefault {
	return &ClusterNdmpModifyDefault{
		_statusCode: code,
	}
}

/*
ClusterNdmpModifyDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterNdmpModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cluster ndmp modify default response
func (o *ClusterNdmpModifyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster ndmp modify default response has a 2xx status code
func (o *ClusterNdmpModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster ndmp modify default response has a 3xx status code
func (o *ClusterNdmpModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster ndmp modify default response has a 4xx status code
func (o *ClusterNdmpModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster ndmp modify default response has a 5xx status code
func (o *ClusterNdmpModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster ndmp modify default response a status code equal to that given
func (o *ClusterNdmpModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterNdmpModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/ndmp][%d] cluster_ndmp_modify default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterNdmpModifyDefault) String() string {
	return fmt.Sprintf("[PATCH /protocols/ndmp][%d] cluster_ndmp_modify default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterNdmpModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNdmpModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}