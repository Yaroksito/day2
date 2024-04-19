// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SanApplicationComponentsTiering application-components.tiering
//
// swagger:model san_application_components_tiering
type SanApplicationComponentsTiering struct {

	// Storage tiering placement rules for the container(s)
	// Enum: [required best_effort disallowed]
	Control *string `json:"control,omitempty"`

	// The storage tiering type of the application component.
	// Enum: [all auto none snapshot_only]
	Policy *string `json:"policy,omitempty"`

	// san application components tiering inline object stores
	SanApplicationComponentsTieringInlineObjectStores []*SanApplicationComponentsTieringInlineObjectStoresInlineArrayItem `json:"object_stores,omitempty"`
}

// Validate validates this san application components tiering
func (m *SanApplicationComponentsTiering) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateControl(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSanApplicationComponentsTieringInlineObjectStores(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sanApplicationComponentsTieringTypeControlPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["required","best_effort","disallowed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sanApplicationComponentsTieringTypeControlPropEnum = append(sanApplicationComponentsTieringTypeControlPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// san_application_components_tiering
	// SanApplicationComponentsTiering
	// control
	// Control
	// required
	// END DEBUGGING
	// SanApplicationComponentsTieringControlRequired captures enum value "required"
	SanApplicationComponentsTieringControlRequired string = "required"

	// BEGIN DEBUGGING
	// san_application_components_tiering
	// SanApplicationComponentsTiering
	// control
	// Control
	// best_effort
	// END DEBUGGING
	// SanApplicationComponentsTieringControlBestEffort captures enum value "best_effort"
	SanApplicationComponentsTieringControlBestEffort string = "best_effort"

	// BEGIN DEBUGGING
	// san_application_components_tiering
	// SanApplicationComponentsTiering
	// control
	// Control
	// disallowed
	// END DEBUGGING
	// SanApplicationComponentsTieringControlDisallowed captures enum value "disallowed"
	SanApplicationComponentsTieringControlDisallowed string = "disallowed"
)

// prop value enum
func (m *SanApplicationComponentsTiering) validateControlEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sanApplicationComponentsTieringTypeControlPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SanApplicationComponentsTiering) validateControl(formats strfmt.Registry) error {
	if swag.IsZero(m.Control) { // not required
		return nil
	}

	// value enum
	if err := m.validateControlEnum("control", "body", *m.Control); err != nil {
		return err
	}

	return nil
}

var sanApplicationComponentsTieringTypePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","auto","none","snapshot_only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sanApplicationComponentsTieringTypePolicyPropEnum = append(sanApplicationComponentsTieringTypePolicyPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// san_application_components_tiering
	// SanApplicationComponentsTiering
	// policy
	// Policy
	// all
	// END DEBUGGING
	// SanApplicationComponentsTieringPolicyAll captures enum value "all"
	SanApplicationComponentsTieringPolicyAll string = "all"

	// BEGIN DEBUGGING
	// san_application_components_tiering
	// SanApplicationComponentsTiering
	// policy
	// Policy
	// auto
	// END DEBUGGING
	// SanApplicationComponentsTieringPolicyAuto captures enum value "auto"
	SanApplicationComponentsTieringPolicyAuto string = "auto"

	// BEGIN DEBUGGING
	// san_application_components_tiering
	// SanApplicationComponentsTiering
	// policy
	// Policy
	// none
	// END DEBUGGING
	// SanApplicationComponentsTieringPolicyNone captures enum value "none"
	SanApplicationComponentsTieringPolicyNone string = "none"

	// BEGIN DEBUGGING
	// san_application_components_tiering
	// SanApplicationComponentsTiering
	// policy
	// Policy
	// snapshot_only
	// END DEBUGGING
	// SanApplicationComponentsTieringPolicySnapshotOnly captures enum value "snapshot_only"
	SanApplicationComponentsTieringPolicySnapshotOnly string = "snapshot_only"
)

// prop value enum
func (m *SanApplicationComponentsTiering) validatePolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sanApplicationComponentsTieringTypePolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SanApplicationComponentsTiering) validatePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	// value enum
	if err := m.validatePolicyEnum("policy", "body", *m.Policy); err != nil {
		return err
	}

	return nil
}

func (m *SanApplicationComponentsTiering) validateSanApplicationComponentsTieringInlineObjectStores(formats strfmt.Registry) error {
	if swag.IsZero(m.SanApplicationComponentsTieringInlineObjectStores) { // not required
		return nil
	}

	for i := 0; i < len(m.SanApplicationComponentsTieringInlineObjectStores); i++ {
		if swag.IsZero(m.SanApplicationComponentsTieringInlineObjectStores[i]) { // not required
			continue
		}

		if m.SanApplicationComponentsTieringInlineObjectStores[i] != nil {
			if err := m.SanApplicationComponentsTieringInlineObjectStores[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("object_stores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this san application components tiering based on the context it is used
func (m *SanApplicationComponentsTiering) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSanApplicationComponentsTieringInlineObjectStores(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SanApplicationComponentsTiering) contextValidateSanApplicationComponentsTieringInlineObjectStores(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SanApplicationComponentsTieringInlineObjectStores); i++ {

		if m.SanApplicationComponentsTieringInlineObjectStores[i] != nil {
			if err := m.SanApplicationComponentsTieringInlineObjectStores[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("object_stores" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SanApplicationComponentsTiering) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SanApplicationComponentsTiering) UnmarshalBinary(b []byte) error {
	var res SanApplicationComponentsTiering
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SanApplicationComponentsTieringInlineObjectStoresInlineArrayItem san application components tiering inline object stores inline array item
//
// swagger:model san_application_components_tiering_inline_object_stores_inline_array_item
type SanApplicationComponentsTieringInlineObjectStoresInlineArrayItem struct {

	// The name of the object-store to use.
	// Max Length: 512
	// Min Length: 1
	Name *string `json:"name,omitempty"`
}

// Validate validates this san application components tiering inline object stores inline array item
func (m *SanApplicationComponentsTieringInlineObjectStoresInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SanApplicationComponentsTieringInlineObjectStoresInlineArrayItem) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 512); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this san application components tiering inline object stores inline array item based on context it is used
func (m *SanApplicationComponentsTieringInlineObjectStoresInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SanApplicationComponentsTieringInlineObjectStoresInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SanApplicationComponentsTieringInlineObjectStoresInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res SanApplicationComponentsTieringInlineObjectStoresInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
