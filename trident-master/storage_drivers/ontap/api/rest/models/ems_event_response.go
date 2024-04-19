// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EmsEventResponse ems event response
//
// swagger:model ems_event_response
type EmsEventResponse struct {

	// links
	Links *EmsEventResponseInlineLinks `json:"_links,omitempty"`

	// ems event response inline records
	EmsEventResponseInlineRecords []*EmsEventResponseInlineRecordsInlineArrayItem `json:"records,omitempty"`

	// Number of records
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`
}

// Validate validates this ems event response
func (m *EmsEventResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmsEventResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *EmsEventResponse) validateEmsEventResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.EmsEventResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.EmsEventResponseInlineRecords); i++ {
		if swag.IsZero(m.EmsEventResponseInlineRecords[i]) { // not required
			continue
		}

		if m.EmsEventResponseInlineRecords[i] != nil {
			if err := m.EmsEventResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ems event response based on the context it is used
func (m *EmsEventResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmsEventResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EmsEventResponse) contextValidateEmsEventResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EmsEventResponseInlineRecords); i++ {

		if m.EmsEventResponseInlineRecords[i] != nil {
			if err := m.EmsEventResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (m *EmsEventResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponse) UnmarshalBinary(b []byte) error {
	var res EmsEventResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseInlineLinks ems event response inline links
//
// swagger:model ems_event_response_inline__links
type EmsEventResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems event response inline links
func (m *EmsEventResponseInlineLinks) Validate(formats strfmt.Registry) error {
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

func (m *EmsEventResponseInlineLinks) validateNext(formats strfmt.Registry) error {
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

func (m *EmsEventResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ems event response inline links based on the context it is used
func (m *EmsEventResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *EmsEventResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EmsEventResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EmsEventResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseInlineRecordsInlineArrayItem ems event response inline records inline array item
//
// swagger:model ems_event_response_inline_records_inline_array_item
type EmsEventResponseInlineRecordsInlineArrayItem struct {

	// links
	Links *EmsEventResponseInlineRecordsInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// Index of the event. Returned by default.
	// Example: 1
	// Read Only: true
	Index *int64 `json:"index,omitempty"`

	// A formatted text string populated with parameter details. Returned by default.
	LogMessage *string `json:"log_message,omitempty"`

	// message
	Message *EmsEventResponseInlineRecordsInlineArrayItemInlineMessage `json:"message,omitempty"`

	// node
	Node *EmsEventResponseInlineRecordsInlineArrayItemInlineNode `json:"node,omitempty"`

	// A list of parameters provided with the EMS event.
	// Read Only: true
	Parameters []*EmsEventResponseRecordsItems0ParametersItems0 `json:"parameters"`

	// Source
	// Read Only: true
	Source *string `json:"source,omitempty"`

	// Timestamp of the event. Returned by default.
	// Read Only: true
	// Format: date-time
	Time *strfmt.DateTime `json:"time,omitempty"`
}

// Validate validates this ems event response inline records inline array item
func (m *EmsEventResponseInlineRecordsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

func (m *EmsEventResponseInlineRecordsInlineArrayItem) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if m.Message != nil {
		if err := m.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItem) validateNode(formats strfmt.Registry) error {
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

func (m *EmsEventResponseInlineRecordsInlineArrayItem) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItem) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ems event response inline records inline array item based on the context it is used
func (m *EmsEventResponseInlineRecordsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIndex(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EmsEventResponseInlineRecordsInlineArrayItem) contextValidateIndex(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItem) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if m.Message != nil {
		if err := m.Message.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItem) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EmsEventResponseInlineRecordsInlineArrayItem) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "parameters", "body", []*EmsEventResponseRecordsItems0ParametersItems0(m.Parameters)); err != nil {
		return err
	}

	for i := 0; i < len(m.Parameters); i++ {

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItem) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItem) contextValidateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventResponseInlineRecordsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseInlineRecordsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseInlineRecordsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseInlineRecordsInlineArrayItemInlineLinks ems event response inline records inline array item inline links
//
// swagger:model ems_event_response_inline_records_inline_array_item_inline__links
type EmsEventResponseInlineRecordsInlineArrayItemInlineLinks struct {

	// self
	Self *EmsEventResponseInlineRecordsInlineArrayItemInlineLinksInlineSelf `json:"self,omitempty"`
}

// Validate validates this ems event response inline records inline array item inline links
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ems event response inline records inline array item inline links based on the context it is used
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseInlineRecordsInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseInlineRecordsInlineArrayItemInlineLinksInlineSelf ems event response inline records inline array item inline links inline self
//
// swagger:model ems_event_response_inline_records_inline_array_item_inline__links_inline_self
type EmsEventResponseInlineRecordsInlineArrayItemInlineLinksInlineSelf struct {

	// href
	// Example: /api/resourcelink
	// Read Only: true
	Href *string `json:"href,omitempty"`
}

// Validate validates this ems event response inline records inline array item inline links inline self
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineLinksInlineSelf) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this ems event response inline records inline array item inline links inline self based on the context it is used
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineLinksInlineSelf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHref(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineLinksInlineSelf) contextValidateHref(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_links"+"."+"self"+"."+"href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineLinksInlineSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineLinksInlineSelf) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseInlineRecordsInlineArrayItemInlineLinksInlineSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseInlineRecordsInlineArrayItemInlineMessage ems event response inline records inline array item inline message
//
// swagger:model ems_event_response_inline_records_inline_array_item_inline_message
type EmsEventResponseInlineRecordsInlineArrayItemInlineMessage struct {

	// links
	Links *EmsEventResponseInlineRecordsInlineArrayItemInlineMessageInlineLinks `json:"_links,omitempty"`

	// Message name of the event. Returned by default.
	// Example: callhome.spares.low
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// Severity of the event. Returned by default.
	// Example: emergency
	// Read Only: true
	// Enum: [emergency alert error notice informational debug]
	Severity *string `json:"severity,omitempty"`
}

// Validate validates this ems event response inline records inline array item inline message
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessage) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

var emsEventResponseInlineRecordsInlineArrayItemInlineMessageTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["emergency","alert","error","notice","informational","debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emsEventResponseInlineRecordsInlineArrayItemInlineMessageTypeSeverityPropEnum = append(emsEventResponseInlineRecordsInlineArrayItemInlineMessageTypeSeverityPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ems_event_response_inline_records_inline_array_item_inline_message
	// EmsEventResponseInlineRecordsInlineArrayItemInlineMessage
	// severity
	// Severity
	// emergency
	// END DEBUGGING
	// EmsEventResponseInlineRecordsInlineArrayItemInlineMessageSeverityEmergency captures enum value "emergency"
	EmsEventResponseInlineRecordsInlineArrayItemInlineMessageSeverityEmergency string = "emergency"

	// BEGIN DEBUGGING
	// ems_event_response_inline_records_inline_array_item_inline_message
	// EmsEventResponseInlineRecordsInlineArrayItemInlineMessage
	// severity
	// Severity
	// alert
	// END DEBUGGING
	// EmsEventResponseInlineRecordsInlineArrayItemInlineMessageSeverityAlert captures enum value "alert"
	EmsEventResponseInlineRecordsInlineArrayItemInlineMessageSeverityAlert string = "alert"

	// BEGIN DEBUGGING
	// ems_event_response_inline_records_inline_array_item_inline_message
	// EmsEventResponseInlineRecordsInlineArrayItemInlineMessage
	// severity
	// Severity
	// error
	// END DEBUGGING
	// EmsEventResponseInlineRecordsInlineArrayItemInlineMessageSeverityError captures enum value "error"
	EmsEventResponseInlineRecordsInlineArrayItemInlineMessageSeverityError string = "error"

	// BEGIN DEBUGGING
	// ems_event_response_inline_records_inline_array_item_inline_message
	// EmsEventResponseInlineRecordsInlineArrayItemInlineMessage
	// severity
	// Severity
	// notice
	// END DEBUGGING
	// EmsEventResponseInlineRecordsInlineArrayItemInlineMessageSeverityNotice captures enum value "notice"
	EmsEventResponseInlineRecordsInlineArrayItemInlineMessageSeverityNotice string = "notice"

	// BEGIN DEBUGGING
	// ems_event_response_inline_records_inline_array_item_inline_message
	// EmsEventResponseInlineRecordsInlineArrayItemInlineMessage
	// severity
	// Severity
	// informational
	// END DEBUGGING
	// EmsEventResponseInlineRecordsInlineArrayItemInlineMessageSeverityInformational captures enum value "informational"
	EmsEventResponseInlineRecordsInlineArrayItemInlineMessageSeverityInformational string = "informational"

	// BEGIN DEBUGGING
	// ems_event_response_inline_records_inline_array_item_inline_message
	// EmsEventResponseInlineRecordsInlineArrayItemInlineMessage
	// severity
	// Severity
	// debug
	// END DEBUGGING
	// EmsEventResponseInlineRecordsInlineArrayItemInlineMessageSeverityDebug captures enum value "debug"
	EmsEventResponseInlineRecordsInlineArrayItemInlineMessageSeverityDebug string = "debug"
)

// prop value enum
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessage) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emsEventResponseInlineRecordsInlineArrayItemInlineMessageTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessage) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("message"+"."+"severity", "body", *m.Severity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ems event response inline records inline array item inline message based on the context it is used
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeverity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessage) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessage) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessage) contextValidateSeverity(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message"+"."+"severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessage) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseInlineRecordsInlineArrayItemInlineMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseInlineRecordsInlineArrayItemInlineMessageInlineLinks ems event response inline records inline array item inline message inline links
//
// swagger:model ems_event_response_inline_records_inline_array_item_inline_message_inline__links
type EmsEventResponseInlineRecordsInlineArrayItemInlineMessageInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems event response inline records inline array item inline message inline links
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessageInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessageInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems event response inline records inline array item inline message inline links based on the context it is used
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessageInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessageInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessageInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineMessageInlineLinks) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseInlineRecordsInlineArrayItemInlineMessageInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseInlineRecordsInlineArrayItemInlineNode ems event response inline records inline array item inline node
//
// swagger:model ems_event_response_inline_records_inline_array_item_inline_node
type EmsEventResponseInlineRecordsInlineArrayItemInlineNode struct {

	// links
	Links *EmsEventResponseInlineRecordsInlineArrayItemInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this ems event response inline records inline array item inline node
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this ems event response inline records inline array item inline node based on the context it is used
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineNode) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseInlineRecordsInlineArrayItemInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseInlineRecordsInlineArrayItemInlineNodeInlineLinks ems event response inline records inline array item inline node inline links
//
// swagger:model ems_event_response_inline_records_inline_array_item_inline_node_inline__links
type EmsEventResponseInlineRecordsInlineArrayItemInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this ems event response inline records inline array item inline node inline links
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this ems event response inline records inline array item inline node inline links based on the context it is used
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseInlineRecordsInlineArrayItemInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseInlineRecordsInlineArrayItemInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsEventResponseRecordsItems0ParametersItems0 ems event response records items0 parameters items0
//
// swagger:model EmsEventResponseRecordsItems0ParametersItems0
type EmsEventResponseRecordsItems0ParametersItems0 struct {

	// Name of parameter
	// Example: numOps
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// Value of parameter
	// Example: 123
	// Read Only: true
	Value *string `json:"value,omitempty"`
}

// Validate validates this ems event response records items0 parameters items0
func (m *EmsEventResponseRecordsItems0ParametersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this ems event response records items0 parameters items0 based on the context it is used
func (m *EmsEventResponseRecordsItems0ParametersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsEventResponseRecordsItems0ParametersItems0) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EmsEventResponseRecordsItems0ParametersItems0) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0ParametersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsEventResponseRecordsItems0ParametersItems0) UnmarshalBinary(b []byte) error {
	var res EmsEventResponseRecordsItems0ParametersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}