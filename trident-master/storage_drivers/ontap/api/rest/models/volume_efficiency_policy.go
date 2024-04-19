// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VolumeEfficiencyPolicy volume efficiency policy
//
// swagger:model volume_efficiency_policy
type VolumeEfficiencyPolicy struct {

	// links
	Links *VolumeEfficiencyPolicyInlineLinks `json:"_links,omitempty"`

	// A comment associated with the volume efficiency policy.
	Comment *string `json:"comment,omitempty"`

	// This field is used with the policy type "scheduled" to indicate the allowed duration for a session, in hours. Possible value is a number between 0 and 999 inclusive. Default is unlimited indicated by value 0.
	// Example: 5
	Duration *int64 `json:"duration,omitempty"`

	// Is the volume efficiency policy enabled?
	// Example: true
	Enabled *bool `json:"enabled,omitempty"`

	// Name of the volume efficiency policy.
	// Example: default
	Name *string `json:"name,omitempty"`

	// QoS policy for the sis operation. Possible values are background and best_effort. In background, sis operation will run in background with minimal or no impact on data serving client operations. In best_effort, sis operations may have some impact on data serving client operations.
	// Enum: [background best_effort]
	QosPolicy *string `json:"qos_policy,omitempty"`

	// schedule
	Schedule *VolumeEfficiencyPolicyInlineSchedule `json:"schedule,omitempty"`

	// This field is used with the policy type "threshold" to indicate the threshold percentage for triggering the volume efficiency policy. It is mutuallly exclusive of the schedule.
	// Example: 30
	StartThresholdPercent *int64 `json:"start_threshold_percent,omitempty"`

	// svm
	Svm *VolumeEfficiencyPolicyInlineSvm `json:"svm,omitempty"`

	// Type of volume efficiency policy.
	// Enum: [threshold scheduled]
	Type *string `json:"type,omitempty"`

	// Unique identifier of volume efficiency policy.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this volume efficiency policy
func (m *VolumeEfficiencyPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQosPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeEfficiencyPolicy) validateLinks(formats strfmt.Registry) error {
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

var volumeEfficiencyPolicyTypeQosPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["background","best_effort"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		volumeEfficiencyPolicyTypeQosPolicyPropEnum = append(volumeEfficiencyPolicyTypeQosPolicyPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// volume_efficiency_policy
	// VolumeEfficiencyPolicy
	// qos_policy
	// QosPolicy
	// background
	// END DEBUGGING
	// VolumeEfficiencyPolicyQosPolicyBackground captures enum value "background"
	VolumeEfficiencyPolicyQosPolicyBackground string = "background"

	// BEGIN DEBUGGING
	// volume_efficiency_policy
	// VolumeEfficiencyPolicy
	// qos_policy
	// QosPolicy
	// best_effort
	// END DEBUGGING
	// VolumeEfficiencyPolicyQosPolicyBestEffort captures enum value "best_effort"
	VolumeEfficiencyPolicyQosPolicyBestEffort string = "best_effort"
)

// prop value enum
func (m *VolumeEfficiencyPolicy) validateQosPolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, volumeEfficiencyPolicyTypeQosPolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VolumeEfficiencyPolicy) validateQosPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.QosPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateQosPolicyEnum("qos_policy", "body", *m.QosPolicy); err != nil {
		return err
	}

	return nil
}

func (m *VolumeEfficiencyPolicy) validateSchedule(formats strfmt.Registry) error {
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

func (m *VolumeEfficiencyPolicy) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

var volumeEfficiencyPolicyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["threshold","scheduled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		volumeEfficiencyPolicyTypeTypePropEnum = append(volumeEfficiencyPolicyTypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// volume_efficiency_policy
	// VolumeEfficiencyPolicy
	// type
	// Type
	// threshold
	// END DEBUGGING
	// VolumeEfficiencyPolicyTypeThreshold captures enum value "threshold"
	VolumeEfficiencyPolicyTypeThreshold string = "threshold"

	// BEGIN DEBUGGING
	// volume_efficiency_policy
	// VolumeEfficiencyPolicy
	// type
	// Type
	// scheduled
	// END DEBUGGING
	// VolumeEfficiencyPolicyTypeScheduled captures enum value "scheduled"
	VolumeEfficiencyPolicyTypeScheduled string = "scheduled"
)

// prop value enum
func (m *VolumeEfficiencyPolicy) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, volumeEfficiencyPolicyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VolumeEfficiencyPolicy) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this volume efficiency policy based on the context it is used
func (m *VolumeEfficiencyPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeEfficiencyPolicy) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VolumeEfficiencyPolicy) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VolumeEfficiencyPolicy) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *VolumeEfficiencyPolicy) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumeEfficiencyPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeEfficiencyPolicy) UnmarshalBinary(b []byte) error {
	var res VolumeEfficiencyPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VolumeEfficiencyPolicyInlineLinks volume efficiency policy inline links
//
// swagger:model volume_efficiency_policy_inline__links
type VolumeEfficiencyPolicyInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this volume efficiency policy inline links
func (m *VolumeEfficiencyPolicyInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeEfficiencyPolicyInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this volume efficiency policy inline links based on the context it is used
func (m *VolumeEfficiencyPolicyInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeEfficiencyPolicyInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VolumeEfficiencyPolicyInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeEfficiencyPolicyInlineLinks) UnmarshalBinary(b []byte) error {
	var res VolumeEfficiencyPolicyInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VolumeEfficiencyPolicyInlineSchedule volume efficiency policy inline schedule
//
// swagger:model volume_efficiency_policy_inline_schedule
type VolumeEfficiencyPolicyInlineSchedule struct {

	// Schedule at which volume efficiency policies are captured on the SVM. Some common schedules already defined in the system are hourly, daily, weekly, at 5 minute intervals, and at 8 hour intervals. Volume efficiency policies with custom schedules can be referenced.
	// Example: daily
	Name *string `json:"name,omitempty"`
}

// Validate validates this volume efficiency policy inline schedule
func (m *VolumeEfficiencyPolicyInlineSchedule) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume efficiency policy inline schedule based on context it is used
func (m *VolumeEfficiencyPolicyInlineSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeEfficiencyPolicyInlineSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeEfficiencyPolicyInlineSchedule) UnmarshalBinary(b []byte) error {
	var res VolumeEfficiencyPolicyInlineSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VolumeEfficiencyPolicyInlineSvm volume efficiency policy inline svm
//
// swagger:model volume_efficiency_policy_inline_svm
type VolumeEfficiencyPolicyInlineSvm struct {

	// links
	Links *VolumeEfficiencyPolicyInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this volume efficiency policy inline svm
func (m *VolumeEfficiencyPolicyInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeEfficiencyPolicyInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this volume efficiency policy inline svm based on the context it is used
func (m *VolumeEfficiencyPolicyInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeEfficiencyPolicyInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumeEfficiencyPolicyInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeEfficiencyPolicyInlineSvm) UnmarshalBinary(b []byte) error {
	var res VolumeEfficiencyPolicyInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VolumeEfficiencyPolicyInlineSvmInlineLinks volume efficiency policy inline svm inline links
//
// swagger:model volume_efficiency_policy_inline_svm_inline__links
type VolumeEfficiencyPolicyInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this volume efficiency policy inline svm inline links
func (m *VolumeEfficiencyPolicyInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeEfficiencyPolicyInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this volume efficiency policy inline svm inline links based on the context it is used
func (m *VolumeEfficiencyPolicyInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeEfficiencyPolicyInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumeEfficiencyPolicyInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeEfficiencyPolicyInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res VolumeEfficiencyPolicyInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}