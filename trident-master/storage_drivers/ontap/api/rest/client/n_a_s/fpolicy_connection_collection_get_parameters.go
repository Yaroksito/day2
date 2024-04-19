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

// NewFpolicyConnectionCollectionGetParams creates a new FpolicyConnectionCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFpolicyConnectionCollectionGetParams() *FpolicyConnectionCollectionGetParams {
	return &FpolicyConnectionCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFpolicyConnectionCollectionGetParamsWithTimeout creates a new FpolicyConnectionCollectionGetParams object
// with the ability to set a timeout on a request.
func NewFpolicyConnectionCollectionGetParamsWithTimeout(timeout time.Duration) *FpolicyConnectionCollectionGetParams {
	return &FpolicyConnectionCollectionGetParams{
		timeout: timeout,
	}
}

// NewFpolicyConnectionCollectionGetParamsWithContext creates a new FpolicyConnectionCollectionGetParams object
// with the ability to set a context for a request.
func NewFpolicyConnectionCollectionGetParamsWithContext(ctx context.Context) *FpolicyConnectionCollectionGetParams {
	return &FpolicyConnectionCollectionGetParams{
		Context: ctx,
	}
}

// NewFpolicyConnectionCollectionGetParamsWithHTTPClient creates a new FpolicyConnectionCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFpolicyConnectionCollectionGetParamsWithHTTPClient(client *http.Client) *FpolicyConnectionCollectionGetParams {
	return &FpolicyConnectionCollectionGetParams{
		HTTPClient: client,
	}
}

/*
FpolicyConnectionCollectionGetParams contains all the parameters to send to the API endpoint

	for the fpolicy connection collection get operation.

	Typically these are written to a http.Request.
*/
type FpolicyConnectionCollectionGetParams struct {

	/* DisconnectedReasonCode.

	   Filter by disconnected_reason.code
	*/
	DisconnectedReasonCode *int64

	/* DisconnectedReasonMessage.

	   Filter by disconnected_reason.message
	*/
	DisconnectedReasonMessage *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

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

	/* PassthroughRead.

	   Whether to view only passthrough-read connections
	*/
	PassthroughRead *bool

	/* PolicyName.

	   Filter by policy.name
	*/
	PolicyName *string

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

	/* SessionUUID.

	   Filter by session_uuid
	*/
	SessionUUID *string

	/* State.

	   Filter by state
	*/
	State *string

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

	/* UpdateTime.

	   Filter by update_time
	*/
	UpdateTime *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fpolicy connection collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyConnectionCollectionGetParams) WithDefaults() *FpolicyConnectionCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fpolicy connection collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyConnectionCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := FpolicyConnectionCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithTimeout(timeout time.Duration) *FpolicyConnectionCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithContext(ctx context.Context) *FpolicyConnectionCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithHTTPClient(client *http.Client) *FpolicyConnectionCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisconnectedReasonCode adds the disconnectedReasonCode to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithDisconnectedReasonCode(disconnectedReasonCode *int64) *FpolicyConnectionCollectionGetParams {
	o.SetDisconnectedReasonCode(disconnectedReasonCode)
	return o
}

// SetDisconnectedReasonCode adds the disconnectedReasonCode to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetDisconnectedReasonCode(disconnectedReasonCode *int64) {
	o.DisconnectedReasonCode = disconnectedReasonCode
}

// WithDisconnectedReasonMessage adds the disconnectedReasonMessage to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithDisconnectedReasonMessage(disconnectedReasonMessage *string) *FpolicyConnectionCollectionGetParams {
	o.SetDisconnectedReasonMessage(disconnectedReasonMessage)
	return o
}

// SetDisconnectedReasonMessage adds the disconnectedReasonMessage to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetDisconnectedReasonMessage(disconnectedReasonMessage *string) {
	o.DisconnectedReasonMessage = disconnectedReasonMessage
}

// WithFields adds the fields to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithFields(fields []string) *FpolicyConnectionCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithMaxRecords(maxRecords *int64) *FpolicyConnectionCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNodeName adds the nodeName to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithNodeName(nodeName *string) *FpolicyConnectionCollectionGetParams {
	o.SetNodeName(nodeName)
	return o
}

// SetNodeName adds the nodeName to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetNodeName(nodeName *string) {
	o.NodeName = nodeName
}

// WithNodeUUID adds the nodeUUID to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithNodeUUID(nodeUUID *string) *FpolicyConnectionCollectionGetParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetNodeUUID(nodeUUID *string) {
	o.NodeUUID = nodeUUID
}

// WithOrderBy adds the orderBy to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithOrderBy(orderBy []string) *FpolicyConnectionCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPassthroughRead adds the passthroughRead to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithPassthroughRead(passthroughRead *bool) *FpolicyConnectionCollectionGetParams {
	o.SetPassthroughRead(passthroughRead)
	return o
}

// SetPassthroughRead adds the passthroughRead to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetPassthroughRead(passthroughRead *bool) {
	o.PassthroughRead = passthroughRead
}

// WithPolicyName adds the policyName to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithPolicyName(policyName *string) *FpolicyConnectionCollectionGetParams {
	o.SetPolicyName(policyName)
	return o
}

// SetPolicyName adds the policyName to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetPolicyName(policyName *string) {
	o.PolicyName = policyName
}

// WithReturnRecords adds the returnRecords to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithReturnRecords(returnRecords *bool) *FpolicyConnectionCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *FpolicyConnectionCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithServer adds the server to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithServer(server *string) *FpolicyConnectionCollectionGetParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetServer(server *string) {
	o.Server = server
}

// WithSessionUUID adds the sessionUUID to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithSessionUUID(sessionUUID *string) *FpolicyConnectionCollectionGetParams {
	o.SetSessionUUID(sessionUUID)
	return o
}

// SetSessionUUID adds the sessionUuid to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetSessionUUID(sessionUUID *string) {
	o.SessionUUID = sessionUUID
}

// WithState adds the state to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithState(state *string) *FpolicyConnectionCollectionGetParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetState(state *string) {
	o.State = state
}

// WithSvmName adds the svmName to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithSvmName(svmName *string) *FpolicyConnectionCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithSvmUUID(svmUUID string) *FpolicyConnectionCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WithType adds the typeVar to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithType(typeVar *string) *FpolicyConnectionCollectionGetParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithUpdateTime adds the updateTime to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithUpdateTime(updateTime *string) *FpolicyConnectionCollectionGetParams {
	o.SetUpdateTime(updateTime)
	return o
}

// SetUpdateTime adds the updateTime to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetUpdateTime(updateTime *string) {
	o.UpdateTime = updateTime
}

// WriteToRequest writes these params to a swagger request
func (o *FpolicyConnectionCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DisconnectedReasonCode != nil {

		// query param disconnected_reason.code
		var qrDisconnectedReasonCode int64

		if o.DisconnectedReasonCode != nil {
			qrDisconnectedReasonCode = *o.DisconnectedReasonCode
		}
		qDisconnectedReasonCode := swag.FormatInt64(qrDisconnectedReasonCode)
		if qDisconnectedReasonCode != "" {

			if err := r.SetQueryParam("disconnected_reason.code", qDisconnectedReasonCode); err != nil {
				return err
			}
		}
	}

	if o.DisconnectedReasonMessage != nil {

		// query param disconnected_reason.message
		var qrDisconnectedReasonMessage string

		if o.DisconnectedReasonMessage != nil {
			qrDisconnectedReasonMessage = *o.DisconnectedReasonMessage
		}
		qDisconnectedReasonMessage := qrDisconnectedReasonMessage
		if qDisconnectedReasonMessage != "" {

			if err := r.SetQueryParam("disconnected_reason.message", qDisconnectedReasonMessage); err != nil {
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

	if o.PassthroughRead != nil {

		// query param passthrough_read
		var qrPassthroughRead bool

		if o.PassthroughRead != nil {
			qrPassthroughRead = *o.PassthroughRead
		}
		qPassthroughRead := swag.FormatBool(qrPassthroughRead)
		if qPassthroughRead != "" {

			if err := r.SetQueryParam("passthrough_read", qPassthroughRead); err != nil {
				return err
			}
		}
	}

	if o.PolicyName != nil {

		// query param policy.name
		var qrPolicyName string

		if o.PolicyName != nil {
			qrPolicyName = *o.PolicyName
		}
		qPolicyName := qrPolicyName
		if qPolicyName != "" {

			if err := r.SetQueryParam("policy.name", qPolicyName); err != nil {
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

	if o.SessionUUID != nil {

		// query param session_uuid
		var qrSessionUUID string

		if o.SessionUUID != nil {
			qrSessionUUID = *o.SessionUUID
		}
		qSessionUUID := qrSessionUUID
		if qSessionUUID != "" {

			if err := r.SetQueryParam("session_uuid", qSessionUUID); err != nil {
				return err
			}
		}
	}

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
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

	if o.UpdateTime != nil {

		// query param update_time
		var qrUpdateTime string

		if o.UpdateTime != nil {
			qrUpdateTime = *o.UpdateTime
		}
		qUpdateTime := qrUpdateTime
		if qUpdateTime != "" {

			if err := r.SetQueryParam("update_time", qUpdateTime); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamFpolicyConnectionCollectionGet binds the parameter fields
func (o *FpolicyConnectionCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamFpolicyConnectionCollectionGet binds the parameter order_by
func (o *FpolicyConnectionCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
