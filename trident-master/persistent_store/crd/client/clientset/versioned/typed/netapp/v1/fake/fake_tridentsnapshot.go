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

// FakeTridentSnapshots implements TridentSnapshotInterface
type FakeTridentSnapshots struct {
	Fake *FakeTridentV1
	ns   string
}

var tridentsnapshotsResource = schema.GroupVersionResource{Group: "trident.netapp.io", Version: "v1", Resource: "tridentsnapshots"}

var tridentsnapshotsKind = schema.GroupVersionKind{Group: "trident.netapp.io", Version: "v1", Kind: "TridentSnapshot"}

// Get takes name of the tridentSnapshot, and returns the corresponding tridentSnapshot object, and an error if there is any.
func (c *FakeTridentSnapshots) Get(ctx context.Context, name string, options v1.GetOptions) (result *netappv1.TridentSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tridentsnapshotsResource, c.ns, name), &netappv1.TridentSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentSnapshot), err
}

// List takes label and field selectors, and returns the list of TridentSnapshots that match those selectors.
func (c *FakeTridentSnapshots) List(ctx context.Context, opts v1.ListOptions) (result *netappv1.TridentSnapshotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tridentsnapshotsResource, tridentsnapshotsKind, c.ns, opts), &netappv1.TridentSnapshotList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &netappv1.TridentSnapshotList{ListMeta: obj.(*netappv1.TridentSnapshotList).ListMeta}
	for _, item := range obj.(*netappv1.TridentSnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tridentSnapshots.
func (c *FakeTridentSnapshots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tridentsnapshotsResource, c.ns, opts))

}

// Create takes the representation of a tridentSnapshot and creates it.  Returns the server's representation of the tridentSnapshot, and an error, if there is any.
func (c *FakeTridentSnapshots) Create(ctx context.Context, tridentSnapshot *netappv1.TridentSnapshot, opts v1.CreateOptions) (result *netappv1.TridentSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tridentsnapshotsResource, c.ns, tridentSnapshot), &netappv1.TridentSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentSnapshot), err
}

// Update takes the representation of a tridentSnapshot and updates it. Returns the server's representation of the tridentSnapshot, and an error, if there is any.
func (c *FakeTridentSnapshots) Update(ctx context.Context, tridentSnapshot *netappv1.TridentSnapshot, opts v1.UpdateOptions) (result *netappv1.TridentSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tridentsnapshotsResource, c.ns, tridentSnapshot), &netappv1.TridentSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentSnapshot), err
}

// Delete takes name of the tridentSnapshot and deletes it. Returns an error if one occurs.
func (c *FakeTridentSnapshots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tridentsnapshotsResource, c.ns, name), &netappv1.TridentSnapshot{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTridentSnapshots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tridentsnapshotsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &netappv1.TridentSnapshotList{})
	return err
}

// Patch applies the patch and returns the patched tridentSnapshot.
func (c *FakeTridentSnapshots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *netappv1.TridentSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tridentsnapshotsResource, c.ns, name, pt, data, subresources...), &netappv1.TridentSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentSnapshot), err
}