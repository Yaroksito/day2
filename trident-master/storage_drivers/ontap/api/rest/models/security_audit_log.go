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

// SecurityAuditLog security audit log
//
// swagger:model security_audit_log
type SecurityAuditLog struct {

	// links
	Links *SecurityAuditLogInlineLinks `json:"_links,omitempty"`

	// This identifies the "application" by which the request was processed.
	//
	// Read Only: true
	// Enum: [internal console rsh telnet ssh ontapi http system]
	Application *string `json:"application,omitempty"`

	// This is the command ID for this request.
	// Each command received on a CLI session is assigned a command ID. This enables you to correlate a request and response.
	//
	// Read Only: true
	CommandID *string `json:"command_id,omitempty"`

	// Internal index for accessing records with same time/node. This is a 64 bit unsigned value.
	// Read Only: true
	Index *int64 `json:"index,omitempty"`

	// The request.
	// Read Only: true
	Input *string `json:"input,omitempty"`

	// This identifies the location of the remote user. This is an IP address or "console".
	// Read Only: true
	Location *string `json:"location,omitempty"`

	// This is an optional field that might contain "error" or "additional information" about the status of a command.
	// Read Only: true
	Message *string `json:"message,omitempty"`

	// node
	Node *SecurityAuditLogInlineNode `json:"node,omitempty"`

	// Set to "svm" when the request is on a data SVM; otherwise set to "cluster".
	// Enum: [svm cluster]
	Scope *string `json:"scope,omitempty"`

	// This is the session ID on which the request is received. Each SSH session is assigned a session ID.
	// Each http/ontapi/snmp request is assigned a unique session ID.
	//
	// Read Only: true
	SessionID *string `json:"session_id,omitempty"`

	// State of of this request.
	// Read Only: true
	// Enum: [pending success error]
	State *string `json:"state,omitempty"`

	// svm
	Svm *SecurityAuditLogInlineSvm `json:"svm,omitempty"`

	// Log entry timestamp. Valid in URL
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// Username of the remote user.
	// Read Only: true
	User *string `json:"user,omitempty"`
}

// Validate validates this security audit log
func (m *SecurityAuditLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAuditLog) validateLinks(formats strfmt.Registry) error {
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

var securityAuditLogTypeApplicationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["internal","console","rsh","telnet","ssh","ontapi","http","system"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityAuditLogTypeApplicationPropEnum = append(securityAuditLogTypeApplicationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// security_audit_log
	// SecurityAuditLog
	// application
	// Application
	// internal
	// END DEBUGGING
	// SecurityAuditLogApplicationInternal captures enum value "internal"
	SecurityAuditLogApplicationInternal string = "internal"

	// BEGIN DEBUGGING
	// security_audit_log
	// SecurityAuditLog
	// application
	// Application
	// console
	// END DEBUGGING
	// SecurityAuditLogApplicationConsole captures enum value "console"
	SecurityAuditLogApplicationConsole string = "console"

	// BEGIN DEBUGGING
	// security_audit_log
	// SecurityAuditLog
	// application
	// Application
	// rsh
	// END DEBUGGING
	// SecurityAuditLogApplicationRsh captures enum value "rsh"
	SecurityAuditLogApplicationRsh string = "rsh"

	// BEGIN DEBUGGING
	// security_audit_log
	// SecurityAuditLog
	// application
	// Application
	// telnet
	// END DEBUGGING
	// SecurityAuditLogApplicationTelnet captures enum value "telnet"
	SecurityAuditLogApplicationTelnet string = "telnet"

	// BEGIN DEBUGGING
	// security_audit_log
	// SecurityAuditLog
	// application
	// Application
	// ssh
	// END DEBUGGING
	// SecurityAuditLogApplicationSSH captures enum value "ssh"
	SecurityAuditLogApplicationSSH string = "ssh"

	// BEGIN DEBUGGING
	// security_audit_log
	// SecurityAuditLog
	// application
	// Application
	// ontapi
	// END DEBUGGING
	// SecurityAuditLogApplicationOntapi captures enum value "ontapi"
	SecurityAuditLogApplicationOntapi string = "ontapi"

	// BEGIN DEBUGGING
	// security_audit_log
	// SecurityAuditLog
	// application
	// Application
	// http
	// END DEBUGGING
	// SecurityAuditLogApplicationHTTP captures enum value "http"
	SecurityAuditLogApplicationHTTP string = "http"

	// BEGIN DEBUGGING
	// security_audit_log
	// SecurityAuditLog
	// application
	// Application
	// system
	// END DEBUGGING
	// SecurityAuditLogApplicationSystem captures enum value "system"
	SecurityAuditLogApplicationSystem string = "system"
)

// prop value enum
func (m *SecurityAuditLog) validateApplicationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityAuditLogTypeApplicationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecurityAuditLog) validateApplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Application) { // not required
		return nil
	}

	// value enum
	if err := m.validateApplicationEnum("application", "body", *m.Application); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAuditLog) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

var securityAuditLogTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["svm","cluster"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityAuditLogTypeScopePropEnum = append(securityAuditLogTypeScopePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// security_audit_log
	// SecurityAuditLog
	// scope
	// Scope
	// svm
	// END DEBUGGING
	// SecurityAuditLogScopeSvm captures enum value "svm"
	SecurityAuditLogScopeSvm string = "svm"

	// BEGIN DEBUGGING
	// security_audit_log
	// SecurityAuditLog
	// scope
	// Scope
	// cluster
	// END DEBUGGING
	// SecurityAuditLogScopeCluster captures enum value "cluster"
	SecurityAuditLogScopeCluster string = "cluster"
)

// prop value enum
func (m *SecurityAuditLog) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityAuditLogTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecurityAuditLog) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", *m.Scope); err != nil {
		return err
	}

	return nil
}

var securityAuditLogTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","success","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityAuditLogTypeStatePropEnum = append(securityAuditLogTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// security_audit_log
	// SecurityAuditLog
	// state
	// State
	// pending
	// END DEBUGGING
	// SecurityAuditLogStatePending captures enum value "pending"
	SecurityAuditLogStatePending string = "pending"

	// BEGIN DEBUGGING
	// security_audit_log
	// SecurityAuditLog
	// state
	// State
	// success
	// END DEBUGGING
	// SecurityAuditLogStateSuccess captures enum value "success"
	SecurityAuditLogStateSuccess string = "success"

	// BEGIN DEBUGGING
	// security_audit_log
	// SecurityAuditLog
	// state
	// State
	// error
	// END DEBUGGING
	// SecurityAuditLogStateError captures enum value "error"
	SecurityAuditLogStateError string = "error"
)

// prop value enum
func (m *SecurityAuditLog) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityAuditLogTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecurityAuditLog) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAuditLog) validateSvm(formats strfmt.Registry) error {
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

func (m *SecurityAuditLog) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this security audit log based on the context it is used
func (m *SecurityAuditLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommandID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIndex(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSessionID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAuditLog) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SecurityAuditLog) contextValidateApplication(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "application", "body", m.Application); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAuditLog) contextValidateCommandID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "command_id", "body", m.CommandID); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAuditLog) contextValidateIndex(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAuditLog) contextValidateInput(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "input", "body", m.Input); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAuditLog) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "location", "body", m.Location); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAuditLog) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAuditLog) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityAuditLog) contextValidateSessionID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "session_id", "body", m.SessionID); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAuditLog) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAuditLog) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SecurityAuditLog) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *SecurityAuditLog) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "user", "body", m.User); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAuditLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAuditLog) UnmarshalBinary(b []byte) error {
	var res SecurityAuditLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityAuditLogInlineLinks security audit log inline links
//
// swagger:model security_audit_log_inline__links
type SecurityAuditLogInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this security audit log inline links
func (m *SecurityAuditLogInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAuditLogInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this security audit log inline links based on the context it is used
func (m *SecurityAuditLogInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAuditLogInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SecurityAuditLogInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAuditLogInlineLinks) UnmarshalBinary(b []byte) error {
	var res SecurityAuditLogInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityAuditLogInlineNode Node where the audit message resides.
//
// swagger:model security_audit_log_inline_node
type SecurityAuditLogInlineNode struct {

	// links
	Links *SecurityAuditLogInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this security audit log inline node
func (m *SecurityAuditLogInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAuditLogInlineNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security audit log inline node based on the context it is used
func (m *SecurityAuditLogInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAuditLogInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAuditLogInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAuditLogInlineNode) UnmarshalBinary(b []byte) error {
	var res SecurityAuditLogInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityAuditLogInlineNodeInlineLinks security audit log inline node inline links
//
// swagger:model security_audit_log_inline_node_inline__links
type SecurityAuditLogInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this security audit log inline node inline links
func (m *SecurityAuditLogInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAuditLogInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security audit log inline node inline links based on the context it is used
func (m *SecurityAuditLogInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityAuditLogInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAuditLogInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAuditLogInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res SecurityAuditLogInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SecurityAuditLogInlineSvm This is the SVM through which the user connected.
//
// swagger:model security_audit_log_inline_svm
type SecurityAuditLogInlineSvm struct {

	// name
	Name *string `json:"name,omitempty"`
}

// Validate validates this security audit log inline svm
func (m *SecurityAuditLogInlineSvm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this security audit log inline svm based on the context it is used
func (m *SecurityAuditLogInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityAuditLogInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityAuditLogInlineSvm) UnmarshalBinary(b []byte) error {
	var res SecurityAuditLogInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
