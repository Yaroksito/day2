// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NetworkIPServicePoliciesGetReader is a Reader for the NetworkIPServicePoliciesGet structure.
type NetworkIPServicePoliciesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPServicePoliciesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkIPServicePoliciesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPServicePoliciesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPServicePoliciesGetOK creates a NetworkIPServicePoliciesGetOK with default headers values
func NewNetworkIPServicePoliciesGetOK() *NetworkIPServicePoliciesGetOK {
	return &NetworkIPServicePoliciesGetOK{}
}

/*
NetworkIPServicePoliciesGetOK describes a response with status code 200, with default header values.

OK
*/
type NetworkIPServicePoliciesGetOK struct {
	Payload *models.IPServicePolicyResponse
}

// IsSuccess returns true when this network Ip service policies get o k response has a 2xx status code
func (o *NetworkIPServicePoliciesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network Ip service policies get o k response has a 3xx status code
func (o *NetworkIPServicePoliciesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network Ip service policies get o k response has a 4xx status code
func (o *NetworkIPServicePoliciesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network Ip service policies get o k response has a 5xx status code
func (o *NetworkIPServicePoliciesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network Ip service policies get o k response a status code equal to that given
func (o *NetworkIPServicePoliciesGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *NetworkIPServicePoliciesGetOK) Error() string {
	return fmt.Sprintf("[GET /network/ip/service-policies][%d] networkIpServicePoliciesGetOK  %+v", 200, o.Payload)
}

func (o *NetworkIPServicePoliciesGetOK) String() string {
	return fmt.Sprintf("[GET /network/ip/service-policies][%d] networkIpServicePoliciesGetOK  %+v", 200, o.Payload)
}

func (o *NetworkIPServicePoliciesGetOK) GetPayload() *models.IPServicePolicyResponse {
	return o.Payload
}

func (o *NetworkIPServicePoliciesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPServicePolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkIPServicePoliciesGetDefault creates a NetworkIPServicePoliciesGetDefault with default headers values
func NewNetworkIPServicePoliciesGetDefault(code int) *NetworkIPServicePoliciesGetDefault {
	return &NetworkIPServicePoliciesGetDefault{
		_statusCode: code,
	}
}

/*
NetworkIPServicePoliciesGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetworkIPServicePoliciesGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the network ip service policies get default response
func (o *NetworkIPServicePoliciesGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this network ip service policies get default response has a 2xx status code
func (o *NetworkIPServicePoliciesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ip service policies get default response has a 3xx status code
func (o *NetworkIPServicePoliciesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ip service policies get default response has a 4xx status code
func (o *NetworkIPServicePoliciesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ip service policies get default response has a 5xx status code
func (o *NetworkIPServicePoliciesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ip service policies get default response a status code equal to that given
func (o *NetworkIPServicePoliciesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NetworkIPServicePoliciesGetDefault) Error() string {
	return fmt.Sprintf("[GET /network/ip/service-policies][%d] network_ip_service_policies_get default  %+v", o._statusCode, o.Payload)
}

func (o *NetworkIPServicePoliciesGetDefault) String() string {
	return fmt.Sprintf("[GET /network/ip/service-policies][%d] network_ip_service_policies_get default  %+v", o._statusCode, o.Payload)
}

func (o *NetworkIPServicePoliciesGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkIPServicePoliciesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}