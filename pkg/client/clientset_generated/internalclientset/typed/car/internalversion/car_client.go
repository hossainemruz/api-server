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
*/package internalversion

import (
	"api-server/pkg/client/clientset_generated/internalclientset/scheme"
	rest "k8s.io/client-go/rest"
)

type CarInterface interface {
	RESTClient() rest.Interface
	CarControllersGetter
}

// CarClient is used to interact with features provided by the car.emruz.example.com group.
type CarClient struct {
	restClient rest.Interface
}

func (c *CarClient) CarControllers(namespace string) CarControllerInterface {
	return newCarControllers(c, namespace)
}

// NewForConfig creates a new CarClient for the given config.
func NewForConfig(c *rest.Config) (*CarClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CarClient{client}, nil
}

// NewForConfigOrDie creates a new CarClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CarClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CarClient for the given RESTClient.
func New(c rest.Interface) *CarClient {
	return &CarClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	g, err := scheme.Registry.Group("car.emruz.example.com")
	if err != nil {
		return err
	}

	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	if config.GroupVersion == nil || config.GroupVersion.Group != g.GroupVersion.Group {
		gv := g.GroupVersion
		config.GroupVersion = &gv
	}
	config.NegotiatedSerializer = scheme.Codecs

	if config.QPS == 0 {
		config.QPS = 5
	}
	if config.Burst == 0 {
		config.Burst = 10
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CarClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
