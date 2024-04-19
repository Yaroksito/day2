// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConsistencyGroupNvmeHost The NVMe host provisioned to access NVMe namespaces mapped to a subsystem.
//
// swagger:model consistency_group_nvme_host
type ConsistencyGroupNvmeHost struct {

	// The NVMe qualified name (NQN) used to identify the NVMe storage target. Not allowed in POST when the `records` property is used.
	//
	// Example: nqn.1992-01.example.com:string
	Nqn *string `json:"nqn,omitempty"`
}

// Validate validates this consistency group nvme host
func (m *ConsistencyGroupNvmeHost) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this consistency group nvme host based on context it is used
func (m *ConsistencyGroupNvmeHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupNvmeHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupNvmeHost) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupNvmeHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
