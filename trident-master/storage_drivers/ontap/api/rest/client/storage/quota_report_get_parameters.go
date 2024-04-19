// Code generated by go-swagger; DO NOT EDIT.

package storage

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
)

// NewQuotaReportGetParams creates a new QuotaReportGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQuotaReportGetParams() *QuotaReportGetParams {
	return &QuotaReportGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQuotaReportGetParamsWithTimeout creates a new QuotaReportGetParams object
// with the ability to set a timeout on a request.
func NewQuotaReportGetParamsWithTimeout(timeout time.Duration) *QuotaReportGetParams {
	return &QuotaReportGetParams{
		timeout: timeout,
	}
}

// NewQuotaReportGetParamsWithContext creates a new QuotaReportGetParams object
// with the ability to set a context for a request.
func NewQuotaReportGetParamsWithContext(ctx context.Context) *QuotaReportGetParams {
	return &QuotaReportGetParams{
		Context: ctx,
	}
}

// NewQuotaReportGetParamsWithHTTPClient creates a new QuotaReportGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewQuotaReportGetParamsWithHTTPClient(client *http.Client) *QuotaReportGetParams {
	return &QuotaReportGetParams{
		HTTPClient: client,
	}
}

/*
QuotaReportGetParams contains all the parameters to send to the API endpoint

	for the quota report get operation.

	Typically these are written to a http.Request.
*/
type QuotaReportGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Index.

	   Quota report index
	*/
	Index int64

	/* VolumeUUID.

	   Volume UUID
	*/
	VolumeUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the quota report get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuotaReportGetParams) WithDefaults() *QuotaReportGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the quota report get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuotaReportGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the quota report get params
func (o *QuotaReportGetParams) WithTimeout(timeout time.Duration) *QuotaReportGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quota report get params
func (o *QuotaReportGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quota report get params
func (o *QuotaReportGetParams) WithContext(ctx context.Context) *QuotaReportGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quota report get params
func (o *QuotaReportGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quota report get params
func (o *QuotaReportGetParams) WithHTTPClient(client *http.Client) *QuotaReportGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quota report get params
func (o *QuotaReportGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the quota report get params
func (o *QuotaReportGetParams) WithFields(fields []string) *QuotaReportGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the quota report get params
func (o *QuotaReportGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIndex adds the index to the quota report get params
func (o *QuotaReportGetParams) WithIndex(index int64) *QuotaReportGetParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the quota report get params
func (o *QuotaReportGetParams) SetIndex(index int64) {
	o.Index = index
}

// WithVolumeUUID adds the volumeUUID to the quota report get params
func (o *QuotaReportGetParams) WithVolumeUUID(volumeUUID string) *QuotaReportGetParams {
	o.SetVolumeUUID(volumeUUID)
	return o
}

// SetVolumeUUID adds the volumeUuid to the quota report get params
func (o *QuotaReportGetParams) SetVolumeUUID(volumeUUID string) {
	o.VolumeUUID = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *QuotaReportGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

	// path param volume.uuid
	if err := r.SetPathParam("volume.uuid", o.VolumeUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamQuotaReportGet binds the parameter fields
func (o *QuotaReportGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}