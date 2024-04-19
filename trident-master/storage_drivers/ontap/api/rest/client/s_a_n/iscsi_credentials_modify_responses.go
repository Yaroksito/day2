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

// IscsiCredentialsModifyReader is a Reader for the IscsiCredentialsModify structure.
type IscsiCredentialsModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IscsiCredentialsModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIscsiCredentialsModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIscsiCredentialsModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIscsiCredentialsModifyOK creates a IscsiCredentialsModifyOK with default headers values
func NewIscsiCredentialsModifyOK() *IscsiCredentialsModifyOK {
	return &IscsiCredentialsModifyOK{}
}

/*
IscsiCredentialsModifyOK describes a response with status code 200, with default header values.

OK
*/
type IscsiCredentialsModifyOK struct {
}

// IsSuccess returns true when this iscsi credentials modify o k response has a 2xx status code
func (o *IscsiCredentialsModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iscsi credentials modify o k response has a 3xx status code
func (o *IscsiCredentialsModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iscsi credentials modify o k response has a 4xx status code
func (o *IscsiCredentialsModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iscsi credentials modify o k response has a 5xx status code
func (o *IscsiCredentialsModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iscsi credentials modify o k response a status code equal to that given
func (o *IscsiCredentialsModifyOK) IsCode(code int) bool {
	return code == 200
}

func (o *IscsiCredentialsModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/san/iscsi/credentials/{svm.uuid}/{initiator}][%d] iscsiCredentialsModifyOK ", 200)
}

func (o *IscsiCredentialsModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/san/iscsi/credentials/{svm.uuid}/{initiator}][%d] iscsiCredentialsModifyOK ", 200)
}

func (o *IscsiCredentialsModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIscsiCredentialsModifyDefault creates a IscsiCredentialsModifyDefault with default headers values
func NewIscsiCredentialsModifyDefault(code int) *IscsiCredentialsModifyDefault {
	return &IscsiCredentialsModifyDefault{
		_statusCode: code,
	}
}

/*
	IscsiCredentialsModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621462 | An SVM with the specified UUID does not exist. |
| 2621706 | Both the SVM UUID and SVM name were supplied, but they do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either `svm.name` or `svm.uuid` must be supplied. |
| 5374145 | The iSCSI security password must contain an even number of valid hex digits. |
| 5374147 | The CHAP inbound and outbound passwords must be different. |
| 5374149 | The inbound user and password properties are required for CHAP authentication. |
| 5374150 | Outbound CHAP authentication requires an outbound password. |
| 5374155 | The functionality is not supported for the default security credential. |
| 5374855 | The value for property `initiator_address.ranges.start` is greater than the value for property `initiator_address.ranges.end`. |
| 5374856 | The value for property `initiator_address.ranges.start` does not belong to the same IP address family as the value for property `initiator_address.ranges.end`. |
| 5374895 | The iSCSI security credential does not exist on the specified SVM. |
| 5374900 | Setting the CHAP authentication properties are not supported with authentication types _none_ or _deny_. |
*/
type IscsiCredentialsModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the iscsi credentials modify default response
func (o *IscsiCredentialsModifyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this iscsi credentials modify default response has a 2xx status code
func (o *IscsiCredentialsModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iscsi credentials modify default response has a 3xx status code
func (o *IscsiCredentialsModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iscsi credentials modify default response has a 4xx status code
func (o *IscsiCredentialsModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iscsi credentials modify default response has a 5xx status code
func (o *IscsiCredentialsModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iscsi credentials modify default response a status code equal to that given
func (o *IscsiCredentialsModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IscsiCredentialsModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/san/iscsi/credentials/{svm.uuid}/{initiator}][%d] iscsi_credentials_modify default  %+v", o._statusCode, o.Payload)
}

func (o *IscsiCredentialsModifyDefault) String() string {
	return fmt.Sprintf("[PATCH /protocols/san/iscsi/credentials/{svm.uuid}/{initiator}][%d] iscsi_credentials_modify default  %+v", o._statusCode, o.Payload)
}

func (o *IscsiCredentialsModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IscsiCredentialsModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
