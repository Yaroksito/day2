// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NewNameMappingCollectionGetParams creates a new NameMappingCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNameMappingCollectionGetParams() *NameMappingCollectionGetParams {
	return &NameMappingCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNameMappingCollectionGetParamsWithTimeout creates a new NameMappingCollectionGetParams object
// with the ability to set a timeout on a request.
func NewNameMappingCollectionGetParamsWithTimeout(timeout time.Duration) *NameMappingCollectionGetParams {
	return &NameMappingCollectionGetParams{
		timeout: timeout,
	}
}

// NewNameMappingCollectionGetParamsWithContext creates a new NameMappingCollectionGetParams object
// with the ability to set a context for a request.
func NewNameMappingCollectionGetParamsWithContext(ctx context.Context) *NameMappingCollectionGetParams {
	return &NameMappingCollectionGetParams{
		Context: ctx,
	}
}

// NewNameMappingCollectionGetParamsWithHTTPClient creates a new NameMappingCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNameMappingCollectionGetParamsWithHTTPClient(client *http.Client) *NameMappingCollectionGetParams {
	return &NameMappingCollectionGetParams{
		HTTPClient: client,
	}
}

/*
NameMappingCollectionGetParams contains all the parameters to send to the API endpoint

	for the name mapping collection get operation.

	Typically these are written to a http.Request.
*/
type NameMappingCollectionGetParams struct {

	/* ClientMatch.

	   Filter by client_match
	*/
	ClientMatch *string

	/* Direction.

	   Filter by direction
	*/
	Direction *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Index.

	   Filter by index
	*/
	Index *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* Pattern.

	   Filter by pattern
	*/
	Pattern *string

	/* Replacement.

	   Filter by replacement
	*/
	Replacement *string

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

	   Filter by svm.uuid
	*/
	SvmUUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the name mapping collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NameMappingCollectionGetParams) WithDefaults() *NameMappingCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the name mapping collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NameMappingCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := NameMappingCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithTimeout(timeout time.Duration) *NameMappingCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithContext(ctx context.Context) *NameMappingCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithHTTPClient(client *http.Client) *NameMappingCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientMatch adds the clientMatch to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithClientMatch(clientMatch *string) *NameMappingCollectionGetParams {
	o.SetClientMatch(clientMatch)
	return o
}

// SetClientMatch adds the clientMatch to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetClientMatch(clientMatch *string) {
	o.ClientMatch = clientMatch
}

// WithDirection adds the direction to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithDirection(direction *string) *NameMappingCollectionGetParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithFields adds the fields to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithFields(fields []string) *NameMappingCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIndex adds the index to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithIndex(index *int64) *NameMappingCollectionGetParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetIndex(index *int64) {
	o.Index = index
}

// WithMaxRecords adds the maxRecords to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithMaxRecords(maxRecords *int64) *NameMappingCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithOrderBy(orderBy []string) *NameMappingCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPattern adds the pattern to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithPattern(pattern *string) *NameMappingCollectionGetParams {
	o.SetPattern(pattern)
	return o
}

// SetPattern adds the pattern to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetPattern(pattern *string) {
	o.Pattern = pattern
}

// WithReplacement adds the replacement to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithReplacement(replacement *string) *NameMappingCollectionGetParams {
	o.SetReplacement(replacement)
	return o
}

// SetReplacement adds the replacement to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetReplacement(replacement *string) {
	o.Replacement = replacement
}

// WithReturnRecords adds the returnRecords to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithReturnRecords(returnRecords *bool) *NameMappingCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *NameMappingCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmName adds the svmName to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithSvmName(svmName *string) *NameMappingCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the name mapping collection get params
func (o *NameMappingCollectionGetParams) WithSvmUUID(svmUUID *string) *NameMappingCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the name mapping collection get params
func (o *NameMappingCollectionGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NameMappingCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientMatch != nil {

		// query param client_match
		var qrClientMatch string

		if o.ClientMatch != nil {
			qrClientMatch = *o.ClientMatch
		}
		qClientMatch := qrClientMatch
		if qClientMatch != "" {

			if err := r.SetQueryParam("client_match", qClientMatch); err != nil {
				return err
			}
		}
	}

	if o.Direction != nil {

		// query param direction
		var qrDirection string

		if o.Direction != nil {
			qrDirection = *o.Direction
		}
		qDirection := qrDirection
		if qDirection != "" {

			if err := r.SetQueryParam("direction", qDirection); err != nil {
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

	if o.Index != nil {

		// query param index
		var qrIndex int64

		if o.Index != nil {
			qrIndex = *o.Index
		}
		qIndex := swag.FormatInt64(qrIndex)
		if qIndex != "" {

			if err := r.SetQueryParam("index", qIndex); err != nil {
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

	if o.Pattern != nil {

		// query param pattern
		var qrPattern string

		if o.Pattern != nil {
			qrPattern = *o.Pattern
		}
		qPattern := qrPattern
		if qPattern != "" {

			if err := r.SetQueryParam("pattern", qPattern); err != nil {
				return err
			}
		}
	}

	if o.Replacement != nil {

		// query param replacement
		var qrReplacement string

		if o.Replacement != nil {
			qrReplacement = *o.Replacement
		}
		qReplacement := qrReplacement
		if qReplacement != "" {

			if err := r.SetQueryParam("replacement", qReplacement); err != nil {
				return err
			}
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

	if o.SvmUUID != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SvmUUID != nil {
			qrSvmUUID = *o.SvmUUID
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNameMappingCollectionGet binds the parameter fields
func (o *NameMappingCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNameMappingCollectionGet binds the parameter order_by
func (o *NameMappingCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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