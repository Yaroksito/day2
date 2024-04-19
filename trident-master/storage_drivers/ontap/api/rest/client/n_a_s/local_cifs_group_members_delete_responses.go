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

// LocalCifsGroupMembersDeleteReader is a Reader for the LocalCifsGroupMembersDelete structure.
type LocalCifsGroupMembersDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalCifsGroupMembersDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalCifsGroupMembersDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalCifsGroupMembersDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalCifsGroupMembersDeleteOK creates a LocalCifsGroupMembersDeleteOK with default headers values
func NewLocalCifsGroupMembersDeleteOK() *LocalCifsGroupMembersDeleteOK {
	return &LocalCifsGroupMembersDeleteOK{}
}

/*
LocalCifsGroupMembersDeleteOK describes a response with status code 200, with default header values.

OK
*/
type LocalCifsGroupMembersDeleteOK struct {
}

// IsSuccess returns true when this local cifs group members delete o k response has a 2xx status code
func (o *LocalCifsGroupMembersDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local cifs group members delete o k response has a 3xx status code
func (o *LocalCifsGroupMembersDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local cifs group members delete o k response has a 4xx status code
func (o *LocalCifsGroupMembersDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this local cifs group members delete o k response has a 5xx status code
func (o *LocalCifsGroupMembersDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this local cifs group members delete o k response a status code equal to that given
func (o *LocalCifsGroupMembersDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *LocalCifsGroupMembersDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/local-groups/{svm.uuid}/{local_cifs_group.sid}/members/{name}][%d] localCifsGroupMembersDeleteOK ", 200)
}

func (o *LocalCifsGroupMembersDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/local-groups/{svm.uuid}/{local_cifs_group.sid}/members/{name}][%d] localCifsGroupMembersDeleteOK ", 200)
}

func (o *LocalCifsGroupMembersDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLocalCifsGroupMembersDeleteDefault creates a LocalCifsGroupMembersDeleteDefault with default headers values
func NewLocalCifsGroupMembersDeleteDefault(code int) *LocalCifsGroupMembersDeleteDefault {
	return &LocalCifsGroupMembersDeleteDefault{
		_statusCode: code,
	}
}

/*
LocalCifsGroupMembersDeleteDefault describes a response with status code -1, with default header values.

Error ONTAP Error Response Codes | Error Code | Description | | ---------- | ----------- | | 655673     | Failed to resolve the member to be deleted from the specified group. | | 655719     | Failed to delete a member from the specified group. The error code returned details the failure along with the reason for the failure. Take corrective actions as per the specified reason. |
*/
type LocalCifsGroupMembersDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the local cifs group members delete default response
func (o *LocalCifsGroupMembersDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this local cifs group members delete default response has a 2xx status code
func (o *LocalCifsGroupMembersDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local cifs group members delete default response has a 3xx status code
func (o *LocalCifsGroupMembersDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local cifs group members delete default response has a 4xx status code
func (o *LocalCifsGroupMembersDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local cifs group members delete default response has a 5xx status code
func (o *LocalCifsGroupMembersDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local cifs group members delete default response a status code equal to that given
func (o *LocalCifsGroupMembersDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *LocalCifsGroupMembersDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/local-groups/{svm.uuid}/{local_cifs_group.sid}/members/{name}][%d] local_cifs_group_members_delete default  %+v", o._statusCode, o.Payload)
}

func (o *LocalCifsGroupMembersDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/local-groups/{svm.uuid}/{local_cifs_group.sid}/members/{name}][%d] local_cifs_group_members_delete default  %+v", o._statusCode, o.Payload)
}

func (o *LocalCifsGroupMembersDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalCifsGroupMembersDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}