// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// AwsKmsCreateReader is a Reader for the AwsKmsCreate structure.
type AwsKmsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AwsKmsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAwsKmsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAwsKmsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAwsKmsCreateCreated creates a AwsKmsCreateCreated with default headers values
func NewAwsKmsCreateCreated() *AwsKmsCreateCreated {
	return &AwsKmsCreateCreated{}
}

/*
AwsKmsCreateCreated describes a response with status code 201, with default header values.

Created
*/
type AwsKmsCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.AwsKmsResponse
}

// IsSuccess returns true when this aws kms create created response has a 2xx status code
func (o *AwsKmsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aws kms create created response has a 3xx status code
func (o *AwsKmsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aws kms create created response has a 4xx status code
func (o *AwsKmsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this aws kms create created response has a 5xx status code
func (o *AwsKmsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this aws kms create created response a status code equal to that given
func (o *AwsKmsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AwsKmsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /security/aws-kms][%d] awsKmsCreateCreated  %+v", 201, o.Payload)
}

func (o *AwsKmsCreateCreated) String() string {
	return fmt.Sprintf("[POST /security/aws-kms][%d] awsKmsCreateCreated  %+v", 201, o.Payload)
}

func (o *AwsKmsCreateCreated) GetPayload() *models.AwsKmsResponse {
	return o.Payload
}

func (o *AwsKmsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.AwsKmsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAwsKmsCreateDefault creates a AwsKmsCreateDefault with default headers values
func NewAwsKmsCreateDefault(code int) *AwsKmsCreateDefault {
	return &AwsKmsCreateDefault{
		_statusCode: code,
	}
}

/*
	AwsKmsCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 3735622 | Certificate type not supported for create operation. |
| 3735645 | You cannot specify a value for serial as it is generated automatically. |
| 3735657 | Specifying \\\"-subtype\\\" when creating a certificate is not supported. |
| 3735664 | Specified key size is not supported in FIPS mode. |
| 3735665 | Specified hash function is not supported in FIPS mode. |
| 3735700 | Specified key size is not supported. |
| 65536600 | Nodes are out of quorum. |
| 65537518 | Failed to find a LIF with Cluster role on node. One or more nodes may be out of quorum. |
| 65537900 | Failed to enable the Amazon Web Service Key Management Service for an SVM due to an invalid secret access key. |
| 65537901 | The Amazon Web Service Key Management Service (AWSKMS) cannot be enabled because all nodes in the cluster are not running a version that supports the AWSKMS feature. |
| 65537906 | Failed to store the secret access key. |
| 65537907 | The Amazon Web Service Key Management Service is disabled on the cluster. For further assistance, contact technical support. |
| 65537908 | The Amazon Web Service Key Management Service is not supported for the admin SVM. |
| 65537910 | Failed to configure Amazon Web Service Key Management Service for an SVM because a key manager has already been configured for the SVM. |
| 65537911 | The Amazon Web Service Key Management Service is not supported in MetroCluster configurations. |
| 65537912 | The Amazon Web Service Key Management Service cannot be configured for an SVM because one or more volume encryption keys of the SVM are stored on the admin SVM. |
| 65537926 | The Amazon Web Service Key Management Service is not configured for this SVM. |
*/
type AwsKmsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the aws kms create default response
func (o *AwsKmsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this aws kms create default response has a 2xx status code
func (o *AwsKmsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this aws kms create default response has a 3xx status code
func (o *AwsKmsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this aws kms create default response has a 4xx status code
func (o *AwsKmsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this aws kms create default response has a 5xx status code
func (o *AwsKmsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this aws kms create default response a status code equal to that given
func (o *AwsKmsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AwsKmsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /security/aws-kms][%d] aws_kms_create default  %+v", o._statusCode, o.Payload)
}

func (o *AwsKmsCreateDefault) String() string {
	return fmt.Sprintf("[POST /security/aws-kms][%d] aws_kms_create default  %+v", o._statusCode, o.Payload)
}

func (o *AwsKmsCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AwsKmsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
