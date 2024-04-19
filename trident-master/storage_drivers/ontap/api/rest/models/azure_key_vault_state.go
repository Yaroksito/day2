// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureKeyVaultState Indicates whether or not the AKV wrapped internal key is available cluster wide.
// This is an advanced property; there is an added computationl cost to retrieving its value. The property is not populated for either a collection GET or an instance GET unless it is explicitly requested using the `fields` query parameter or GET for all advanced properties is enabled.
//
// swagger:model azure_key_vault_state
type AzureKeyVaultState struct {

	// Set to true when an AKV wrapped internal key is present on all nodes of the cluster.
	Available *bool `json:"available,omitempty"`

	// Code corresponding to the status message. Returns a 0 if AKV wrapped key is available on all nodes in the cluster.
	// Example: 346758
	Code *string `json:"code,omitempty"`

	// Error message set when top-level internal key protection key (KEK) availability on cluster is false.
	// Example: Top-level internal key protection key (KEK) is unavailable on the following nodes with the associated reasons: Node: node1. Reason: No volumes created yet for the SVM. Wrapped KEK status will be available after creating encrypted volumes.
	Message *string `json:"message,omitempty"`
}

// Validate validates this azure key vault state
func (m *AzureKeyVaultState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure key vault state based on context it is used
func (m *AzureKeyVaultState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureKeyVaultState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureKeyVaultState) UnmarshalBinary(b []byte) error {
	var res AzureKeyVaultState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}