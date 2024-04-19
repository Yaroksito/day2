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

// NewApplicationTemplateCollectionGetParams creates a new ApplicationTemplateCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationTemplateCollectionGetParams() *ApplicationTemplateCollectionGetParams {
	return &ApplicationTemplateCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationTemplateCollectionGetParamsWithTimeout creates a new ApplicationTemplateCollectionGetParams object
// with the ability to set a timeout on a request.
func NewApplicationTemplateCollectionGetParamsWithTimeout(timeout time.Duration) *ApplicationTemplateCollectionGetParams {
	return &ApplicationTemplateCollectionGetParams{
		timeout: timeout,
	}
}

// NewApplicationTemplateCollectionGetParamsWithContext creates a new ApplicationTemplateCollectionGetParams object
// with the ability to set a context for a request.
func NewApplicationTemplateCollectionGetParamsWithContext(ctx context.Context) *ApplicationTemplateCollectionGetParams {
	return &ApplicationTemplateCollectionGetParams{
		Context: ctx,
	}
}

// NewApplicationTemplateCollectionGetParamsWithHTTPClient creates a new ApplicationTemplateCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationTemplateCollectionGetParamsWithHTTPClient(client *http.Client) *ApplicationTemplateCollectionGetParams {
	return &ApplicationTemplateCollectionGetParams{
		HTTPClient: client,
	}
}

/*
ApplicationTemplateCollectionGetParams contains all the parameters to send to the API endpoint

	for the application template collection get operation.

	Typically these are written to a http.Request.
*/
type ApplicationTemplateCollectionGetParams struct {

	/* Description.

	   Filter by description
	*/
	Description *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* MissingPrerequisites.

	   Filter by missing_prerequisites
	*/
	MissingPrerequisites *string

	/* Name.

	   Filter by name
	*/
	Name *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* Protocol.

	   Filter by protocol
	*/
	Protocol *string

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

// WithDefaults hydrates default values in the application template collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationTemplateCollectionGetParams) WithDefaults() *ApplicationTemplateCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application template collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationTemplateCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := ApplicationTemplateCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) WithTimeout(timeout time.Duration) *ApplicationTemplateCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) WithContext(ctx context.Context) *ApplicationTemplateCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) WithHTTPClient(client *http.Client) *ApplicationTemplateCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) WithDescription(description *string) *ApplicationTemplateCollectionGetParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) SetDescription(description *string) {
	o.Description = description
}

// WithFields adds the fields to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) WithFields(fields []string) *ApplicationTemplateCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) WithMaxRecords(maxRecords *int64) *ApplicationTemplateCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithMissingPrerequisites adds the missingPrerequisites to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) WithMissingPrerequisites(missingPrerequisites *string) *ApplicationTemplateCollectionGetParams {
	o.SetMissingPrerequisites(missingPrerequisites)
	return o
}

// SetMissingPrerequisites adds the missingPrerequisites to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) SetMissingPrerequisites(missingPrerequisites *string) {
	o.MissingPrerequisites = missingPrerequisites
}

// WithName adds the name to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) WithName(name *string) *ApplicationTemplateCollectionGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) SetName(name *string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) WithOrderBy(orderBy []string) *ApplicationTemplateCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithProtocol adds the protocol to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) WithProtocol(protocol *string) *ApplicationTemplateCollectionGetParams {
	o.SetProtocol(protocol)
	return o
}

// SetProtocol adds the protocol to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) SetProtocol(protocol *string) {
	o.Protocol = protocol
}

// WithReturnRecords adds the returnRecords to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) WithReturnRecords(returnRecords *bool) *ApplicationTemplateCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *ApplicationTemplateCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the application template collection get params
func (o *ApplicationTemplateCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationTemplateCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
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

	if o.MissingPrerequisites != nil {

		// query param missing_prerequisites
		var qrMissingPrerequisites string

		if o.MissingPrerequisites != nil {
			qrMissingPrerequisites = *o.MissingPrerequisites
		}
		qMissingPrerequisites := qrMissingPrerequisites
		if qMissingPrerequisites != "" {

			if err := r.SetQueryParam("missing_prerequisites", qMissingPrerequisites); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
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

	if o.Protocol != nil {

		// query param protocol
		var qrProtocol string

		if o.Protocol != nil {
			qrProtocol = *o.Protocol
		}
		qProtocol := qrProtocol
		if qProtocol != "" {

			if err := r.SetQueryParam("protocol", qProtocol); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamApplicationTemplateCollectionGet binds the parameter fields
func (o *ApplicationTemplateCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamApplicationTemplateCollectionGet binds the parameter order_by
func (o *ApplicationTemplateCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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