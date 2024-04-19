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

// AggregateSpare aggregate spare
//
// swagger:model aggregate_spare
type AggregateSpare struct {

	// Available RAID protections and their restrictions.
	AggregateSpareInlineLayoutRequirements []*LayoutRequirement `json:"layout_requirements,omitempty"`

	// The checksum type that has been assigned to the spares.
	// Read Only: true
	// Enum: [block advanced_zoned]
	ChecksumStyle *string `json:"checksum_style,omitempty"`

	// Disk class of spares.
	// Example: solid_state
	// Read Only: true
	// Enum: [unknown capacity performance archive solid_state array virtual data_center capacity_flash]
	DiskClass *string `json:"disk_class,omitempty"`

	// Type of disk.
	// Read Only: true
	// Enum: [fc lun nl_sas nvme_ssd sas sata scsi ssd ssd_cap ssd_zns vm_disk]
	DiskType *string `json:"disk_type,omitempty"`

	// Indicates whether a disk is partitioned (true) or whole (false)
	// Example: true
	// Read Only: true
	IsPartition *bool `json:"is_partition,omitempty"`

	// node
	Node *AggregateSpareInlineNode `json:"node,omitempty"`

	// Usable size of each spare, in bytes.
	// Example: 10156769280
	// Read Only: true
	Size *int64 `json:"size,omitempty"`

	// SyncMirror spare pool.
	// Read Only: true
	// Enum: [pool0 pool1]
	SyncmirrorPool *string `json:"syncmirror_pool,omitempty"`

	// Total number of spares in the bucket. The total spare count for each class of spares also includes reserved spare capacity recommended by ONTAP best practices.
	// Example: 10
	// Read Only: true
	Total *int64 `json:"total,omitempty"`

	// Total number of usable spares in the bucket. The usable count for each class of spares does not include reserved spare capacity recommended by ONTAP best practices.
	// Example: 9
	// Read Only: true
	Usable *int64 `json:"usable,omitempty"`
}

// Validate validates this aggregate spare
func (m *AggregateSpare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregateSpareInlineLayoutRequirements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChecksumStyle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncmirrorPool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregateSpare) validateAggregateSpareInlineLayoutRequirements(formats strfmt.Registry) error {
	if swag.IsZero(m.AggregateSpareInlineLayoutRequirements) { // not required
		return nil
	}

	for i := 0; i < len(m.AggregateSpareInlineLayoutRequirements); i++ {
		if swag.IsZero(m.AggregateSpareInlineLayoutRequirements[i]) { // not required
			continue
		}

		if m.AggregateSpareInlineLayoutRequirements[i] != nil {
			if err := m.AggregateSpareInlineLayoutRequirements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("layout_requirements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var aggregateSpareTypeChecksumStylePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["block","advanced_zoned"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aggregateSpareTypeChecksumStylePropEnum = append(aggregateSpareTypeChecksumStylePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// checksum_style
	// ChecksumStyle
	// block
	// END DEBUGGING
	// AggregateSpareChecksumStyleBlock captures enum value "block"
	AggregateSpareChecksumStyleBlock string = "block"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// checksum_style
	// ChecksumStyle
	// advanced_zoned
	// END DEBUGGING
	// AggregateSpareChecksumStyleAdvancedZoned captures enum value "advanced_zoned"
	AggregateSpareChecksumStyleAdvancedZoned string = "advanced_zoned"
)

// prop value enum
func (m *AggregateSpare) validateChecksumStyleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, aggregateSpareTypeChecksumStylePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AggregateSpare) validateChecksumStyle(formats strfmt.Registry) error {
	if swag.IsZero(m.ChecksumStyle) { // not required
		return nil
	}

	// value enum
	if err := m.validateChecksumStyleEnum("checksum_style", "body", *m.ChecksumStyle); err != nil {
		return err
	}

	return nil
}

var aggregateSpareTypeDiskClassPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","capacity","performance","archive","solid_state","array","virtual","data_center","capacity_flash"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aggregateSpareTypeDiskClassPropEnum = append(aggregateSpareTypeDiskClassPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_class
	// DiskClass
	// unknown
	// END DEBUGGING
	// AggregateSpareDiskClassUnknown captures enum value "unknown"
	AggregateSpareDiskClassUnknown string = "unknown"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_class
	// DiskClass
	// capacity
	// END DEBUGGING
	// AggregateSpareDiskClassCapacity captures enum value "capacity"
	AggregateSpareDiskClassCapacity string = "capacity"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_class
	// DiskClass
	// performance
	// END DEBUGGING
	// AggregateSpareDiskClassPerformance captures enum value "performance"
	AggregateSpareDiskClassPerformance string = "performance"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_class
	// DiskClass
	// archive
	// END DEBUGGING
	// AggregateSpareDiskClassArchive captures enum value "archive"
	AggregateSpareDiskClassArchive string = "archive"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_class
	// DiskClass
	// solid_state
	// END DEBUGGING
	// AggregateSpareDiskClassSolidState captures enum value "solid_state"
	AggregateSpareDiskClassSolidState string = "solid_state"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_class
	// DiskClass
	// array
	// END DEBUGGING
	// AggregateSpareDiskClassArray captures enum value "array"
	AggregateSpareDiskClassArray string = "array"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_class
	// DiskClass
	// virtual
	// END DEBUGGING
	// AggregateSpareDiskClassVirtual captures enum value "virtual"
	AggregateSpareDiskClassVirtual string = "virtual"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_class
	// DiskClass
	// data_center
	// END DEBUGGING
	// AggregateSpareDiskClassDataCenter captures enum value "data_center"
	AggregateSpareDiskClassDataCenter string = "data_center"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_class
	// DiskClass
	// capacity_flash
	// END DEBUGGING
	// AggregateSpareDiskClassCapacityFlash captures enum value "capacity_flash"
	AggregateSpareDiskClassCapacityFlash string = "capacity_flash"
)

// prop value enum
func (m *AggregateSpare) validateDiskClassEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, aggregateSpareTypeDiskClassPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AggregateSpare) validateDiskClass(formats strfmt.Registry) error {
	if swag.IsZero(m.DiskClass) { // not required
		return nil
	}

	// value enum
	if err := m.validateDiskClassEnum("disk_class", "body", *m.DiskClass); err != nil {
		return err
	}

	return nil
}

var aggregateSpareTypeDiskTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fc","lun","nl_sas","nvme_ssd","sas","sata","scsi","ssd","ssd_cap","ssd_zns","vm_disk"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aggregateSpareTypeDiskTypePropEnum = append(aggregateSpareTypeDiskTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_type
	// DiskType
	// fc
	// END DEBUGGING
	// AggregateSpareDiskTypeFc captures enum value "fc"
	AggregateSpareDiskTypeFc string = "fc"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_type
	// DiskType
	// lun
	// END DEBUGGING
	// AggregateSpareDiskTypeLun captures enum value "lun"
	AggregateSpareDiskTypeLun string = "lun"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_type
	// DiskType
	// nl_sas
	// END DEBUGGING
	// AggregateSpareDiskTypeNlSas captures enum value "nl_sas"
	AggregateSpareDiskTypeNlSas string = "nl_sas"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_type
	// DiskType
	// nvme_ssd
	// END DEBUGGING
	// AggregateSpareDiskTypeNvmeSsd captures enum value "nvme_ssd"
	AggregateSpareDiskTypeNvmeSsd string = "nvme_ssd"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_type
	// DiskType
	// sas
	// END DEBUGGING
	// AggregateSpareDiskTypeSas captures enum value "sas"
	AggregateSpareDiskTypeSas string = "sas"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_type
	// DiskType
	// sata
	// END DEBUGGING
	// AggregateSpareDiskTypeSata captures enum value "sata"
	AggregateSpareDiskTypeSata string = "sata"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_type
	// DiskType
	// scsi
	// END DEBUGGING
	// AggregateSpareDiskTypeScsi captures enum value "scsi"
	AggregateSpareDiskTypeScsi string = "scsi"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_type
	// DiskType
	// ssd
	// END DEBUGGING
	// AggregateSpareDiskTypeSsd captures enum value "ssd"
	AggregateSpareDiskTypeSsd string = "ssd"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_type
	// DiskType
	// ssd_cap
	// END DEBUGGING
	// AggregateSpareDiskTypeSsdCap captures enum value "ssd_cap"
	AggregateSpareDiskTypeSsdCap string = "ssd_cap"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_type
	// DiskType
	// ssd_zns
	// END DEBUGGING
	// AggregateSpareDiskTypeSsdZns captures enum value "ssd_zns"
	AggregateSpareDiskTypeSsdZns string = "ssd_zns"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// disk_type
	// DiskType
	// vm_disk
	// END DEBUGGING
	// AggregateSpareDiskTypeVMDisk captures enum value "vm_disk"
	AggregateSpareDiskTypeVMDisk string = "vm_disk"
)

// prop value enum
func (m *AggregateSpare) validateDiskTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, aggregateSpareTypeDiskTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AggregateSpare) validateDiskType(formats strfmt.Registry) error {
	if swag.IsZero(m.DiskType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDiskTypeEnum("disk_type", "body", *m.DiskType); err != nil {
		return err
	}

	return nil
}

func (m *AggregateSpare) validateNode(formats strfmt.Registry) error {
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

var aggregateSpareTypeSyncmirrorPoolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pool0","pool1"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aggregateSpareTypeSyncmirrorPoolPropEnum = append(aggregateSpareTypeSyncmirrorPoolPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// syncmirror_pool
	// SyncmirrorPool
	// pool0
	// END DEBUGGING
	// AggregateSpareSyncmirrorPoolPool0 captures enum value "pool0"
	AggregateSpareSyncmirrorPoolPool0 string = "pool0"

	// BEGIN DEBUGGING
	// aggregate_spare
	// AggregateSpare
	// syncmirror_pool
	// SyncmirrorPool
	// pool1
	// END DEBUGGING
	// AggregateSpareSyncmirrorPoolPool1 captures enum value "pool1"
	AggregateSpareSyncmirrorPoolPool1 string = "pool1"
)

// prop value enum
func (m *AggregateSpare) validateSyncmirrorPoolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, aggregateSpareTypeSyncmirrorPoolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AggregateSpare) validateSyncmirrorPool(formats strfmt.Registry) error {
	if swag.IsZero(m.SyncmirrorPool) { // not required
		return nil
	}

	// value enum
	if err := m.validateSyncmirrorPoolEnum("syncmirror_pool", "body", *m.SyncmirrorPool); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this aggregate spare based on the context it is used
func (m *AggregateSpare) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregateSpareInlineLayoutRequirements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChecksumStyle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDiskClass(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDiskType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsPartition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSyncmirrorPool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregateSpare) contextValidateAggregateSpareInlineLayoutRequirements(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AggregateSpareInlineLayoutRequirements); i++ {

		if m.AggregateSpareInlineLayoutRequirements[i] != nil {
			if err := m.AggregateSpareInlineLayoutRequirements[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("layout_requirements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AggregateSpare) contextValidateChecksumStyle(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "checksum_style", "body", m.ChecksumStyle); err != nil {
		return err
	}

	return nil
}

func (m *AggregateSpare) contextValidateDiskClass(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "disk_class", "body", m.DiskClass); err != nil {
		return err
	}

	return nil
}

func (m *AggregateSpare) contextValidateDiskType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "disk_type", "body", m.DiskType); err != nil {
		return err
	}

	return nil
}

func (m *AggregateSpare) contextValidateIsPartition(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_partition", "body", m.IsPartition); err != nil {
		return err
	}

	return nil
}

func (m *AggregateSpare) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AggregateSpare) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *AggregateSpare) contextValidateSyncmirrorPool(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "syncmirror_pool", "body", m.SyncmirrorPool); err != nil {
		return err
	}

	return nil
}

func (m *AggregateSpare) contextValidateTotal(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

func (m *AggregateSpare) contextValidateUsable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "usable", "body", m.Usable); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AggregateSpare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregateSpare) UnmarshalBinary(b []byte) error {
	var res AggregateSpare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AggregateSpareInlineNode Node where the spares are assigned.
//
// swagger:model aggregate_spare_inline_node
type AggregateSpareInlineNode struct {

	// links
	Links *AggregateSpareInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this aggregate spare inline node
func (m *AggregateSpareInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregateSpareInlineNode) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this aggregate spare inline node based on the context it is used
func (m *AggregateSpareInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregateSpareInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AggregateSpareInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregateSpareInlineNode) UnmarshalBinary(b []byte) error {
	var res AggregateSpareInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AggregateSpareInlineNodeInlineLinks aggregate spare inline node inline links
//
// swagger:model aggregate_spare_inline_node_inline__links
type AggregateSpareInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this aggregate spare inline node inline links
func (m *AggregateSpareInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregateSpareInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this aggregate spare inline node inline links based on the context it is used
func (m *AggregateSpareInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregateSpareInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AggregateSpareInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregateSpareInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res AggregateSpareInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}