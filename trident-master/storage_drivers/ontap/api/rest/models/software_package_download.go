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

// SoftwarePackageDownload software package download
//
// swagger:model software_package_download
type SoftwarePackageDownload struct {

	// Password for download
	// Example: admin_password
	// Format: password
	Password *strfmt.Password `json:"password,omitempty"`

	// HTTP or FTP URL of the package through a server
	// Example: http://server/package
	URL *string `json:"url,omitempty"`

	// Username for download
	// Example: admin
	Username *string `json:"username,omitempty"`
}

// Validate validates this software package download
func (m *SoftwarePackageDownload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwarePackageDownload) validatePassword(formats strfmt.Registry) error {
	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.FormatOf("password", "body", "password", m.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this software package download based on context it is used
func (m *SoftwarePackageDownload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SoftwarePackageDownload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwarePackageDownload) UnmarshalBinary(b []byte) error {
	var res SoftwarePackageDownload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
