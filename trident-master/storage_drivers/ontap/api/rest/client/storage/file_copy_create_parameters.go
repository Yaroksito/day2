// Code generated by go-swagger; DO NOT EDIT.

package storage

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewFileCopyCreateParams creates a new FileCopyCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFileCopyCreateParams() *FileCopyCreateParams {
	return &FileCopyCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFileCopyCreateParamsWithTimeout creates a new FileCopyCreateParams object
// with the ability to set a timeout on a request.
func NewFileCopyCreateParamsWithTimeout(timeout time.Duration) *FileCopyCreateParams {
	return &FileCopyCreateParams{
		timeout: timeout,
	}
}

// NewFileCopyCreateParamsWithContext creates a new FileCopyCreateParams object
// with the ability to set a context for a request.
func NewFileCopyCreateParamsWithContext(ctx context.Context) *FileCopyCreateParams {
	return &FileCopyCreateParams{
		Context: ctx,
	}
}

// NewFileCopyCreateParamsWithHTTPClient creates a new FileCopyCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewFileCopyCreateParamsWithHTTPClient(client *http.Client) *FileCopyCreateParams {
	return &FileCopyCreateParams{
		HTTPClient: client,
	}
}

/*
FileCopyCreateParams contains all the parameters to send to the API endpoint

	for the file copy create operation.

	Typically these are written to a http.Request.
*/
type FileCopyCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.FileCopy

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the file copy create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileCopyCreateParams) WithDefaults() *FileCopyCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the file copy create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileCopyCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := FileCopyCreateParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the file copy create params
func (o *FileCopyCreateParams) WithTimeout(timeout time.Duration) *FileCopyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the file copy create params
func (o *FileCopyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the file copy create params
func (o *FileCopyCreateParams) WithContext(ctx context.Context) *FileCopyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the file copy create params
func (o *FileCopyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the file copy create params
func (o *FileCopyCreateParams) WithHTTPClient(client *http.Client) *FileCopyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the file copy create params
func (o *FileCopyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the file copy create params
func (o *FileCopyCreateParams) WithInfo(info *models.FileCopy) *FileCopyCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the file copy create params
func (o *FileCopyCreateParams) SetInfo(info *models.FileCopy) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the file copy create params
func (o *FileCopyCreateParams) WithReturnRecords(returnRecords *bool) *FileCopyCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the file copy create params
func (o *FileCopyCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the file copy create params
func (o *FileCopyCreateParams) WithReturnTimeout(returnTimeout *int64) *FileCopyCreateParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the file copy create params
func (o *FileCopyCreateParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *FileCopyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
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