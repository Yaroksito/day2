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

// NewEmsFilterCollectionGetParams creates a new EmsFilterCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmsFilterCollectionGetParams() *EmsFilterCollectionGetParams {
	return &EmsFilterCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmsFilterCollectionGetParamsWithTimeout creates a new EmsFilterCollectionGetParams object
// with the ability to set a timeout on a request.
func NewEmsFilterCollectionGetParamsWithTimeout(timeout time.Duration) *EmsFilterCollectionGetParams {
	return &EmsFilterCollectionGetParams{
		timeout: timeout,
	}
}

// NewEmsFilterCollectionGetParamsWithContext creates a new EmsFilterCollectionGetParams object
// with the ability to set a context for a request.
func NewEmsFilterCollectionGetParamsWithContext(ctx context.Context) *EmsFilterCollectionGetParams {
	return &EmsFilterCollectionGetParams{
		Context: ctx,
	}
}

// NewEmsFilterCollectionGetParamsWithHTTPClient creates a new EmsFilterCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmsFilterCollectionGetParamsWithHTTPClient(client *http.Client) *EmsFilterCollectionGetParams {
	return &EmsFilterCollectionGetParams{
		HTTPClient: client,
	}
}

/*
EmsFilterCollectionGetParams contains all the parameters to send to the API endpoint

	for the ems filter collection get operation.

	Typically these are written to a http.Request.
*/
type EmsFilterCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	Name *string

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

	/* RulesIndex.

	   Filter by rules.index
	*/
	RulesIndex *int64

	/* RulesMessageCriteriaNamePattern.

	   Filter by rules.message_criteria.name_pattern
	*/
	RulesMessageCriteriaNamePattern *string

	/* RulesMessageCriteriaSeverities.

	   Filter by rules.message_criteria.severities
	*/
	RulesMessageCriteriaSeverities *string

	/* RulesMessageCriteriaSnmpTrapTypes.

	   Filter by rules.message_criteria.snmp_trap_types
	*/
	RulesMessageCriteriaSnmpTrapTypes *string

	/* RulesType.

	   Filter by rules.type
	*/
	RulesType *string

	/* SystemDefined.

	   Filter by system_defined
	*/
	SystemDefined *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ems filter collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsFilterCollectionGetParams) WithDefaults() *EmsFilterCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ems filter collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsFilterCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := EmsFilterCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithTimeout(timeout time.Duration) *EmsFilterCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithContext(ctx context.Context) *EmsFilterCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithHTTPClient(client *http.Client) *EmsFilterCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithFields(fields []string) *EmsFilterCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithMaxRecords(maxRecords *int64) *EmsFilterCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithName adds the name to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithName(name *string) *EmsFilterCollectionGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetName(name *string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithOrderBy(orderBy []string) *EmsFilterCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithReturnRecords(returnRecords *bool) *EmsFilterCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *EmsFilterCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithRulesIndex adds the rulesIndex to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithRulesIndex(rulesIndex *int64) *EmsFilterCollectionGetParams {
	o.SetRulesIndex(rulesIndex)
	return o
}

// SetRulesIndex adds the rulesIndex to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetRulesIndex(rulesIndex *int64) {
	o.RulesIndex = rulesIndex
}

// WithRulesMessageCriteriaNamePattern adds the rulesMessageCriteriaNamePattern to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithRulesMessageCriteriaNamePattern(rulesMessageCriteriaNamePattern *string) *EmsFilterCollectionGetParams {
	o.SetRulesMessageCriteriaNamePattern(rulesMessageCriteriaNamePattern)
	return o
}

// SetRulesMessageCriteriaNamePattern adds the rulesMessageCriteriaNamePattern to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetRulesMessageCriteriaNamePattern(rulesMessageCriteriaNamePattern *string) {
	o.RulesMessageCriteriaNamePattern = rulesMessageCriteriaNamePattern
}

// WithRulesMessageCriteriaSeverities adds the rulesMessageCriteriaSeverities to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithRulesMessageCriteriaSeverities(rulesMessageCriteriaSeverities *string) *EmsFilterCollectionGetParams {
	o.SetRulesMessageCriteriaSeverities(rulesMessageCriteriaSeverities)
	return o
}

// SetRulesMessageCriteriaSeverities adds the rulesMessageCriteriaSeverities to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetRulesMessageCriteriaSeverities(rulesMessageCriteriaSeverities *string) {
	o.RulesMessageCriteriaSeverities = rulesMessageCriteriaSeverities
}

// WithRulesMessageCriteriaSnmpTrapTypes adds the rulesMessageCriteriaSnmpTrapTypes to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithRulesMessageCriteriaSnmpTrapTypes(rulesMessageCriteriaSnmpTrapTypes *string) *EmsFilterCollectionGetParams {
	o.SetRulesMessageCriteriaSnmpTrapTypes(rulesMessageCriteriaSnmpTrapTypes)
	return o
}

// SetRulesMessageCriteriaSnmpTrapTypes adds the rulesMessageCriteriaSnmpTrapTypes to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetRulesMessageCriteriaSnmpTrapTypes(rulesMessageCriteriaSnmpTrapTypes *string) {
	o.RulesMessageCriteriaSnmpTrapTypes = rulesMessageCriteriaSnmpTrapTypes
}

// WithRulesType adds the rulesType to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithRulesType(rulesType *string) *EmsFilterCollectionGetParams {
	o.SetRulesType(rulesType)
	return o
}

// SetRulesType adds the rulesType to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetRulesType(rulesType *string) {
	o.RulesType = rulesType
}

// WithSystemDefined adds the systemDefined to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) WithSystemDefined(systemDefined *bool) *EmsFilterCollectionGetParams {
	o.SetSystemDefined(systemDefined)
	return o
}

// SetSystemDefined adds the systemDefined to the ems filter collection get params
func (o *EmsFilterCollectionGetParams) SetSystemDefined(systemDefined *bool) {
	o.SystemDefined = systemDefined
}

// WriteToRequest writes these params to a swagger request
func (o *EmsFilterCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
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

	if o.RulesIndex != nil {

		// query param rules.index
		var qrRulesIndex int64

		if o.RulesIndex != nil {
			qrRulesIndex = *o.RulesIndex
		}
		qRulesIndex := swag.FormatInt64(qrRulesIndex)
		if qRulesIndex != "" {

			if err := r.SetQueryParam("rules.index", qRulesIndex); err != nil {
				return err
			}
		}
	}

	if o.RulesMessageCriteriaNamePattern != nil {

		// query param rules.message_criteria.name_pattern
		var qrRulesMessageCriteriaNamePattern string

		if o.RulesMessageCriteriaNamePattern != nil {
			qrRulesMessageCriteriaNamePattern = *o.RulesMessageCriteriaNamePattern
		}
		qRulesMessageCriteriaNamePattern := qrRulesMessageCriteriaNamePattern
		if qRulesMessageCriteriaNamePattern != "" {

			if err := r.SetQueryParam("rules.message_criteria.name_pattern", qRulesMessageCriteriaNamePattern); err != nil {
				return err
			}
		}
	}

	if o.RulesMessageCriteriaSeverities != nil {

		// query param rules.message_criteria.severities
		var qrRulesMessageCriteriaSeverities string

		if o.RulesMessageCriteriaSeverities != nil {
			qrRulesMessageCriteriaSeverities = *o.RulesMessageCriteriaSeverities
		}
		qRulesMessageCriteriaSeverities := qrRulesMessageCriteriaSeverities
		if qRulesMessageCriteriaSeverities != "" {

			if err := r.SetQueryParam("rules.message_criteria.severities", qRulesMessageCriteriaSeverities); err != nil {
				return err
			}
		}
	}

	if o.RulesMessageCriteriaSnmpTrapTypes != nil {

		// query param rules.message_criteria.snmp_trap_types
		var qrRulesMessageCriteriaSnmpTrapTypes string

		if o.RulesMessageCriteriaSnmpTrapTypes != nil {
			qrRulesMessageCriteriaSnmpTrapTypes = *o.RulesMessageCriteriaSnmpTrapTypes
		}
		qRulesMessageCriteriaSnmpTrapTypes := qrRulesMessageCriteriaSnmpTrapTypes
		if qRulesMessageCriteriaSnmpTrapTypes != "" {

			if err := r.SetQueryParam("rules.message_criteria.snmp_trap_types", qRulesMessageCriteriaSnmpTrapTypes); err != nil {
				return err
			}
		}
	}

	if o.RulesType != nil {

		// query param rules.type
		var qrRulesType string

		if o.RulesType != nil {
			qrRulesType = *o.RulesType
		}
		qRulesType := qrRulesType
		if qRulesType != "" {

			if err := r.SetQueryParam("rules.type", qRulesType); err != nil {
				return err
			}
		}
	}

	if o.SystemDefined != nil {

		// query param system_defined
		var qrSystemDefined bool

		if o.SystemDefined != nil {
			qrSystemDefined = *o.SystemDefined
		}
		qSystemDefined := swag.FormatBool(qrSystemDefined)
		if qSystemDefined != "" {

			if err := r.SetQueryParam("system_defined", qSystemDefined); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamEmsFilterCollectionGet binds the parameter fields
func (o *EmsFilterCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamEmsFilterCollectionGet binds the parameter order_by
func (o *EmsFilterCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
