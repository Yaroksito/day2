// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MultiAdminVerifyConfig multi admin verify config
//
// swagger:model multi_admin_verify_config
type MultiAdminVerifyConfig struct {

	// Default time for requests to be approved, in ISO-8601 duration format.
	ApprovalExpiry *string `json:"approval_expiry,omitempty"`

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Default time for requests to be executed once approved, in ISO-8601 duration format.
	ExecutionExpiry *string `json:"execution_expiry,omitempty"`

	// List of approval groups that are allowed to approve requests for rules that don't have approval groups.
	MultiAdminVerifyConfigInlineApprovalGroups []*string `json:"approval_groups,omitempty"`

	// The number of required approvers, excluding the user that made the request.
	RequiredApprovers *int64 `json:"required_approvers,omitempty"`
}

// Validate validates this multi admin verify config
func (m *MultiAdminVerifyConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this multi admin verify config based on context it is used
func (m *MultiAdminVerifyConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MultiAdminVerifyConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultiAdminVerifyConfig) UnmarshalBinary(b []byte) error {
	var res MultiAdminVerifyConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}