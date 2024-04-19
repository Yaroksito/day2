// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewSoftwarePackagesCollectionGetParams creates a new SoftwarePackagesCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSoftwarePackagesCollectionGetParams() *SoftwarePackagesCollectionGetParams {
	return &SoftwarePackagesCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSoftwarePackagesCollectionGetParamsWithTimeout creates a new SoftwarePackagesCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSoftwarePackagesCollectionGetParamsWithTimeout(timeout time.Duration) *SoftwarePackagesCollectionGetParams {
	return &SoftwarePackagesCollectionGetParams{
		timeout: timeout,
	}
}

// NewSoftwarePackagesCollectionGetParamsWithContext creates a new SoftwarePackagesCollectionGetParams object
// with the ability to set a context for a request.
func NewSoftwarePackagesCollectionGetParamsWithContext(ctx context.Context) *SoftwarePackagesCollectionGetParams {
	return &SoftwarePackagesCollectionGetParams{
		Context: ctx,
	}
}

// NewSoftwarePackagesCollectionGetParamsWithHTTPClient creates a new SoftwarePackagesCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSoftwarePackagesCollectionGetParamsWithHTTPClient(client *http.Client) *SoftwarePackagesCollectionGetParams {
	return &SoftwarePackagesCollectionGetParams{
		HTTPClient: client,
	}
}

/*
SoftwarePackagesCollectionGetParams contains all the parameters to send to the API endpoint

	for the software packages collection get operation.

	Typically these are written to a http.Request.
*/
type SoftwarePackagesCollectionGetParams struct {

	/* CreateTime.

	   Filter by create_time
	*/
	CreateTime *string

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

	/* Version.

	   Filter by version
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the software packages collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SoftwarePackagesCollectionGetParams) WithDefaults() *SoftwarePackagesCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the software packages collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SoftwarePackagesCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SoftwarePackagesCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) WithTimeout(timeout time.Duration) *SoftwarePackagesCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) WithContext(ctx context.Context) *SoftwarePackagesCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) WithHTTPClient(client *http.Client) *SoftwarePackagesCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreateTime adds the createTime to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) WithCreateTime(createTime *string) *SoftwarePackagesCollectionGetParams {
	o.SetCreateTime(createTime)
	return o
}

// SetCreateTime adds the createTime to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) SetCreateTime(createTime *string) {
	o.CreateTime = createTime
}

// WithFields adds the fields to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) WithFields(fields []string) *SoftwarePackagesCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) WithMaxRecords(maxRecords *int64) *SoftwarePackagesCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) WithOrderBy(orderBy []string) *SoftwarePackagesCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) WithReturnRecords(returnRecords *bool) *SoftwarePackagesCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *SoftwarePackagesCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithVersion adds the version to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) WithVersion(version *string) *SoftwarePackagesCollectionGetParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the software packages collection get params
func (o *SoftwarePackagesCollectionGetParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *SoftwarePackagesCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CreateTime != nil {

		// query param create_time
		var qrCreateTime string

		if o.CreateTime != nil {
			qrCreateTime = *o.CreateTime
		}
		qCreateTime := qrCreateTime
		if qCreateTime != "" {

			if err := r.SetQueryParam("create_time", qCreateTime); err != nil {
				return err
			}
		}
	}

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

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSoftwarePackagesCollectionGet binds the parameter fields
func (o *SoftwarePackagesCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSoftwarePackagesCollectionGet binds the parameter order_by
func (o *SoftwarePackagesCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
