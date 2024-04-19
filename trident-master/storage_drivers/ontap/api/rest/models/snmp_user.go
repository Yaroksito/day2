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

// SnmpUser An SNMP user can be an SNMPv1/SNMPv2c user or an SNMPv3 user. SNMPv1/SNMPv2c user is also called a "community" user. An SNMPv3 user, also called a User-based Security Model (USM) user, can be a local SNMPv3 user or a remote SNMPv3 user. A local SNMPv3 user can be used for querying ONTAP SNMP server over SNMPv3 and/or for sending SNMPv3 traps. The local SNMPv3 user used for sending SNMPv3 traps must be configured with the same authentication and privacy credentials on the traphost receiver as well. A remote SNMPv3 user is also configured on a remote switch and used by ONTAP SNMP client functionality to query the remote switch over SNMPv3. An SNMP user is scoped to its owning Storage Virtual Machine (SVM). Owning SVM could be a data SVM or the administrative SVM.
//
// swagger:model snmp_user
type SnmpUser struct {

	// links
	Links *SnmpUserInlineLinks `json:"_links,omitempty"`

	// Optional authentication method.
	// Example: usm
	// Enum: [community usm both]
	AuthenticationMethod *string `json:"authentication_method,omitempty"`

	// Optional comment text.
	// Example: This is a comment.
	// Max Length: 128
	// Min Length: 0
	Comment *string `json:"comment,omitempty"`

	// Optional SNMPv3 engine identifier. For a local SNMP user belonging to the administrative Storage Virtual Machine (SVM), the default value of this parameter is the SNMPv3 engine identifier for the administrative SVM. For a local SNMP user belonging to a data SVM, the default value of this parameter is the SNMPv3 engine identifier for that data SVM. For an SNMPv1/SNMPv2c community, this parameter should not be specified in "POST" method. For a remote switch SNMPv3 user, this parameter specifies the SNMPv3 engine identifier for the remote switch. This parameter can also optionally specify a custom engine identifier.
	// Example: 80000315055415ab26d4aae811ac4d005056bb792e
	EngineID *string `json:"engine_id,omitempty"`

	// SNMP user name.
	// Example: snmpv3user2
	// Max Length: 32
	Name *string `json:"name,omitempty"`

	// owner
	Owner *SnmpUserInlineOwner `json:"owner,omitempty"`

	// Set to "svm" for data Storage Virtual Machine (SVM) SNMP users and to "cluster" for administrative SVM SNMP users.
	// Example: svm
	// Read Only: true
	// Enum: [svm cluster]
	Scope *string `json:"scope,omitempty"`

	// snmpv3
	Snmpv3 *SnmpUserInlineSnmpv3 `json:"snmpv3,omitempty"`

	// Optional remote switch address. It can be an IPv4 address or an IPv6 address. A remote switch can be queried over SNMPv3 using ONTAP SNMP client functionality. Querying such a switch requires an SNMPv3 user (remote switch user) to be configured on the switch. Since ONTAP requires remote switch user's SNMPv3 credentials (to query it), this user must be configured in ONTAP as well. This parameter is specified when configuring such a user.
	// Example: 10.23.34.45
	SwitchAddress *string `json:"switch_address,omitempty"`
}

// Validate validates this snmp user
func (m *SnmpUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnmpv3(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpUser) validateLinks(formats strfmt.Registry) error {
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

var snmpUserTypeAuthenticationMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["community","usm","both"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snmpUserTypeAuthenticationMethodPropEnum = append(snmpUserTypeAuthenticationMethodPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// snmp_user
	// SnmpUser
	// authentication_method
	// AuthenticationMethod
	// community
	// END DEBUGGING
	// SnmpUserAuthenticationMethodCommunity captures enum value "community"
	SnmpUserAuthenticationMethodCommunity string = "community"

	// BEGIN DEBUGGING
	// snmp_user
	// SnmpUser
	// authentication_method
	// AuthenticationMethod
	// usm
	// END DEBUGGING
	// SnmpUserAuthenticationMethodUsm captures enum value "usm"
	SnmpUserAuthenticationMethodUsm string = "usm"

	// BEGIN DEBUGGING
	// snmp_user
	// SnmpUser
	// authentication_method
	// AuthenticationMethod
	// both
	// END DEBUGGING
	// SnmpUserAuthenticationMethodBoth captures enum value "both"
	SnmpUserAuthenticationMethodBoth string = "both"
)

// prop value enum
func (m *SnmpUser) validateAuthenticationMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snmpUserTypeAuthenticationMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnmpUser) validateAuthenticationMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationMethodEnum("authentication_method", "body", *m.AuthenticationMethod); err != nil {
		return err
	}

	return nil
}

func (m *SnmpUser) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("comment", "body", *m.Comment, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("comment", "body", *m.Comment, 128); err != nil {
		return err
	}

	return nil
}

func (m *SnmpUser) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", *m.Name, 32); err != nil {
		return err
	}

	return nil
}

func (m *SnmpUser) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

var snmpUserTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["svm","cluster"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snmpUserTypeScopePropEnum = append(snmpUserTypeScopePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// snmp_user
	// SnmpUser
	// scope
	// Scope
	// svm
	// END DEBUGGING
	// SnmpUserScopeSvm captures enum value "svm"
	SnmpUserScopeSvm string = "svm"

	// BEGIN DEBUGGING
	// snmp_user
	// SnmpUser
	// scope
	// Scope
	// cluster
	// END DEBUGGING
	// SnmpUserScopeCluster captures enum value "cluster"
	SnmpUserScopeCluster string = "cluster"
)

// prop value enum
func (m *SnmpUser) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snmpUserTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnmpUser) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", *m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *SnmpUser) validateSnmpv3(formats strfmt.Registry) error {
	if swag.IsZero(m.Snmpv3) { // not required
		return nil
	}

	if m.Snmpv3 != nil {
		if err := m.Snmpv3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snmpv3")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snmp user based on the context it is used
func (m *SnmpUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnmpv3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpUser) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SnmpUser) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {
		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpUser) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *SnmpUser) contextValidateSnmpv3(ctx context.Context, formats strfmt.Registry) error {

	if m.Snmpv3 != nil {
		if err := m.Snmpv3.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snmpv3")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnmpUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnmpUser) UnmarshalBinary(b []byte) error {
	var res SnmpUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnmpUserInlineLinks snmp user inline links
//
// swagger:model snmp_user_inline__links
type SnmpUserInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snmp user inline links
func (m *SnmpUserInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpUserInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this snmp user inline links based on the context it is used
func (m *SnmpUserInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpUserInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SnmpUserInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnmpUserInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnmpUserInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnmpUserInlineOwner Optional name and UUID of owning Storage Virtual Machine (SVM).
//
// swagger:model snmp_user_inline_owner
type SnmpUserInlineOwner struct {

	// links
	Links *SnmpUserInlineOwnerInlineLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this snmp user inline owner
func (m *SnmpUserInlineOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpUserInlineOwner) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snmp user inline owner based on the context it is used
func (m *SnmpUserInlineOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpUserInlineOwner) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnmpUserInlineOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnmpUserInlineOwner) UnmarshalBinary(b []byte) error {
	var res SnmpUserInlineOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnmpUserInlineOwnerInlineLinks snmp user inline owner inline links
//
// swagger:model snmp_user_inline_owner_inline__links
type SnmpUserInlineOwnerInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snmp user inline owner inline links
func (m *SnmpUserInlineOwnerInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpUserInlineOwnerInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snmp user inline owner inline links based on the context it is used
func (m *SnmpUserInlineOwnerInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpUserInlineOwnerInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnmpUserInlineOwnerInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnmpUserInlineOwnerInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnmpUserInlineOwnerInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnmpUserInlineSnmpv3 Optional parameter that can be specified only for an SNMPv3 user i.e. when 'authentication_method' is either 'usm' or 'both'. This parameter defines the SNMPv3 credentials for an SNMPv3 user.
//
// swagger:model snmp_user_inline_snmpv3
type SnmpUserInlineSnmpv3 struct {

	// links
	Links *SnmpUserInlineSnmpv3InlineLinks `json:"_links,omitempty"`

	// Authentication protocol password.
	// Example: humTdumt*@t0nAwa11
	// Min Length: 8
	AuthenticationPassword *string `json:"authentication_password,omitempty"`

	// Authentication protocol.
	// Example: sha2_256
	// Enum: [none md5 sha sha2_256]
	AuthenticationProtocol *string `json:"authentication_protocol,omitempty"`

	// Privacy protocol password.
	// Example: p@**GOandCLCt*200
	// Min Length: 8
	PrivacyPassword *string `json:"privacy_password,omitempty"`

	// Privacy protocol.
	// Example: aes128
	// Enum: [none des aes128]
	PrivacyProtocol *string `json:"privacy_protocol,omitempty"`
}

// Validate validates this snmp user inline snmpv3
func (m *SnmpUserInlineSnmpv3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivacyPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivacyProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpUserInlineSnmpv3) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snmpv3" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpUserInlineSnmpv3) validateAuthenticationPassword(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationPassword) { // not required
		return nil
	}

	if err := validate.MinLength("snmpv3"+"."+"authentication_password", "body", *m.AuthenticationPassword, 8); err != nil {
		return err
	}

	return nil
}

var snmpUserInlineSnmpv3TypeAuthenticationProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","md5","sha","sha2_256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snmpUserInlineSnmpv3TypeAuthenticationProtocolPropEnum = append(snmpUserInlineSnmpv3TypeAuthenticationProtocolPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// snmp_user_inline_snmpv3
	// SnmpUserInlineSnmpv3
	// authentication_protocol
	// AuthenticationProtocol
	// none
	// END DEBUGGING
	// SnmpUserInlineSnmpv3AuthenticationProtocolNone captures enum value "none"
	SnmpUserInlineSnmpv3AuthenticationProtocolNone string = "none"

	// BEGIN DEBUGGING
	// snmp_user_inline_snmpv3
	// SnmpUserInlineSnmpv3
	// authentication_protocol
	// AuthenticationProtocol
	// md5
	// END DEBUGGING
	// SnmpUserInlineSnmpv3AuthenticationProtocolMd5 captures enum value "md5"
	SnmpUserInlineSnmpv3AuthenticationProtocolMd5 string = "md5"

	// BEGIN DEBUGGING
	// snmp_user_inline_snmpv3
	// SnmpUserInlineSnmpv3
	// authentication_protocol
	// AuthenticationProtocol
	// sha
	// END DEBUGGING
	// SnmpUserInlineSnmpv3AuthenticationProtocolSha captures enum value "sha"
	SnmpUserInlineSnmpv3AuthenticationProtocolSha string = "sha"

	// BEGIN DEBUGGING
	// snmp_user_inline_snmpv3
	// SnmpUserInlineSnmpv3
	// authentication_protocol
	// AuthenticationProtocol
	// sha2_256
	// END DEBUGGING
	// SnmpUserInlineSnmpv3AuthenticationProtocolSha2256 captures enum value "sha2_256"
	SnmpUserInlineSnmpv3AuthenticationProtocolSha2256 string = "sha2_256"
)

// prop value enum
func (m *SnmpUserInlineSnmpv3) validateAuthenticationProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snmpUserInlineSnmpv3TypeAuthenticationProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnmpUserInlineSnmpv3) validateAuthenticationProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationProtocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationProtocolEnum("snmpv3"+"."+"authentication_protocol", "body", *m.AuthenticationProtocol); err != nil {
		return err
	}

	return nil
}

func (m *SnmpUserInlineSnmpv3) validatePrivacyPassword(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivacyPassword) { // not required
		return nil
	}

	if err := validate.MinLength("snmpv3"+"."+"privacy_password", "body", *m.PrivacyPassword, 8); err != nil {
		return err
	}

	return nil
}

var snmpUserInlineSnmpv3TypePrivacyProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","des","aes128"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snmpUserInlineSnmpv3TypePrivacyProtocolPropEnum = append(snmpUserInlineSnmpv3TypePrivacyProtocolPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// snmp_user_inline_snmpv3
	// SnmpUserInlineSnmpv3
	// privacy_protocol
	// PrivacyProtocol
	// none
	// END DEBUGGING
	// SnmpUserInlineSnmpv3PrivacyProtocolNone captures enum value "none"
	SnmpUserInlineSnmpv3PrivacyProtocolNone string = "none"

	// BEGIN DEBUGGING
	// snmp_user_inline_snmpv3
	// SnmpUserInlineSnmpv3
	// privacy_protocol
	// PrivacyProtocol
	// des
	// END DEBUGGING
	// SnmpUserInlineSnmpv3PrivacyProtocolDes captures enum value "des"
	SnmpUserInlineSnmpv3PrivacyProtocolDes string = "des"

	// BEGIN DEBUGGING
	// snmp_user_inline_snmpv3
	// SnmpUserInlineSnmpv3
	// privacy_protocol
	// PrivacyProtocol
	// aes128
	// END DEBUGGING
	// SnmpUserInlineSnmpv3PrivacyProtocolAes128 captures enum value "aes128"
	SnmpUserInlineSnmpv3PrivacyProtocolAes128 string = "aes128"
)

// prop value enum
func (m *SnmpUserInlineSnmpv3) validatePrivacyProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snmpUserInlineSnmpv3TypePrivacyProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnmpUserInlineSnmpv3) validatePrivacyProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivacyProtocol) { // not required
		return nil
	}

	// value enum
	if err := m.validatePrivacyProtocolEnum("snmpv3"+"."+"privacy_protocol", "body", *m.PrivacyProtocol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this snmp user inline snmpv3 based on the context it is used
func (m *SnmpUserInlineSnmpv3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpUserInlineSnmpv3) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snmpv3" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnmpUserInlineSnmpv3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnmpUserInlineSnmpv3) UnmarshalBinary(b []byte) error {
	var res SnmpUserInlineSnmpv3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnmpUserInlineSnmpv3InlineLinks snmp user inline snmpv3 inline links
//
// swagger:model snmp_user_inline_snmpv3_inline__links
type SnmpUserInlineSnmpv3InlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snmp user inline snmpv3 inline links
func (m *SnmpUserInlineSnmpv3InlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpUserInlineSnmpv3InlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snmpv3" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snmp user inline snmpv3 inline links based on the context it is used
func (m *SnmpUserInlineSnmpv3InlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpUserInlineSnmpv3InlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snmpv3" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnmpUserInlineSnmpv3InlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnmpUserInlineSnmpv3InlineLinks) UnmarshalBinary(b []byte) error {
	var res SnmpUserInlineSnmpv3InlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
