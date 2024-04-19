// Code generated by go-swagger; DO NOT EDIT.

package security

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewClusterLdapModifyParams creates a new ClusterLdapModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterLdapModifyParams() *ClusterLdapModifyParams {
	return &ClusterLdapModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterLdapModifyParamsWithTimeout creates a new ClusterLdapModifyParams object
// with the ability to set a timeout on a request.
func NewClusterLdapModifyParamsWithTimeout(timeout time.Duration) *ClusterLdapModifyParams {
	return &ClusterLdapModifyParams{
		timeout: timeout,
	}
}

// NewClusterLdapModifyParamsWithContext creates a new ClusterLdapModifyParams object
// with the ability to set a context for a request.
func NewClusterLdapModifyParamsWithContext(ctx context.Context) *ClusterLdapModifyParams {
	return &ClusterLdapModifyParams{
		Context: ctx,
	}
}

// NewClusterLdapModifyParamsWithHTTPClient creates a new ClusterLdapModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterLdapModifyParamsWithHTTPClient(client *http.Client) *ClusterLdapModifyParams {
	return &ClusterLdapModifyParams{
		HTTPClient: client,
	}
}

/*
ClusterLdapModifyParams contains all the parameters to send to the API endpoint

	for the cluster ldap modify operation.

	Typically these are written to a http.Request.
*/
type ClusterLdapModifyParams struct {

	/* Info.

	   Information specification
	*/
	Info *models.ClusterLdap

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster ldap modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterLdapModifyParams) WithDefaults() *ClusterLdapModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster ldap modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterLdapModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster ldap modify params
func (o *ClusterLdapModifyParams) WithTimeout(timeout time.Duration) *ClusterLdapModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster ldap modify params
func (o *ClusterLdapModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster ldap modify params
func (o *ClusterLdapModifyParams) WithContext(ctx context.Context) *ClusterLdapModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster ldap modify params
func (o *ClusterLdapModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster ldap modify params
func (o *ClusterLdapModifyParams) WithHTTPClient(client *http.Client) *ClusterLdapModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster ldap modify params
func (o *ClusterLdapModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the cluster ldap modify params
func (o *ClusterLdapModifyParams) WithInfo(info *models.ClusterLdap) *ClusterLdapModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cluster ldap modify params
func (o *ClusterLdapModifyParams) SetInfo(info *models.ClusterLdap) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterLdapModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
