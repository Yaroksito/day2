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
	"github.com/go-openapi/validate"
)

// TopMetricsSvmUser Aggregated information about a user's IO activity at a SVM scope.
// Example: {"iops":{"error":{"lower_bound":"10","upper_bound":"11"},"read":"10","write":"5"},"svm":{"name":"vserver_3","uuid":"42ee3002-67dd-11ea-8508-005056a7b8ac"},"throughput":{"error":{"lower_bound":"2","upper_bound":"3"},"read":"2","write":"1"},"user_id":"1876","user_name":"Ryan","volumes":[{"name":"vol_6","uuid":"02178914-5f67-11eb-b987-005056ac5da5"},{"name":"vol_8","uuid":"c05eb66a-685f-11ea-8508-005056a7b8ac"}]}
//
// swagger:model top_metrics_svm_user
type TopMetricsSvmUser struct {

	// iops
	Iops *TopMetricsSvmUserInlineIops `json:"iops,omitempty"`

	// svm
	Svm *TopMetricsSvmUserInlineSvm `json:"svm,omitempty"`

	// throughput
	Throughput *TopMetricsSvmUserInlineThroughput `json:"throughput,omitempty"`

	// List of volumes where the user is generating traffic.
	TopMetricsSvmUserInlineVolumes []*TopMetricsSvmUserInlineVolumesInlineArrayItem `json:"volumes,omitempty"`

	// User ID of the user.
	// Example: S-1-5-21-256008430-3394229847-3930036330-1001
	// Read Only: true
	UserID *string `json:"user_id,omitempty"`

	// Name of the user.
	// Example: James
	// Read Only: true
	UserName *string `json:"user_name,omitempty"`
}

// Validate validates this top metrics svm user
func (m *TopMetricsSvmUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopMetricsSvmUserInlineVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmUser) validateIops(formats strfmt.Registry) error {
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

func (m *TopMetricsSvmUser) validateSvm(formats strfmt.Registry) error {
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

func (m *TopMetricsSvmUser) validateThroughput(formats strfmt.Registry) error {
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

func (m *TopMetricsSvmUser) validateTopMetricsSvmUserInlineVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.TopMetricsSvmUserInlineVolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.TopMetricsSvmUserInlineVolumes); i++ {
		if swag.IsZero(m.TopMetricsSvmUserInlineVolumes[i]) { // not required
			continue
		}

		if m.TopMetricsSvmUserInlineVolumes[i] != nil {
			if err := m.TopMetricsSvmUserInlineVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this top metrics svm user based on the context it is used
func (m *TopMetricsSvmUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTopMetricsSvmUserInlineVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmUser) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TopMetricsSvmUser) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TopMetricsSvmUser) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TopMetricsSvmUser) contextValidateTopMetricsSvmUserInlineVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TopMetricsSvmUserInlineVolumes); i++ {

		if m.TopMetricsSvmUserInlineVolumes[i] != nil {
			if err := m.TopMetricsSvmUserInlineVolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TopMetricsSvmUser) contextValidateUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsSvmUser) contextValidateUserName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "user_name", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmUser) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmUserInlineIops top metrics svm user inline iops
//
// swagger:model top_metrics_svm_user_inline_iops
type TopMetricsSvmUserInlineIops struct {

	// error
	Error *TopMetricValueErrorBounds `json:"error,omitempty"`

	// Average number of read operations per second.
	// Example: 4
	// Read Only: true
	Read *int64 `json:"read,omitempty"`

	// Average number of write operations per second.
	// Example: 8
	// Read Only: true
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this top metrics svm user inline iops
func (m *TopMetricsSvmUserInlineIops) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmUserInlineIops) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics svm user inline iops based on the context it is used
func (m *TopMetricsSvmUserInlineIops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRead(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWrite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmUserInlineIops) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {
		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmUserInlineIops) contextValidateRead(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "iops"+"."+"read", "body", m.Read); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsSvmUserInlineIops) contextValidateWrite(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "iops"+"."+"write", "body", m.Write); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmUserInlineIops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmUserInlineIops) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmUserInlineIops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmUserInlineSvm top metrics svm user inline svm
//
// swagger:model top_metrics_svm_user_inline_svm
type TopMetricsSvmUserInlineSvm struct {

	// links
	Links *TopMetricsSvmUserInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this top metrics svm user inline svm
func (m *TopMetricsSvmUserInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmUserInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics svm user inline svm based on the context it is used
func (m *TopMetricsSvmUserInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmUserInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsSvmUserInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmUserInlineSvm) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmUserInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmUserInlineSvmInlineLinks top metrics svm user inline svm inline links
//
// swagger:model top_metrics_svm_user_inline_svm_inline__links
type TopMetricsSvmUserInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this top metrics svm user inline svm inline links
func (m *TopMetricsSvmUserInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmUserInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics svm user inline svm inline links based on the context it is used
func (m *TopMetricsSvmUserInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmUserInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsSvmUserInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmUserInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmUserInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmUserInlineThroughput top metrics svm user inline throughput
//
// swagger:model top_metrics_svm_user_inline_throughput
type TopMetricsSvmUserInlineThroughput struct {

	// error
	Error *TopMetricValueErrorBounds `json:"error,omitempty"`

	// Average number of read bytes received per second.
	// Example: 10
	// Read Only: true
	Read *int64 `json:"read,omitempty"`

	// Average number of write bytes received per second.
	// Example: 7
	// Read Only: true
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this top metrics svm user inline throughput
func (m *TopMetricsSvmUserInlineThroughput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmUserInlineThroughput) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this top metrics svm user inline throughput based on the context it is used
func (m *TopMetricsSvmUserInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRead(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWrite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmUserInlineThroughput) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {
		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (m *TopMetricsSvmUserInlineThroughput) contextValidateRead(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "throughput"+"."+"read", "body", m.Read); err != nil {
		return err
	}

	return nil
}

func (m *TopMetricsSvmUserInlineThroughput) contextValidateWrite(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "throughput"+"."+"write", "body", m.Write); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TopMetricsSvmUserInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmUserInlineThroughput) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmUserInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmUserInlineVolumesInlineArrayItem top metrics svm user inline volumes inline array item
//
// swagger:model top_metrics_svm_user_inline_volumes_inline_array_item
type TopMetricsSvmUserInlineVolumesInlineArrayItem struct {

	// links
	Links *TopMetricsSvmUserInlineVolumesInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// The name of the volume.
	// Example: volume1
	Name *string `json:"name,omitempty"`

	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this top metrics svm user inline volumes inline array item
func (m *TopMetricsSvmUserInlineVolumesInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmUserInlineVolumesInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics svm user inline volumes inline array item based on the context it is used
func (m *TopMetricsSvmUserInlineVolumesInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmUserInlineVolumesInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsSvmUserInlineVolumesInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmUserInlineVolumesInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmUserInlineVolumesInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TopMetricsSvmUserInlineVolumesInlineArrayItemInlineLinks top metrics svm user inline volumes inline array item inline links
//
// swagger:model top_metrics_svm_user_inline_volumes_inline_array_item_inline__links
type TopMetricsSvmUserInlineVolumesInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this top metrics svm user inline volumes inline array item inline links
func (m *TopMetricsSvmUserInlineVolumesInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmUserInlineVolumesInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this top metrics svm user inline volumes inline array item inline links based on the context it is used
func (m *TopMetricsSvmUserInlineVolumesInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopMetricsSvmUserInlineVolumesInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TopMetricsSvmUserInlineVolumesInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopMetricsSvmUserInlineVolumesInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res TopMetricsSvmUserInlineVolumesInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
