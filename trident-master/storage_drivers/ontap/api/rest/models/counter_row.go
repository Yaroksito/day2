// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CounterRow A single row of counter and property counter data.
//
// swagger:model counter_row
type CounterRow struct {

	// links
	Links *CounterRowInlineLinks `json:"_links,omitempty"`

	// aggregation
	Aggregation *InstanceCounterAggregation `json:"aggregation,omitempty"`

	// Array of counter name/value pairs.
	CounterRowInlineCounters []*Counter `json:"counters,omitempty"`

	// Array of property name/value pairs.
	CounterRowInlineProperties []*CounterProperty `json:"properties,omitempty"`

	// counter table
	CounterTable *CounterTableReference `json:"counter_table,omitempty"`

	// Unique row idenfier.
	ID *string `json:"id,omitempty"`
}

// Validate validates this counter row
func (m *CounterRow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAggregation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCounterRowInlineCounters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCounterRowInlineProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCounterTable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CounterRow) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *CounterRow) validateAggregation(formats strfmt.Registry) error {
	if swag.IsZero(m.Aggregation) { // not required
		return nil
	}

	if m.Aggregation != nil {
		if err := m.Aggregation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregation")
			}
			return err
		}
	}

	return nil
}

func (m *CounterRow) validateCounterRowInlineCounters(formats strfmt.Registry) error {
	if swag.IsZero(m.CounterRowInlineCounters) { // not required
		return nil
	}

	for i := 0; i < len(m.CounterRowInlineCounters); i++ {
		if swag.IsZero(m.CounterRowInlineCounters[i]) { // not required
			continue
		}

		if m.CounterRowInlineCounters[i] != nil {
			if err := m.CounterRowInlineCounters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("counters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CounterRow) validateCounterRowInlineProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.CounterRowInlineProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.CounterRowInlineProperties); i++ {
		if swag.IsZero(m.CounterRowInlineProperties[i]) { // not required
			continue
		}

		if m.CounterRowInlineProperties[i] != nil {
			if err := m.CounterRowInlineProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CounterRow) validateCounterTable(formats strfmt.Registry) error {
	if swag.IsZero(m.CounterTable) { // not required
		return nil
	}

	if m.CounterTable != nil {
		if err := m.CounterTable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counter_table")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this counter row based on the context it is used
func (m *CounterRow) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAggregation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCounterRowInlineCounters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCounterRowInlineProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCounterTable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CounterRow) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *CounterRow) contextValidateAggregation(ctx context.Context, formats strfmt.Registry) error {

	if m.Aggregation != nil {
		if err := m.Aggregation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregation")
			}
			return err
		}
	}

	return nil
}

func (m *CounterRow) contextValidateCounterRowInlineCounters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CounterRowInlineCounters); i++ {

		if m.CounterRowInlineCounters[i] != nil {
			if err := m.CounterRowInlineCounters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("counters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CounterRow) contextValidateCounterRowInlineProperties(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CounterRowInlineProperties); i++ {

		if m.CounterRowInlineProperties[i] != nil {
			if err := m.CounterRowInlineProperties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CounterRow) contextValidateCounterTable(ctx context.Context, formats strfmt.Registry) error {

	if m.CounterTable != nil {
		if err := m.CounterTable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counter_table")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CounterRow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CounterRow) UnmarshalBinary(b []byte) error {
	var res CounterRow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CounterRowInlineLinks counter row inline links
//
// swagger:model counter_row_inline__links
type CounterRowInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this counter row inline links
func (m *CounterRowInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CounterRowInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this counter row inline links based on the context it is used
func (m *CounterRowInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CounterRowInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CounterRowInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CounterRowInlineLinks) UnmarshalBinary(b []byte) error {
	var res CounterRowInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
