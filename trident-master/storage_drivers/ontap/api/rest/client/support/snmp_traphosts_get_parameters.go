// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewSnmpTraphostsGetParams creates a new SnmpTraphostsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnmpTraphostsGetParams() *SnmpTraphostsGetParams {
	return &SnmpTraphostsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnmpTraphostsGetParamsWithTimeout creates a new SnmpTraphostsGetParams object
// with the ability to set a timeout on a request.
func NewSnmpTraphostsGetParamsWithTimeout(timeout time.Duration) *SnmpTraphostsGetParams {
	return &SnmpTraphostsGetParams{
		timeout: timeout,
	}
}

// NewSnmpTraphostsGetParamsWithContext creates a new SnmpTraphostsGetParams object
// with the ability to set a context for a request.
func NewSnmpTraphostsGetParamsWithContext(ctx context.Context) *SnmpTraphostsGetParams {
	return &SnmpTraphostsGetParams{
		Context: ctx,
	}
}

// NewSnmpTraphostsGetParamsWithHTTPClient creates a new SnmpTraphostsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnmpTraphostsGetParamsWithHTTPClient(client *http.Client) *SnmpTraphostsGetParams {
	return &SnmpTraphostsGetParams{
		HTTPClient: client,
	}
}

/*
SnmpTraphostsGetParams contains all the parameters to send to the API endpoint

	for the snmp traphosts get operation.

	Typically these are written to a http.Request.
*/
type SnmpTraphostsGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Host.

	   Fully Qualified Domain Name (FQDN), IPv4 address or IPv6 address of SNMP traphost.
	*/
	Host string

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snmp traphosts get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnmpTraphostsGetParams) WithDefaults() *SnmpTraphostsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snmp traphosts get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnmpTraphostsGetParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(15)
	)

	val := SnmpTraphostsGetParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snmp traphosts get params
func (o *SnmpTraphostsGetParams) WithTimeout(timeout time.Duration) *SnmpTraphostsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snmp traphosts get params
func (o *SnmpTraphostsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snmp traphosts get params
func (o *SnmpTraphostsGetParams) WithContext(ctx context.Context) *SnmpTraphostsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snmp traphosts get params
func (o *SnmpTraphostsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snmp traphosts get params
func (o *SnmpTraphostsGetParams) WithHTTPClient(client *http.Client) *SnmpTraphostsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snmp traphosts get params
func (o *SnmpTraphostsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the snmp traphosts get params
func (o *SnmpTraphostsGetParams) WithFields(fields []string) *SnmpTraphostsGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the snmp traphosts get params
func (o *SnmpTraphostsGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithHost adds the host to the snmp traphosts get params
func (o *SnmpTraphostsGetParams) WithHost(host string) *SnmpTraphostsGetParams {
	o.SetHost(host)
	return o
}

// SetHost adds the host to the snmp traphosts get params
func (o *SnmpTraphostsGetParams) SetHost(host string) {
	o.Host = host
}

// WithReturnTimeout adds the returnTimeout to the snmp traphosts get params
func (o *SnmpTraphostsGetParams) WithReturnTimeout(returnTimeout *int64) *SnmpTraphostsGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the snmp traphosts get params
func (o *SnmpTraphostsGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *SnmpTraphostsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param host
	if err := r.SetPathParam("host", o.Host); err != nil {
		return err
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

// bindParamSnmpTraphostsGet binds the parameter fields
func (o *SnmpTraphostsGetParams) bindParamFields(formats strfmt.Registry) []string {
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