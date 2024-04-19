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

// CifsSymlinkMappingGetReader is a Reader for the CifsSymlinkMappingGet structure.
type CifsSymlinkMappingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsSymlinkMappingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsSymlinkMappingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsSymlinkMappingGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsSymlinkMappingGetOK creates a CifsSymlinkMappingGetOK with default headers values
func NewCifsSymlinkMappingGetOK() *CifsSymlinkMappingGetOK {
	return &CifsSymlinkMappingGetOK{}
}

/*
CifsSymlinkMappingGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsSymlinkMappingGetOK struct {
	Payload *models.CifsSymlinkMapping
}

// IsSuccess returns true when this cifs symlink mapping get o k response has a 2xx status code
func (o *CifsSymlinkMappingGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs symlink mapping get o k response has a 3xx status code
func (o *CifsSymlinkMappingGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs symlink mapping get o k response has a 4xx status code
func (o *CifsSymlinkMappingGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs symlink mapping get o k response has a 5xx status code
func (o *CifsSymlinkMappingGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs symlink mapping get o k response a status code equal to that given
func (o *CifsSymlinkMappingGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *CifsSymlinkMappingGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/cifs/unix-symlink-mapping/{svm.uuid}/{unix_path}][%d] cifsSymlinkMappingGetOK  %+v", 200, o.Payload)
}

func (o *CifsSymlinkMappingGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/cifs/unix-symlink-mapping/{svm.uuid}/{unix_path}][%d] cifsSymlinkMappingGetOK  %+v", 200, o.Payload)
}

func (o *CifsSymlinkMappingGetOK) GetPayload() *models.CifsSymlinkMapping {
	return o.Payload
}

func (o *CifsSymlinkMappingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsSymlinkMapping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsSymlinkMappingGetDefault creates a CifsSymlinkMappingGetDefault with default headers values
func NewCifsSymlinkMappingGetDefault(code int) *CifsSymlinkMappingGetDefault {
	return &CifsSymlinkMappingGetDefault{
		_statusCode: code,
	}
}

/*
CifsSymlinkMappingGetDefault describes a response with status code -1, with default header values.

Error
*/
type CifsSymlinkMappingGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs symlink mapping get default response
func (o *CifsSymlinkMappingGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cifs symlink mapping get default response has a 2xx status code
func (o *CifsSymlinkMappingGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs symlink mapping get default response has a 3xx status code
func (o *CifsSymlinkMappingGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs symlink mapping get default response has a 4xx status code
func (o *CifsSymlinkMappingGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs symlink mapping get default response has a 5xx status code
func (o *CifsSymlinkMappingGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs symlink mapping get default response a status code equal to that given
func (o *CifsSymlinkMappingGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CifsSymlinkMappingGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/cifs/unix-symlink-mapping/{svm.uuid}/{unix_path}][%d] cifs_symlink_mapping_get default  %+v", o._statusCode, o.Payload)
}

func (o *CifsSymlinkMappingGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/cifs/unix-symlink-mapping/{svm.uuid}/{unix_path}][%d] cifs_symlink_mapping_get default  %+v", o._statusCode, o.Payload)
}

func (o *CifsSymlinkMappingGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsSymlinkMappingGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}