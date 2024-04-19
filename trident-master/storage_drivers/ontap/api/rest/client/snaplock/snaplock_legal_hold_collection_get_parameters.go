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
)

// NewSnaplockLegalHoldCollectionGetParams creates a new SnaplockLegalHoldCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockLegalHoldCollectionGetParams() *SnaplockLegalHoldCollectionGetParams {
	return &SnaplockLegalHoldCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockLegalHoldCollectionGetParamsWithTimeout creates a new SnaplockLegalHoldCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSnaplockLegalHoldCollectionGetParamsWithTimeout(timeout time.Duration) *SnaplockLegalHoldCollectionGetParams {
	return &SnaplockLegalHoldCollectionGetParams{
		timeout: timeout,
	}
}

// NewSnaplockLegalHoldCollectionGetParamsWithContext creates a new SnaplockLegalHoldCollectionGetParams object
// with the ability to set a context for a request.
func NewSnaplockLegalHoldCollectionGetParamsWithContext(ctx context.Context) *SnaplockLegalHoldCollectionGetParams {
	return &SnaplockLegalHoldCollectionGetParams{
		Context: ctx,
	}
}

// NewSnaplockLegalHoldCollectionGetParamsWithHTTPClient creates a new SnaplockLegalHoldCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockLegalHoldCollectionGetParamsWithHTTPClient(client *http.Client) *SnaplockLegalHoldCollectionGetParams {
	return &SnaplockLegalHoldCollectionGetParams{
		HTTPClient: client,
	}
}

/*
SnaplockLegalHoldCollectionGetParams contains all the parameters to send to the API endpoint

	for the snaplock legal hold collection get operation.

	Typically these are written to a http.Request.
*/
type SnaplockLegalHoldCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock legal hold collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldCollectionGetParams) WithDefaults() *SnaplockLegalHoldCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock legal hold collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SnaplockLegalHoldCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) WithTimeout(timeout time.Duration) *SnaplockLegalHoldCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) WithContext(ctx context.Context) *SnaplockLegalHoldCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) WithHTTPClient(client *http.Client) *SnaplockLegalHoldCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) WithFields(fields []string) *SnaplockLegalHoldCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) WithMaxRecords(maxRecords *int64) *SnaplockLegalHoldCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) WithOrderBy(orderBy []string) *SnaplockLegalHoldCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) WithReturnRecords(returnRecords *bool) *SnaplockLegalHoldCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *SnaplockLegalHoldCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the snaplock legal hold collection get params
func (o *SnaplockLegalHoldCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockLegalHoldCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
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

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnaplockLegalHoldCollectionGet binds the parameter fields
func (o *SnaplockLegalHoldCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSnaplockLegalHoldCollectionGet binds the parameter order_by
func (o *SnaplockLegalHoldCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}