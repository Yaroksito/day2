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

// NewSplitLoadGetParams creates a new SplitLoadGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSplitLoadGetParams() *SplitLoadGetParams {
	return &SplitLoadGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSplitLoadGetParamsWithTimeout creates a new SplitLoadGetParams object
// with the ability to set a timeout on a request.
func NewSplitLoadGetParamsWithTimeout(timeout time.Duration) *SplitLoadGetParams {
	return &SplitLoadGetParams{
		timeout: timeout,
	}
}

// NewSplitLoadGetParamsWithContext creates a new SplitLoadGetParams object
// with the ability to set a context for a request.
func NewSplitLoadGetParamsWithContext(ctx context.Context) *SplitLoadGetParams {
	return &SplitLoadGetParams{
		Context: ctx,
	}
}

// NewSplitLoadGetParamsWithHTTPClient creates a new SplitLoadGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSplitLoadGetParamsWithHTTPClient(client *http.Client) *SplitLoadGetParams {
	return &SplitLoadGetParams{
		HTTPClient: client,
	}
}

/*
SplitLoadGetParams contains all the parameters to send to the API endpoint

	for the split load get operation.

	Typically these are written to a http.Request.
*/
type SplitLoadGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* NodeUUID.

	   Node Uuid
	*/
	NodeUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the split load get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SplitLoadGetParams) WithDefaults() *SplitLoadGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the split load get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SplitLoadGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the split load get params
func (o *SplitLoadGetParams) WithTimeout(timeout time.Duration) *SplitLoadGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the split load get params
func (o *SplitLoadGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the split load get params
func (o *SplitLoadGetParams) WithContext(ctx context.Context) *SplitLoadGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the split load get params
func (o *SplitLoadGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the split load get params
func (o *SplitLoadGetParams) WithHTTPClient(client *http.Client) *SplitLoadGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the split load get params
func (o *SplitLoadGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the split load get params
func (o *SplitLoadGetParams) WithFields(fields []string) *SplitLoadGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the split load get params
func (o *SplitLoadGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithNodeUUID adds the nodeUUID to the split load get params
func (o *SplitLoadGetParams) WithNodeUUID(nodeUUID string) *SplitLoadGetParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the split load get params
func (o *SplitLoadGetParams) SetNodeUUID(nodeUUID string) {
	o.NodeUUID = nodeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SplitLoadGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param node.uuid
	if err := r.SetPathParam("node.uuid", o.NodeUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSplitLoadGet binds the parameter fields
func (o *SplitLoadGetParams) bindParamFields(formats strfmt.Registry) []string {
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
