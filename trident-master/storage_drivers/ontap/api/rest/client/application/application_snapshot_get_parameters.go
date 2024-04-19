// Code generated by go-swagger; DO NOT EDIT.

package application

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

// NewApplicationSnapshotGetParams creates a new ApplicationSnapshotGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationSnapshotGetParams() *ApplicationSnapshotGetParams {
	return &ApplicationSnapshotGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationSnapshotGetParamsWithTimeout creates a new ApplicationSnapshotGetParams object
// with the ability to set a timeout on a request.
func NewApplicationSnapshotGetParamsWithTimeout(timeout time.Duration) *ApplicationSnapshotGetParams {
	return &ApplicationSnapshotGetParams{
		timeout: timeout,
	}
}

// NewApplicationSnapshotGetParamsWithContext creates a new ApplicationSnapshotGetParams object
// with the ability to set a context for a request.
func NewApplicationSnapshotGetParamsWithContext(ctx context.Context) *ApplicationSnapshotGetParams {
	return &ApplicationSnapshotGetParams{
		Context: ctx,
	}
}

// NewApplicationSnapshotGetParamsWithHTTPClient creates a new ApplicationSnapshotGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationSnapshotGetParamsWithHTTPClient(client *http.Client) *ApplicationSnapshotGetParams {
	return &ApplicationSnapshotGetParams{
		HTTPClient: client,
	}
}

/*
ApplicationSnapshotGetParams contains all the parameters to send to the API endpoint

	for the application snapshot get operation.

	Typically these are written to a http.Request.
*/
type ApplicationSnapshotGetParams struct {

	/* ApplicationUUID.

	   Application UUID
	*/
	ApplicationUUID string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* UUID.

	   Snapshot copy UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application snapshot get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationSnapshotGetParams) WithDefaults() *ApplicationSnapshotGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application snapshot get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationSnapshotGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the application snapshot get params
func (o *ApplicationSnapshotGetParams) WithTimeout(timeout time.Duration) *ApplicationSnapshotGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application snapshot get params
func (o *ApplicationSnapshotGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application snapshot get params
func (o *ApplicationSnapshotGetParams) WithContext(ctx context.Context) *ApplicationSnapshotGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application snapshot get params
func (o *ApplicationSnapshotGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application snapshot get params
func (o *ApplicationSnapshotGetParams) WithHTTPClient(client *http.Client) *ApplicationSnapshotGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application snapshot get params
func (o *ApplicationSnapshotGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationUUID adds the applicationUUID to the application snapshot get params
func (o *ApplicationSnapshotGetParams) WithApplicationUUID(applicationUUID string) *ApplicationSnapshotGetParams {
	o.SetApplicationUUID(applicationUUID)
	return o
}

// SetApplicationUUID adds the applicationUuid to the application snapshot get params
func (o *ApplicationSnapshotGetParams) SetApplicationUUID(applicationUUID string) {
	o.ApplicationUUID = applicationUUID
}

// WithFields adds the fields to the application snapshot get params
func (o *ApplicationSnapshotGetParams) WithFields(fields []string) *ApplicationSnapshotGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the application snapshot get params
func (o *ApplicationSnapshotGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithUUID adds the uuid to the application snapshot get params
func (o *ApplicationSnapshotGetParams) WithUUID(uuid string) *ApplicationSnapshotGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the application snapshot get params
func (o *ApplicationSnapshotGetParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationSnapshotGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application.uuid
	if err := r.SetPathParam("application.uuid", o.ApplicationUUID); err != nil {
		return err
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamApplicationSnapshotGet binds the parameter fields
func (o *ApplicationSnapshotGetParams) bindParamFields(formats strfmt.Registry) []string {
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