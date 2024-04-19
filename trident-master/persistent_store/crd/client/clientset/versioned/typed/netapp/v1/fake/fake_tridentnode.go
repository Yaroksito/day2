// Copyright 2023 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	netappv1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTridentNodes implements TridentNodeInterface
type FakeTridentNodes struct {
	Fake *FakeTridentV1
	ns   string
}

var tridentnodesResource = schema.GroupVersionResource{Group: "trident.netapp.io", Version: "v1", Resource: "tridentnodes"}

var tridentnodesKind = schema.GroupVersionKind{Group: "trident.netapp.io", Version: "v1", Kind: "TridentNode"}

// Get takes name of the tridentNode, and returns the corresponding tridentNode object, and an error if there is any.
func (c *FakeTridentNodes) Get(ctx context.Context, name string, options v1.GetOptions) (result *netappv1.TridentNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tridentnodesResource, c.ns, name), &netappv1.TridentNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentNode), err
}

// List takes label and field selectors, and returns the list of TridentNodes that match those selectors.
func (c *FakeTridentNodes) List(ctx context.Context, opts v1.ListOptions) (result *netappv1.TridentNodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tridentnodesResource, tridentnodesKind, c.ns, opts), &netappv1.TridentNodeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &netappv1.TridentNodeList{ListMeta: obj.(*netappv1.TridentNodeList).ListMeta}
	for _, item := range obj.(*netappv1.TridentNodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tridentNodes.
func (c *FakeTridentNodes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tridentnodesResource, c.ns, opts))

}

// Create takes the representation of a tridentNode and creates it.  Returns the server's representation of the tridentNode, and an error, if there is any.
func (c *FakeTridentNodes) Create(ctx context.Context, tridentNode *netappv1.TridentNode, opts v1.CreateOptions) (result *netappv1.TridentNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tridentnodesResource, c.ns, tridentNode), &netappv1.TridentNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentNode), err
}

// Update takes the representation of a tridentNode and updates it. Returns the server's representation of the tridentNode, and an error, if there is any.
func (c *FakeTridentNodes) Update(ctx context.Context, tridentNode *netappv1.TridentNode, opts v1.UpdateOptions) (result *netappv1.TridentNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tridentnodesResource, c.ns, tridentNode), &netappv1.TridentNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentNode), err
}

// Delete takes name of the tridentNode and deletes it. Returns an error if one occurs.
func (c *FakeTridentNodes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tridentnodesResource, c.ns, name), &netappv1.TridentNode{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTridentNodes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tridentnodesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &netappv1.TridentNodeList{})
	return err
}

// Patch applies the patch and returns the patched tridentNode.
func (c *FakeTridentNodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *netappv1.TridentNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tridentnodesResource, c.ns, name, pt, data, subresources...), &netappv1.TridentNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentNode), err
}