// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewFpolicyCreateParams creates a new FpolicyCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFpolicyCreateParams() *FpolicyCreateParams {
	return &FpolicyCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFpolicyCreateParamsWithTimeout creates a new FpolicyCreateParams object
// with the ability to set a timeout on a request.
func NewFpolicyCreateParamsWithTimeout(timeout time.Duration) *FpolicyCreateParams {
	return &FpolicyCreateParams{
		timeout: timeout,
	}
}

// NewFpolicyCreateParamsWithContext creates a new FpolicyCreateParams object
// with the ability to set a context for a request.
func NewFpolicyCreateParamsWithContext(ctx context.Context) *FpolicyCreateParams {
	return &FpolicyCreateParams{
		Context: ctx,
	}
}

// NewFpolicyCreateParamsWithHTTPClient creates a new FpolicyCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewFpolicyCreateParamsWithHTTPClient(client *http.Client) *FpolicyCreateParams {
	return &FpolicyCreateParams{
		HTTPClient: client,
	}
}

/*
FpolicyCreateParams contains all the parameters to send to the API endpoint

	for the fpolicy create operation.

	Typically these are written to a http.Request.
*/
type FpolicyCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.Fpolicy

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fpolicy create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyCreateParams) WithDefaults() *FpolicyCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fpolicy create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := FpolicyCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the fpolicy create params
func (o *FpolicyCreateParams) WithTimeout(timeout time.Duration) *FpolicyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fpolicy create params
func (o *FpolicyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fpolicy create params
func (o *FpolicyCreateParams) WithContext(ctx context.Context) *FpolicyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fpolicy create params
func (o *FpolicyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fpolicy create params
func (o *FpolicyCreateParams) WithHTTPClient(client *http.Client) *FpolicyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fpolicy create params
func (o *FpolicyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the fpolicy create params
func (o *FpolicyCreateParams) WithInfo(info *models.Fpolicy) *FpolicyCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the fpolicy create params
func (o *FpolicyCreateParams) SetInfo(info *models.Fpolicy) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the fpolicy create params
func (o *FpolicyCreateParams) WithReturnRecords(returnRecords *bool) *FpolicyCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the fpolicy create params
func (o *FpolicyCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *FpolicyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}