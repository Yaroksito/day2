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

// VscanOnAccessModifyReader is a Reader for the VscanOnAccessModify structure.
type VscanOnAccessModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanOnAccessModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanOnAccessModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanOnAccessModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanOnAccessModifyOK creates a VscanOnAccessModifyOK with default headers values
func NewVscanOnAccessModifyOK() *VscanOnAccessModifyOK {
	return &VscanOnAccessModifyOK{}
}

/*
VscanOnAccessModifyOK describes a response with status code 200, with default header values.

OK
*/
type VscanOnAccessModifyOK struct {
}

// IsSuccess returns true when this vscan on access modify o k response has a 2xx status code
func (o *VscanOnAccessModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan on access modify o k response has a 3xx status code
func (o *VscanOnAccessModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan on access modify o k response has a 4xx status code
func (o *VscanOnAccessModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan on access modify o k response has a 5xx status code
func (o *VscanOnAccessModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan on access modify o k response a status code equal to that given
func (o *VscanOnAccessModifyOK) IsCode(code int) bool {
	return code == 200
}

func (o *VscanOnAccessModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/vscan/{svm.uuid}/on-access-policies/{name}][%d] vscanOnAccessModifyOK ", 200)
}

func (o *VscanOnAccessModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/vscan/{svm.uuid}/on-access-policies/{name}][%d] vscanOnAccessModifyOK ", 200)
}

func (o *VscanOnAccessModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVscanOnAccessModifyDefault creates a VscanOnAccessModifyDefault with default headers values
func NewVscanOnAccessModifyDefault(code int) *VscanOnAccessModifyDefault {
	return &VscanOnAccessModifyDefault{
		_statusCode: code,
	}
}

/*
	VscanOnAccessModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 10027033   | Configurations for an On-Access policy associated with an administrative SVM cannot be modified. However, the policy can be enabled or disabled. |
| 10027046   | The specified SVM is not the owner of the specified policy. Check for the correct SVM who owns the policy. |
| 10027101   | The file size must be in the range 1KB to 1TB |
| 10027107   | The include extensions list cannot be empty. Specify at least one extension for inclusion. |
| 10027109   | The specified CIFS path is invalid. It must be in the form \"\\dir1\\dir2\" or \"\\dir1\\dir2\\\". |
| 10027249   | The On-Access policy updated successfully but failed to enable/disable the policy. The reason for an enable policy operation failure might be that another policy is enabled. Disable the already enabled policy and then enable the policy. The reason for a disable policy operation failure might be that Vscan is enabled on the SVM. Disable the Vscan first and then disable the policy. |
| 10027250   | The On-Access policy cannot be enabled/disabled. The reason for an enable policy operation failure might be that another policy is enabled. Disable the already enabled policy and then enable the policy. The reason for a disable policy operation failure might be that Vscan is enabled on the SVM. Disable the Vscan and then disable the policy. |
| 10027253   | The number of paths specified exceeds the configured maximum number of paths. You cannot specify more than the maximum number of configured paths. |
| 10027254   | The number of extensions specified exceeds the configured maximum number of extensions. You cannot specify more than the maximum number of configured extensions. |
*/
type VscanOnAccessModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the vscan on access modify default response
func (o *VscanOnAccessModifyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this vscan on access modify default response has a 2xx status code
func (o *VscanOnAccessModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan on access modify default response has a 3xx status code
func (o *VscanOnAccessModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan on access modify default response has a 4xx status code
func (o *VscanOnAccessModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan on access modify default response has a 5xx status code
func (o *VscanOnAccessModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan on access modify default response a status code equal to that given
func (o *VscanOnAccessModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VscanOnAccessModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/vscan/{svm.uuid}/on-access-policies/{name}][%d] vscan_on_access_modify default  %+v", o._statusCode, o.Payload)
}

func (o *VscanOnAccessModifyDefault) String() string {
	return fmt.Sprintf("[PATCH /protocols/vscan/{svm.uuid}/on-access-policies/{name}][%d] vscan_on_access_modify default  %+v", o._statusCode, o.Payload)
}

func (o *VscanOnAccessModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanOnAccessModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}