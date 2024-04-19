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

// NvmeSubsystemMap An NVMe subsystem map is an association of an NVMe namespace with an NVMe subsystem. When an NVMe namespace is mapped to an NVMe subsystem, the NVMe subsystem's hosts are granted access to the NVMe namespace. The relationship between an NVMe subsystem and an NVMe namespace is one subsystem to many namespaces.
//
// swagger:model nvme_subsystem_map
type NvmeSubsystemMap struct {

	// links
	Links *NvmeSubsystemMapInlineLinks `json:"_links,omitempty"`

	// The Asymmetric Namespace Access Group ID (ANAGRPID) of the NVMe namespace.<br/>
	// The format for an ANAGRPID is 8 hexadecimal digits (zero-filled) followed by a lower case "h".<br/>
	// There is an added computational cost to retrieving this property's value. It is not populated for either a collection GET or an instance GET unless it is explicitly requested using the `fields` query parameter. See [`Requesting specific fields`](#Requesting_specific_fields) to learn more.
	//
	// Example: 00103050h
	// Read Only: true
	Anagrpid *string `json:"anagrpid,omitempty"`

	// namespace
	Namespace *NvmeSubsystemMapInlineNamespace `json:"namespace,omitempty"`

	// The NVMe namespace identifier. This is an identifier used by an NVMe controller to provide access to the NVMe namespace.<br/>
	// The format for an NVMe namespace identifier is 8 hexadecimal digits (zero-filled) followed by a lower case "h".
	//
	// Example: 00000001h
	// Read Only: true
	Nsid *string `json:"nsid,omitempty"`

	// subsystem
	Subsystem *NvmeSubsystemMapInlineSubsystem `json:"subsystem,omitempty"`

	// svm
	Svm *NvmeSubsystemMapInlineSvm `json:"svm,omitempty"`
}

// Validate validates this nvme subsystem map
func (m *NvmeSubsystemMap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubsystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMap) validateLinks(formats strfmt.Registry) error {
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

func (m *NvmeSubsystemMap) validateNamespace(formats strfmt.Registry) error {
	if swag.IsZero(m.Namespace) { // not required
		return nil
	}

	if m.Namespace != nil {
		if err := m.Namespace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemMap) validateSubsystem(formats strfmt.Registry) error {
	if swag.IsZero(m.Subsystem) { // not required
		return nil
	}

	if m.Subsystem != nil {
		if err := m.Subsystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemMap) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this nvme subsystem map based on the context it is used
func (m *NvmeSubsystemMap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAnagrpid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamespace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNsid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubsystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMap) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NvmeSubsystemMap) contextValidateAnagrpid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "anagrpid", "body", m.Anagrpid); err != nil {
		return err
	}

	return nil
}

func (m *NvmeSubsystemMap) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

	if m.Namespace != nil {
		if err := m.Namespace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemMap) contextValidateNsid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nsid", "body", m.Nsid); err != nil {
		return err
	}

	return nil
}

func (m *NvmeSubsystemMap) contextValidateSubsystem(ctx context.Context, formats strfmt.Registry) error {

	if m.Subsystem != nil {
		if err := m.Subsystem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemMap) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *NvmeSubsystemMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemMap) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemMapInlineLinks nvme subsystem map inline links
//
// swagger:model nvme_subsystem_map_inline__links
type NvmeSubsystemMapInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this nvme subsystem map inline links
func (m *NvmeSubsystemMapInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this nvme subsystem map inline links based on the context it is used
func (m *NvmeSubsystemMapInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NvmeSubsystemMapInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineLinks) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemMapInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemMapInlineNamespace The NVMe namespace to which the NVMe subsystem is mapped. Required in POST by supplying either the UUID, name, or both.
//
// swagger:model nvme_subsystem_map_inline_namespace
type NvmeSubsystemMapInlineNamespace struct {

	// links
	Links *NvmeSubsystemMapInlineNamespaceInlineLinks `json:"_links,omitempty"`

	// The fully qualified path name of the NVMe namespace composed from the volume name, qtree name, and file name of the NVMe namespace. Valid in POST.
	//
	// Example: /vol/vol1/namespace1
	Name *string `json:"name,omitempty"`

	// node
	Node *NvmeSubsystemMapInlineNamespaceInlineNode `json:"node,omitempty"`

	// The unique identifier of the NVMe namespace. Valid in POST.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this nvme subsystem map inline namespace
func (m *NvmeSubsystemMapInlineNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineNamespace) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemMapInlineNamespace) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "node")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem map inline namespace based on the context it is used
func (m *NvmeSubsystemMapInlineNamespace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineNamespace) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemMapInlineNamespace) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineNamespace) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemMapInlineNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemMapInlineNamespaceInlineLinks nvme subsystem map inline namespace inline links
//
// swagger:model nvme_subsystem_map_inline_namespace_inline__links
type NvmeSubsystemMapInlineNamespaceInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this nvme subsystem map inline namespace inline links
func (m *NvmeSubsystemMapInlineNamespaceInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineNamespaceInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem map inline namespace inline links based on the context it is used
func (m *NvmeSubsystemMapInlineNamespaceInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineNamespaceInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineNamespaceInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineNamespaceInlineLinks) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemMapInlineNamespaceInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemMapInlineNamespaceInlineNode nvme subsystem map inline namespace inline node
//
// swagger:model nvme_subsystem_map_inline_namespace_inline_node
type NvmeSubsystemMapInlineNamespaceInlineNode struct {

	// links
	Links *NvmeSubsystemMapInlineNamespaceInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this nvme subsystem map inline namespace inline node
func (m *NvmeSubsystemMapInlineNamespaceInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineNamespaceInlineNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem map inline namespace inline node based on the context it is used
func (m *NvmeSubsystemMapInlineNamespaceInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineNamespaceInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineNamespaceInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineNamespaceInlineNode) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemMapInlineNamespaceInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemMapInlineNamespaceInlineNodeInlineLinks nvme subsystem map inline namespace inline node inline links
//
// swagger:model nvme_subsystem_map_inline_namespace_inline_node_inline__links
type NvmeSubsystemMapInlineNamespaceInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this nvme subsystem map inline namespace inline node inline links
func (m *NvmeSubsystemMapInlineNamespaceInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineNamespaceInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem map inline namespace inline node inline links based on the context it is used
func (m *NvmeSubsystemMapInlineNamespaceInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineNamespaceInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespace" + "." + "node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineNamespaceInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineNamespaceInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemMapInlineNamespaceInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemMapInlineSubsystem The NVMe subsystem to which the NVMe namespace is mapped. Required in POST by supplying either `subsystem.uuid`, `subsystem.name`  or both.
//
// swagger:model nvme_subsystem_map_inline_subsystem
type NvmeSubsystemMapInlineSubsystem struct {

	// links
	Links *NvmeSubsystemMapInlineSubsystemInlineLinks `json:"_links,omitempty"`

	// The name of the NVMe subsystem.
	//
	// Example: subsystem1
	// Max Length: 96
	// Min Length: 1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the NVMe subsystem.
	//
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this nvme subsystem map inline subsystem
func (m *NvmeSubsystemMapInlineSubsystem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineSubsystem) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *NvmeSubsystemMapInlineSubsystem) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("subsystem"+"."+"name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("subsystem"+"."+"name", "body", *m.Name, 96); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nvme subsystem map inline subsystem based on the context it is used
func (m *NvmeSubsystemMapInlineSubsystem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineSubsystem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineSubsystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineSubsystem) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemMapInlineSubsystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemMapInlineSubsystemInlineLinks nvme subsystem map inline subsystem inline links
//
// swagger:model nvme_subsystem_map_inline_subsystem_inline__links
type NvmeSubsystemMapInlineSubsystemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this nvme subsystem map inline subsystem inline links
func (m *NvmeSubsystemMapInlineSubsystemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineSubsystemInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvme subsystem map inline subsystem inline links based on the context it is used
func (m *NvmeSubsystemMapInlineSubsystemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineSubsystemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineSubsystemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineSubsystemInlineLinks) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemMapInlineSubsystemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemMapInlineSvm nvme subsystem map inline svm
//
// swagger:model nvme_subsystem_map_inline_svm
type NvmeSubsystemMapInlineSvm struct {

	// links
	Links *NvmeSubsystemMapInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this nvme subsystem map inline svm
func (m *NvmeSubsystemMapInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this nvme subsystem map inline svm based on the context it is used
func (m *NvmeSubsystemMapInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NvmeSubsystemMapInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineSvm) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemMapInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NvmeSubsystemMapInlineSvmInlineLinks nvme subsystem map inline svm inline links
//
// swagger:model nvme_subsystem_map_inline_svm_inline__links
type NvmeSubsystemMapInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this nvme subsystem map inline svm inline links
func (m *NvmeSubsystemMapInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this nvme subsystem map inline svm inline links based on the context it is used
func (m *NvmeSubsystemMapInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeSubsystemMapInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NvmeSubsystemMapInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeSubsystemMapInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res NvmeSubsystemMapInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}