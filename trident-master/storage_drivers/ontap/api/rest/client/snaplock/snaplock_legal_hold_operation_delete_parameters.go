// Code generated by go-swagger; DO NOT EDIT.

package snaplock

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

// NewSnaplockLegalHoldOperationDeleteParams creates a new SnaplockLegalHoldOperationDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockLegalHoldOperationDeleteParams() *SnaplockLegalHoldOperationDeleteParams {
	return &SnaplockLegalHoldOperationDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockLegalHoldOperationDeleteParamsWithTimeout creates a new SnaplockLegalHoldOperationDeleteParams object
// with the ability to set a timeout on a request.
func NewSnaplockLegalHoldOperationDeleteParamsWithTimeout(timeout time.Duration) *SnaplockLegalHoldOperationDeleteParams {
	return &SnaplockLegalHoldOperationDeleteParams{
		timeout: timeout,
	}
}

// NewSnaplockLegalHoldOperationDeleteParamsWithContext creates a new SnaplockLegalHoldOperationDeleteParams object
// with the ability to set a context for a request.
func NewSnaplockLegalHoldOperationDeleteParamsWithContext(ctx context.Context) *SnaplockLegalHoldOperationDeleteParams {
	return &SnaplockLegalHoldOperationDeleteParams{
		Context: ctx,
	}
}

// NewSnaplockLegalHoldOperationDeleteParamsWithHTTPClient creates a new SnaplockLegalHoldOperationDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockLegalHoldOperationDeleteParamsWithHTTPClient(client *http.Client) *SnaplockLegalHoldOperationDeleteParams {
	return &SnaplockLegalHoldOperationDeleteParams{
		HTTPClient: client,
	}
}

/*
SnaplockLegalHoldOperationDeleteParams contains all the parameters to send to the API endpoint

	for the snaplock legal hold operation delete operation.

	Typically these are written to a http.Request.
*/
type SnaplockLegalHoldOperationDeleteParams struct {

	/* ID.

	   Operation ID.
	*/
	ID string

	/* LitigationID.

	   Litigation ID
	*/
	LitigationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock legal hold operation delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldOperationDeleteParams) WithDefaults() *SnaplockLegalHoldOperationDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock legal hold operation delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldOperationDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snaplock legal hold operation delete params
func (o *SnaplockLegalHoldOperationDeleteParams) WithTimeout(timeout time.Duration) *SnaplockLegalHoldOperationDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock legal hold operation delete params
func (o *SnaplockLegalHoldOperationDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock legal hold operation delete params
func (o *SnaplockLegalHoldOperationDeleteParams) WithContext(ctx context.Context) *SnaplockLegalHoldOperationDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock legal hold operation delete params
func (o *SnaplockLegalHoldOperationDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock legal hold operation delete params
func (o *SnaplockLegalHoldOperationDeleteParams) WithHTTPClient(client *http.Client) *SnaplockLegalHoldOperationDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock legal hold operation delete params
func (o *SnaplockLegalHoldOperationDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the snaplock legal hold operation delete params
func (o *SnaplockLegalHoldOperationDeleteParams) WithID(id string) *SnaplockLegalHoldOperationDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the snaplock legal hold operation delete params
func (o *SnaplockLegalHoldOperationDeleteParams) SetID(id string) {
	o.ID = id
}

// WithLitigationID adds the litigationID to the snaplock legal hold operation delete params
func (o *SnaplockLegalHoldOperationDeleteParams) WithLitigationID(litigationID string) *SnaplockLegalHoldOperationDeleteParams {
	o.SetLitigationID(litigationID)
	return o
}

// SetLitigationID adds the litigationId to the snaplock legal hold operation delete params
func (o *SnaplockLegalHoldOperationDeleteParams) SetLitigationID(litigationID string) {
	o.LitigationID = litigationID
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockLegalHoldOperationDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param litigation.id
	if err := r.SetPathParam("litigation.id", o.LitigationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
