// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LicenseManagerURI License manager URI.
//
// swagger:model license_manager_uri
type LicenseManagerURI struct {

	// License manager host name, IPv4 or IPv6 address.
	// Example: 10.1.1.1
	Host *string `json:"host,omitempty"`
}

// Validate validates this license manager uri
func (m *LicenseManagerURI) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this license manager uri based on context it is used
func (m *LicenseManagerURI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LicenseManagerURI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicenseManagerURI) UnmarshalBinary(b []byte) error {
	var res LicenseManagerURI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
