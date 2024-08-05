/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha3

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha3 "sigs.k8s.io/gateway-api/apis/v1alpha3"
)

// BackendTLSPolicyLister helps list BackendTLSPolicies.
// All objects returned here must be treated as read-only.
type BackendTLSPolicyLister interface {
	// List lists all BackendTLSPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha3.BackendTLSPolicy, err error)
	// BackendTLSPolicies returns an object that can list and get BackendTLSPolicies.
	BackendTLSPolicies(namespace string) BackendTLSPolicyNamespaceLister
	BackendTLSPolicyListerExpansion
}

// backendTLSPolicyLister implements the BackendTLSPolicyLister interface.
type backendTLSPolicyLister struct {
	indexer cache.Indexer
}

// NewBackendTLSPolicyLister returns a new BackendTLSPolicyLister.
func NewBackendTLSPolicyLister(indexer cache.Indexer) BackendTLSPolicyLister {
	return &backendTLSPolicyLister{indexer: indexer}
}

// List lists all BackendTLSPolicies in the indexer.
func (s *backendTLSPolicyLister) List(selector labels.Selector) (ret []*v1alpha3.BackendTLSPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.BackendTLSPolicy))
	})
	return ret, err
}

// BackendTLSPolicies returns an object that can list and get BackendTLSPolicies.
func (s *backendTLSPolicyLister) BackendTLSPolicies(namespace string) BackendTLSPolicyNamespaceLister {
	return backendTLSPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackendTLSPolicyNamespaceLister helps list and get BackendTLSPolicies.
// All objects returned here must be treated as read-only.
type BackendTLSPolicyNamespaceLister interface {
	// List lists all BackendTLSPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha3.BackendTLSPolicy, err error)
	// Get retrieves the BackendTLSPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha3.BackendTLSPolicy, error)
	BackendTLSPolicyNamespaceListerExpansion
}

// backendTLSPolicyNamespaceLister implements the BackendTLSPolicyNamespaceLister
// interface.
type backendTLSPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BackendTLSPolicies in the indexer for a given namespace.
func (s backendTLSPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha3.BackendTLSPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.BackendTLSPolicy))
	})
	return ret, err
}

// Get retrieves the BackendTLSPolicy from the indexer for a given namespace and name.
func (s backendTLSPolicyNamespaceLister) Get(name string) (*v1alpha3.BackendTLSPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha3.Resource("backendtlspolicy"), name)
	}
	return obj.(*v1alpha3.BackendTLSPolicy), nil
}
