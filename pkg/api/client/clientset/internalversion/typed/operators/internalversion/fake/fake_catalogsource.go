/*
Copyright Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	operators "github.com/operator-framework/api/pkg/operators"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCatalogSources implements CatalogSourceInterface
type FakeCatalogSources struct {
	Fake *FakeOperators
	ns   string
}

var catalogsourcesResource = schema.GroupVersionResource{Group: "operators.coreos.com", Version: "", Resource: "catalogsources"}

var catalogsourcesKind = schema.GroupVersionKind{Group: "operators.coreos.com", Version: "", Kind: "CatalogSource"}

// Get takes name of the catalogSource, and returns the corresponding catalogSource object, and an error if there is any.
func (c *FakeCatalogSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *operators.CatalogSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(catalogsourcesResource, c.ns, name), &operators.CatalogSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operators.CatalogSource), err
}

// List takes label and field selectors, and returns the list of CatalogSources that match those selectors.
func (c *FakeCatalogSources) List(ctx context.Context, opts v1.ListOptions) (result *operators.CatalogSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(catalogsourcesResource, catalogsourcesKind, c.ns, opts), &operators.CatalogSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operators.CatalogSourceList{ListMeta: obj.(*operators.CatalogSourceList).ListMeta}
	for _, item := range obj.(*operators.CatalogSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested catalogSources.
func (c *FakeCatalogSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(catalogsourcesResource, c.ns, opts))

}

// Create takes the representation of a catalogSource and creates it.  Returns the server's representation of the catalogSource, and an error, if there is any.
func (c *FakeCatalogSources) Create(ctx context.Context, catalogSource *operators.CatalogSource, opts v1.CreateOptions) (result *operators.CatalogSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(catalogsourcesResource, c.ns, catalogSource), &operators.CatalogSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operators.CatalogSource), err
}

// Update takes the representation of a catalogSource and updates it. Returns the server's representation of the catalogSource, and an error, if there is any.
func (c *FakeCatalogSources) Update(ctx context.Context, catalogSource *operators.CatalogSource, opts v1.UpdateOptions) (result *operators.CatalogSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(catalogsourcesResource, c.ns, catalogSource), &operators.CatalogSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operators.CatalogSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCatalogSources) UpdateStatus(ctx context.Context, catalogSource *operators.CatalogSource, opts v1.UpdateOptions) (*operators.CatalogSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(catalogsourcesResource, "status", c.ns, catalogSource), &operators.CatalogSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operators.CatalogSource), err
}

// Delete takes name of the catalogSource and deletes it. Returns an error if one occurs.
func (c *FakeCatalogSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(catalogsourcesResource, c.ns, name), &operators.CatalogSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCatalogSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(catalogsourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &operators.CatalogSourceList{})
	return err
}

// Patch applies the patch and returns the patched catalogSource.
func (c *FakeCatalogSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operators.CatalogSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(catalogsourcesResource, c.ns, name, pt, data, subresources...), &operators.CatalogSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operators.CatalogSource), err
}
