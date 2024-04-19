// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DetailedStatusCodeMessage detailed status code message
//
// swagger:model detailed_status_code_message
type DetailedStatusCodeMessage struct {

	// Code corresponding to the import status failure.
	//
	// Example: 6684732
	Code *string `json:"code,omitempty"`

	// Detailed description of the import status.
	//
	Message *string `json:"message,omitempty"`
}

// Validate validates this detailed status code message
func (m *DetailedStatusCodeMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this detailed status code message based on context it is used
func (m *DetailedStatusCodeMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DetailedStatusCodeMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailedStatusCodeMessage) UnmarshalBinary(b []byte) error {
	var res DetailedStatusCodeMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
