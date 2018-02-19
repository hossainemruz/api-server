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
// This file was automatically generated by informer-gen

package internalversion

import (
	customrc "api-server/pkg/apis/customrc"
	internalclientset "api-server/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "api-server/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "api-server/pkg/client/listers_generated/customrc/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// CustomReplicationControllerInformer provides access to a shared informer and lister for
// CustomReplicationControllers.
type CustomReplicationControllerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.CustomReplicationControllerLister
}

type customReplicationControllerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCustomReplicationControllerInformer constructs a new informer for CustomReplicationController type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCustomReplicationControllerInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCustomReplicationControllerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCustomReplicationControllerInformer constructs a new informer for CustomReplicationController type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCustomReplicationControllerInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Customrc().CustomReplicationControllers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Customrc().CustomReplicationControllers(namespace).Watch(options)
			},
		},
		&customrc.CustomReplicationController{},
		resyncPeriod,
		indexers,
	)
}

func (f *customReplicationControllerInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCustomReplicationControllerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *customReplicationControllerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&customrc.CustomReplicationController{}, f.defaultInformer)
}

func (f *customReplicationControllerInformer) Lister() internalversion.CustomReplicationControllerLister {
	return internalversion.NewCustomReplicationControllerLister(f.Informer().GetIndexer())
}