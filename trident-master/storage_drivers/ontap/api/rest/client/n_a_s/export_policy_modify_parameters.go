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

// NewExportPolicyModifyParams creates a new ExportPolicyModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportPolicyModifyParams() *ExportPolicyModifyParams {
	return &ExportPolicyModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportPolicyModifyParamsWithTimeout creates a new ExportPolicyModifyParams object
// with the ability to set a timeout on a request.
func NewExportPolicyModifyParamsWithTimeout(timeout time.Duration) *ExportPolicyModifyParams {
	return &ExportPolicyModifyParams{
		timeout: timeout,
	}
}

// NewExportPolicyModifyParamsWithContext creates a new ExportPolicyModifyParams object
// with the ability to set a context for a request.
func NewExportPolicyModifyParamsWithContext(ctx context.Context) *ExportPolicyModifyParams {
	return &ExportPolicyModifyParams{
		Context: ctx,
	}
}

// NewExportPolicyModifyParamsWithHTTPClient creates a new ExportPolicyModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportPolicyModifyParamsWithHTTPClient(client *http.Client) *ExportPolicyModifyParams {
	return &ExportPolicyModifyParams{
		HTTPClient: client,
	}
}

/*
ExportPolicyModifyParams contains all the parameters to send to the API endpoint

	for the export policy modify operation.

	Typically these are written to a http.Request.
*/
type ExportPolicyModifyParams struct {

	/* ID.

	   Export Policy ID
	*/
	ID int64

	/* Info.

	   Info specification
	*/
	Info *models.ExportPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the export policy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportPolicyModifyParams) WithDefaults() *ExportPolicyModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export policy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportPolicyModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the export policy modify params
func (o *ExportPolicyModifyParams) WithTimeout(timeout time.Duration) *ExportPolicyModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export policy modify params
func (o *ExportPolicyModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export policy modify params
func (o *ExportPolicyModifyParams) WithContext(ctx context.Context) *ExportPolicyModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export policy modify params
func (o *ExportPolicyModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export policy modify params
func (o *ExportPolicyModifyParams) WithHTTPClient(client *http.Client) *ExportPolicyModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export policy modify params
func (o *ExportPolicyModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the export policy modify params
func (o *ExportPolicyModifyParams) WithID(id int64) *ExportPolicyModifyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the export policy modify params
func (o *ExportPolicyModifyParams) SetID(id int64) {
	o.ID = id
}

// WithInfo adds the info to the export policy modify params
func (o *ExportPolicyModifyParams) WithInfo(info *models.ExportPolicy) *ExportPolicyModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the export policy modify params
func (o *ExportPolicyModifyParams) SetInfo(info *models.ExportPolicy) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *ExportPolicyModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}
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