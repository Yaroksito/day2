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

// PortStatisticsDevice These are raw counters for the device associated with the Ethernet port.
//
// swagger:model port_statistics_device
type PortStatisticsDevice struct {

	// The number of link state changes from up to down seen on the device.
	// Example: 3
	LinkDownCountRaw *int64 `json:"link_down_count_raw,omitempty"`

	// receive raw
	ReceiveRaw *PortStatisticsDeviceInlineReceiveRaw `json:"receive_raw,omitempty"`

	// The timestamp when the device specific counters were collected.
	// Example: 2017-01-25T11:20:13Z
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// transmit raw
	TransmitRaw *PortStatisticsDeviceInlineTransmitRaw `json:"transmit_raw,omitempty"`
}

// Validate validates this port statistics device
func (m *PortStatisticsDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReceiveRaw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransmitRaw(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortStatisticsDevice) validateReceiveRaw(formats strfmt.Registry) error {
	if swag.IsZero(m.ReceiveRaw) { // not required
		return nil
	}

	if m.ReceiveRaw != nil {
		if err := m.ReceiveRaw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("receive_raw")
			}
			return err
		}
	}

	return nil
}

func (m *PortStatisticsDevice) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PortStatisticsDevice) validateTransmitRaw(formats strfmt.Registry) error {
	if swag.IsZero(m.TransmitRaw) { // not required
		return nil
	}

	if m.TransmitRaw != nil {
		if err := m.TransmitRaw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transmit_raw")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this port statistics device based on the context it is used
func (m *PortStatisticsDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReceiveRaw(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransmitRaw(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortStatisticsDevice) contextValidateReceiveRaw(ctx context.Context, formats strfmt.Registry) error {

	if m.ReceiveRaw != nil {
		if err := m.ReceiveRaw.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("receive_raw")
			}
			return err
		}
	}

	return nil
}

func (m *PortStatisticsDevice) contextValidateTransmitRaw(ctx context.Context, formats strfmt.Registry) error {

	if m.TransmitRaw != nil {
		if err := m.TransmitRaw.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transmit_raw")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortStatisticsDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortStatisticsDevice) UnmarshalBinary(b []byte) error {
	var res PortStatisticsDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortStatisticsDeviceInlineReceiveRaw Packet receive counters for the Ethernet port.
//
// swagger:model port_statistics_device_inline_receive_raw
type PortStatisticsDeviceInlineReceiveRaw struct {

	// Total number of discarded packets.
	// Example: 100
	Discards *int64 `json:"discards,omitempty"`

	// Number of packet errors.
	// Example: 200
	Errors *int64 `json:"errors,omitempty"`

	// Total packet count.
	// Example: 500
	Packets *int64 `json:"packets,omitempty"`
}

// Validate validates this port statistics device inline receive raw
func (m *PortStatisticsDeviceInlineReceiveRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this port statistics device inline receive raw based on context it is used
func (m *PortStatisticsDeviceInlineReceiveRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortStatisticsDeviceInlineReceiveRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortStatisticsDeviceInlineReceiveRaw) UnmarshalBinary(b []byte) error {
	var res PortStatisticsDeviceInlineReceiveRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortStatisticsDeviceInlineTransmitRaw Packet transmit counters for the Ethernet port.
//
// swagger:model port_statistics_device_inline_transmit_raw
type PortStatisticsDeviceInlineTransmitRaw struct {

	// Total number of discarded packets.
	// Example: 100
	Discards *int64 `json:"discards,omitempty"`

	// Number of packet errors.
	// Example: 200
	Errors *int64 `json:"errors,omitempty"`

	// Total packet count.
	// Example: 500
	Packets *int64 `json:"packets,omitempty"`
}

// Validate validates this port statistics device inline transmit raw
func (m *PortStatisticsDeviceInlineTransmitRaw) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this port statistics device inline transmit raw based on context it is used
func (m *PortStatisticsDeviceInlineTransmitRaw) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortStatisticsDeviceInlineTransmitRaw) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortStatisticsDeviceInlineTransmitRaw) UnmarshalBinary(b []byte) error {
	var res PortStatisticsDeviceInlineTransmitRaw
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}