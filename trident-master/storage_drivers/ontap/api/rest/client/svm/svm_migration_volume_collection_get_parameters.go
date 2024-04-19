// Code generated by go-swagger; DO NOT EDIT.

package svm

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

// NewSvmMigrationVolumeCollectionGetParams creates a new SvmMigrationVolumeCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmMigrationVolumeCollectionGetParams() *SvmMigrationVolumeCollectionGetParams {
	return &SvmMigrationVolumeCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmMigrationVolumeCollectionGetParamsWithTimeout creates a new SvmMigrationVolumeCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSvmMigrationVolumeCollectionGetParamsWithTimeout(timeout time.Duration) *SvmMigrationVolumeCollectionGetParams {
	return &SvmMigrationVolumeCollectionGetParams{
		timeout: timeout,
	}
}

// NewSvmMigrationVolumeCollectionGetParamsWithContext creates a new SvmMigrationVolumeCollectionGetParams object
// with the ability to set a context for a request.
func NewSvmMigrationVolumeCollectionGetParamsWithContext(ctx context.Context) *SvmMigrationVolumeCollectionGetParams {
	return &SvmMigrationVolumeCollectionGetParams{
		Context: ctx,
	}
}

// NewSvmMigrationVolumeCollectionGetParamsWithHTTPClient creates a new SvmMigrationVolumeCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmMigrationVolumeCollectionGetParamsWithHTTPClient(client *http.Client) *SvmMigrationVolumeCollectionGetParams {
	return &SvmMigrationVolumeCollectionGetParams{
		HTTPClient: client,
	}
}

/*
SvmMigrationVolumeCollectionGetParams contains all the parameters to send to the API endpoint

	for the svm migration volume collection get operation.

	Typically these are written to a http.Request.
*/
type SvmMigrationVolumeCollectionGetParams struct {

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

	/* SvmMigrationUUID.

	   Migration UUID
	*/
	SvmMigrationUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the svm migration volume collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmMigrationVolumeCollectionGetParams) WithDefaults() *SvmMigrationVolumeCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm migration volume collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmMigrationVolumeCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SvmMigrationVolumeCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) WithTimeout(timeout time.Duration) *SvmMigrationVolumeCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) WithContext(ctx context.Context) *SvmMigrationVolumeCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) WithHTTPClient(client *http.Client) *SvmMigrationVolumeCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) WithFields(fields []string) *SvmMigrationVolumeCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) WithMaxRecords(maxRecords *int64) *SvmMigrationVolumeCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) WithOrderBy(orderBy []string) *SvmMigrationVolumeCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) WithReturnRecords(returnRecords *bool) *SvmMigrationVolumeCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *SvmMigrationVolumeCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmMigrationUUID adds the svmMigrationUUID to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) WithSvmMigrationUUID(svmMigrationUUID string) *SvmMigrationVolumeCollectionGetParams {
	o.SetSvmMigrationUUID(svmMigrationUUID)
	return o
}

// SetSvmMigrationUUID adds the svmMigrationUuid to the svm migration volume collection get params
func (o *SvmMigrationVolumeCollectionGetParams) SetSvmMigrationUUID(svmMigrationUUID string) {
	o.SvmMigrationUUID = svmMigrationUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SvmMigrationVolumeCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param svm_migration.uuid
	if err := r.SetPathParam("svm_migration.uuid", o.SvmMigrationUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSvmMigrationVolumeCollectionGet binds the parameter fields
func (o *SvmMigrationVolumeCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSvmMigrationVolumeCollectionGet binds the parameter order_by
func (o *SvmMigrationVolumeCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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