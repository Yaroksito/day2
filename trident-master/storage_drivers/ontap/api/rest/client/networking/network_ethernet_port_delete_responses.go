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

// NetworkEthernetPortDeleteReader is a Reader for the NetworkEthernetPortDelete structure.
type NetworkEthernetPortDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkEthernetPortDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkEthernetPortDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkEthernetPortDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkEthernetPortDeleteOK creates a NetworkEthernetPortDeleteOK with default headers values
func NewNetworkEthernetPortDeleteOK() *NetworkEthernetPortDeleteOK {
	return &NetworkEthernetPortDeleteOK{}
}

/*
NetworkEthernetPortDeleteOK describes a response with status code 200, with default header values.

OK
*/
type NetworkEthernetPortDeleteOK struct {
}

// IsSuccess returns true when this network ethernet port delete o k response has a 2xx status code
func (o *NetworkEthernetPortDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network ethernet port delete o k response has a 3xx status code
func (o *NetworkEthernetPortDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network ethernet port delete o k response has a 4xx status code
func (o *NetworkEthernetPortDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network ethernet port delete o k response has a 5xx status code
func (o *NetworkEthernetPortDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network ethernet port delete o k response a status code equal to that given
func (o *NetworkEthernetPortDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *NetworkEthernetPortDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /network/ethernet/ports/{uuid}][%d] networkEthernetPortDeleteOK ", 200)
}

func (o *NetworkEthernetPortDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /network/ethernet/ports/{uuid}][%d] networkEthernetPortDeleteOK ", 200)
}

func (o *NetworkEthernetPortDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkEthernetPortDeleteDefault creates a NetworkEthernetPortDeleteDefault with default headers values
func NewNetworkEthernetPortDeleteDefault(code int) *NetworkEthernetPortDeleteDefault {
	return &NetworkEthernetPortDeleteDefault{
		_statusCode: code,
	}
}

/*
	NetworkEthernetPortDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1376858 | Port already has an interface bound. |
| 1966189 | Port is the home port or current port of an interface. |
*/
type NetworkEthernetPortDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the network ethernet port delete default response
func (o *NetworkEthernetPortDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this network ethernet port delete default response has a 2xx status code
func (o *NetworkEthernetPortDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ethernet port delete default response has a 3xx status code
func (o *NetworkEthernetPortDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ethernet port delete default response has a 4xx status code
func (o *NetworkEthernetPortDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ethernet port delete default response has a 5xx status code
func (o *NetworkEthernetPortDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ethernet port delete default response a status code equal to that given
func (o *NetworkEthernetPortDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NetworkEthernetPortDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /network/ethernet/ports/{uuid}][%d] network_ethernet_port_delete default  %+v", o._statusCode, o.Payload)
}

func (o *NetworkEthernetPortDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /network/ethernet/ports/{uuid}][%d] network_ethernet_port_delete default  %+v", o._statusCode, o.Payload)
}

func (o *NetworkEthernetPortDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkEthernetPortDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
