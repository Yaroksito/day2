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

// IscsiConnection An active iSCSI connection.
//
// swagger:model iscsi_connection
type IscsiConnection struct {

	// links
	Links *IscsiConnectionInlineLinks `json:"_links,omitempty"`

	// The iSCSI authentication type used to establish the connection.
	//
	// Read Only: true
	// Enum: [chap none]
	AuthenticationType *string `json:"authentication_type,omitempty"`

	// The identifier of the connection within the session.
	//
	// Read Only: true
	Cid *int64 `json:"cid,omitempty"`

	// initiator address
	InitiatorAddress *IscsiConnectionInlineInitiatorAddress `json:"initiator_address,omitempty"`

	// interface
	Interface *IscsiConnectionInlineInterface `json:"interface,omitempty"`
}

// Validate validates this iscsi connection
func (m *IscsiConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatorAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterface(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiConnection) validateLinks(formats strfmt.Registry) error {
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

var iscsiConnectionTypeAuthenticationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["chap","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iscsiConnectionTypeAuthenticationTypePropEnum = append(iscsiConnectionTypeAuthenticationTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// iscsi_connection
	// IscsiConnection
	// authentication_type
	// AuthenticationType
	// chap
	// END DEBUGGING
	// IscsiConnectionAuthenticationTypeChap captures enum value "chap"
	IscsiConnectionAuthenticationTypeChap string = "chap"

	// BEGIN DEBUGGING
	// iscsi_connection
	// IscsiConnection
	// authentication_type
	// AuthenticationType
	// none
	// END DEBUGGING
	// IscsiConnectionAuthenticationTypeNone captures enum value "none"
	IscsiConnectionAuthenticationTypeNone string = "none"
)

// prop value enum
func (m *IscsiConnection) validateAuthenticationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, iscsiConnectionTypeAuthenticationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IscsiConnection) validateAuthenticationType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationTypeEnum("authentication_type", "body", *m.AuthenticationType); err != nil {
		return err
	}

	return nil
}

func (m *IscsiConnection) validateInitiatorAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.InitiatorAddress) { // not required
		return nil
	}

	if m.InitiatorAddress != nil {
		if err := m.InitiatorAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiator_address")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiConnection) validateInterface(formats strfmt.Registry) error {
	if swag.IsZero(m.Interface) { // not required
		return nil
	}

	if m.Interface != nil {
		if err := m.Interface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this iscsi connection based on the context it is used
func (m *IscsiConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthenticationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitiatorAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiConnection) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IscsiConnection) contextValidateAuthenticationType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "authentication_type", "body", m.AuthenticationType); err != nil {
		return err
	}

	return nil
}

func (m *IscsiConnection) contextValidateCid(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *IscsiConnection) contextValidateInitiatorAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.InitiatorAddress != nil {
		if err := m.InitiatorAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiator_address")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiConnection) contextValidateInterface(ctx context.Context, formats strfmt.Registry) error {

	if m.Interface != nil {
		if err := m.Interface.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiConnection) UnmarshalBinary(b []byte) error {
	var res IscsiConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiConnectionInlineInitiatorAddress The TCP socket information for the initiator end of the connection. This is useful for network packet debugging.
//
// swagger:model iscsi_connection_inline_initiator_address
type IscsiConnectionInlineInitiatorAddress struct {

	// The TCP IPv4 or IPv6 address of the initiator end of the iSCSI connection.
	//
	// Example: 10.10.10.7
	// Read Only: true
	Address *string `json:"address,omitempty"`

	// The TCP port number of the initiator end of the iSCSI connection.
	//
	// Example: 55432
	// Read Only: true
	Port *int64 `json:"port,omitempty"`
}

// Validate validates this iscsi connection inline initiator address
func (m *IscsiConnectionInlineInitiatorAddress) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this iscsi connection inline initiator address based on the context it is used
func (m *IscsiConnectionInlineInitiatorAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiConnectionInlineInitiatorAddress) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "initiator_address"+"."+"address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *IscsiConnectionInlineInitiatorAddress) contextValidatePort(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "initiator_address"+"."+"port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiConnectionInlineInitiatorAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiConnectionInlineInitiatorAddress) UnmarshalBinary(b []byte) error {
	var res IscsiConnectionInlineInitiatorAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiConnectionInlineInterface The network interface information for the target end of the connection.
//
// swagger:model iscsi_connection_inline_interface
type IscsiConnectionInlineInterface struct {

	// links
	Links *IscsiConnectionInlineInterfaceInlineLinks `json:"_links,omitempty"`

	// ip
	IP *IscsiConnectionInlineInterfaceInlineIP `json:"ip,omitempty"`

	// The name of the interface.
	// Example: lif1
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// The UUID that uniquely identifies the interface.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this iscsi connection inline interface
func (m *IscsiConnectionInlineInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiConnectionInlineInterface) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiConnectionInlineInterface) validateIP(formats strfmt.Registry) error {
	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if m.IP != nil {
		if err := m.IP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "ip")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this iscsi connection inline interface based on the context it is used
func (m *IscsiConnectionInlineInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
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

func (m *IscsiConnectionInlineInterface) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiConnectionInlineInterface) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if m.IP != nil {
		if err := m.IP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "ip")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiConnectionInlineInterface) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "interface"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IscsiConnectionInlineInterface) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "interface"+"."+"uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiConnectionInlineInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiConnectionInlineInterface) UnmarshalBinary(b []byte) error {
	var res IscsiConnectionInlineInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiConnectionInlineInterfaceInlineIP The IP information. ONTAP only supports port 3260.
//
// swagger:model iscsi_connection_inline_interface_inline_ip
type IscsiConnectionInlineInterfaceInlineIP struct {

	// address
	Address *IPAddressReadonly `json:"address,omitempty"`

	// The TCP port number of the iSCSI access endpoint.
	// Example: 3260
	// Read Only: true
	// Maximum: 65536
	// Minimum: 1
	Port *int64 `json:"port,omitempty"`
}

// Validate validates this iscsi connection inline interface inline ip
func (m *IscsiConnectionInlineInterfaceInlineIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiConnectionInlineInterfaceInlineIP) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "ip" + "." + "address")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiConnectionInlineInterfaceInlineIP) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("interface"+"."+"ip"+"."+"port", "body", *m.Port, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("interface"+"."+"ip"+"."+"port", "body", *m.Port, 65536, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this iscsi connection inline interface inline ip based on the context it is used
func (m *IscsiConnectionInlineInterfaceInlineIP) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiConnectionInlineInterfaceInlineIP) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.Address != nil {
		if err := m.Address.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "ip" + "." + "address")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiConnectionInlineInterfaceInlineIP) contextValidatePort(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "interface"+"."+"ip"+"."+"port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiConnectionInlineInterfaceInlineIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiConnectionInlineInterfaceInlineIP) UnmarshalBinary(b []byte) error {
	var res IscsiConnectionInlineInterfaceInlineIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiConnectionInlineInterfaceInlineLinks iscsi connection inline interface inline links
//
// swagger:model iscsi_connection_inline_interface_inline__links
type IscsiConnectionInlineInterfaceInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this iscsi connection inline interface inline links
func (m *IscsiConnectionInlineInterfaceInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiConnectionInlineInterfaceInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this iscsi connection inline interface inline links based on the context it is used
func (m *IscsiConnectionInlineInterfaceInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiConnectionInlineInterfaceInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiConnectionInlineInterfaceInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiConnectionInlineInterfaceInlineLinks) UnmarshalBinary(b []byte) error {
	var res IscsiConnectionInlineInterfaceInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiConnectionInlineLinks iscsi connection inline links
//
// swagger:model iscsi_connection_inline__links
type IscsiConnectionInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this iscsi connection inline links
func (m *IscsiConnectionInlineLinks) Validate(formats strfmt.Registry) error {
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

func (m *IscsiConnectionInlineLinks) validateNext(formats strfmt.Registry) error {
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

func (m *IscsiConnectionInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this iscsi connection inline links based on the context it is used
func (m *IscsiConnectionInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *IscsiConnectionInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IscsiConnectionInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IscsiConnectionInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiConnectionInlineLinks) UnmarshalBinary(b []byte) error {
	var res IscsiConnectionInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}