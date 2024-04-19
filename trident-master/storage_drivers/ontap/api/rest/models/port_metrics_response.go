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

// PortMetricsResponse port metrics response
//
// swagger:model port_metrics_response
type PortMetricsResponse struct {

	// links
	Links *PortMetricsResponseInlineLinks `json:"_links,omitempty"`

	// Number of records
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`

	// port metrics response inline records
	PortMetricsResponseInlineRecords []*PortMetricsResponseInlineRecordsInlineArrayItem `json:"records,omitempty"`
}

// Validate validates this port metrics response
func (m *PortMetricsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortMetricsResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortMetricsResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *PortMetricsResponse) validatePortMetricsResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.PortMetricsResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.PortMetricsResponseInlineRecords); i++ {
		if swag.IsZero(m.PortMetricsResponseInlineRecords[i]) { // not required
			continue
		}

		if m.PortMetricsResponseInlineRecords[i] != nil {
			if err := m.PortMetricsResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this port metrics response based on the context it is used
func (m *PortMetricsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePortMetricsResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortMetricsResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortMetricsResponse) contextValidatePortMetricsResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PortMetricsResponseInlineRecords); i++ {

		if m.PortMetricsResponseInlineRecords[i] != nil {
			if err := m.PortMetricsResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (m *PortMetricsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortMetricsResponse) UnmarshalBinary(b []byte) error {
	var res PortMetricsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortMetricsResponseInlineLinks port metrics response inline links
//
// swagger:model port_metrics_response_inline__links
type PortMetricsResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this port metrics response inline links
func (m *PortMetricsResponseInlineLinks) Validate(formats strfmt.Registry) error {
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

func (m *PortMetricsResponseInlineLinks) validateNext(formats strfmt.Registry) error {
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

func (m *PortMetricsResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this port metrics response inline links based on the context it is used
func (m *PortMetricsResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PortMetricsResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortMetricsResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PortMetricsResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortMetricsResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res PortMetricsResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortMetricsResponseInlineRecordsInlineArrayItem Throughput performance for the Ethernet port.
//
// swagger:model port_metrics_response_inline_records_inline_array_item
type PortMetricsResponseInlineRecordsInlineArrayItem struct {

	// links
	Links *PortMetricsResponseInlineRecordsInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Enum: [PT15S PT4M PT30M PT2H P1D PT5M]
	Duration *string `json:"duration,omitempty"`

	// Errors associated with the sample. For example, if the aggregation of data over multiple nodes fails, then any partial errors might return "ok" on success or "error" on an internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "inconsistent_delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Enum: [ok error partial_no_data partial_no_uuid partial_no_response partial_other_error negative_delta backfilled_data inconsistent_delta_time inconsistent_old_data]
	Status *string `json:"status,omitempty"`

	// throughput
	Throughput *PortMetricsResponseInlineRecordsInlineArrayItemInlineThroughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25T11:20:13Z
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// Port UUID
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this port metrics response inline records inline array item
func (m *PortMetricsResponseInlineRecordsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughput(formats); err != nil {
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

func (m *PortMetricsResponseInlineRecordsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

var portMetricsResponseInlineRecordsInlineArrayItemTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portMetricsResponseInlineRecordsInlineArrayItemTypeDurationPropEnum = append(portMetricsResponseInlineRecordsInlineArrayItemTypeDurationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT15S
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemDurationPT15S captures enum value "PT15S"
	PortMetricsResponseInlineRecordsInlineArrayItemDurationPT15S string = "PT15S"

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT4M
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemDurationPT4M captures enum value "PT4M"
	PortMetricsResponseInlineRecordsInlineArrayItemDurationPT4M string = "PT4M"

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT30M
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemDurationPT30M captures enum value "PT30M"
	PortMetricsResponseInlineRecordsInlineArrayItemDurationPT30M string = "PT30M"

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT2H
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemDurationPT2H captures enum value "PT2H"
	PortMetricsResponseInlineRecordsInlineArrayItemDurationPT2H string = "PT2H"

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// P1D
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemDurationP1D captures enum value "P1D"
	PortMetricsResponseInlineRecordsInlineArrayItemDurationP1D string = "P1D"

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT5M
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemDurationPT5M captures enum value "PT5M"
	PortMetricsResponseInlineRecordsInlineArrayItemDurationPT5M string = "PT5M"
)

// prop value enum
func (m *PortMetricsResponseInlineRecordsInlineArrayItem) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, portMetricsResponseInlineRecordsInlineArrayItemTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PortMetricsResponseInlineRecordsInlineArrayItem) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", *m.Duration); err != nil {
		return err
	}

	return nil
}

var portMetricsResponseInlineRecordsInlineArrayItemTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		portMetricsResponseInlineRecordsInlineArrayItemTypeStatusPropEnum = append(portMetricsResponseInlineRecordsInlineArrayItemTypeStatusPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// ok
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemStatusOk captures enum value "ok"
	PortMetricsResponseInlineRecordsInlineArrayItemStatusOk string = "ok"

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// error
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemStatusError captures enum value "error"
	PortMetricsResponseInlineRecordsInlineArrayItemStatusError string = "error"

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_data
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoData captures enum value "partial_no_data"
	PortMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoData string = "partial_no_data"

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_uuid
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoUUID captures enum value "partial_no_uuid"
	PortMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoUUID string = "partial_no_uuid"

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_response
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoResponse captures enum value "partial_no_response"
	PortMetricsResponseInlineRecordsInlineArrayItemStatusPartialNoResponse string = "partial_no_response"

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_other_error
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemStatusPartialOtherError captures enum value "partial_other_error"
	PortMetricsResponseInlineRecordsInlineArrayItemStatusPartialOtherError string = "partial_other_error"

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// negative_delta
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemStatusNegativeDelta captures enum value "negative_delta"
	PortMetricsResponseInlineRecordsInlineArrayItemStatusNegativeDelta string = "negative_delta"

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// backfilled_data
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemStatusBackfilledData captures enum value "backfilled_data"
	PortMetricsResponseInlineRecordsInlineArrayItemStatusBackfilledData string = "backfilled_data"

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// inconsistent_delta_time
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PortMetricsResponseInlineRecordsInlineArrayItemStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN DEBUGGING
	// port_metrics_response_inline_records_inline_array_item
	// PortMetricsResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// inconsistent_old_data
	// END DEBUGGING
	// PortMetricsResponseInlineRecordsInlineArrayItemStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PortMetricsResponseInlineRecordsInlineArrayItemStatusInconsistentOldData string = "inconsistent_old_data"
)

// prop value enum
func (m *PortMetricsResponseInlineRecordsInlineArrayItem) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, portMetricsResponseInlineRecordsInlineArrayItemTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PortMetricsResponseInlineRecordsInlineArrayItem) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PortMetricsResponseInlineRecordsInlineArrayItem) validateThroughput(formats strfmt.Registry) error {
	if swag.IsZero(m.Throughput) { // not required
		return nil
	}

	if m.Throughput != nil {
		if err := m.Throughput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PortMetricsResponseInlineRecordsInlineArrayItem) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this port metrics response inline records inline array item based on the context it is used
func (m *PortMetricsResponseInlineRecordsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
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

func (m *PortMetricsResponseInlineRecordsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PortMetricsResponseInlineRecordsInlineArrayItem) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

	if m.Throughput != nil {
		if err := m.Throughput.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PortMetricsResponseInlineRecordsInlineArrayItem) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PortMetricsResponseInlineRecordsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortMetricsResponseInlineRecordsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res PortMetricsResponseInlineRecordsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortMetricsResponseInlineRecordsInlineArrayItemInlineLinks port metrics response inline records inline array item inline links
//
// swagger:model port_metrics_response_inline_records_inline_array_item_inline__links
type PortMetricsResponseInlineRecordsInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this port metrics response inline records inline array item inline links
func (m *PortMetricsResponseInlineRecordsInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortMetricsResponseInlineRecordsInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this port metrics response inline records inline array item inline links based on the context it is used
func (m *PortMetricsResponseInlineRecordsInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortMetricsResponseInlineRecordsInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PortMetricsResponseInlineRecordsInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortMetricsResponseInlineRecordsInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res PortMetricsResponseInlineRecordsInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PortMetricsResponseInlineRecordsInlineArrayItemInlineThroughput The rate of throughput bytes per second observed at the interface.
//
// swagger:model port_metrics_response_inline_records_inline_array_item_inline_throughput
type PortMetricsResponseInlineRecordsInlineArrayItemInlineThroughput struct {

	// Performance metric for read I/O operations.
	// Example: 200
	Read *int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total *int64 `json:"total,omitempty"`

	// Peformance metric for write I/O operations.
	// Example: 100
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this port metrics response inline records inline array item inline throughput
func (m *PortMetricsResponseInlineRecordsInlineArrayItemInlineThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this port metrics response inline records inline array item inline throughput based on context it is used
func (m *PortMetricsResponseInlineRecordsInlineArrayItemInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortMetricsResponseInlineRecordsInlineArrayItemInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortMetricsResponseInlineRecordsInlineArrayItemInlineThroughput) UnmarshalBinary(b []byte) error {
	var res PortMetricsResponseInlineRecordsInlineArrayItemInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}