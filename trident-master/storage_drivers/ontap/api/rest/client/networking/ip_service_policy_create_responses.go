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

// IPServicePolicyCreateReader is a Reader for the IPServicePolicyCreate structure.
type IPServicePolicyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPServicePolicyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIPServicePolicyCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIPServicePolicyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIPServicePolicyCreateOK creates a IPServicePolicyCreateOK with default headers values
func NewIPServicePolicyCreateOK() *IPServicePolicyCreateOK {
	return &IPServicePolicyCreateOK{}
}

/*
IPServicePolicyCreateOK describes a response with status code 200, with default header values.

OK
*/
type IPServicePolicyCreateOK struct {
}

// IsSuccess returns true when this ip service policy create o k response has a 2xx status code
func (o *IPServicePolicyCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ip service policy create o k response has a 3xx status code
func (o *IPServicePolicyCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ip service policy create o k response has a 4xx status code
func (o *IPServicePolicyCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ip service policy create o k response has a 5xx status code
func (o *IPServicePolicyCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ip service policy create o k response a status code equal to that given
func (o *IPServicePolicyCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IPServicePolicyCreateOK) Error() string {
	return fmt.Sprintf("[POST /network/ip/service-policies][%d] ipServicePolicyCreateOK ", 200)
}

func (o *IPServicePolicyCreateOK) String() string {
	return fmt.Sprintf("[POST /network/ip/service-policies][%d] ipServicePolicyCreateOK ", 200)
}

func (o *IPServicePolicyCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIPServicePolicyCreateDefault creates a IPServicePolicyCreateDefault with default headers values
func NewIPServicePolicyCreateDefault(code int) *IPServicePolicyCreateDefault {
	return &IPServicePolicyCreateDefault{
		_statusCode: code,
	}
}

/*
	IPServicePolicyCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1966373 | Port must reside in the same IPspace as the interface's SVM. |
| 1967146 | Svm.name does not exist. |
| 1967147 | Svm.uuid does not exist. |
| 53281929 | Service policies cannot combine block and file services. |
| 53281931 | Service policy names cannot start with "default-". |
| 53281932 | Service cannot be added because the service does not exist for the specified SVM or IPspace. |
| 53281933 | A Cluster-scoped service cannot be added to a SVM-scoped service policy. |
| 53281934 | An SVM-scoped service cannot be added to a Cluster-scoped service policy. |
| 53281935 | Scope is set to "svm" and svm.uuid or svm.name have not been specified. |
| 53281936 | The SVM is not in the specified IPspace. |
| 53281937 | Svm.uuid and svm.name are not valid parameters when scope is cluster. |
| 53281938 | Svm.uuid or svm.name specify a vserver that does not exist. |
| 53281939 | One or more of the svm.uuid, svm.name, ipspace.uuid, and ipspace.name have invalid values. |
| 53281940 | SVM or IPspace has not been specified. |
| 53281941 | SVM does not exist. |
| 53281944 | Ipspace.name does not exist. |
| 53281945 | Ipspace.uuid is not an IPspace. |
| 53281946 | Service policy already exists. |
| 53281958 | Service policies cannot contain multiple block-oriented services. |
*/
type IPServicePolicyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ip service policy create default response
func (o *IPServicePolicyCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ip service policy create default response has a 2xx status code
func (o *IPServicePolicyCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ip service policy create default response has a 3xx status code
func (o *IPServicePolicyCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ip service policy create default response has a 4xx status code
func (o *IPServicePolicyCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ip service policy create default response has a 5xx status code
func (o *IPServicePolicyCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ip service policy create default response a status code equal to that given
func (o *IPServicePolicyCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IPServicePolicyCreateDefault) Error() string {
	return fmt.Sprintf("[POST /network/ip/service-policies][%d] ip_service_policy_create default  %+v", o._statusCode, o.Payload)
}

func (o *IPServicePolicyCreateDefault) String() string {
	return fmt.Sprintf("[POST /network/ip/service-policies][%d] ip_service_policy_create default  %+v", o._statusCode, o.Payload)
}

func (o *IPServicePolicyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IPServicePolicyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}