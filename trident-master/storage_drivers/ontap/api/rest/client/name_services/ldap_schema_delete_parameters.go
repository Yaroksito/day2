// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewLdapSchemaDeleteParams creates a new LdapSchemaDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLdapSchemaDeleteParams() *LdapSchemaDeleteParams {
	return &LdapSchemaDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLdapSchemaDeleteParamsWithTimeout creates a new LdapSchemaDeleteParams object
// with the ability to set a timeout on a request.
func NewLdapSchemaDeleteParamsWithTimeout(timeout time.Duration) *LdapSchemaDeleteParams {
	return &LdapSchemaDeleteParams{
		timeout: timeout,
	}
}

// NewLdapSchemaDeleteParamsWithContext creates a new LdapSchemaDeleteParams object
// with the ability to set a context for a request.
func NewLdapSchemaDeleteParamsWithContext(ctx context.Context) *LdapSchemaDeleteParams {
	return &LdapSchemaDeleteParams{
		Context: ctx,
	}
}

// NewLdapSchemaDeleteParamsWithHTTPClient creates a new LdapSchemaDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewLdapSchemaDeleteParamsWithHTTPClient(client *http.Client) *LdapSchemaDeleteParams {
	return &LdapSchemaDeleteParams{
		HTTPClient: client,
	}
}

/*
LdapSchemaDeleteParams contains all the parameters to send to the API endpoint

	for the ldap schema delete operation.

	Typically these are written to a http.Request.
*/
type LdapSchemaDeleteParams struct {

	/* Name.

	   LDAP schema name.
	*/
	Name string

	/* OwnerUUID.

	   UUID of the owner to which this object belongs.
	*/
	OwnerUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ldap schema delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LdapSchemaDeleteParams) WithDefaults() *LdapSchemaDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ldap schema delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LdapSchemaDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ldap schema delete params
func (o *LdapSchemaDeleteParams) WithTimeout(timeout time.Duration) *LdapSchemaDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ldap schema delete params
func (o *LdapSchemaDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ldap schema delete params
func (o *LdapSchemaDeleteParams) WithContext(ctx context.Context) *LdapSchemaDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ldap schema delete params
func (o *LdapSchemaDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ldap schema delete params
func (o *LdapSchemaDeleteParams) WithHTTPClient(client *http.Client) *LdapSchemaDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ldap schema delete params
func (o *LdapSchemaDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the ldap schema delete params
func (o *LdapSchemaDeleteParams) WithName(name string) *LdapSchemaDeleteParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ldap schema delete params
func (o *LdapSchemaDeleteParams) SetName(name string) {
	o.Name = name
}

// WithOwnerUUID adds the ownerUUID to the ldap schema delete params
func (o *LdapSchemaDeleteParams) WithOwnerUUID(ownerUUID string) *LdapSchemaDeleteParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the ldap schema delete params
func (o *LdapSchemaDeleteParams) SetOwnerUUID(ownerUUID string) {
	o.OwnerUUID = ownerUUID
}

// WriteToRequest writes these params to a swagger request
func (o *LdapSchemaDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param owner.uuid
	if err := r.SetPathParam("owner.uuid", o.OwnerUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}