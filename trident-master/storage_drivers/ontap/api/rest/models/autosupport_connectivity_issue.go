// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AutosupportConnectivityIssue autosupport connectivity issue
//
// swagger:model autosupport_connectivity_issue
type AutosupportConnectivityIssue struct {

	// Error code
	// Example: 53149746
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// Error message
	// Example: SMTP connectivity check failed for destination: mailhost. Error: Could not resolve host - 'mailhost'
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this autosupport connectivity issue
func (m *AutosupportConnectivityIssue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this autosupport connectivity issue based on the context it is used
func (m *AutosupportConnectivityIssue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutosupportConnectivityIssue) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *AutosupportConnectivityIssue) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutosupportConnectivityIssue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutosupportConnectivityIssue) UnmarshalBinary(b []byte) error {
	var res AutosupportConnectivityIssue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}