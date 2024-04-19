// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewIscsiCredentialsCollectionGetParams creates a new IscsiCredentialsCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIscsiCredentialsCollectionGetParams() *IscsiCredentialsCollectionGetParams {
	return &IscsiCredentialsCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIscsiCredentialsCollectionGetParamsWithTimeout creates a new IscsiCredentialsCollectionGetParams object
// with the ability to set a timeout on a request.
func NewIscsiCredentialsCollectionGetParamsWithTimeout(timeout time.Duration) *IscsiCredentialsCollectionGetParams {
	return &IscsiCredentialsCollectionGetParams{
		timeout: timeout,
	}
}

// NewIscsiCredentialsCollectionGetParamsWithContext creates a new IscsiCredentialsCollectionGetParams object
// with the ability to set a context for a request.
func NewIscsiCredentialsCollectionGetParamsWithContext(ctx context.Context) *IscsiCredentialsCollectionGetParams {
	return &IscsiCredentialsCollectionGetParams{
		Context: ctx,
	}
}

// NewIscsiCredentialsCollectionGetParamsWithHTTPClient creates a new IscsiCredentialsCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewIscsiCredentialsCollectionGetParamsWithHTTPClient(client *http.Client) *IscsiCredentialsCollectionGetParams {
	return &IscsiCredentialsCollectionGetParams{
		HTTPClient: client,
	}
}

/*
IscsiCredentialsCollectionGetParams contains all the parameters to send to the API endpoint

	for the iscsi credentials collection get operation.

	Typically these are written to a http.Request.
*/
type IscsiCredentialsCollectionGetParams struct {

	/* AuthenticationType.

	   Filter by authentication_type
	*/
	AuthenticationType *string

	/* ChapInboundUser.

	   Filter by chap.inbound.user
	*/
	ChapInboundUser *string

	/* ChapOutboundUser.

	   Filter by chap.outbound.user
	*/
	ChapOutboundUser *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Initiator.

	   Filter by initiator
	*/
	Initiator *string

	/* InitiatorAddressMasksAddress.

	   Filter by initiator_address.masks.address
	*/
	InitiatorAddressMasksAddress *string

	/* InitiatorAddressMasksFamily.

	   Filter by initiator_address.masks.family
	*/
	InitiatorAddressMasksFamily *string

	/* InitiatorAddressMasksNetmask.

	   Filter by initiator_address.masks.netmask
	*/
	InitiatorAddressMasksNetmask *string

	/* InitiatorAddressRangesEnd.

	   Filter by initiator_address.ranges.end
	*/
	InitiatorAddressRangesEnd *string

	/* InitiatorAddressRangesFamily.

	   Filter by initiator_address.ranges.family
	*/
	InitiatorAddressRangesFamily *string

	/* InitiatorAddressRangesStart.

	   Filter by initiator_address.ranges.start
	*/
	InitiatorAddressRangesStart *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

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

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iscsi credentials collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiCredentialsCollectionGetParams) WithDefaults() *IscsiCredentialsCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iscsi credentials collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiCredentialsCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := IscsiCredentialsCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithTimeout(timeout time.Duration) *IscsiCredentialsCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithContext(ctx context.Context) *IscsiCredentialsCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithHTTPClient(client *http.Client) *IscsiCredentialsCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthenticationType adds the authenticationType to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithAuthenticationType(authenticationType *string) *IscsiCredentialsCollectionGetParams {
	o.SetAuthenticationType(authenticationType)
	return o
}

// SetAuthenticationType adds the authenticationType to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetAuthenticationType(authenticationType *string) {
	o.AuthenticationType = authenticationType
}

// WithChapInboundUser adds the chapInboundUser to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithChapInboundUser(chapInboundUser *string) *IscsiCredentialsCollectionGetParams {
	o.SetChapInboundUser(chapInboundUser)
	return o
}

// SetChapInboundUser adds the chapInboundUser to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetChapInboundUser(chapInboundUser *string) {
	o.ChapInboundUser = chapInboundUser
}

// WithChapOutboundUser adds the chapOutboundUser to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithChapOutboundUser(chapOutboundUser *string) *IscsiCredentialsCollectionGetParams {
	o.SetChapOutboundUser(chapOutboundUser)
	return o
}

// SetChapOutboundUser adds the chapOutboundUser to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetChapOutboundUser(chapOutboundUser *string) {
	o.ChapOutboundUser = chapOutboundUser
}

// WithFields adds the fields to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithFields(fields []string) *IscsiCredentialsCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithInitiator adds the initiator to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithInitiator(initiator *string) *IscsiCredentialsCollectionGetParams {
	o.SetInitiator(initiator)
	return o
}

// SetInitiator adds the initiator to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetInitiator(initiator *string) {
	o.Initiator = initiator
}

// WithInitiatorAddressMasksAddress adds the initiatorAddressMasksAddress to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithInitiatorAddressMasksAddress(initiatorAddressMasksAddress *string) *IscsiCredentialsCollectionGetParams {
	o.SetInitiatorAddressMasksAddress(initiatorAddressMasksAddress)
	return o
}

// SetInitiatorAddressMasksAddress adds the initiatorAddressMasksAddress to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetInitiatorAddressMasksAddress(initiatorAddressMasksAddress *string) {
	o.InitiatorAddressMasksAddress = initiatorAddressMasksAddress
}

// WithInitiatorAddressMasksFamily adds the initiatorAddressMasksFamily to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithInitiatorAddressMasksFamily(initiatorAddressMasksFamily *string) *IscsiCredentialsCollectionGetParams {
	o.SetInitiatorAddressMasksFamily(initiatorAddressMasksFamily)
	return o
}

// SetInitiatorAddressMasksFamily adds the initiatorAddressMasksFamily to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetInitiatorAddressMasksFamily(initiatorAddressMasksFamily *string) {
	o.InitiatorAddressMasksFamily = initiatorAddressMasksFamily
}

// WithInitiatorAddressMasksNetmask adds the initiatorAddressMasksNetmask to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithInitiatorAddressMasksNetmask(initiatorAddressMasksNetmask *string) *IscsiCredentialsCollectionGetParams {
	o.SetInitiatorAddressMasksNetmask(initiatorAddressMasksNetmask)
	return o
}

// SetInitiatorAddressMasksNetmask adds the initiatorAddressMasksNetmask to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetInitiatorAddressMasksNetmask(initiatorAddressMasksNetmask *string) {
	o.InitiatorAddressMasksNetmask = initiatorAddressMasksNetmask
}

// WithInitiatorAddressRangesEnd adds the initiatorAddressRangesEnd to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithInitiatorAddressRangesEnd(initiatorAddressRangesEnd *string) *IscsiCredentialsCollectionGetParams {
	o.SetInitiatorAddressRangesEnd(initiatorAddressRangesEnd)
	return o
}

// SetInitiatorAddressRangesEnd adds the initiatorAddressRangesEnd to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetInitiatorAddressRangesEnd(initiatorAddressRangesEnd *string) {
	o.InitiatorAddressRangesEnd = initiatorAddressRangesEnd
}

// WithInitiatorAddressRangesFamily adds the initiatorAddressRangesFamily to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithInitiatorAddressRangesFamily(initiatorAddressRangesFamily *string) *IscsiCredentialsCollectionGetParams {
	o.SetInitiatorAddressRangesFamily(initiatorAddressRangesFamily)
	return o
}

// SetInitiatorAddressRangesFamily adds the initiatorAddressRangesFamily to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetInitiatorAddressRangesFamily(initiatorAddressRangesFamily *string) {
	o.InitiatorAddressRangesFamily = initiatorAddressRangesFamily
}

// WithInitiatorAddressRangesStart adds the initiatorAddressRangesStart to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithInitiatorAddressRangesStart(initiatorAddressRangesStart *string) *IscsiCredentialsCollectionGetParams {
	o.SetInitiatorAddressRangesStart(initiatorAddressRangesStart)
	return o
}

// SetInitiatorAddressRangesStart adds the initiatorAddressRangesStart to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetInitiatorAddressRangesStart(initiatorAddressRangesStart *string) {
	o.InitiatorAddressRangesStart = initiatorAddressRangesStart
}

// WithMaxRecords adds the maxRecords to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithMaxRecords(maxRecords *int64) *IscsiCredentialsCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithOrderBy(orderBy []string) *IscsiCredentialsCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithReturnRecords(returnRecords *bool) *IscsiCredentialsCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *IscsiCredentialsCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmName adds the svmName to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithSvmName(svmName *string) *IscsiCredentialsCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) WithSvmUUID(svmUUID *string) *IscsiCredentialsCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the iscsi credentials collection get params
func (o *IscsiCredentialsCollectionGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *IscsiCredentialsCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AuthenticationType != nil {

		// query param authentication_type
		var qrAuthenticationType string

		if o.AuthenticationType != nil {
			qrAuthenticationType = *o.AuthenticationType
		}
		qAuthenticationType := qrAuthenticationType
		if qAuthenticationType != "" {

			if err := r.SetQueryParam("authentication_type", qAuthenticationType); err != nil {
				return err
			}
		}
	}

	if o.ChapInboundUser != nil {

		// query param chap.inbound.user
		var qrChapInboundUser string

		if o.ChapInboundUser != nil {
			qrChapInboundUser = *o.ChapInboundUser
		}
		qChapInboundUser := qrChapInboundUser
		if qChapInboundUser != "" {

			if err := r.SetQueryParam("chap.inbound.user", qChapInboundUser); err != nil {
				return err
			}
		}
	}

	if o.ChapOutboundUser != nil {

		// query param chap.outbound.user
		var qrChapOutboundUser string

		if o.ChapOutboundUser != nil {
			qrChapOutboundUser = *o.ChapOutboundUser
		}
		qChapOutboundUser := qrChapOutboundUser
		if qChapOutboundUser != "" {

			if err := r.SetQueryParam("chap.outbound.user", qChapOutboundUser); err != nil {
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

	if o.Initiator != nil {

		// query param initiator
		var qrInitiator string

		if o.Initiator != nil {
			qrInitiator = *o.Initiator
		}
		qInitiator := qrInitiator
		if qInitiator != "" {

			if err := r.SetQueryParam("initiator", qInitiator); err != nil {
				return err
			}
		}
	}

	if o.InitiatorAddressMasksAddress != nil {

		// query param initiator_address.masks.address
		var qrInitiatorAddressMasksAddress string

		if o.InitiatorAddressMasksAddress != nil {
			qrInitiatorAddressMasksAddress = *o.InitiatorAddressMasksAddress
		}
		qInitiatorAddressMasksAddress := qrInitiatorAddressMasksAddress
		if qInitiatorAddressMasksAddress != "" {

			if err := r.SetQueryParam("initiator_address.masks.address", qInitiatorAddressMasksAddress); err != nil {
				return err
			}
		}
	}

	if o.InitiatorAddressMasksFamily != nil {

		// query param initiator_address.masks.family
		var qrInitiatorAddressMasksFamily string

		if o.InitiatorAddressMasksFamily != nil {
			qrInitiatorAddressMasksFamily = *o.InitiatorAddressMasksFamily
		}
		qInitiatorAddressMasksFamily := qrInitiatorAddressMasksFamily
		if qInitiatorAddressMasksFamily != "" {

			if err := r.SetQueryParam("initiator_address.masks.family", qInitiatorAddressMasksFamily); err != nil {
				return err
			}
		}
	}

	if o.InitiatorAddressMasksNetmask != nil {

		// query param initiator_address.masks.netmask
		var qrInitiatorAddressMasksNetmask string

		if o.InitiatorAddressMasksNetmask != nil {
			qrInitiatorAddressMasksNetmask = *o.InitiatorAddressMasksNetmask
		}
		qInitiatorAddressMasksNetmask := qrInitiatorAddressMasksNetmask
		if qInitiatorAddressMasksNetmask != "" {

			if err := r.SetQueryParam("initiator_address.masks.netmask", qInitiatorAddressMasksNetmask); err != nil {
				return err
			}
		}
	}

	if o.InitiatorAddressRangesEnd != nil {

		// query param initiator_address.ranges.end
		var qrInitiatorAddressRangesEnd string

		if o.InitiatorAddressRangesEnd != nil {
			qrInitiatorAddressRangesEnd = *o.InitiatorAddressRangesEnd
		}
		qInitiatorAddressRangesEnd := qrInitiatorAddressRangesEnd
		if qInitiatorAddressRangesEnd != "" {

			if err := r.SetQueryParam("initiator_address.ranges.end", qInitiatorAddressRangesEnd); err != nil {
				return err
			}
		}
	}

	if o.InitiatorAddressRangesFamily != nil {

		// query param initiator_address.ranges.family
		var qrInitiatorAddressRangesFamily string

		if o.InitiatorAddressRangesFamily != nil {
			qrInitiatorAddressRangesFamily = *o.InitiatorAddressRangesFamily
		}
		qInitiatorAddressRangesFamily := qrInitiatorAddressRangesFamily
		if qInitiatorAddressRangesFamily != "" {

			if err := r.SetQueryParam("initiator_address.ranges.family", qInitiatorAddressRangesFamily); err != nil {
				return err
			}
		}
	}

	if o.InitiatorAddressRangesStart != nil {

		// query param initiator_address.ranges.start
		var qrInitiatorAddressRangesStart string

		if o.InitiatorAddressRangesStart != nil {
			qrInitiatorAddressRangesStart = *o.InitiatorAddressRangesStart
		}
		qInitiatorAddressRangesStart := qrInitiatorAddressRangesStart
		if qInitiatorAddressRangesStart != "" {

			if err := r.SetQueryParam("initiator_address.ranges.start", qInitiatorAddressRangesStart); err != nil {
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

	if o.SvmUUID != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SvmUUID != nil {
			qrSvmUUID = *o.SvmUUID
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamIscsiCredentialsCollectionGet binds the parameter fields
func (o *IscsiCredentialsCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamIscsiCredentialsCollectionGet binds the parameter order_by
func (o *IscsiCredentialsCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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