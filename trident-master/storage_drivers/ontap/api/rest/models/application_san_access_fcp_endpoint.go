// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationSanAccessFcpEndpoint A Fibre Channel Protocol (FCP) access endpoint for the LUN.
//
// swagger:model application_san_access_fcp_endpoint
type ApplicationSanAccessFcpEndpoint struct {

	// interface
	// Read Only: true
	Interface *FcInterfaceReference `json:"interface,omitempty"`
}

// Validate validates this application san access fcp endpoint
func (m *ApplicationSanAccessFcpEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterface(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSanAccessFcpEndpoint) validateInterface(formats strfmt.Registry) error {
	if swag.IsZero(m.Interface) { // not required
		return nil
	}

	if m.Interface != nil {
		if err := m.Interface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application san access fcp endpoint based on the context it is used
func (m *ApplicationSanAccessFcpEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInterface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSanAccessFcpEndpoint) contextValidateInterface(ctx context.Context, formats strfmt.Registry) error {

	if m.Interface != nil {
		if err := m.Interface.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSanAccessFcpEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSanAccessFcpEndpoint) UnmarshalBinary(b []byte) error {
	var res ApplicationSanAccessFcpEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
