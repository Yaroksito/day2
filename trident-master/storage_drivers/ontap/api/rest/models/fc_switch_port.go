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

// FcSwitchPort A Fibre Channel switch port.
//
// swagger:model fc_switch_port
type FcSwitchPort struct {

	// attached device
	AttachedDevice *FcSwitchPortInlineAttachedDevice `json:"attached_device,omitempty"`

	// The slot of the Fibre Channel switch port.
	//
	// Example: 1
	// Read Only: true
	Slot *string `json:"slot,omitempty"`

	// The state of the Fibre Channel switch port.
	//
	// Example: online
	// Read Only: true
	// Enum: [unknown online offline testing fault]
	State *string `json:"state,omitempty"`

	// The type of the Fibre Channel switch port.
	//
	// Read Only: true
	// Enum: [b_port e_port f_port fl_port fnl_port fv_port n_port nl_port nv_port nx_port sd_port te_port tf_port tl_port tnp_port none]
	Type *string `json:"type,omitempty"`

	// The world wide port name (WWPN) of the Fibre Channel switch port.
	//
	// Example: 50:0a:31:32:33:34:35:36
	// Read Only: true
	Wwpn *string `json:"wwpn,omitempty"`
}

// Validate validates this fc switch port
func (m *FcSwitchPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachedDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *FcSwitchPort) validateAttachedDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.AttachedDevice) { // not required
		return nil
	}

	if m.AttachedDevice != nil {
		if err := m.AttachedDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attached_device")
			}
			return err
		}
	}

	return nil
}

var fcSwitchPortTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","online","offline","testing","fault"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcSwitchPortTypeStatePropEnum = append(fcSwitchPortTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// state
	// State
	// unknown
	// END DEBUGGING
	// FcSwitchPortStateUnknown captures enum value "unknown"
	FcSwitchPortStateUnknown string = "unknown"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// state
	// State
	// online
	// END DEBUGGING
	// FcSwitchPortStateOnline captures enum value "online"
	FcSwitchPortStateOnline string = "online"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// state
	// State
	// offline
	// END DEBUGGING
	// FcSwitchPortStateOffline captures enum value "offline"
	FcSwitchPortStateOffline string = "offline"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// state
	// State
	// testing
	// END DEBUGGING
	// FcSwitchPortStateTesting captures enum value "testing"
	FcSwitchPortStateTesting string = "testing"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// state
	// State
	// fault
	// END DEBUGGING
	// FcSwitchPortStateFault captures enum value "fault"
	FcSwitchPortStateFault string = "fault"
)

// prop value enum
func (m *FcSwitchPort) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcSwitchPortTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FcSwitchPort) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

var fcSwitchPortTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["b_port","e_port","f_port","fl_port","fnl_port","fv_port","n_port","nl_port","nv_port","nx_port","sd_port","te_port","tf_port","tl_port","tnp_port","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcSwitchPortTypeTypePropEnum = append(fcSwitchPortTypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// b_port
	// END DEBUGGING
	// FcSwitchPortTypeBPort captures enum value "b_port"
	FcSwitchPortTypeBPort string = "b_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// e_port
	// END DEBUGGING
	// FcSwitchPortTypeEPort captures enum value "e_port"
	FcSwitchPortTypeEPort string = "e_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// f_port
	// END DEBUGGING
	// FcSwitchPortTypeFPort captures enum value "f_port"
	FcSwitchPortTypeFPort string = "f_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// fl_port
	// END DEBUGGING
	// FcSwitchPortTypeFlPort captures enum value "fl_port"
	FcSwitchPortTypeFlPort string = "fl_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// fnl_port
	// END DEBUGGING
	// FcSwitchPortTypeFnlPort captures enum value "fnl_port"
	FcSwitchPortTypeFnlPort string = "fnl_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// fv_port
	// END DEBUGGING
	// FcSwitchPortTypeFvPort captures enum value "fv_port"
	FcSwitchPortTypeFvPort string = "fv_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// n_port
	// END DEBUGGING
	// FcSwitchPortTypeNPort captures enum value "n_port"
	FcSwitchPortTypeNPort string = "n_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// nl_port
	// END DEBUGGING
	// FcSwitchPortTypeNlPort captures enum value "nl_port"
	FcSwitchPortTypeNlPort string = "nl_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// nv_port
	// END DEBUGGING
	// FcSwitchPortTypeNvPort captures enum value "nv_port"
	FcSwitchPortTypeNvPort string = "nv_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// nx_port
	// END DEBUGGING
	// FcSwitchPortTypeNxPort captures enum value "nx_port"
	FcSwitchPortTypeNxPort string = "nx_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// sd_port
	// END DEBUGGING
	// FcSwitchPortTypeSdPort captures enum value "sd_port"
	FcSwitchPortTypeSdPort string = "sd_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// te_port
	// END DEBUGGING
	// FcSwitchPortTypeTePort captures enum value "te_port"
	FcSwitchPortTypeTePort string = "te_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// tf_port
	// END DEBUGGING
	// FcSwitchPortTypeTfPort captures enum value "tf_port"
	FcSwitchPortTypeTfPort string = "tf_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// tl_port
	// END DEBUGGING
	// FcSwitchPortTypeTlPort captures enum value "tl_port"
	FcSwitchPortTypeTlPort string = "tl_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// tnp_port
	// END DEBUGGING
	// FcSwitchPortTypeTnpPort captures enum value "tnp_port"
	FcSwitchPortTypeTnpPort string = "tnp_port"

	// BEGIN DEBUGGING
	// fc_switch_port
	// FcSwitchPort
	// type
	// Type
	// none
	// END DEBUGGING
	// FcSwitchPortTypeNone captures enum value "none"
	FcSwitchPortTypeNone string = "none"
)

// prop value enum
func (m *FcSwitchPort) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcSwitchPortTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FcSwitchPort) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fc switch port based on the context it is used
func (m *FcSwitchPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachedDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchPort) contextValidateAttachedDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.AttachedDevice != nil {
		if err := m.AttachedDevice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attached_device")
			}
			return err
		}
	}

	return nil
}

func (m *FcSwitchPort) contextValidateSlot(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "slot", "body", m.Slot); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchPort) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchPort) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchPort) contextValidateWwpn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "wwpn", "body", m.Wwpn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitchPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitchPort) UnmarshalBinary(b []byte) error {
	var res FcSwitchPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcSwitchPortInlineAttachedDevice The Fibre Channel (FC) device attached to the FC switch port.
//
// swagger:model fc_switch_port_inline_attached_device
type FcSwitchPortInlineAttachedDevice struct {

	// The Fibre Channel port identifier of the attach device.
	//
	// Example: 0x011300
	// Read Only: true
	PortID *string `json:"port_id,omitempty"`

	// The world-wide port name (WWPN) of the attached device.
	//
	// Example: 50:0a:21:22:23:24:25:26
	// Read Only: true
	Wwpn *string `json:"wwpn,omitempty"`
}

// Validate validates this fc switch port inline attached device
func (m *FcSwitchPortInlineAttachedDevice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this fc switch port inline attached device based on the context it is used
func (m *FcSwitchPortInlineAttachedDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePortID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcSwitchPortInlineAttachedDevice) contextValidatePortID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "attached_device"+"."+"port_id", "body", m.PortID); err != nil {
		return err
	}

	return nil
}

func (m *FcSwitchPortInlineAttachedDevice) contextValidateWwpn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "attached_device"+"."+"wwpn", "body", m.Wwpn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcSwitchPortInlineAttachedDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcSwitchPortInlineAttachedDevice) UnmarshalBinary(b []byte) error {
	var res FcSwitchPortInlineAttachedDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
