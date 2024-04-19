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

// MetroclusterDrGroupReference DR group reference.
//
// swagger:model metrocluster_dr_group_reference
type MetroclusterDrGroupReference struct {

	// DR Group ID
	// Read Only: true
	ID *int64 `json:"id,omitempty"`
}

// Validate validates this metrocluster dr group reference
func (m *MetroclusterDrGroupReference) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this metrocluster dr group reference based on the context it is used
func (m *MetroclusterDrGroupReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDrGroupReference) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterDrGroupReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterDrGroupReference) UnmarshalBinary(b []byte) error {
	var res MetroclusterDrGroupReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}