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

// NewTopMetricsSvmClientCollectionGetParams creates a new TopMetricsSvmClientCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTopMetricsSvmClientCollectionGetParams() *TopMetricsSvmClientCollectionGetParams {
	return &TopMetricsSvmClientCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTopMetricsSvmClientCollectionGetParamsWithTimeout creates a new TopMetricsSvmClientCollectionGetParams object
// with the ability to set a timeout on a request.
func NewTopMetricsSvmClientCollectionGetParamsWithTimeout(timeout time.Duration) *TopMetricsSvmClientCollectionGetParams {
	return &TopMetricsSvmClientCollectionGetParams{
		timeout: timeout,
	}
}

// NewTopMetricsSvmClientCollectionGetParamsWithContext creates a new TopMetricsSvmClientCollectionGetParams object
// with the ability to set a context for a request.
func NewTopMetricsSvmClientCollectionGetParamsWithContext(ctx context.Context) *TopMetricsSvmClientCollectionGetParams {
	return &TopMetricsSvmClientCollectionGetParams{
		Context: ctx,
	}
}

// NewTopMetricsSvmClientCollectionGetParamsWithHTTPClient creates a new TopMetricsSvmClientCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewTopMetricsSvmClientCollectionGetParamsWithHTTPClient(client *http.Client) *TopMetricsSvmClientCollectionGetParams {
	return &TopMetricsSvmClientCollectionGetParams{
		HTTPClient: client,
	}
}

/*
TopMetricsSvmClientCollectionGetParams contains all the parameters to send to the API endpoint

	for the top metrics svm client collection get operation.

	Typically these are written to a http.Request.
*/
type TopMetricsSvmClientCollectionGetParams struct {

	/* ClientIP.

	   Filter by client_ip
	*/
	ClientIP *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* IopsErrorLowerBound.

	   Filter by iops.error.lower_bound
	*/
	IopsErrorLowerBound *int64

	/* IopsErrorUpperBound.

	   Filter by iops.error.upper_bound
	*/
	IopsErrorUpperBound *int64

	/* IopsRead.

	   Filter by iops.read
	*/
	IopsRead *int64

	/* IopsWrite.

	   Filter by iops.write
	*/
	IopsWrite *int64

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

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   SVM UUID
	*/
	SvmUUID string

	/* ThroughputErrorLowerBound.

	   Filter by throughput.error.lower_bound
	*/
	ThroughputErrorLowerBound *int64

	/* ThroughputErrorUpperBound.

	   Filter by throughput.error.upper_bound
	*/
	ThroughputErrorUpperBound *int64

	/* ThroughputRead.

	   Filter by throughput.read
	*/
	ThroughputRead *int64

	/* ThroughputWrite.

	   Filter by throughput.write
	*/
	ThroughputWrite *int64

	/* TopMetric.

	   IO activity type

	   Default: "iops.read"
	*/
	TopMetric *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the top metrics svm client collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TopMetricsSvmClientCollectionGetParams) WithDefaults() *TopMetricsSvmClientCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the top metrics svm client collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TopMetricsSvmClientCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		topMetricDefault = string("iops.read")
	)

	val := TopMetricsSvmClientCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
		TopMetric:     &topMetricDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithTimeout(timeout time.Duration) *TopMetricsSvmClientCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithContext(ctx context.Context) *TopMetricsSvmClientCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithHTTPClient(client *http.Client) *TopMetricsSvmClientCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientIP adds the clientIP to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithClientIP(clientIP *string) *TopMetricsSvmClientCollectionGetParams {
	o.SetClientIP(clientIP)
	return o
}

// SetClientIP adds the clientIp to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetClientIP(clientIP *string) {
	o.ClientIP = clientIP
}

// WithFields adds the fields to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithFields(fields []string) *TopMetricsSvmClientCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIopsErrorLowerBound adds the iopsErrorLowerBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithIopsErrorLowerBound(iopsErrorLowerBound *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetIopsErrorLowerBound(iopsErrorLowerBound)
	return o
}

// SetIopsErrorLowerBound adds the iopsErrorLowerBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetIopsErrorLowerBound(iopsErrorLowerBound *int64) {
	o.IopsErrorLowerBound = iopsErrorLowerBound
}

// WithIopsErrorUpperBound adds the iopsErrorUpperBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithIopsErrorUpperBound(iopsErrorUpperBound *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetIopsErrorUpperBound(iopsErrorUpperBound)
	return o
}

// SetIopsErrorUpperBound adds the iopsErrorUpperBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetIopsErrorUpperBound(iopsErrorUpperBound *int64) {
	o.IopsErrorUpperBound = iopsErrorUpperBound
}

// WithIopsRead adds the iopsRead to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithIopsRead(iopsRead *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetIopsRead(iopsRead)
	return o
}

// SetIopsRead adds the iopsRead to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetIopsRead(iopsRead *int64) {
	o.IopsRead = iopsRead
}

// WithIopsWrite adds the iopsWrite to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithIopsWrite(iopsWrite *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetIopsWrite(iopsWrite)
	return o
}

// SetIopsWrite adds the iopsWrite to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetIopsWrite(iopsWrite *int64) {
	o.IopsWrite = iopsWrite
}

// WithMaxRecords adds the maxRecords to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithMaxRecords(maxRecords *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithOrderBy(orderBy []string) *TopMetricsSvmClientCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithReturnRecords(returnRecords *bool) *TopMetricsSvmClientCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmName adds the svmName to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithSvmName(svmName *string) *TopMetricsSvmClientCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithSvmUUID(svmUUID string) *TopMetricsSvmClientCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WithThroughputErrorLowerBound adds the throughputErrorLowerBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithThroughputErrorLowerBound(throughputErrorLowerBound *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetThroughputErrorLowerBound(throughputErrorLowerBound)
	return o
}

// SetThroughputErrorLowerBound adds the throughputErrorLowerBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetThroughputErrorLowerBound(throughputErrorLowerBound *int64) {
	o.ThroughputErrorLowerBound = throughputErrorLowerBound
}

// WithThroughputErrorUpperBound adds the throughputErrorUpperBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithThroughputErrorUpperBound(throughputErrorUpperBound *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetThroughputErrorUpperBound(throughputErrorUpperBound)
	return o
}

// SetThroughputErrorUpperBound adds the throughputErrorUpperBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetThroughputErrorUpperBound(throughputErrorUpperBound *int64) {
	o.ThroughputErrorUpperBound = throughputErrorUpperBound
}

// WithThroughputRead adds the throughputRead to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithThroughputRead(throughputRead *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetThroughputRead(throughputRead)
	return o
}

// SetThroughputRead adds the throughputRead to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetThroughputRead(throughputRead *int64) {
	o.ThroughputRead = throughputRead
}

// WithThroughputWrite adds the throughputWrite to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithThroughputWrite(throughputWrite *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetThroughputWrite(throughputWrite)
	return o
}

// SetThroughputWrite adds the throughputWrite to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetThroughputWrite(throughputWrite *int64) {
	o.ThroughputWrite = throughputWrite
}

// WithTopMetric adds the topMetric to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithTopMetric(topMetric *string) *TopMetricsSvmClientCollectionGetParams {
	o.SetTopMetric(topMetric)
	return o
}

// SetTopMetric adds the topMetric to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetTopMetric(topMetric *string) {
	o.TopMetric = topMetric
}

// WriteToRequest writes these params to a swagger request
func (o *TopMetricsSvmClientCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientIP != nil {

		// query param client_ip
		var qrClientIP string

		if o.ClientIP != nil {
			qrClientIP = *o.ClientIP
		}
		qClientIP := qrClientIP
		if qClientIP != "" {

			if err := r.SetQueryParam("client_ip", qClientIP); err != nil {
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

	if o.IopsErrorLowerBound != nil {

		// query param iops.error.lower_bound
		var qrIopsErrorLowerBound int64

		if o.IopsErrorLowerBound != nil {
			qrIopsErrorLowerBound = *o.IopsErrorLowerBound
		}
		qIopsErrorLowerBound := swag.FormatInt64(qrIopsErrorLowerBound)
		if qIopsErrorLowerBound != "" {

			if err := r.SetQueryParam("iops.error.lower_bound", qIopsErrorLowerBound); err != nil {
				return err
			}
		}
	}

	if o.IopsErrorUpperBound != nil {

		// query param iops.error.upper_bound
		var qrIopsErrorUpperBound int64

		if o.IopsErrorUpperBound != nil {
			qrIopsErrorUpperBound = *o.IopsErrorUpperBound
		}
		qIopsErrorUpperBound := swag.FormatInt64(qrIopsErrorUpperBound)
		if qIopsErrorUpperBound != "" {

			if err := r.SetQueryParam("iops.error.upper_bound", qIopsErrorUpperBound); err != nil {
				return err
			}
		}
	}

	if o.IopsRead != nil {

		// query param iops.read
		var qrIopsRead int64

		if o.IopsRead != nil {
			qrIopsRead = *o.IopsRead
		}
		qIopsRead := swag.FormatInt64(qrIopsRead)
		if qIopsRead != "" {

			if err := r.SetQueryParam("iops.read", qIopsRead); err != nil {
				return err
			}
		}
	}

	if o.IopsWrite != nil {

		// query param iops.write
		var qrIopsWrite int64

		if o.IopsWrite != nil {
			qrIopsWrite = *o.IopsWrite
		}
		qIopsWrite := swag.FormatInt64(qrIopsWrite)
		if qIopsWrite != "" {

			if err := r.SetQueryParam("iops.write", qIopsWrite); err != nil {
				return err
			}
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

	if o.SvmName != nil {

		// query param svm.name
		var qrSvmName string

		if o.SvmName != nil {
			qrSvmName = *o.SvmName
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if o.ThroughputErrorLowerBound != nil {

		// query param throughput.error.lower_bound
		var qrThroughputErrorLowerBound int64

		if o.ThroughputErrorLowerBound != nil {
			qrThroughputErrorLowerBound = *o.ThroughputErrorLowerBound
		}
		qThroughputErrorLowerBound := swag.FormatInt64(qrThroughputErrorLowerBound)
		if qThroughputErrorLowerBound != "" {

			if err := r.SetQueryParam("throughput.error.lower_bound", qThroughputErrorLowerBound); err != nil {
				return err
			}
		}
	}

	if o.ThroughputErrorUpperBound != nil {

		// query param throughput.error.upper_bound
		var qrThroughputErrorUpperBound int64

		if o.ThroughputErrorUpperBound != nil {
			qrThroughputErrorUpperBound = *o.ThroughputErrorUpperBound
		}
		qThroughputErrorUpperBound := swag.FormatInt64(qrThroughputErrorUpperBound)
		if qThroughputErrorUpperBound != "" {

			if err := r.SetQueryParam("throughput.error.upper_bound", qThroughputErrorUpperBound); err != nil {
				return err
			}
		}
	}

	if o.ThroughputRead != nil {

		// query param throughput.read
		var qrThroughputRead int64

		if o.ThroughputRead != nil {
			qrThroughputRead = *o.ThroughputRead
		}
		qThroughputRead := swag.FormatInt64(qrThroughputRead)
		if qThroughputRead != "" {

			if err := r.SetQueryParam("throughput.read", qThroughputRead); err != nil {
				return err
			}
		}
	}

	if o.ThroughputWrite != nil {

		// query param throughput.write
		var qrThroughputWrite int64

		if o.ThroughputWrite != nil {
			qrThroughputWrite = *o.ThroughputWrite
		}
		qThroughputWrite := swag.FormatInt64(qrThroughputWrite)
		if qThroughputWrite != "" {

			if err := r.SetQueryParam("throughput.write", qThroughputWrite); err != nil {
				return err
			}
		}
	}

	if o.TopMetric != nil {

		// query param top_metric
		var qrTopMetric string

		if o.TopMetric != nil {
			qrTopMetric = *o.TopMetric
		}
		qTopMetric := qrTopMetric
		if qTopMetric != "" {

			if err := r.SetQueryParam("top_metric", qTopMetric); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamTopMetricsSvmClientCollectionGet binds the parameter fields
func (o *TopMetricsSvmClientCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamTopMetricsSvmClientCollectionGet binds the parameter order_by
func (o *TopMetricsSvmClientCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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