/*
  Copyright 2017 The Kubernetes Authors.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/package fake

import (
	v1alpha1 "api-server/pkg/apis/car/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCarControllers implements CarControllerInterface
type FakeCarControllers struct {
	Fake *FakeCarV1alpha1
	ns   string
}

var carcontrollersResource = schema.GroupVersionResource{Group: "car.emruz.example.com", Version: "v1alpha1", Resource: "carcontrollers"}

var carcontrollersKind = schema.GroupVersionKind{Group: "car.emruz.example.com", Version: "v1alpha1", Kind: "CarController"}

// Get takes name of the carController, and returns the corresponding carController object, and an error if there is any.
func (c *FakeCarControllers) Get(name string, options v1.GetOptions) (result *v1alpha1.CarController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(carcontrollersResource, c.ns, name), &v1alpha1.CarController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CarController), err
}

// List takes label and field selectors, and returns the list of CarControllers that match those selectors.
func (c *FakeCarControllers) List(opts v1.ListOptions) (result *v1alpha1.CarControllerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(carcontrollersResource, carcontrollersKind, c.ns, opts), &v1alpha1.CarControllerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CarControllerList{}
	for _, item := range obj.(*v1alpha1.CarControllerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested carControllers.
func (c *FakeCarControllers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(carcontrollersResource, c.ns, opts))

}

// Create takes the representation of a carController and creates it.  Returns the server's representation of the carController, and an error, if there is any.
func (c *FakeCarControllers) Create(carController *v1alpha1.CarController) (result *v1alpha1.CarController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(carcontrollersResource, c.ns, carController), &v1alpha1.CarController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CarController), err
}

// Update takes the representation of a carController and updates it. Returns the server's representation of the carController, and an error, if there is any.
func (c *FakeCarControllers) Update(carController *v1alpha1.CarController) (result *v1alpha1.CarController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(carcontrollersResource, c.ns, carController), &v1alpha1.CarController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CarController), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCarControllers) UpdateStatus(carController *v1alpha1.CarController) (*v1alpha1.CarController, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(carcontrollersResource, "status", c.ns, carController), &v1alpha1.CarController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CarController), err
}

// Delete takes name of the carController and deletes it. Returns an error if one occurs.
func (c *FakeCarControllers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(carcontrollersResource, c.ns, name), &v1alpha1.CarController{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCarControllers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(carcontrollersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CarControllerList{})
	return err
}

// Patch applies the patch and returns the patched carController.
func (c *FakeCarControllers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CarController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(carcontrollersResource, c.ns, name, data, subresources...), &v1alpha1.CarController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CarController), err
}
