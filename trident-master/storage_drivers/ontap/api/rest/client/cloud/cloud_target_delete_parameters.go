// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

// NewCloudTargetDeleteParams creates a new CloudTargetDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloudTargetDeleteParams() *CloudTargetDeleteParams {
	return &CloudTargetDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloudTargetDeleteParamsWithTimeout creates a new CloudTargetDeleteParams object
// with the ability to set a timeout on a request.
func NewCloudTargetDeleteParamsWithTimeout(timeout time.Duration) *CloudTargetDeleteParams {
	return &CloudTargetDeleteParams{
		timeout: timeout,
	}
}

// NewCloudTargetDeleteParamsWithContext creates a new CloudTargetDeleteParams object
// with the ability to set a context for a request.
func NewCloudTargetDeleteParamsWithContext(ctx context.Context) *CloudTargetDeleteParams {
	return &CloudTargetDeleteParams{
		Context: ctx,
	}
}

// NewCloudTargetDeleteParamsWithHTTPClient creates a new CloudTargetDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloudTargetDeleteParamsWithHTTPClient(client *http.Client) *CloudTargetDeleteParams {
	return &CloudTargetDeleteParams{
		HTTPClient: client,
	}
}

/*
CloudTargetDeleteParams contains all the parameters to send to the API endpoint

	for the cloud target delete operation.

	Typically these are written to a http.Request.
*/
type CloudTargetDeleteParams struct {

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   Cloud target UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cloud target delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudTargetDeleteParams) WithDefaults() *CloudTargetDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cloud target delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudTargetDeleteParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := CloudTargetDeleteParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cloud target delete params
func (o *CloudTargetDeleteParams) WithTimeout(timeout time.Duration) *CloudTargetDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud target delete params
func (o *CloudTargetDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud target delete params
func (o *CloudTargetDeleteParams) WithContext(ctx context.Context) *CloudTargetDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud target delete params
func (o *CloudTargetDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud target delete params
func (o *CloudTargetDeleteParams) WithHTTPClient(client *http.Client) *CloudTargetDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud target delete params
func (o *CloudTargetDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnTimeout adds the returnTimeout to the cloud target delete params
func (o *CloudTargetDeleteParams) WithReturnTimeout(returnTimeout *int64) *CloudTargetDeleteParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the cloud target delete params
func (o *CloudTargetDeleteParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the cloud target delete params
func (o *CloudTargetDeleteParams) WithUUID(uuid string) *CloudTargetDeleteParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the cloud target delete params
func (o *CloudTargetDeleteParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *CloudTargetDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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