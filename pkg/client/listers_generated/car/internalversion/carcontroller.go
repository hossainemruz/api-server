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
*/
// This file was automatically generated by lister-gen

package internalversion

import (
	car "api-server/pkg/apis/car"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CarControllerLister helps list CarControllers.
type CarControllerLister interface {
	// List lists all CarControllers in the indexer.
	List(selector labels.Selector) (ret []*car.CarController, err error)
	// CarControllers returns an object that can list and get CarControllers.
	CarControllers(namespace string) CarControllerNamespaceLister
	CarControllerListerExpansion
}

// carControllerLister implements the CarControllerLister interface.
type carControllerLister struct {
	indexer cache.Indexer
}

// NewCarControllerLister returns a new CarControllerLister.
func NewCarControllerLister(indexer cache.Indexer) CarControllerLister {
	return &carControllerLister{indexer: indexer}
}

// List lists all CarControllers in the indexer.
func (s *carControllerLister) List(selector labels.Selector) (ret []*car.CarController, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*car.CarController))
	})
	return ret, err
}

// CarControllers returns an object that can list and get CarControllers.
func (s *carControllerLister) CarControllers(namespace string) CarControllerNamespaceLister {
	return carControllerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CarControllerNamespaceLister helps list and get CarControllers.
type CarControllerNamespaceLister interface {
	// List lists all CarControllers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*car.CarController, err error)
	// Get retrieves the CarController from the indexer for a given namespace and name.
	Get(name string) (*car.CarController, error)
	CarControllerNamespaceListerExpansion
}

// carControllerNamespaceLister implements the CarControllerNamespaceLister
// interface.
type carControllerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CarControllers in the indexer for a given namespace.
func (s carControllerNamespaceLister) List(selector labels.Selector) (ret []*car.CarController, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*car.CarController))
	})
	return ret, err
}

// Get retrieves the CarController from the indexer for a given namespace and name.
func (s carControllerNamespaceLister) Get(name string) (*car.CarController, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(car.Resource("carcontroller"), name)
	}
	return obj.(*car.CarController), nil
}
