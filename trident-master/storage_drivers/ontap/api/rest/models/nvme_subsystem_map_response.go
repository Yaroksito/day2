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

// NvmeSubsystemMapResponse nvme subsystem map response
//
// swagger:model nvme_subsystem_map_response
type NvmeSubsystemMapResponse struct {

	// links
	Links *NvmeSubsystemMapResponseInlineLinks `json:"_links,omitempty"`

	// The number of records in the response.
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`

	// nvme subsystem map response inline records
	NvmeSubsystemMapResponseInlineRecords []*NvmeSubsystemMap `json:"records,omitempty"`
}

// Validate validates this nvme subsystem map response
func (m *NvmeSubsystemMapResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvmeSubsystemMapResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *NvmeSubsystemMapResponse) validateNvmeSubsystemMapResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.NvmeSubsystemMapResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.NvmeSubsystemMapResponseInlineRecords); i++ {
		if swag.IsZero(m.NvmeSubsystemMapResponseInlineRecords[i]) { // not required
			continue
		}

		if m.NvmeSubsystemMapResponseInlineRecords[i] != nil {
			if err := m.NvmeSubsystemMapResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this nvme subsystem map response based on the context it is used
func (m *NvmeSubsystemMapResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNvmeSubsystemMapResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NvmeSubsystemMapResponse) contextValidateNvmeSubsystemMapResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NvmeSubsystemMapResponseInlineRecords); i++ {

		if m.NvmeSubsystemMapResponseInlineRecords[i] != nil {
			if err := m.NvmeSubsystemMapResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemMapResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemMapResponse) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemMapResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemMapResponseInlineLinks nvme subsystem map response inline links
//
// swagger:model nvme_subsystem_map_response_inline__links
type NvmeSubsystemMapResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this nvme subsystem map response inline links
func (m *NvmeSubsystemMapResponseInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapResponseInlineLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemMapResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this nvme subsystem map response inline links based on the context it is used
func (m *NvmeSubsystemMapResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemMapResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NvmeSubsystemMapResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemMapResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemMapResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}