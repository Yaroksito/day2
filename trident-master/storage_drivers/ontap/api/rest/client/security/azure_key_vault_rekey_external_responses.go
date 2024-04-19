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

// AzureKeyVaultRekeyExternalReader is a Reader for the AzureKeyVaultRekeyExternal structure.
type AzureKeyVaultRekeyExternalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzureKeyVaultRekeyExternalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewAzureKeyVaultRekeyExternalAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAzureKeyVaultRekeyExternalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAzureKeyVaultRekeyExternalAccepted creates a AzureKeyVaultRekeyExternalAccepted with default headers values
func NewAzureKeyVaultRekeyExternalAccepted() *AzureKeyVaultRekeyExternalAccepted {
	return &AzureKeyVaultRekeyExternalAccepted{}
}

/*
AzureKeyVaultRekeyExternalAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AzureKeyVaultRekeyExternalAccepted struct {
}

// IsSuccess returns true when this azure key vault rekey external accepted response has a 2xx status code
func (o *AzureKeyVaultRekeyExternalAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure key vault rekey external accepted response has a 3xx status code
func (o *AzureKeyVaultRekeyExternalAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure key vault rekey external accepted response has a 4xx status code
func (o *AzureKeyVaultRekeyExternalAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure key vault rekey external accepted response has a 5xx status code
func (o *AzureKeyVaultRekeyExternalAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this azure key vault rekey external accepted response a status code equal to that given
func (o *AzureKeyVaultRekeyExternalAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *AzureKeyVaultRekeyExternalAccepted) Error() string {
	return fmt.Sprintf("[POST /security/azure-key-vaults/{azure_key_vault.uuid}/rekey-external][%d] azureKeyVaultRekeyExternalAccepted ", 202)
}

func (o *AzureKeyVaultRekeyExternalAccepted) String() string {
	return fmt.Sprintf("[POST /security/azure-key-vaults/{azure_key_vault.uuid}/rekey-external][%d] azureKeyVaultRekeyExternalAccepted ", 202)
}

func (o *AzureKeyVaultRekeyExternalAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAzureKeyVaultRekeyExternalDefault creates a AzureKeyVaultRekeyExternalDefault with default headers values
func NewAzureKeyVaultRekeyExternalDefault(code int) *AzureKeyVaultRekeyExternalDefault {
	return &AzureKeyVaultRekeyExternalDefault{
		_statusCode: code,
	}
}

/*
	AzureKeyVaultRekeyExternalDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65537120 | Azure Key Vault is not configured for the given SVM. |
| 65537547 | One or more volume encryption keys for encrypted volumes of this data SVM are stored in the key manager configured for the admin SVM. Use the REST API POST method to migrate this data SVM's keys from the admin SVM's key manager to this data SVM's key manager before running the rekey operation. |
*/
type AzureKeyVaultRekeyExternalDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the azure key vault rekey external default response
func (o *AzureKeyVaultRekeyExternalDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this azure key vault rekey external default response has a 2xx status code
func (o *AzureKeyVaultRekeyExternalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this azure key vault rekey external default response has a 3xx status code
func (o *AzureKeyVaultRekeyExternalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this azure key vault rekey external default response has a 4xx status code
func (o *AzureKeyVaultRekeyExternalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this azure key vault rekey external default response has a 5xx status code
func (o *AzureKeyVaultRekeyExternalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this azure key vault rekey external default response a status code equal to that given
func (o *AzureKeyVaultRekeyExternalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AzureKeyVaultRekeyExternalDefault) Error() string {
	return fmt.Sprintf("[POST /security/azure-key-vaults/{azure_key_vault.uuid}/rekey-external][%d] azure_key_vault_rekey_external default  %+v", o._statusCode, o.Payload)
}

func (o *AzureKeyVaultRekeyExternalDefault) String() string {
	return fmt.Sprintf("[POST /security/azure-key-vaults/{azure_key_vault.uuid}/rekey-external][%d] azure_key_vault_rekey_external default  %+v", o._statusCode, o.Payload)
}

func (o *AzureKeyVaultRekeyExternalDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AzureKeyVaultRekeyExternalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
