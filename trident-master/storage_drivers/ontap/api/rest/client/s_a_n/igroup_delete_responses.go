// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// IgroupDeleteReader is a Reader for the IgroupDelete structure.
type IgroupDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IgroupDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIgroupDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIgroupDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIgroupDeleteOK creates a IgroupDeleteOK with default headers values
func NewIgroupDeleteOK() *IgroupDeleteOK {
	return &IgroupDeleteOK{}
}

/*
IgroupDeleteOK describes a response with status code 200, with default header values.

OK
*/
type IgroupDeleteOK struct {
}

// IsSuccess returns true when this igroup delete o k response has a 2xx status code
func (o *IgroupDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this igroup delete o k response has a 3xx status code
func (o *IgroupDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this igroup delete o k response has a 4xx status code
func (o *IgroupDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this igroup delete o k response has a 5xx status code
func (o *IgroupDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this igroup delete o k response a status code equal to that given
func (o *IgroupDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *IgroupDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/san/igroups/{uuid}][%d] igroupDeleteOK ", 200)
}

func (o *IgroupDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/san/igroups/{uuid}][%d] igroupDeleteOK ", 200)
}

func (o *IgroupDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIgroupDeleteDefault creates a IgroupDeleteDefault with default headers values
func NewIgroupDeleteDefault(code int) *IgroupDeleteDefault {
	return &IgroupDeleteDefault{
		_statusCode: code,
	}
}

/*
	IgroupDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1254213 | The initiator group is mapped to one or more LUNs and `allow_delete_while_mapped` has not been specified. |
| 5374852 | The initiator group does not exist. |
*/
type IgroupDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the igroup delete default response
func (o *IgroupDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this igroup delete default response has a 2xx status code
func (o *IgroupDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this igroup delete default response has a 3xx status code
func (o *IgroupDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this igroup delete default response has a 4xx status code
func (o *IgroupDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this igroup delete default response has a 5xx status code
func (o *IgroupDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this igroup delete default response a status code equal to that given
func (o *IgroupDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IgroupDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/san/igroups/{uuid}][%d] igroup_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IgroupDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /protocols/san/igroups/{uuid}][%d] igroup_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IgroupDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IgroupDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}