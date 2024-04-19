// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/application"
	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/cloud"
	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/cluster"
	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/n_a_s"
	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/n_v_me"
	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/name_services"
	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/ndmp"
	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/networking"
	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/object_store"
	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/s_a_n"
	securityops "github.com/netapp/trident/storage_drivers/ontap/api/rest/client/security"
	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/snaplock"
	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/snapmirror"
	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/storage"
	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/support"
	"github.com/netapp/trident/storage_drivers/ontap/api/rest/client/svm"
)

// Default o n t a p r e s t API online reference HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new o n t a p r e s t API online reference HTTP client.
func NewHTTPClient(formats strfmt.Registry) *ONTAPRESTAPIOnlineReference {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new o n t a p r e s t API online reference HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *ONTAPRESTAPIOnlineReference {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)

	// BEGIN TRIDENT
	const HALPlusJSONMime = "application/hal+json"
	transport.Consumers[HALPlusJSONMime] = runtime.JSONConsumer()
	transport.Producers[HALPlusJSONMime] = runtime.JSONProducer()
	// END TRIDENT

	return New(transport, formats)
}

// New creates a new o n t a p r e s t API online reference client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *ONTAPRESTAPIOnlineReference {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(ONTAPRESTAPIOnlineReference)
	cli.Transport = transport
	cli.Application = application.New(transport, formats)
	cli.Cloud = cloud.New(transport, formats)
	cli.Cluster = cluster.New(transport, formats)
	cli.Nas = n_a_s.New(transport, formats)
	cli.NvMe = n_v_me.New(transport, formats)
	cli.NameServices = name_services.New(transport, formats)
	cli.Ndmp = ndmp.New(transport, formats)
	cli.Networking = networking.New(transport, formats)
	cli.ObjectStore = object_store.New(transport, formats)
	cli.San = s_a_n.New(transport, formats)
	cli.Security = securityops.New(transport, formats)
	cli.Snaplock = snaplock.New(transport, formats)
	cli.Snapmirror = snapmirror.New(transport, formats)
	cli.Storage = storage.New(transport, formats)
	cli.Support = support.New(transport, formats)
	cli.Svm = svm.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// ONTAPRESTAPIOnlineReference is a client for o n t a p r e s t API online reference
type ONTAPRESTAPIOnlineReference struct {
	Application application.ClientService

	Cloud cloud.ClientService

	Cluster cluster.ClientService

	Nas n_a_s.ClientService

	NvMe n_v_me.ClientService

	NameServices name_services.ClientService

	Ndmp ndmp.ClientService

	Networking networking.ClientService

	ObjectStore object_store.ClientService

	San s_a_n.ClientService

	Security securityops.ClientService

	Snaplock snaplock.ClientService

	Snapmirror snapmirror.ClientService

	Storage storage.ClientService

	Support support.ClientService

	Svm svm.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *ONTAPRESTAPIOnlineReference) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Application.SetTransport(transport)
	c.Cloud.SetTransport(transport)
	c.Cluster.SetTransport(transport)
	c.Nas.SetTransport(transport)
	c.NvMe.SetTransport(transport)
	c.NameServices.SetTransport(transport)
	c.Ndmp.SetTransport(transport)
	c.Networking.SetTransport(transport)
	c.ObjectStore.SetTransport(transport)
	c.San.SetTransport(transport)
	c.Security.SetTransport(transport)
	c.Snaplock.SetTransport(transport)
	c.Snapmirror.SetTransport(transport)
	c.Storage.SetTransport(transport)
	c.Support.SetTransport(transport)
	c.Svm.SetTransport(transport)
}
