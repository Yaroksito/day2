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

// IgroupModifyReader is a Reader for the IgroupModify structure.
type IgroupModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IgroupModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIgroupModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIgroupModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIgroupModifyOK creates a IgroupModifyOK with default headers values
func NewIgroupModifyOK() *IgroupModifyOK {
	return &IgroupModifyOK{}
}

/*
IgroupModifyOK describes a response with status code 200, with default header values.

OK
*/
type IgroupModifyOK struct {
}

// IsSuccess returns true when this igroup modify o k response has a 2xx status code
func (o *IgroupModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this igroup modify o k response has a 3xx status code
func (o *IgroupModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this igroup modify o k response has a 4xx status code
func (o *IgroupModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this igroup modify o k response has a 5xx status code
func (o *IgroupModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this igroup modify o k response a status code equal to that given
func (o *IgroupModifyOK) IsCode(code int) bool {
	return code == 200
}

func (o *IgroupModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/san/igroups/{uuid}][%d] igroupModifyOK ", 200)
}

func (o *IgroupModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/san/igroups/{uuid}][%d] igroupModifyOK ", 200)
}

func (o *IgroupModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIgroupModifyDefault creates a IgroupModifyDefault with default headers values
func NewIgroupModifyDefault(code int) *IgroupModifyDefault {
	return &IgroupModifyDefault{
		_statusCode: code,
	}
}

/*
	IgroupModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1254264 | An attempt was made to bind a portset to an initiator group that is already bound to a portset. |
| 5373958 | An invalid initiator group name was supplied for a rename operation. |
| 5374023 | A rename operation failed because an initiator group with the same name already exists. |
| 5374027 | An attempt was made to bind a portset with no member network interfaces to the initiator group. |
| 5374028 | An attempt was made to bind a portset with an incompatible protocol to the initiator group. |
| 5374733 | An initiator is already in another initiator group with a conflicting operating system type. |
| 5374745 | An attempt was made to add an initiator group as a child to itself. |
| 5374746 | The cluster is currently running in a mixed version and nested initiator groups cannot be created until the effective cluster version reaches 9.9.1. |
| 5374747 | The cluster is currently running in a mixed version and initiator group comments cannot be created until the effective cluster version reaches 9.9.1. |
| 5374852 | The initiator group does not exist. |
| 5374868 | The initiator group was partially modified before an error was encountered while renaming the initiator group. |
*/
type IgroupModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the igroup modify default response
func (o *IgroupModifyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this igroup modify default response has a 2xx status code
func (o *IgroupModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this igroup modify default response has a 3xx status code
func (o *IgroupModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this igroup modify default response has a 4xx status code
func (o *IgroupModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this igroup modify default response has a 5xx status code
func (o *IgroupModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this igroup modify default response a status code equal to that given
func (o *IgroupModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IgroupModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/san/igroups/{uuid}][%d] igroup_modify default  %+v", o._statusCode, o.Payload)
}

func (o *IgroupModifyDefault) String() string {
	return fmt.Sprintf("[PATCH /protocols/san/igroups/{uuid}][%d] igroup_modify default  %+v", o._statusCode, o.Payload)
}

func (o *IgroupModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IgroupModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}