// Copyright 2023 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	scheme "github.com/netapp/trident/persistent_store/crd/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TridentVersionsGetter has a method to return a TridentVersionInterface.
// A group's client should implement this interface.
type TridentVersionsGetter interface {
	TridentVersions(namespace string) TridentVersionInterface
}

// TridentVersionInterface has methods to work with TridentVersion resources.
type TridentVersionInterface interface {
	Create(ctx context.Context, tridentVersion *v1.TridentVersion, opts metav1.CreateOptions) (*v1.TridentVersion, error)
	Update(ctx context.Context, tridentVersion *v1.TridentVersion, opts metav1.UpdateOptions) (*v1.TridentVersion, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TridentVersion, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TridentVersionList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentVersion, err error)
	TridentVersionExpansion
}

// tridentVersions implements TridentVersionInterface
type tridentVersions struct {
	client rest.Interface
	ns     string
}

// newTridentVersions returns a TridentVersions
func newTridentVersions(c *TridentV1Client, namespace string) *tridentVersions {
	return &tridentVersions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tridentVersion, and returns the corresponding tridentVersion object, and an error if there is any.
func (c *tridentVersions) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.TridentVersion, err error) {
	result = &v1.TridentVersion{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TridentVersions that match those selectors.
func (c *tridentVersions) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TridentVersionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TridentVersionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tridentversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tridentVersions.
func (c *tridentVersions) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tridentversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tridentVersion and creates it.  Returns the server's representation of the tridentVersion, and an error, if there is any.
func (c *tridentVersions) Create(ctx context.Context, tridentVersion *v1.TridentVersion, opts metav1.CreateOptions) (result *v1.TridentVersion, err error) {
	result = &v1.TridentVersion{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tridentversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentVersion).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tridentVersion and updates it. Returns the server's representation of the tridentVersion, and an error, if there is any.
func (c *tridentVersions) Update(ctx context.Context, tridentVersion *v1.TridentVersion, opts metav1.UpdateOptions) (result *v1.TridentVersion, err error) {
	result = &v1.TridentVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tridentversions").
		Name(tridentVersion.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentVersion).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tridentVersion and deletes it. Returns an error if one occurs.
func (c *tridentVersions) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentversions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tridentVersions) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tridentversions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tridentVersion.
func (c *tridentVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentVersion, err error) {
	result = &v1.TridentVersion{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tridentversions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}