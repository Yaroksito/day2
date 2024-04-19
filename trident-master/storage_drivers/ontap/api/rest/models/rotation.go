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

// Rotation Audit event log files are rotated when they reach a configured threshold log size or are on a configured schedule. When an event log file is rotated, the scheduled consolidation task first renames the active converted file to a time-stamped archive file, and then creates a new active converted event log file.
//
// swagger:model rotation
type Rotation struct {

	// Manually rotates the audit logs. Optional in PATCH only. Not available in POST.
	Now *bool `json:"now,omitempty"`

	// schedule
	Schedule *AuditSchedule `json:"schedule,omitempty"`

	// Rotates logs based on log size in bytes.
	Size *int64 `json:"size,omitempty"`
}

// Validate validates this rotation
func (m *Rotation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rotation) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rotation based on the context it is used
func (m *Rotation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rotation) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {
		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Rotation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rotation) UnmarshalBinary(b []byte) error {
	var res Rotation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
