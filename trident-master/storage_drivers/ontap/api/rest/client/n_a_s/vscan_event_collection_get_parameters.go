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

// NewVscanEventCollectionGetParams creates a new VscanEventCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVscanEventCollectionGetParams() *VscanEventCollectionGetParams {
	return &VscanEventCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVscanEventCollectionGetParamsWithTimeout creates a new VscanEventCollectionGetParams object
// with the ability to set a timeout on a request.
func NewVscanEventCollectionGetParamsWithTimeout(timeout time.Duration) *VscanEventCollectionGetParams {
	return &VscanEventCollectionGetParams{
		timeout: timeout,
	}
}

// NewVscanEventCollectionGetParamsWithContext creates a new VscanEventCollectionGetParams object
// with the ability to set a context for a request.
func NewVscanEventCollectionGetParamsWithContext(ctx context.Context) *VscanEventCollectionGetParams {
	return &VscanEventCollectionGetParams{
		Context: ctx,
	}
}

// NewVscanEventCollectionGetParamsWithHTTPClient creates a new VscanEventCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewVscanEventCollectionGetParamsWithHTTPClient(client *http.Client) *VscanEventCollectionGetParams {
	return &VscanEventCollectionGetParams{
		HTTPClient: client,
	}
}

/*
VscanEventCollectionGetParams contains all the parameters to send to the API endpoint

	for the vscan event collection get operation.

	Typically these are written to a http.Request.
*/
type VscanEventCollectionGetParams struct {

	/* DisconnectReason.

	   Filter by disconnect_reason
	*/
	DisconnectReason *string

	/* EventTime.

	   Filter by event_time
	*/
	EventTime *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* FilePath.

	   Filter by file_path
	*/
	FilePath *string

	/* InterfaceIPAddress.

	   Filter by interface.ip.address
	*/
	InterfaceIPAddress *string

	/* InterfaceName.

	   Filter by interface.name
	*/
	InterfaceName *string

	/* InterfaceUUID.

	   Filter by interface.uuid
	*/
	InterfaceUUID *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* NodeName.

	   Filter by node.name
	*/
	NodeName *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUID *string

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

	/* Server.

	   Filter by server
	*/
	Server *string

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	/* Type.

	   Filter by type
	*/
	Type *string

	/* Vendor.

	   Filter by vendor
	*/
	Vendor *string

	/* Version.

	   Filter by version
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vscan event collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanEventCollectionGetParams) WithDefaults() *VscanEventCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vscan event collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanEventCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := VscanEventCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithTimeout(timeout time.Duration) *VscanEventCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithContext(ctx context.Context) *VscanEventCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithHTTPClient(client *http.Client) *VscanEventCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisconnectReason adds the disconnectReason to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithDisconnectReason(disconnectReason *string) *VscanEventCollectionGetParams {
	o.SetDisconnectReason(disconnectReason)
	return o
}

// SetDisconnectReason adds the disconnectReason to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetDisconnectReason(disconnectReason *string) {
	o.DisconnectReason = disconnectReason
}

// WithEventTime adds the eventTime to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithEventTime(eventTime *string) *VscanEventCollectionGetParams {
	o.SetEventTime(eventTime)
	return o
}

// SetEventTime adds the eventTime to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetEventTime(eventTime *string) {
	o.EventTime = eventTime
}

// WithFields adds the fields to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithFields(fields []string) *VscanEventCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithFilePath adds the filePath to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithFilePath(filePath *string) *VscanEventCollectionGetParams {
	o.SetFilePath(filePath)
	return o
}

// SetFilePath adds the filePath to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetFilePath(filePath *string) {
	o.FilePath = filePath
}

// WithInterfaceIPAddress adds the interfaceIPAddress to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithInterfaceIPAddress(interfaceIPAddress *string) *VscanEventCollectionGetParams {
	o.SetInterfaceIPAddress(interfaceIPAddress)
	return o
}

// SetInterfaceIPAddress adds the interfaceIpAddress to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetInterfaceIPAddress(interfaceIPAddress *string) {
	o.InterfaceIPAddress = interfaceIPAddress
}

// WithInterfaceName adds the interfaceName to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithInterfaceName(interfaceName *string) *VscanEventCollectionGetParams {
	o.SetInterfaceName(interfaceName)
	return o
}

// SetInterfaceName adds the interfaceName to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetInterfaceName(interfaceName *string) {
	o.InterfaceName = interfaceName
}

// WithInterfaceUUID adds the interfaceUUID to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithInterfaceUUID(interfaceUUID *string) *VscanEventCollectionGetParams {
	o.SetInterfaceUUID(interfaceUUID)
	return o
}

// SetInterfaceUUID adds the interfaceUuid to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetInterfaceUUID(interfaceUUID *string) {
	o.InterfaceUUID = interfaceUUID
}

// WithMaxRecords adds the maxRecords to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithMaxRecords(maxRecords *int64) *VscanEventCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNodeName adds the nodeName to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithNodeName(nodeName *string) *VscanEventCollectionGetParams {
	o.SetNodeName(nodeName)
	return o
}

// SetNodeName adds the nodeName to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetNodeName(nodeName *string) {
	o.NodeName = nodeName
}

// WithNodeUUID adds the nodeUUID to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithNodeUUID(nodeUUID *string) *VscanEventCollectionGetParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetNodeUUID(nodeUUID *string) {
	o.NodeUUID = nodeUUID
}

// WithOrderBy adds the orderBy to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithOrderBy(orderBy []string) *VscanEventCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithReturnRecords(returnRecords *bool) *VscanEventCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *VscanEventCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithServer adds the server to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithServer(server *string) *VscanEventCollectionGetParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetServer(server *string) {
	o.Server = server
}

// WithSvmName adds the svmName to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithSvmName(svmName *string) *VscanEventCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithSvmUUID(svmUUID string) *VscanEventCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WithType adds the typeVar to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithType(typeVar *string) *VscanEventCollectionGetParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithVendor adds the vendor to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithVendor(vendor *string) *VscanEventCollectionGetParams {
	o.SetVendor(vendor)
	return o
}

// SetVendor adds the vendor to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetVendor(vendor *string) {
	o.Vendor = vendor
}

// WithVersion adds the version to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithVersion(version *string) *VscanEventCollectionGetParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *VscanEventCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DisconnectReason != nil {

		// query param disconnect_reason
		var qrDisconnectReason string

		if o.DisconnectReason != nil {
			qrDisconnectReason = *o.DisconnectReason
		}
		qDisconnectReason := qrDisconnectReason
		if qDisconnectReason != "" {

			if err := r.SetQueryParam("disconnect_reason", qDisconnectReason); err != nil {
				return err
			}
		}
	}

	if o.EventTime != nil {

		// query param event_time
		var qrEventTime string

		if o.EventTime != nil {
			qrEventTime = *o.EventTime
		}
		qEventTime := qrEventTime
		if qEventTime != "" {

			if err := r.SetQueryParam("event_time", qEventTime); err != nil {
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

	if o.FilePath != nil {

		// query param file_path
		var qrFilePath string

		if o.FilePath != nil {
			qrFilePath = *o.FilePath
		}
		qFilePath := qrFilePath
		if qFilePath != "" {

			if err := r.SetQueryParam("file_path", qFilePath); err != nil {
				return err
			}
		}
	}

	if o.InterfaceIPAddress != nil {

		// query param interface.ip.address
		var qrInterfaceIPAddress string

		if o.InterfaceIPAddress != nil {
			qrInterfaceIPAddress = *o.InterfaceIPAddress
		}
		qInterfaceIPAddress := qrInterfaceIPAddress
		if qInterfaceIPAddress != "" {

			if err := r.SetQueryParam("interface.ip.address", qInterfaceIPAddress); err != nil {
				return err
			}
		}
	}

	if o.InterfaceName != nil {

		// query param interface.name
		var qrInterfaceName string

		if o.InterfaceName != nil {
			qrInterfaceName = *o.InterfaceName
		}
		qInterfaceName := qrInterfaceName
		if qInterfaceName != "" {

			if err := r.SetQueryParam("interface.name", qInterfaceName); err != nil {
				return err
			}
		}
	}

	if o.InterfaceUUID != nil {

		// query param interface.uuid
		var qrInterfaceUUID string

		if o.InterfaceUUID != nil {
			qrInterfaceUUID = *o.InterfaceUUID
		}
		qInterfaceUUID := qrInterfaceUUID
		if qInterfaceUUID != "" {

			if err := r.SetQueryParam("interface.uuid", qInterfaceUUID); err != nil {
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

	if o.NodeName != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeName != nil {
			qrNodeName = *o.NodeName
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
				return err
			}
		}
	}

	if o.NodeUUID != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUID != nil {
			qrNodeUUID = *o.NodeUUID
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
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

	if o.Server != nil {

		// query param server
		var qrServer string

		if o.Server != nil {
			qrServer = *o.Server
		}
		qServer := qrServer
		if qServer != "" {

			if err := r.SetQueryParam("server", qServer); err != nil {
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

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.Vendor != nil {

		// query param vendor
		var qrVendor string

		if o.Vendor != nil {
			qrVendor = *o.Vendor
		}
		qVendor := qrVendor
		if qVendor != "" {

			if err := r.SetQueryParam("vendor", qVendor); err != nil {
				return err
			}
		}
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamVscanEventCollectionGet binds the parameter fields
func (o *VscanEventCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamVscanEventCollectionGet binds the parameter order_by
func (o *VscanEventCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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