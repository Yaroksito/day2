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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewStoragePortModifyParams creates a new StoragePortModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStoragePortModifyParams() *StoragePortModifyParams {
	return &StoragePortModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStoragePortModifyParamsWithTimeout creates a new StoragePortModifyParams object
// with the ability to set a timeout on a request.
func NewStoragePortModifyParamsWithTimeout(timeout time.Duration) *StoragePortModifyParams {
	return &StoragePortModifyParams{
		timeout: timeout,
	}
}

// NewStoragePortModifyParamsWithContext creates a new StoragePortModifyParams object
// with the ability to set a context for a request.
func NewStoragePortModifyParamsWithContext(ctx context.Context) *StoragePortModifyParams {
	return &StoragePortModifyParams{
		Context: ctx,
	}
}

// NewStoragePortModifyParamsWithHTTPClient creates a new StoragePortModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewStoragePortModifyParamsWithHTTPClient(client *http.Client) *StoragePortModifyParams {
	return &StoragePortModifyParams{
		HTTPClient: client,
	}
}

/*
StoragePortModifyParams contains all the parameters to send to the API endpoint

	for the storage port modify operation.

	Typically these are written to a http.Request.
*/
type StoragePortModifyParams struct {

	/* Info.

	   The new property values for the port.
	*/
	Info *models.StoragePort

	/* Name.

	   Port name
	*/
	Name string

	/* NodeUUID.

	   Node UUID
	*/
	NodeUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage port modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StoragePortModifyParams) WithDefaults() *StoragePortModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage port modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StoragePortModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage port modify params
func (o *StoragePortModifyParams) WithTimeout(timeout time.Duration) *StoragePortModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage port modify params
func (o *StoragePortModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage port modify params
func (o *StoragePortModifyParams) WithContext(ctx context.Context) *StoragePortModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage port modify params
func (o *StoragePortModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage port modify params
func (o *StoragePortModifyParams) WithHTTPClient(client *http.Client) *StoragePortModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage port modify params
func (o *StoragePortModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the storage port modify params
func (o *StoragePortModifyParams) WithInfo(info *models.StoragePort) *StoragePortModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the storage port modify params
func (o *StoragePortModifyParams) SetInfo(info *models.StoragePort) {
	o.Info = info
}

// WithName adds the name to the storage port modify params
func (o *StoragePortModifyParams) WithName(name string) *StoragePortModifyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the storage port modify params
func (o *StoragePortModifyParams) SetName(name string) {
	o.Name = name
}

// WithNodeUUID adds the nodeUUID to the storage port modify params
func (o *StoragePortModifyParams) WithNodeUUID(nodeUUID string) *StoragePortModifyParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the storage port modify params
func (o *StoragePortModifyParams) SetNodeUUID(nodeUUID string) {
	o.NodeUUID = nodeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *StoragePortModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param node.uuid
	if err := r.SetPathParam("node.uuid", o.NodeUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
