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
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewSnaplockRetentionOperationCreateParams creates a new SnaplockRetentionOperationCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockRetentionOperationCreateParams() *SnaplockRetentionOperationCreateParams {
	return &SnaplockRetentionOperationCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockRetentionOperationCreateParamsWithTimeout creates a new SnaplockRetentionOperationCreateParams object
// with the ability to set a timeout on a request.
func NewSnaplockRetentionOperationCreateParamsWithTimeout(timeout time.Duration) *SnaplockRetentionOperationCreateParams {
	return &SnaplockRetentionOperationCreateParams{
		timeout: timeout,
	}
}

// NewSnaplockRetentionOperationCreateParamsWithContext creates a new SnaplockRetentionOperationCreateParams object
// with the ability to set a context for a request.
func NewSnaplockRetentionOperationCreateParamsWithContext(ctx context.Context) *SnaplockRetentionOperationCreateParams {
	return &SnaplockRetentionOperationCreateParams{
		Context: ctx,
	}
}

// NewSnaplockRetentionOperationCreateParamsWithHTTPClient creates a new SnaplockRetentionOperationCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockRetentionOperationCreateParamsWithHTTPClient(client *http.Client) *SnaplockRetentionOperationCreateParams {
	return &SnaplockRetentionOperationCreateParams{
		HTTPClient: client,
	}
}

/*
SnaplockRetentionOperationCreateParams contains all the parameters to send to the API endpoint

	for the snaplock retention operation create operation.

	Typically these are written to a http.Request.
*/
type SnaplockRetentionOperationCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.EbrOperation

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock retention operation create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockRetentionOperationCreateParams) WithDefaults() *SnaplockRetentionOperationCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock retention operation create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockRetentionOperationCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := SnaplockRetentionOperationCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snaplock retention operation create params
func (o *SnaplockRetentionOperationCreateParams) WithTimeout(timeout time.Duration) *SnaplockRetentionOperationCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock retention operation create params
func (o *SnaplockRetentionOperationCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock retention operation create params
func (o *SnaplockRetentionOperationCreateParams) WithContext(ctx context.Context) *SnaplockRetentionOperationCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock retention operation create params
func (o *SnaplockRetentionOperationCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock retention operation create params
func (o *SnaplockRetentionOperationCreateParams) WithHTTPClient(client *http.Client) *SnaplockRetentionOperationCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock retention operation create params
func (o *SnaplockRetentionOperationCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the snaplock retention operation create params
func (o *SnaplockRetentionOperationCreateParams) WithInfo(info *models.EbrOperation) *SnaplockRetentionOperationCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snaplock retention operation create params
func (o *SnaplockRetentionOperationCreateParams) SetInfo(info *models.EbrOperation) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the snaplock retention operation create params
func (o *SnaplockRetentionOperationCreateParams) WithReturnRecords(returnRecords *bool) *SnaplockRetentionOperationCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the snaplock retention operation create params
func (o *SnaplockRetentionOperationCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockRetentionOperationCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
