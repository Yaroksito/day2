// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// QosPolicyModifyReader is a Reader for the QosPolicyModify structure.
type QosPolicyModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QosPolicyModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewQosPolicyModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQosPolicyModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQosPolicyModifyAccepted creates a QosPolicyModifyAccepted with default headers values
func NewQosPolicyModifyAccepted() *QosPolicyModifyAccepted {
	return &QosPolicyModifyAccepted{}
}

/*
QosPolicyModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type QosPolicyModifyAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this qos policy modify accepted response has a 2xx status code
func (o *QosPolicyModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qos policy modify accepted response has a 3xx status code
func (o *QosPolicyModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qos policy modify accepted response has a 4xx status code
func (o *QosPolicyModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this qos policy modify accepted response has a 5xx status code
func (o *QosPolicyModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this qos policy modify accepted response a status code equal to that given
func (o *QosPolicyModifyAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *QosPolicyModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /storage/qos/policies/{uuid}][%d] qosPolicyModifyAccepted  %+v", 202, o.Payload)
}

func (o *QosPolicyModifyAccepted) String() string {
	return fmt.Sprintf("[PATCH /storage/qos/policies/{uuid}][%d] qosPolicyModifyAccepted  %+v", 202, o.Payload)
}

func (o *QosPolicyModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *QosPolicyModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQosPolicyModifyDefault creates a QosPolicyModifyDefault with default headers values
func NewQosPolicyModifyDefault(code int) *QosPolicyModifyDefault {
	return &QosPolicyModifyDefault{
		_statusCode: code,
	}
}

/*
	QosPolicyModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 8454147 | The maximum limit for QoS policies has been reached. |
| 8454154 | The name specified for creating conflicts with an existing QoS policy name. |
| 8454260 | Invalid value for maximum and minimum fields. Valid values for max_throughput_iops and max_throughput_mbps combination is for the ratio of max_throughput_mbps and max_throughput_iops to be within 1 to 4096. |
| 8454273 | Invalid value for an adaptive field. Value should be non-zero. |
| 8454277 | The name specified for creating an adaptive QoS policy conflicts with an existing fixed QoS policy name. |
| 8454278 | The name specified for creating a fixed QoS policy conflicts with an existing adaptive QoS policy name. |
| 8454286 | Modifications on these cluster scoped preset policies is prohibited. |
| 8454327 | The existing fixed QoS policy cannot be modified to an adaptive QoS policy. |
| 8454328 | The existing adaptive QoS policy cannot be modified to a fixed QoS policy. |
*/
type QosPolicyModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the qos policy modify default response
func (o *QosPolicyModifyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this qos policy modify default response has a 2xx status code
func (o *QosPolicyModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this qos policy modify default response has a 3xx status code
func (o *QosPolicyModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this qos policy modify default response has a 4xx status code
func (o *QosPolicyModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this qos policy modify default response has a 5xx status code
func (o *QosPolicyModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this qos policy modify default response a status code equal to that given
func (o *QosPolicyModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *QosPolicyModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/qos/policies/{uuid}][%d] qos_policy_modify default  %+v", o._statusCode, o.Payload)
}

func (o *QosPolicyModifyDefault) String() string {
	return fmt.Sprintf("[PATCH /storage/qos/policies/{uuid}][%d] qos_policy_modify default  %+v", o._statusCode, o.Payload)
}

func (o *QosPolicyModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QosPolicyModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}