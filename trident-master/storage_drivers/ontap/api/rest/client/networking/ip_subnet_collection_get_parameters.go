// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewIPSubnetCollectionGetParams creates a new IPSubnetCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIPSubnetCollectionGetParams() *IPSubnetCollectionGetParams {
	return &IPSubnetCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIPSubnetCollectionGetParamsWithTimeout creates a new IPSubnetCollectionGetParams object
// with the ability to set a timeout on a request.
func NewIPSubnetCollectionGetParamsWithTimeout(timeout time.Duration) *IPSubnetCollectionGetParams {
	return &IPSubnetCollectionGetParams{
		timeout: timeout,
	}
}

// NewIPSubnetCollectionGetParamsWithContext creates a new IPSubnetCollectionGetParams object
// with the ability to set a context for a request.
func NewIPSubnetCollectionGetParamsWithContext(ctx context.Context) *IPSubnetCollectionGetParams {
	return &IPSubnetCollectionGetParams{
		Context: ctx,
	}
}

// NewIPSubnetCollectionGetParamsWithHTTPClient creates a new IPSubnetCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewIPSubnetCollectionGetParamsWithHTTPClient(client *http.Client) *IPSubnetCollectionGetParams {
	return &IPSubnetCollectionGetParams{
		HTTPClient: client,
	}
}

/*
IPSubnetCollectionGetParams contains all the parameters to send to the API endpoint

	for the ip subnet collection get operation.

	Typically these are written to a http.Request.
*/
type IPSubnetCollectionGetParams struct {

	/* AvailableCount.

	   Filter by available_count
	*/
	AvailableCount *int64

	/* AvailableIPRangesEnd.

	   Filter by available_ip_ranges.end
	*/
	AvailableIPRangesEnd *string

	/* AvailableIPRangesFamily.

	   Filter by available_ip_ranges.family
	*/
	AvailableIPRangesFamily *string

	/* AvailableIPRangesStart.

	   Filter by available_ip_ranges.start
	*/
	AvailableIPRangesStart *string

	/* BroadcastDomainName.

	   Filter by broadcast_domain.name
	*/
	BroadcastDomainName *string

	/* BroadcastDomainUUID.

	   Filter by broadcast_domain.uuid
	*/
	BroadcastDomainUUID *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Gateway.

	   Filter by gateway
	*/
	Gateway *string

	/* IPRangesEnd.

	   Filter by ip_ranges.end
	*/
	IPRangesEnd *string

	/* IPRangesFamily.

	   Filter by ip_ranges.family
	*/
	IPRangesFamily *string

	/* IPRangesStart.

	   Filter by ip_ranges.start
	*/
	IPRangesStart *string

	/* IpspaceName.

	   Filter by ipspace.name
	*/
	IpspaceName *string

	/* IpspaceUUID.

	   Filter by ipspace.uuid
	*/
	IpspaceUUID *string

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

	/* SubnetAddress.

	   Filter by subnet.address
	*/
	SubnetAddress *string

	/* SubnetFamily.

	   Filter by subnet.family
	*/
	SubnetFamily *string

	/* SubnetNetmask.

	   Filter by subnet.netmask
	*/
	SubnetNetmask *string

	/* TotalCount.

	   Filter by total_count
	*/
	TotalCount *int64

	/* UsedCount.

	   Filter by used_count
	*/
	UsedCount *int64

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ip subnet collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IPSubnetCollectionGetParams) WithDefaults() *IPSubnetCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ip subnet collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IPSubnetCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := IPSubnetCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithTimeout(timeout time.Duration) *IPSubnetCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithContext(ctx context.Context) *IPSubnetCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithHTTPClient(client *http.Client) *IPSubnetCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailableCount adds the availableCount to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithAvailableCount(availableCount *int64) *IPSubnetCollectionGetParams {
	o.SetAvailableCount(availableCount)
	return o
}

// SetAvailableCount adds the availableCount to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetAvailableCount(availableCount *int64) {
	o.AvailableCount = availableCount
}

// WithAvailableIPRangesEnd adds the availableIPRangesEnd to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithAvailableIPRangesEnd(availableIPRangesEnd *string) *IPSubnetCollectionGetParams {
	o.SetAvailableIPRangesEnd(availableIPRangesEnd)
	return o
}

// SetAvailableIPRangesEnd adds the availableIpRangesEnd to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetAvailableIPRangesEnd(availableIPRangesEnd *string) {
	o.AvailableIPRangesEnd = availableIPRangesEnd
}

// WithAvailableIPRangesFamily adds the availableIPRangesFamily to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithAvailableIPRangesFamily(availableIPRangesFamily *string) *IPSubnetCollectionGetParams {
	o.SetAvailableIPRangesFamily(availableIPRangesFamily)
	return o
}

// SetAvailableIPRangesFamily adds the availableIpRangesFamily to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetAvailableIPRangesFamily(availableIPRangesFamily *string) {
	o.AvailableIPRangesFamily = availableIPRangesFamily
}

// WithAvailableIPRangesStart adds the availableIPRangesStart to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithAvailableIPRangesStart(availableIPRangesStart *string) *IPSubnetCollectionGetParams {
	o.SetAvailableIPRangesStart(availableIPRangesStart)
	return o
}

// SetAvailableIPRangesStart adds the availableIpRangesStart to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetAvailableIPRangesStart(availableIPRangesStart *string) {
	o.AvailableIPRangesStart = availableIPRangesStart
}

// WithBroadcastDomainName adds the broadcastDomainName to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithBroadcastDomainName(broadcastDomainName *string) *IPSubnetCollectionGetParams {
	o.SetBroadcastDomainName(broadcastDomainName)
	return o
}

// SetBroadcastDomainName adds the broadcastDomainName to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetBroadcastDomainName(broadcastDomainName *string) {
	o.BroadcastDomainName = broadcastDomainName
}

// WithBroadcastDomainUUID adds the broadcastDomainUUID to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithBroadcastDomainUUID(broadcastDomainUUID *string) *IPSubnetCollectionGetParams {
	o.SetBroadcastDomainUUID(broadcastDomainUUID)
	return o
}

// SetBroadcastDomainUUID adds the broadcastDomainUuid to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetBroadcastDomainUUID(broadcastDomainUUID *string) {
	o.BroadcastDomainUUID = broadcastDomainUUID
}

// WithFields adds the fields to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithFields(fields []string) *IPSubnetCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithGateway adds the gateway to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithGateway(gateway *string) *IPSubnetCollectionGetParams {
	o.SetGateway(gateway)
	return o
}

// SetGateway adds the gateway to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetGateway(gateway *string) {
	o.Gateway = gateway
}

// WithIPRangesEnd adds the iPRangesEnd to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithIPRangesEnd(iPRangesEnd *string) *IPSubnetCollectionGetParams {
	o.SetIPRangesEnd(iPRangesEnd)
	return o
}

// SetIPRangesEnd adds the ipRangesEnd to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetIPRangesEnd(iPRangesEnd *string) {
	o.IPRangesEnd = iPRangesEnd
}

// WithIPRangesFamily adds the iPRangesFamily to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithIPRangesFamily(iPRangesFamily *string) *IPSubnetCollectionGetParams {
	o.SetIPRangesFamily(iPRangesFamily)
	return o
}

// SetIPRangesFamily adds the ipRangesFamily to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetIPRangesFamily(iPRangesFamily *string) {
	o.IPRangesFamily = iPRangesFamily
}

// WithIPRangesStart adds the iPRangesStart to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithIPRangesStart(iPRangesStart *string) *IPSubnetCollectionGetParams {
	o.SetIPRangesStart(iPRangesStart)
	return o
}

// SetIPRangesStart adds the ipRangesStart to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetIPRangesStart(iPRangesStart *string) {
	o.IPRangesStart = iPRangesStart
}

// WithIpspaceName adds the ipspaceName to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithIpspaceName(ipspaceName *string) *IPSubnetCollectionGetParams {
	o.SetIpspaceName(ipspaceName)
	return o
}

// SetIpspaceName adds the ipspaceName to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetIpspaceName(ipspaceName *string) {
	o.IpspaceName = ipspaceName
}

// WithIpspaceUUID adds the ipspaceUUID to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithIpspaceUUID(ipspaceUUID *string) *IPSubnetCollectionGetParams {
	o.SetIpspaceUUID(ipspaceUUID)
	return o
}

// SetIpspaceUUID adds the ipspaceUuid to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetIpspaceUUID(ipspaceUUID *string) {
	o.IpspaceUUID = ipspaceUUID
}

// WithMaxRecords adds the maxRecords to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithMaxRecords(maxRecords *int64) *IPSubnetCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithName adds the name to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithName(name *string) *IPSubnetCollectionGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetName(name *string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithOrderBy(orderBy []string) *IPSubnetCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithReturnRecords(returnRecords *bool) *IPSubnetCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *IPSubnetCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSubnetAddress adds the subnetAddress to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithSubnetAddress(subnetAddress *string) *IPSubnetCollectionGetParams {
	o.SetSubnetAddress(subnetAddress)
	return o
}

// SetSubnetAddress adds the subnetAddress to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetSubnetAddress(subnetAddress *string) {
	o.SubnetAddress = subnetAddress
}

// WithSubnetFamily adds the subnetFamily to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithSubnetFamily(subnetFamily *string) *IPSubnetCollectionGetParams {
	o.SetSubnetFamily(subnetFamily)
	return o
}

// SetSubnetFamily adds the subnetFamily to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetSubnetFamily(subnetFamily *string) {
	o.SubnetFamily = subnetFamily
}

// WithSubnetNetmask adds the subnetNetmask to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithSubnetNetmask(subnetNetmask *string) *IPSubnetCollectionGetParams {
	o.SetSubnetNetmask(subnetNetmask)
	return o
}

// SetSubnetNetmask adds the subnetNetmask to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetSubnetNetmask(subnetNetmask *string) {
	o.SubnetNetmask = subnetNetmask
}

// WithTotalCount adds the totalCount to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithTotalCount(totalCount *int64) *IPSubnetCollectionGetParams {
	o.SetTotalCount(totalCount)
	return o
}

// SetTotalCount adds the totalCount to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetTotalCount(totalCount *int64) {
	o.TotalCount = totalCount
}

// WithUsedCount adds the usedCount to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithUsedCount(usedCount *int64) *IPSubnetCollectionGetParams {
	o.SetUsedCount(usedCount)
	return o
}

// SetUsedCount adds the usedCount to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetUsedCount(usedCount *int64) {
	o.UsedCount = usedCount
}

// WithUUID adds the uuid to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) WithUUID(uuid *string) *IPSubnetCollectionGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the ip subnet collection get params
func (o *IPSubnetCollectionGetParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *IPSubnetCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AvailableCount != nil {

		// query param available_count
		var qrAvailableCount int64

		if o.AvailableCount != nil {
			qrAvailableCount = *o.AvailableCount
		}
		qAvailableCount := swag.FormatInt64(qrAvailableCount)
		if qAvailableCount != "" {

			if err := r.SetQueryParam("available_count", qAvailableCount); err != nil {
				return err
			}
		}
	}

	if o.AvailableIPRangesEnd != nil {

		// query param available_ip_ranges.end
		var qrAvailableIPRangesEnd string

		if o.AvailableIPRangesEnd != nil {
			qrAvailableIPRangesEnd = *o.AvailableIPRangesEnd
		}
		qAvailableIPRangesEnd := qrAvailableIPRangesEnd
		if qAvailableIPRangesEnd != "" {

			if err := r.SetQueryParam("available_ip_ranges.end", qAvailableIPRangesEnd); err != nil {
				return err
			}
		}
	}

	if o.AvailableIPRangesFamily != nil {

		// query param available_ip_ranges.family
		var qrAvailableIPRangesFamily string

		if o.AvailableIPRangesFamily != nil {
			qrAvailableIPRangesFamily = *o.AvailableIPRangesFamily
		}
		qAvailableIPRangesFamily := qrAvailableIPRangesFamily
		if qAvailableIPRangesFamily != "" {

			if err := r.SetQueryParam("available_ip_ranges.family", qAvailableIPRangesFamily); err != nil {
				return err
			}
		}
	}

	if o.AvailableIPRangesStart != nil {

		// query param available_ip_ranges.start
		var qrAvailableIPRangesStart string

		if o.AvailableIPRangesStart != nil {
			qrAvailableIPRangesStart = *o.AvailableIPRangesStart
		}
		qAvailableIPRangesStart := qrAvailableIPRangesStart
		if qAvailableIPRangesStart != "" {

			if err := r.SetQueryParam("available_ip_ranges.start", qAvailableIPRangesStart); err != nil {
				return err
			}
		}
	}

	if o.BroadcastDomainName != nil {

		// query param broadcast_domain.name
		var qrBroadcastDomainName string

		if o.BroadcastDomainName != nil {
			qrBroadcastDomainName = *o.BroadcastDomainName
		}
		qBroadcastDomainName := qrBroadcastDomainName
		if qBroadcastDomainName != "" {

			if err := r.SetQueryParam("broadcast_domain.name", qBroadcastDomainName); err != nil {
				return err
			}
		}
	}

	if o.BroadcastDomainUUID != nil {

		// query param broadcast_domain.uuid
		var qrBroadcastDomainUUID string

		if o.BroadcastDomainUUID != nil {
			qrBroadcastDomainUUID = *o.BroadcastDomainUUID
		}
		qBroadcastDomainUUID := qrBroadcastDomainUUID
		if qBroadcastDomainUUID != "" {

			if err := r.SetQueryParam("broadcast_domain.uuid", qBroadcastDomainUUID); err != nil {
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

	if o.Gateway != nil {

		// query param gateway
		var qrGateway string

		if o.Gateway != nil {
			qrGateway = *o.Gateway
		}
		qGateway := qrGateway
		if qGateway != "" {

			if err := r.SetQueryParam("gateway", qGateway); err != nil {
				return err
			}
		}
	}

	if o.IPRangesEnd != nil {

		// query param ip_ranges.end
		var qrIPRangesEnd string

		if o.IPRangesEnd != nil {
			qrIPRangesEnd = *o.IPRangesEnd
		}
		qIPRangesEnd := qrIPRangesEnd
		if qIPRangesEnd != "" {

			if err := r.SetQueryParam("ip_ranges.end", qIPRangesEnd); err != nil {
				return err
			}
		}
	}

	if o.IPRangesFamily != nil {

		// query param ip_ranges.family
		var qrIPRangesFamily string

		if o.IPRangesFamily != nil {
			qrIPRangesFamily = *o.IPRangesFamily
		}
		qIPRangesFamily := qrIPRangesFamily
		if qIPRangesFamily != "" {

			if err := r.SetQueryParam("ip_ranges.family", qIPRangesFamily); err != nil {
				return err
			}
		}
	}

	if o.IPRangesStart != nil {

		// query param ip_ranges.start
		var qrIPRangesStart string

		if o.IPRangesStart != nil {
			qrIPRangesStart = *o.IPRangesStart
		}
		qIPRangesStart := qrIPRangesStart
		if qIPRangesStart != "" {

			if err := r.SetQueryParam("ip_ranges.start", qIPRangesStart); err != nil {
				return err
			}
		}
	}

	if o.IpspaceName != nil {

		// query param ipspace.name
		var qrIpspaceName string

		if o.IpspaceName != nil {
			qrIpspaceName = *o.IpspaceName
		}
		qIpspaceName := qrIpspaceName
		if qIpspaceName != "" {

			if err := r.SetQueryParam("ipspace.name", qIpspaceName); err != nil {
				return err
			}
		}
	}

	if o.IpspaceUUID != nil {

		// query param ipspace.uuid
		var qrIpspaceUUID string

		if o.IpspaceUUID != nil {
			qrIpspaceUUID = *o.IpspaceUUID
		}
		qIpspaceUUID := qrIpspaceUUID
		if qIpspaceUUID != "" {

			if err := r.SetQueryParam("ipspace.uuid", qIpspaceUUID); err != nil {
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

	if o.SubnetAddress != nil {

		// query param subnet.address
		var qrSubnetAddress string

		if o.SubnetAddress != nil {
			qrSubnetAddress = *o.SubnetAddress
		}
		qSubnetAddress := qrSubnetAddress
		if qSubnetAddress != "" {

			if err := r.SetQueryParam("subnet.address", qSubnetAddress); err != nil {
				return err
			}
		}
	}

	if o.SubnetFamily != nil {

		// query param subnet.family
		var qrSubnetFamily string

		if o.SubnetFamily != nil {
			qrSubnetFamily = *o.SubnetFamily
		}
		qSubnetFamily := qrSubnetFamily
		if qSubnetFamily != "" {

			if err := r.SetQueryParam("subnet.family", qSubnetFamily); err != nil {
				return err
			}
		}
	}

	if o.SubnetNetmask != nil {

		// query param subnet.netmask
		var qrSubnetNetmask string

		if o.SubnetNetmask != nil {
			qrSubnetNetmask = *o.SubnetNetmask
		}
		qSubnetNetmask := qrSubnetNetmask
		if qSubnetNetmask != "" {

			if err := r.SetQueryParam("subnet.netmask", qSubnetNetmask); err != nil {
				return err
			}
		}
	}

	if o.TotalCount != nil {

		// query param total_count
		var qrTotalCount int64

		if o.TotalCount != nil {
			qrTotalCount = *o.TotalCount
		}
		qTotalCount := swag.FormatInt64(qrTotalCount)
		if qTotalCount != "" {

			if err := r.SetQueryParam("total_count", qTotalCount); err != nil {
				return err
			}
		}
	}

	if o.UsedCount != nil {

		// query param used_count
		var qrUsedCount int64

		if o.UsedCount != nil {
			qrUsedCount = *o.UsedCount
		}
		qUsedCount := swag.FormatInt64(qrUsedCount)
		if qUsedCount != "" {

			if err := r.SetQueryParam("used_count", qUsedCount); err != nil {
				return err
			}
		}
	}

	if o.UUID != nil {

		// query param uuid
		var qrUUID string

		if o.UUID != nil {
			qrUUID = *o.UUID
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamIPSubnetCollectionGet binds the parameter fields
func (o *IPSubnetCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamIPSubnetCollectionGet binds the parameter order_by
func (o *IPSubnetCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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