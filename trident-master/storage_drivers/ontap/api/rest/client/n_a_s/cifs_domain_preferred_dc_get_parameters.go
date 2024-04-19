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
)

// NewCifsDomainPreferredDcGetParams creates a new CifsDomainPreferredDcGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsDomainPreferredDcGetParams() *CifsDomainPreferredDcGetParams {
	return &CifsDomainPreferredDcGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsDomainPreferredDcGetParamsWithTimeout creates a new CifsDomainPreferredDcGetParams object
// with the ability to set a timeout on a request.
func NewCifsDomainPreferredDcGetParamsWithTimeout(timeout time.Duration) *CifsDomainPreferredDcGetParams {
	return &CifsDomainPreferredDcGetParams{
		timeout: timeout,
	}
}

// NewCifsDomainPreferredDcGetParamsWithContext creates a new CifsDomainPreferredDcGetParams object
// with the ability to set a context for a request.
func NewCifsDomainPreferredDcGetParamsWithContext(ctx context.Context) *CifsDomainPreferredDcGetParams {
	return &CifsDomainPreferredDcGetParams{
		Context: ctx,
	}
}

// NewCifsDomainPreferredDcGetParamsWithHTTPClient creates a new CifsDomainPreferredDcGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsDomainPreferredDcGetParamsWithHTTPClient(client *http.Client) *CifsDomainPreferredDcGetParams {
	return &CifsDomainPreferredDcGetParams{
		HTTPClient: client,
	}
}

/*
CifsDomainPreferredDcGetParams contains all the parameters to send to the API endpoint

	for the cifs domain preferred dc get operation.

	Typically these are written to a http.Request.
*/
type CifsDomainPreferredDcGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Fqdn.

	   Fully Qualified Domain Name
	*/
	Fqdn string

	/* NeedStatus.

	   Retrieves the status of the specified preferred DC.
	*/
	NeedStatus *bool

	/* ServerIP.

	   Domain Controller IP address
	*/
	ServerIP string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs domain preferred dc get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsDomainPreferredDcGetParams) WithDefaults() *CifsDomainPreferredDcGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs domain preferred dc get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsDomainPreferredDcGetParams) SetDefaults() {
	var (
		needStatusDefault = bool(false)
	)

	val := CifsDomainPreferredDcGetParams{
		NeedStatus: &needStatusDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) WithTimeout(timeout time.Duration) *CifsDomainPreferredDcGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) WithContext(ctx context.Context) *CifsDomainPreferredDcGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) WithHTTPClient(client *http.Client) *CifsDomainPreferredDcGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) WithFields(fields []string) *CifsDomainPreferredDcGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithFqdn adds the fqdn to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) WithFqdn(fqdn string) *CifsDomainPreferredDcGetParams {
	o.SetFqdn(fqdn)
	return o
}

// SetFqdn adds the fqdn to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) SetFqdn(fqdn string) {
	o.Fqdn = fqdn
}

// WithNeedStatus adds the needStatus to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) WithNeedStatus(needStatus *bool) *CifsDomainPreferredDcGetParams {
	o.SetNeedStatus(needStatus)
	return o
}

// SetNeedStatus adds the needStatus to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) SetNeedStatus(needStatus *bool) {
	o.NeedStatus = needStatus
}

// WithServerIP adds the serverIP to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) WithServerIP(serverIP string) *CifsDomainPreferredDcGetParams {
	o.SetServerIP(serverIP)
	return o
}

// SetServerIP adds the serverIp to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) SetServerIP(serverIP string) {
	o.ServerIP = serverIP
}

// WithSvmUUID adds the svmUUID to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) WithSvmUUID(svmUUID string) *CifsDomainPreferredDcGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the cifs domain preferred dc get params
func (o *CifsDomainPreferredDcGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsDomainPreferredDcGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param fqdn
	if err := r.SetPathParam("fqdn", o.Fqdn); err != nil {
		return err
	}

	if o.NeedStatus != nil {

		// query param need_status
		var qrNeedStatus bool

		if o.NeedStatus != nil {
			qrNeedStatus = *o.NeedStatus
		}
		qNeedStatus := swag.FormatBool(qrNeedStatus)
		if qNeedStatus != "" {

			if err := r.SetQueryParam("need_status", qNeedStatus); err != nil {
				return err
			}
		}
	}

	// path param server_ip
	if err := r.SetPathParam("server_ip", o.ServerIP); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCifsDomainPreferredDcGet binds the parameter fields
func (o *CifsDomainPreferredDcGetParams) bindParamFields(formats strfmt.Registry) []string {
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