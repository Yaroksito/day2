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

// SvmMigrationVolumePlacement Volume selection information
//
// swagger:model svm_migration_volume_placement
type SvmMigrationVolumePlacement struct {

	// Optional property used to specify the list of desired aggregates to use for volume creation in the destination.
	SvmMigrationVolumePlacementInlineAggregates []*SvmMigrationVolumePlacementInlineAggregatesInlineArrayItem `json:"aggregates,omitempty"`

	// Optional property used to specify the list of desired volume-aggregate pairs in the destination.
	SvmMigrationVolumePlacementInlineVolumeAggregatePairs []*SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItem `json:"volume_aggregate_pairs,omitempty"`
}

// Validate validates this svm migration volume placement
func (m *SvmMigrationVolumePlacement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSvmMigrationVolumePlacementInlineAggregates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvmMigrationVolumePlacementInlineVolumeAggregatePairs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacement) validateSvmMigrationVolumePlacementInlineAggregates(formats strfmt.Registry) error {
	if swag.IsZero(m.SvmMigrationVolumePlacementInlineAggregates) { // not required
		return nil
	}

	for i := 0; i < len(m.SvmMigrationVolumePlacementInlineAggregates); i++ {
		if swag.IsZero(m.SvmMigrationVolumePlacementInlineAggregates[i]) { // not required
			continue
		}

		if m.SvmMigrationVolumePlacementInlineAggregates[i] != nil {
			if err := m.SvmMigrationVolumePlacementInlineAggregates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aggregates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SvmMigrationVolumePlacement) validateSvmMigrationVolumePlacementInlineVolumeAggregatePairs(formats strfmt.Registry) error {
	if swag.IsZero(m.SvmMigrationVolumePlacementInlineVolumeAggregatePairs) { // not required
		return nil
	}

	for i := 0; i < len(m.SvmMigrationVolumePlacementInlineVolumeAggregatePairs); i++ {
		if swag.IsZero(m.SvmMigrationVolumePlacementInlineVolumeAggregatePairs[i]) { // not required
			continue
		}

		if m.SvmMigrationVolumePlacementInlineVolumeAggregatePairs[i] != nil {
			if err := m.SvmMigrationVolumePlacementInlineVolumeAggregatePairs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volume_aggregate_pairs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this svm migration volume placement based on the context it is used
func (m *SvmMigrationVolumePlacement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSvmMigrationVolumePlacementInlineAggregates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvmMigrationVolumePlacementInlineVolumeAggregatePairs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacement) contextValidateSvmMigrationVolumePlacementInlineAggregates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SvmMigrationVolumePlacementInlineAggregates); i++ {

		if m.SvmMigrationVolumePlacementInlineAggregates[i] != nil {
			if err := m.SvmMigrationVolumePlacementInlineAggregates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aggregates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SvmMigrationVolumePlacement) contextValidateSvmMigrationVolumePlacementInlineVolumeAggregatePairs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SvmMigrationVolumePlacementInlineVolumeAggregatePairs); i++ {

		if m.SvmMigrationVolumePlacementInlineVolumeAggregatePairs[i] != nil {
			if err := m.SvmMigrationVolumePlacementInlineVolumeAggregatePairs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volume_aggregate_pairs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationVolumePlacement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumePlacement) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumePlacement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumePlacementInlineAggregatesInlineArrayItem Aggregate
//
// swagger:model svm_migration_volume_placement_inline_aggregates_inline_array_item
type SvmMigrationVolumePlacementInlineAggregatesInlineArrayItem struct {

	// links
	Links *SvmMigrationVolumePlacementInlineAggregatesInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// name
	// Example: aggr1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this svm migration volume placement inline aggregates inline array item
func (m *SvmMigrationVolumePlacementInlineAggregatesInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementInlineAggregatesInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this svm migration volume placement inline aggregates inline array item based on the context it is used
func (m *SvmMigrationVolumePlacementInlineAggregatesInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementInlineAggregatesInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementInlineAggregatesInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementInlineAggregatesInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumePlacementInlineAggregatesInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumePlacementInlineAggregatesInlineArrayItemInlineLinks svm migration volume placement inline aggregates inline array item inline links
//
// swagger:model svm_migration_volume_placement_inline_aggregates_inline_array_item_inline__links
type SvmMigrationVolumePlacementInlineAggregatesInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm migration volume placement inline aggregates inline array item inline links
func (m *SvmMigrationVolumePlacementInlineAggregatesInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementInlineAggregatesInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this svm migration volume placement inline aggregates inline array item inline links based on the context it is used
func (m *SvmMigrationVolumePlacementInlineAggregatesInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementInlineAggregatesInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SvmMigrationVolumePlacementInlineAggregatesInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementInlineAggregatesInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumePlacementInlineAggregatesInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItem Volume-aggregate pair information.
//
// swagger:model svm_migration_volume_placement_inline_volume_aggregate_pairs_inline_array_item
type SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItem struct {

	// aggregate
	Aggregate *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregate `json:"aggregate,omitempty"`

	// volume
	Volume *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolume `json:"volume,omitempty"`
}

// Validate validates this svm migration volume placement inline volume aggregate pairs inline array item
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItem) validateAggregate(formats strfmt.Registry) error {
	if swag.IsZero(m.Aggregate) { // not required
		return nil
	}

	if m.Aggregate != nil {
		if err := m.Aggregate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate")
			}
			return err
		}
	}

	return nil
}

func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItem) validateVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration volume placement inline volume aggregate pairs inline array item based on the context it is used
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItem) contextValidateAggregate(ctx context.Context, formats strfmt.Registry) error {

	if m.Aggregate != nil {
		if err := m.Aggregate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate")
			}
			return err
		}
	}

	return nil
}

func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItem) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {
		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregate Aggregate to use for volume creation.
//
// swagger:model svm_migration_volume_placement_inline_volume_aggregate_pairs_inline_array_item_inline_aggregate
type SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregate struct {

	// links
	Links *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregateInlineLinks `json:"_links,omitempty"`

	// name
	// Example: aggr1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this svm migration volume placement inline volume aggregate pairs inline array item inline aggregate
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregate) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration volume placement inline volume aggregate pairs inline array item inline aggregate based on the context it is used
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregate) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregate) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregateInlineLinks svm migration volume placement inline volume aggregate pairs inline array item inline aggregate inline links
//
// swagger:model svm_migration_volume_placement_inline_volume_aggregate_pairs_inline_array_item_inline_aggregate_inline__links
type SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregateInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm migration volume placement inline volume aggregate pairs inline array item inline aggregate inline links
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregateInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregateInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration volume placement inline volume aggregate pairs inline array item inline aggregate inline links based on the context it is used
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregateInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregateInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregateInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregateInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineAggregateInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolume Property indicating the source volume.
//
// swagger:model svm_migration_volume_placement_inline_volume_aggregate_pairs_inline_array_item_inline_volume
type SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolume struct {

	// links
	Links *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolumeInlineLinks `json:"_links,omitempty"`

	// The name of the volume.
	// Example: volume1
	Name *string `json:"name,omitempty"`

	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this svm migration volume placement inline volume aggregate pairs inline array item inline volume
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolume) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration volume placement inline volume aggregate pairs inline array item inline volume based on the context it is used
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolume) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolume) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolumeInlineLinks svm migration volume placement inline volume aggregate pairs inline array item inline volume inline links
//
// swagger:model svm_migration_volume_placement_inline_volume_aggregate_pairs_inline_array_item_inline_volume_inline__links
type SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolumeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this svm migration volume placement inline volume aggregate pairs inline array item inline volume inline links
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolumeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolumeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this svm migration volume placement inline volume aggregate pairs inline array item inline volume inline links based on the context it is used
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolumeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolumeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolumeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolumeInlineLinks) UnmarshalBinary(b []byte) error {
	var res SvmMigrationVolumePlacementInlineVolumeAggregatePairsInlineArrayItemInlineVolumeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
