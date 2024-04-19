// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewGcpKmsRekeyInternalParams creates a new GcpKmsRekeyInternalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGcpKmsRekeyInternalParams() *GcpKmsRekeyInternalParams {
	return &GcpKmsRekeyInternalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGcpKmsRekeyInternalParamsWithTimeout creates a new GcpKmsRekeyInternalParams object
// with the ability to set a timeout on a request.
func NewGcpKmsRekeyInternalParamsWithTimeout(timeout time.Duration) *GcpKmsRekeyInternalParams {
	return &GcpKmsRekeyInternalParams{
		timeout: timeout,
	}
}

// NewGcpKmsRekeyInternalParamsWithContext creates a new GcpKmsRekeyInternalParams object
// with the ability to set a context for a request.
func NewGcpKmsRekeyInternalParamsWithContext(ctx context.Context) *GcpKmsRekeyInternalParams {
	return &GcpKmsRekeyInternalParams{
		Context: ctx,
	}
}

// NewGcpKmsRekeyInternalParamsWithHTTPClient creates a new GcpKmsRekeyInternalParams object
// with the ability to set a custom HTTPClient for a request.
func NewGcpKmsRekeyInternalParamsWithHTTPClient(client *http.Client) *GcpKmsRekeyInternalParams {
	return &GcpKmsRekeyInternalParams{
		HTTPClient: client,
	}
}

/*
GcpKmsRekeyInternalParams contains all the parameters to send to the API endpoint

	for the gcp kms rekey internal operation.

	Typically these are written to a http.Request.
*/
type GcpKmsRekeyInternalParams struct {

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   UUID of the existing Google Cloud KMS configuration.
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gcp kms rekey internal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GcpKmsRekeyInternalParams) WithDefaults() *GcpKmsRekeyInternalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gcp kms rekey internal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GcpKmsRekeyInternalParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := GcpKmsRekeyInternalParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the gcp kms rekey internal params
func (o *GcpKmsRekeyInternalParams) WithTimeout(timeout time.Duration) *GcpKmsRekeyInternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gcp kms rekey internal params
func (o *GcpKmsRekeyInternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gcp kms rekey internal params
func (o *GcpKmsRekeyInternalParams) WithContext(ctx context.Context) *GcpKmsRekeyInternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gcp kms rekey internal params
func (o *GcpKmsRekeyInternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gcp kms rekey internal params
func (o *GcpKmsRekeyInternalParams) WithHTTPClient(client *http.Client) *GcpKmsRekeyInternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gcp kms rekey internal params
func (o *GcpKmsRekeyInternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnRecords adds the returnRecords to the gcp kms rekey internal params
func (o *GcpKmsRekeyInternalParams) WithReturnRecords(returnRecords *bool) *GcpKmsRekeyInternalParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the gcp kms rekey internal params
func (o *GcpKmsRekeyInternalParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the gcp kms rekey internal params
func (o *GcpKmsRekeyInternalParams) WithReturnTimeout(returnTimeout *int64) *GcpKmsRekeyInternalParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the gcp kms rekey internal params
func (o *GcpKmsRekeyInternalParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the gcp kms rekey internal params
func (o *GcpKmsRekeyInternalParams) WithUUID(uuid string) *GcpKmsRekeyInternalParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the gcp kms rekey internal params
func (o *GcpKmsRekeyInternalParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GcpKmsRekeyInternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}