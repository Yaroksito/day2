// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// LocalCifsGroupModifyReader is a Reader for the LocalCifsGroupModify structure.
type LocalCifsGroupModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalCifsGroupModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalCifsGroupModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalCifsGroupModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalCifsGroupModifyOK creates a LocalCifsGroupModifyOK with default headers values
func NewLocalCifsGroupModifyOK() *LocalCifsGroupModifyOK {
	return &LocalCifsGroupModifyOK{}
}

/*
LocalCifsGroupModifyOK describes a response with status code 200, with default header values.

OK
*/
type LocalCifsGroupModifyOK struct {
}

// IsSuccess returns true when this local cifs group modify o k response has a 2xx status code
func (o *LocalCifsGroupModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local cifs group modify o k response has a 3xx status code
func (o *LocalCifsGroupModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local cifs group modify o k response has a 4xx status code
func (o *LocalCifsGroupModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this local cifs group modify o k response has a 5xx status code
func (o *LocalCifsGroupModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this local cifs group modify o k response a status code equal to that given
func (o *LocalCifsGroupModifyOK) IsCode(code int) bool {
	return code == 200
}

func (o *LocalCifsGroupModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/local-groups/{svm.uuid}/{sid}][%d] localCifsGroupModifyOK ", 200)
}

func (o *LocalCifsGroupModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/local-groups/{svm.uuid}/{sid}][%d] localCifsGroupModifyOK ", 200)
}

func (o *LocalCifsGroupModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLocalCifsGroupModifyDefault creates a LocalCifsGroupModifyDefault with default headers values
func NewLocalCifsGroupModifyDefault(code int) *LocalCifsGroupModifyDefault {
	return &LocalCifsGroupModifyDefault{
		_statusCode: code,
	}
}

/*
	LocalCifsGroupModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655661     | The group name and description should not exceed 256 characters. |
| 655668     | The specified group name contains illegal characters. |
| 655675     | The local domain name specified in the group name does not exist. |
| 655682     | The group name cannot be blank. |
| 655712     | To rename an existing group, the local domain specified in name must match the local domain of the group to be renamed. |
| 655713     | Failed to rename a group. The error code returned details the failure along with the reason for the failure. Take corrective actions as per the specified reason. |
*/
type LocalCifsGroupModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the local cifs group modify default response
func (o *LocalCifsGroupModifyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this local cifs group modify default response has a 2xx status code
func (o *LocalCifsGroupModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local cifs group modify default response has a 3xx status code
func (o *LocalCifsGroupModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local cifs group modify default response has a 4xx status code
func (o *LocalCifsGroupModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local cifs group modify default response has a 5xx status code
func (o *LocalCifsGroupModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local cifs group modify default response a status code equal to that given
func (o *LocalCifsGroupModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *LocalCifsGroupModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/local-groups/{svm.uuid}/{sid}][%d] local_cifs_group_modify default  %+v", o._statusCode, o.Payload)
}

func (o *LocalCifsGroupModifyDefault) String() string {
	return fmt.Sprintf("[PATCH /protocols/cifs/local-groups/{svm.uuid}/{sid}][%d] local_cifs_group_modify default  %+v", o._statusCode, o.Payload)
}

func (o *LocalCifsGroupModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalCifsGroupModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
