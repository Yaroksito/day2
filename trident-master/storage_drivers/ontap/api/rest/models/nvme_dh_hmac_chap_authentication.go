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

// NvmeDhHmacChapAuthentication A container for properties of NVMe in-band authentication with the DH-HMAC-CHAP protocol.
//
// swagger:model nvme_dh_hmac_chap_authentication
type NvmeDhHmacChapAuthentication struct {

	// The controller secret for NVMe in-band authentication. The value of this property is used by the NVMe host to authenticate the NVMe controller while establishing a connection. If unset, the controller is not authenticated. When supplied, the property `host_secret_key` must also be supplied. Optional in POST.<br/>
	// This property is write-only. The `mode` property can be used to identify if a controller secret has been set for the host, but the controller secret value cannot be read. To change the value, the host must be deleted from the subsystem and re-added.
	//
	// Example: DHHC-1:00:ia6zGodOr4SEG0Zzaw398rpY0wqipUWj4jWjUh4HWUz6aQ2n:
	ControllerSecretKey *string `json:"controller_secret_key,omitempty"`

	// The Diffie-Hellman group size for NVMe in-band authentication. When property `host_secret_key` is provided, this property defaults to `2048_bit`. When supplied, the property `host_secret_key` must also be supplied. Optional in POST.
	//
	// Enum: [none 2048_bit 3072_bit 4096_bit 6144_bit 8192_bit]
	GroupSize *string `json:"group_size,omitempty"`

	// The hash function for NVMe in-band authentication. When property `host_secret_key` is provided, this property defaults to `sha_256`. When supplied, the property `host_secret_key` must also be supplied. Optional in POST.
	//
	// Enum: [sha_256 sha_512]
	HashFunction *string `json:"hash_function,omitempty"`

	// The host secret for NVMe in-band authentication. The value of this property is used by the NVMe controller to authenticate the NVMe host while establishing a connection. If unset, no authentication is performed by the host or controller. This property must be supplied if any other NVMe in-band authentication properties are supplied. Optional in POST.<br/>
	// This property is write-only. The `mode` property can be used to identify if a host secret has been set for the host, but the host secret value cannot be read. To change the value, the host must be deleted from the subsystem and re-added.
	//
	// Example: DHHC-1:00:ia6zGodOr4SEG0Zzaw398rpY0wqipUWj4jWjUh4HWUz6aQ2n:
	HostSecretKey *string `json:"host_secret_key,omitempty"`

	// The expected NVMe in-band authentication mode for the host. This property is an indication of which secrets are configured for the host. When set to:
	// - none: The host has neither the host nor controller secret configured, and no authentication is performed.
	// - unidirectional: The host has a host secret configured. The controller will authenticate the host.
	// - bidirectional: The host has both a host and controller secret configured. The controller will authenticate the host and the host will authenticate the controller.
	//
	// Example: bidirectional
	// Read Only: true
	// Enum: [none unidirectional bidirectional]
	Mode *string `json:"mode,omitempty"`
}

// Validate validates this nvme dh hmac chap authentication
func (m *NvmeDhHmacChapAuthentication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHashFunction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var nvmeDhHmacChapAuthenticationTypeGroupSizePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","2048_bit","3072_bit","4096_bit","6144_bit","8192_bit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nvmeDhHmacChapAuthenticationTypeGroupSizePropEnum = append(nvmeDhHmacChapAuthenticationTypeGroupSizePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// nvme_dh_hmac_chap_authentication
	// NvmeDhHmacChapAuthentication
	// group_size
	// GroupSize
	// none
	// END DEBUGGING
	// NvmeDhHmacChapAuthenticationGroupSizeNone captures enum value "none"
	NvmeDhHmacChapAuthenticationGroupSizeNone string = "none"

	// BEGIN DEBUGGING
	// nvme_dh_hmac_chap_authentication
	// NvmeDhHmacChapAuthentication
	// group_size
	// GroupSize
	// 2048_bit
	// END DEBUGGING
	// NvmeDhHmacChapAuthenticationGroupSizeNr2048Bit captures enum value "2048_bit"
	NvmeDhHmacChapAuthenticationGroupSizeNr2048Bit string = "2048_bit"

	// BEGIN DEBUGGING
	// nvme_dh_hmac_chap_authentication
	// NvmeDhHmacChapAuthentication
	// group_size
	// GroupSize
	// 3072_bit
	// END DEBUGGING
	// NvmeDhHmacChapAuthenticationGroupSizeNr3072Bit captures enum value "3072_bit"
	NvmeDhHmacChapAuthenticationGroupSizeNr3072Bit string = "3072_bit"

	// BEGIN DEBUGGING
	// nvme_dh_hmac_chap_authentication
	// NvmeDhHmacChapAuthentication
	// group_size
	// GroupSize
	// 4096_bit
	// END DEBUGGING
	// NvmeDhHmacChapAuthenticationGroupSizeNr4096Bit captures enum value "4096_bit"
	NvmeDhHmacChapAuthenticationGroupSizeNr4096Bit string = "4096_bit"

	// BEGIN DEBUGGING
	// nvme_dh_hmac_chap_authentication
	// NvmeDhHmacChapAuthentication
	// group_size
	// GroupSize
	// 6144_bit
	// END DEBUGGING
	// NvmeDhHmacChapAuthenticationGroupSizeNr6144Bit captures enum value "6144_bit"
	NvmeDhHmacChapAuthenticationGroupSizeNr6144Bit string = "6144_bit"

	// BEGIN DEBUGGING
	// nvme_dh_hmac_chap_authentication
	// NvmeDhHmacChapAuthentication
	// group_size
	// GroupSize
	// 8192_bit
	// END DEBUGGING
	// NvmeDhHmacChapAuthenticationGroupSizeNr8192Bit captures enum value "8192_bit"
	NvmeDhHmacChapAuthenticationGroupSizeNr8192Bit string = "8192_bit"
)

// prop value enum
func (m *NvmeDhHmacChapAuthentication) validateGroupSizeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nvmeDhHmacChapAuthenticationTypeGroupSizePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NvmeDhHmacChapAuthentication) validateGroupSize(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupSize) { // not required
		return nil
	}

	// value enum
	if err := m.validateGroupSizeEnum("group_size", "body", *m.GroupSize); err != nil {
		return err
	}

	return nil
}

var nvmeDhHmacChapAuthenticationTypeHashFunctionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha_256","sha_512"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nvmeDhHmacChapAuthenticationTypeHashFunctionPropEnum = append(nvmeDhHmacChapAuthenticationTypeHashFunctionPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// nvme_dh_hmac_chap_authentication
	// NvmeDhHmacChapAuthentication
	// hash_function
	// HashFunction
	// sha_256
	// END DEBUGGING
	// NvmeDhHmacChapAuthenticationHashFunctionSha256 captures enum value "sha_256"
	NvmeDhHmacChapAuthenticationHashFunctionSha256 string = "sha_256"

	// BEGIN DEBUGGING
	// nvme_dh_hmac_chap_authentication
	// NvmeDhHmacChapAuthentication
	// hash_function
	// HashFunction
	// sha_512
	// END DEBUGGING
	// NvmeDhHmacChapAuthenticationHashFunctionSha512 captures enum value "sha_512"
	NvmeDhHmacChapAuthenticationHashFunctionSha512 string = "sha_512"
)

// prop value enum
func (m *NvmeDhHmacChapAuthentication) validateHashFunctionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nvmeDhHmacChapAuthenticationTypeHashFunctionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NvmeDhHmacChapAuthentication) validateHashFunction(formats strfmt.Registry) error {
	if swag.IsZero(m.HashFunction) { // not required
		return nil
	}

	// value enum
	if err := m.validateHashFunctionEnum("hash_function", "body", *m.HashFunction); err != nil {
		return err
	}

	return nil
}

var nvmeDhHmacChapAuthenticationTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","unidirectional","bidirectional"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nvmeDhHmacChapAuthenticationTypeModePropEnum = append(nvmeDhHmacChapAuthenticationTypeModePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// nvme_dh_hmac_chap_authentication
	// NvmeDhHmacChapAuthentication
	// mode
	// Mode
	// none
	// END DEBUGGING
	// NvmeDhHmacChapAuthenticationModeNone captures enum value "none"
	NvmeDhHmacChapAuthenticationModeNone string = "none"

	// BEGIN DEBUGGING
	// nvme_dh_hmac_chap_authentication
	// NvmeDhHmacChapAuthentication
	// mode
	// Mode
	// unidirectional
	// END DEBUGGING
	// NvmeDhHmacChapAuthenticationModeUnidirectional captures enum value "unidirectional"
	NvmeDhHmacChapAuthenticationModeUnidirectional string = "unidirectional"

	// BEGIN DEBUGGING
	// nvme_dh_hmac_chap_authentication
	// NvmeDhHmacChapAuthentication
	// mode
	// Mode
	// bidirectional
	// END DEBUGGING
	// NvmeDhHmacChapAuthenticationModeBidirectional captures enum value "bidirectional"
	NvmeDhHmacChapAuthenticationModeBidirectional string = "bidirectional"
)

// prop value enum
func (m *NvmeDhHmacChapAuthentication) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nvmeDhHmacChapAuthenticationTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NvmeDhHmacChapAuthentication) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nvme dh hmac chap authentication based on the context it is used
func (m *NvmeDhHmacChapAuthentication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmeDhHmacChapAuthentication) contextValidateMode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmeDhHmacChapAuthentication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmeDhHmacChapAuthentication) UnmarshalBinary(b []byte) error {
	var res NvmeDhHmacChapAuthentication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
