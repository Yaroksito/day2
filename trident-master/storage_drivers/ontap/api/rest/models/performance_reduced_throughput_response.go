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

// PerformanceReducedThroughputResponse performance reduced throughput response
//
// swagger:model performance_reduced_throughput_response
type PerformanceReducedThroughputResponse struct {

	// links
	Links *PerformanceReducedThroughputResponseInlineLinks `json:"_links,omitempty"`

	// Number of records
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`

	// performance reduced throughput response inline records
	PerformanceReducedThroughputResponseInlineRecords []*PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem `json:"records,omitempty"`
}

// Validate validates this performance reduced throughput response
func (m *PerformanceReducedThroughputResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceReducedThroughputResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceReducedThroughputResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *PerformanceReducedThroughputResponse) validatePerformanceReducedThroughputResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceReducedThroughputResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.PerformanceReducedThroughputResponseInlineRecords); i++ {
		if swag.IsZero(m.PerformanceReducedThroughputResponseInlineRecords[i]) { // not required
			continue
		}

		if m.PerformanceReducedThroughputResponseInlineRecords[i] != nil {
			if err := m.PerformanceReducedThroughputResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this performance reduced throughput response based on the context it is used
func (m *PerformanceReducedThroughputResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerformanceReducedThroughputResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceReducedThroughputResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceReducedThroughputResponse) contextValidatePerformanceReducedThroughputResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PerformanceReducedThroughputResponseInlineRecords); i++ {

		if m.PerformanceReducedThroughputResponseInlineRecords[i] != nil {
			if err := m.PerformanceReducedThroughputResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (m *PerformanceReducedThroughputResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceReducedThroughputResponse) UnmarshalBinary(b []byte) error {
	var res PerformanceReducedThroughputResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceReducedThroughputResponseInlineLinks performance reduced throughput response inline links
//
// swagger:model performance_reduced_throughput_response_inline__links
type PerformanceReducedThroughputResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance reduced throughput response inline links
func (m *PerformanceReducedThroughputResponseInlineLinks) Validate(formats strfmt.Registry) error {
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

func (m *PerformanceReducedThroughputResponseInlineLinks) validateNext(formats strfmt.Registry) error {
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

func (m *PerformanceReducedThroughputResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this performance reduced throughput response inline links based on the context it is used
func (m *PerformanceReducedThroughputResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PerformanceReducedThroughputResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceReducedThroughputResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PerformanceReducedThroughputResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceReducedThroughputResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceReducedThroughputResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem Performance numbers, such as IOPS latency and throughput
//
// swagger:model performance_reduced_throughput_response_inline_records_inline_array_item
type PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem struct {

	// links
	Links *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Read Only: true
	// Enum: [PT15S PT4M PT30M PT2H P1D PT5M]
	Duration *string `json:"duration,omitempty"`

	// iops
	Iops *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineIops `json:"iops,omitempty"`

	// latency
	Latency *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLatency `json:"latency,omitempty"`

	// Any errors associated with the sample. For example, if the aggregation of data over multiple nodes fails then any of the partial errors might be returned, "ok" on success, or "error" on any internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_ delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: [ok error partial_no_data partial_no_response partial_other_error negative_delta not_found backfilled_data inconsistent_delta_time inconsistent_old_data partial_no_uuid]
	Status *string `json:"status,omitempty"`

	// throughput
	Throughput *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineThroughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25T11:20:13Z
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this performance reduced throughput response inline records inline array item
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatency(formats); err != nil {
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

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

var performanceReducedThroughputResponseInlineRecordsInlineArrayItemTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceReducedThroughputResponseInlineRecordsInlineArrayItemTypeDurationPropEnum = append(performanceReducedThroughputResponseInlineRecordsInlineArrayItemTypeDurationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT15S
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemDurationPT15S captures enum value "PT15S"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemDurationPT15S string = "PT15S"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT4M
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemDurationPT4M captures enum value "PT4M"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemDurationPT4M string = "PT4M"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT30M
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemDurationPT30M captures enum value "PT30M"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemDurationPT30M string = "PT30M"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT2H
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemDurationPT2H captures enum value "PT2H"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemDurationPT2H string = "PT2H"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// P1D
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemDurationP1D captures enum value "P1D"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemDurationP1D string = "P1D"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT5M
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemDurationPT5M captures enum value "PT5M"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemDurationPT5M string = "PT5M"
)

// prop value enum
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceReducedThroughputResponseInlineRecordsInlineArrayItemTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", *m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) validateIops(formats strfmt.Registry) error {
	if swag.IsZero(m.Iops) { // not required
		return nil
	}

	if m.Iops != nil {
		if err := m.Iops.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) validateLatency(formats strfmt.Registry) error {
	if swag.IsZero(m.Latency) { // not required
		return nil
	}

	if m.Latency != nil {
		if err := m.Latency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

var performanceReducedThroughputResponseInlineRecordsInlineArrayItemTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceReducedThroughputResponseInlineRecordsInlineArrayItemTypeStatusPropEnum = append(performanceReducedThroughputResponseInlineRecordsInlineArrayItemTypeStatusPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// ok
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusOk captures enum value "ok"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusOk string = "ok"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// error
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusError captures enum value "error"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusError string = "error"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_data
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusPartialNoData captures enum value "partial_no_data"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusPartialNoData string = "partial_no_data"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_response
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusPartialNoResponse string = "partial_no_response"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_other_error
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusPartialOtherError captures enum value "partial_other_error"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusPartialOtherError string = "partial_other_error"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// negative_delta
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusNegativeDelta captures enum value "negative_delta"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusNegativeDelta string = "negative_delta"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// not_found
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusNotFound captures enum value "not_found"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusNotFound string = "not_found"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// backfilled_data
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusBackfilledData captures enum value "backfilled_data"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusBackfilledData string = "backfilled_data"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// inconsistent_delta_time
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// inconsistent_old_data
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusInconsistentOldData string = "inconsistent_old_data"

	// BEGIN DEBUGGING
	// performance_reduced_throughput_response_inline_records_inline_array_item
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_uuid
	// END DEBUGGING
	// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemStatusPartialNoUUID string = "partial_no_uuid"
)

// prop value enum
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceReducedThroughputResponseInlineRecordsInlineArrayItemTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) validateThroughput(formats strfmt.Registry) error {
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

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance reduced throughput response inline records inline array item based on the context it is used
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

	if m.Iops != nil {
		if err := m.Iops.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) contextValidateLatency(ctx context.Context, formats strfmt.Registry) error {

	if m.Latency != nil {
		if err := m.Latency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res PerformanceReducedThroughputResponseInlineRecordsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineIops The rate of I/O operations observed at the storage object.
//
// swagger:model performance_reduced_throughput_response_inline_records_inline_array_item_inline_iops
type PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineIops struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other *int64 `json:"other,omitempty"`

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

// Validate validates this performance reduced throughput response inline records inline array item inline iops
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineIops) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance reduced throughput response inline records inline array item inline iops based on the context it is used
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineIops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineIops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineIops) UnmarshalBinary(b []byte) error {
	var res PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineIops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLatency The round trip latency in microseconds observed at the storage object.
//
// swagger:model performance_reduced_throughput_response_inline_records_inline_array_item_inline_latency
type PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLatency struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other *int64 `json:"other,omitempty"`

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

// Validate validates this performance reduced throughput response inline records inline array item inline latency
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLatency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance reduced throughput response inline records inline array item inline latency based on the context it is used
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLatency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLatency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLatency) UnmarshalBinary(b []byte) error {
	var res PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLatency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLinks performance reduced throughput response inline records inline array item inline links
//
// swagger:model performance_reduced_throughput_response_inline_records_inline_array_item_inline__links
type PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance reduced throughput response inline records inline array item inline links
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this performance reduced throughput response inline records inline array item inline links based on the context it is used
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineThroughput The rate of throughput bytes per second observed at the storage object.
//
// swagger:model performance_reduced_throughput_response_inline_records_inline_array_item_inline_throughput
type PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineThroughput struct {

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

// Validate validates this performance reduced throughput response inline records inline array item inline throughput
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance reduced throughput response inline records inline array item inline throughput based on the context it is used
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineThroughput) UnmarshalBinary(b []byte) error {
	var res PerformanceReducedThroughputResponseInlineRecordsInlineArrayItemInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
