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

// NewNisDeleteParams creates a new NisDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNisDeleteParams() *NisDeleteParams {
	return &NisDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNisDeleteParamsWithTimeout creates a new NisDeleteParams object
// with the ability to set a timeout on a request.
func NewNisDeleteParamsWithTimeout(timeout time.Duration) *NisDeleteParams {
	return &NisDeleteParams{
		timeout: timeout,
	}
}

// NewNisDeleteParamsWithContext creates a new NisDeleteParams object
// with the ability to set a context for a request.
func NewNisDeleteParamsWithContext(ctx context.Context) *NisDeleteParams {
	return &NisDeleteParams{
		Context: ctx,
	}
}

// NewNisDeleteParamsWithHTTPClient creates a new NisDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewNisDeleteParamsWithHTTPClient(client *http.Client) *NisDeleteParams {
	return &NisDeleteParams{
		HTTPClient: client,
	}
}

/*
NisDeleteParams contains all the parameters to send to the API endpoint

	for the nis delete operation.

	Typically these are written to a http.Request.
*/
type NisDeleteParams struct {

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nis delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NisDeleteParams) WithDefaults() *NisDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nis delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NisDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nis delete params
func (o *NisDeleteParams) WithTimeout(timeout time.Duration) *NisDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nis delete params
func (o *NisDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nis delete params
func (o *NisDeleteParams) WithContext(ctx context.Context) *NisDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nis delete params
func (o *NisDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nis delete params
func (o *NisDeleteParams) WithHTTPClient(client *http.Client) *NisDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nis delete params
func (o *NisDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSvmUUID adds the svmUUID to the nis delete params
func (o *NisDeleteParams) WithSvmUUID(svmUUID string) *NisDeleteParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the nis delete params
func (o *NisDeleteParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NisDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
