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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewMetroclusterInterconnectModifyParams creates a new MetroclusterInterconnectModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMetroclusterInterconnectModifyParams() *MetroclusterInterconnectModifyParams {
	return &MetroclusterInterconnectModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMetroclusterInterconnectModifyParamsWithTimeout creates a new MetroclusterInterconnectModifyParams object
// with the ability to set a timeout on a request.
func NewMetroclusterInterconnectModifyParamsWithTimeout(timeout time.Duration) *MetroclusterInterconnectModifyParams {
	return &MetroclusterInterconnectModifyParams{
		timeout: timeout,
	}
}

// NewMetroclusterInterconnectModifyParamsWithContext creates a new MetroclusterInterconnectModifyParams object
// with the ability to set a context for a request.
func NewMetroclusterInterconnectModifyParamsWithContext(ctx context.Context) *MetroclusterInterconnectModifyParams {
	return &MetroclusterInterconnectModifyParams{
		Context: ctx,
	}
}

// NewMetroclusterInterconnectModifyParamsWithHTTPClient creates a new MetroclusterInterconnectModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewMetroclusterInterconnectModifyParamsWithHTTPClient(client *http.Client) *MetroclusterInterconnectModifyParams {
	return &MetroclusterInterconnectModifyParams{
		HTTPClient: client,
	}
}

/*
MetroclusterInterconnectModifyParams contains all the parameters to send to the API endpoint

	for the metrocluster interconnect modify operation.

	Typically these are written to a http.Request.
*/
type MetroclusterInterconnectModifyParams struct {

	/* Adapter.

	   Interconnect adapter
	*/
	Adapter string

	/* Info.

	   MetroCluster interconnect interface information
	*/
	Info *models.MetroclusterInterconnect

	/* NodeUUID.

	   Node UUID
	*/
	NodeUUID string

	/* PartnerType.

	   DR Partner type
	*/
	PartnerType string

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the metrocluster interconnect modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroclusterInterconnectModifyParams) WithDefaults() *MetroclusterInterconnectModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the metrocluster interconnect modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroclusterInterconnectModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := MetroclusterInterconnectModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithTimeout(timeout time.Duration) *MetroclusterInterconnectModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithContext(ctx context.Context) *MetroclusterInterconnectModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithHTTPClient(client *http.Client) *MetroclusterInterconnectModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdapter adds the adapter to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithAdapter(adapter string) *MetroclusterInterconnectModifyParams {
	o.SetAdapter(adapter)
	return o
}

// SetAdapter adds the adapter to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetAdapter(adapter string) {
	o.Adapter = adapter
}

// WithInfo adds the info to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithInfo(info *models.MetroclusterInterconnect) *MetroclusterInterconnectModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetInfo(info *models.MetroclusterInterconnect) {
	o.Info = info
}

// WithNodeUUID adds the nodeUUID to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithNodeUUID(nodeUUID string) *MetroclusterInterconnectModifyParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetNodeUUID(nodeUUID string) {
	o.NodeUUID = nodeUUID
}

// WithPartnerType adds the partnerType to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithPartnerType(partnerType string) *MetroclusterInterconnectModifyParams {
	o.SetPartnerType(partnerType)
	return o
}

// SetPartnerType adds the partnerType to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetPartnerType(partnerType string) {
	o.PartnerType = partnerType
}

// WithReturnTimeout adds the returnTimeout to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) WithReturnTimeout(returnTimeout *int64) *MetroclusterInterconnectModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the metrocluster interconnect modify params
func (o *MetroclusterInterconnectModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *MetroclusterInterconnectModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param adapter
	if err := r.SetPathParam("adapter", o.Adapter); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param node.uuid
	if err := r.SetPathParam("node.uuid", o.NodeUUID); err != nil {
		return err
	}

	// path param partner_type
	if err := r.SetPathParam("partner_type", o.PartnerType); err != nil {
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
