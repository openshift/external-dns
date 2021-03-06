// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by kubetype-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha1 "istio.io/api/meta/v1alpha1"
	client "istio.io/api/mixer/v1/config/client"
	v1beta1 "istio.io/api/policy/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//
// +kubetype-gen:lowerCaseScheme
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AttributeManifest describes a set of Attributes produced by some component
// of an Istio deployment.
//
// <!-- crd generation tags
// +cue-gen:attributemanifest:schema:istio.policy.v1beta1.AttributeManifest
// +cue-gen:attributemanifest:groupName:config.istio.io
// +cue-gen:attributemanifest:version:v1alpha2
// +cue-gen:attributemanifest:storageVersion
// +cue-gen:attributemanifest:annotations:helm.sh/resource-policy=keep
// +cue-gen:attributemanifest:labels:app=mixer,chart=istio,heritage=Tiller,istio=core,package=istio.io.mixer,release=istio
// +cue-gen:attributemanifest:subresource:status
// +cue-gen:attributemanifest:scope:Namespaced
// +cue-gen:attributemanifest:resource:categories=istio-io,policy-istio-io
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=config.istio.io/v1alpha2
// +kubetype-gen:kubeType=AttributeManifest
// +kubetype-gen:AttributeManifest:tag=kubetype-gen:lowerCaseScheme
// +genclient
// +k8s:deepcopy-gen=true
// -->
type AttributeManifest struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1beta1.AttributeManifest `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AttributeManifestList is a collection of AttributeManifests.
type AttributeManifestList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []AttributeManifest `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// please upgrade the proto package
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPAPISpec defines the canonical configuration for generating
// API-related attributes from HTTP requests based on the method and
// uri templated path matches. It is sufficient for defining the API
// surface of a service for the purposes of API attribute
// generation. It is not intended to represent auth, quota,
// documentation, or other information commonly found in other API
// specifications, e.g. OpenAPI.
//
// Existing standards that define operations (or methods) in terms of
// HTTP methods and paths can be normalized to this format for use in
// Istio. For example, a simple petstore API described by OpenAPIv2
// [here](https://github.com/googleapis/gnostic/blob/master/examples/v2.0/yaml/petstore-simple.yaml)
// can be represented with the following HTTPAPISpec.
//
// ```yaml
// apiVersion: config.istio.io/v1alpha2
// kind: HTTPAPISpec
// metadata:
//   name: petstore
//   namespace: default
// spec:
//   attributes:
//     attributes:
//       api.service:
//         stringValue: petstore.swagger.io
//       api.version:
//         stringValue: 1.0.0
//   patterns:
//   - attributes:
//       attributes:
//         api.operation:
//           stringValue: findPets
//     httpMethod: GET
//     uriTemplate: /api/pets
//   - attributes:
//       attributes:
//         api.operation:
//           stringValue: addPet
//     httpMethod: POST
//     uriTemplate: /api/pets
//   - attributes:
//       attributes:
//         api.operation:
//           stringValue: findPetById
//     httpMethod: GET
//     uriTemplate: /api/pets/{id}
//   - attributes:
//       attributes:
//         api.operation:
//           stringValue: deletePet
//     httpMethod: DELETE
//     uriTemplate: /api/pets/{id}
//   apiKeys:
//   - query: api-key
// ```
//
// <!-- crd generation tags
// +cue-gen:HTTPAPISpec:schema:istio.mixer.v1.config.client.HTTPAPISpec
// +cue-gen:HTTPAPISpec:groupName:config.istio.io
// +cue-gen:HTTPAPISpec:version:v1alpha2
// +cue-gen:HTTPAPISpec:storageVersion
// +cue-gen:HTTPAPISpec:annotations:helm.sh/resource-policy=keep
// +cue-gen:HTTPAPISpec:labels:app=istio-mixer,chart=istio,heritage=Tiller,release=istio
// +cue-gen:HTTPAPISpec:subresource:status
// +cue-gen:HTTPAPISpec:scope:Namespaced
// +cue-gen:HTTPAPISpec:resource:categories=istio-io,apim-istio-io
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=config.istio.io/v1alpha2
// +genclient
// +k8s:deepcopy-gen=true
// -->
type HTTPAPISpec struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec client.HTTPAPISpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPAPISpecList is a collection of HTTPAPISpecs.
type HTTPAPISpecList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []HTTPAPISpec `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPAPISpecBinding defines the binding between HTTPAPISpecs and one or more
// IstioService. For example, the following establishes a binding
// between the HTTPAPISpec `petstore` and service `foo` in namespace `bar`.
//
// ```yaml
// apiVersion: config.istio.io/v1alpha2
// kind: HTTPAPISpecBinding
// metadata:
//   name: my-binding
//   namespace: default
// spec:
//   services:
//   - name: foo
//     namespace: bar
//   apiSpecs:
//   - name: petstore
//     namespace: default
// ```
//
// <!-- crd generation tags
// +cue-gen:HTTPAPISpecBinding:schema:istio.mixer.v1.config.client.HTTPAPISpecBinding
// +cue-gen:HTTPAPISpecBinding:groupName:config.istio.io
// +cue-gen:HTTPAPISpecBinding:version:v1alpha2
// +cue-gen:HTTPAPISpecBinding:storageVersion
// +cue-gen:HTTPAPISpecBinding:annotations:helm.sh/resource-policy=keep
// +cue-gen:HTTPAPISpecBinding:labels:app=istio-mixer,chart=istio,heritage=Tiller,release=istio
// +cue-gen:HTTPAPISpecBinding:subresource:status
// +cue-gen:HTTPAPISpecBinding:scope:Namespaced
// +cue-gen:HTTPAPISpecBinding:resource:categories=istio-io,apim-istio-io
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=config.istio.io/v1alpha2
// +genclient
// +k8s:deepcopy-gen=true
// -->
type HTTPAPISpecBinding struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec client.HTTPAPISpecBinding `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPAPISpecBindingList is a collection of HTTPAPISpecBindings.
type HTTPAPISpecBindingList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []HTTPAPISpecBinding `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +kubetype-gen:lowerCaseScheme
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Handler allows the operator to configure a specific adapter implementation.
// Each adapter implementation defines its own `params` proto.
//
// In the following example we define a `metrics` handler for the `prometheus` adapter.
// The example is in the form of a Kubernetes resource:
// * The `metadata.name` is the name of the handler
// * The `kind` refers to the adapter name
// * The `spec` block represents adapter-specific configuration as well as the connection information
//
// ```yaml
// # Sample-1: No connection specified (for compiled in adapters)
// # Note: if connection information is not specified, the adapter configuration is directly inside
// # `spec` block. This is going to be DEPRECATED in favor of Sample-2
// apiVersion: "config.istio.io/v1alpha2"
// kind: handler
// metadata:
//   name: requestcount
//   namespace: istio-system
// spec:
//   compiledAdapter: prometheus
//   params:
//     metrics:
//     - name: request_count
//       instance_name: requestcount.metric.istio-system
//       kind: COUNTER
//       label_names:
//       - source_service
//       - source_version
//       - destination_service
//       - destination_version
// ---
// # Sample-2: With connection information (for out-of-process adapters)
// # Note: Unlike sample-1, the adapter configuration is parallel to `connection` and is nested inside `param` block.
// apiVersion: "config.istio.io/v1alpha2"
// kind: handler
// metadata:
//   name: requestcount
//   namespace: istio-system
// spec:
//   compiledAdapter: prometheus
//   params:
//     param:
//       metrics:
//       - name: request_count
//         instance_name: requestcount.metric.istio-system
//         kind: COUNTER
//         label_names:
//         - source_service
//         - source_version
//         - destination_service
//         - destination_version
//     connection:
//       address: localhost:8090
// ---
// ```
//
// <!-- crd generation tags
// +cue-gen:handler:schema:istio.policy.v1beta1.Handler
// +cue-gen:handler:groupName:config.istio.io
// +cue-gen:handler:version:v1alpha2
// +cue-gen:handler:storageVersion
// +cue-gen:handler:annotations:helm.sh/resource-policy=keep
// +cue-gen:handler:labels:app=mixer,chart=istio,heritage=Tiller,istio=mixer-handler,package=handler,release=istio
// +cue-gen:handler:subresource:status
// +cue-gen:handler:scope:Namespaced
// +cue-gen:handler:resource:categories=istio-io,policy-istio-io
// +cue-gen:handler:preserveUnknownFields:params
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=config.istio.io/v1alpha2
// +kubetype-gen:kubeType=Handler
// +kubetype-gen:Handler:tag=kubetype-gen:lowerCaseScheme
// +genclient
// +k8s:deepcopy-gen=true
// -->
type Handler struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1beta1.Handler `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HandlerList is a collection of Handlers.
type HandlerList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Handler `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +kubetype-gen:lowerCaseScheme
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// An Instance tells Mixer how to create instances for particular template.
//
// Instance is defined by the operator. Instance is defined relative to a known
// template. Their purpose is to tell Mixer how to use attributes or literals to produce
// instances of the specified template at runtime.
//
// The following example instructs Mixer to construct an instance associated with template
// 'istio.mixer.adapter.metric.Metric'. It provides a mapping from the template's fields to expressions.
// Instances produced with this instance can be referenced by [Actions][istio.policy.v1beta1.Action] using name
// 'RequestCountByService'
//
// ```yaml
// - name: RequestCountByService
//   template: istio.mixer.adapter.metric.Metric
//   params:
//     value: 1
//     dimensions:
//       source: source.name
//       destination_ip: destination.ip
// ```
//
// <!-- crd generation tags
// +cue-gen:instance:schema:istio.policy.v1beta1.Instance
// +cue-gen:instance:groupName:config.istio.io
// +cue-gen:instance:version:v1alpha2
// +cue-gen:instance:storageVersion
// +cue-gen:instance:annotations:helm.sh/resource-policy=keep
// +cue-gen:instance:labels:app=mixer,chart=istio,heritage=Tiller,istio=mixer-instance,package=instance,release=istio
// +cue-gen:instance:subresource:status
// +cue-gen:instance:scope:Namespaced
// +cue-gen:instance:resource:categories=istio-io,policy-istio-io
// +cue-gen:instance:preserveUnknownFields:params
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=config.istio.io/v1alpha2
// +kubetype-gen:kubeType=Instance
// +kubetype-gen:Instance:tag=kubetype-gen:lowerCaseScheme
// +genclient
// +k8s:deepcopy-gen=true
// -->
type Instance struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1beta1.Instance `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InstanceList is a collection of Instances.
type InstanceList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Instance `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// please upgrade the proto package
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Determines the quotas used for individual requests.
//
// <!-- crd generation tags
// +cue-gen:QuotaSpec:schema:istio.mixer.v1.config.client.QuotaSpec
// +cue-gen:QuotaSpec:groupName:config.istio.io
// +cue-gen:QuotaSpec:version:v1alpha2
// +cue-gen:QuotaSpec:storageVersion
// +cue-gen:QuotaSpec:annotations:helm.sh/resource-policy=keep
// +cue-gen:QuotaSpec:labels:app=istio-mixer,chart=istio,heritage=Tiller,release=istio
// +cue-gen:QuotaSpec:subresource:status
// +cue-gen:QuotaSpec:scope:Namespaced
// +cue-gen:QuotaSpec:resource:categories=istio-io,apim-istio-io
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=config.istio.io/v1alpha2
// +genclient
// +k8s:deepcopy-gen=true
// -->
type QuotaSpec struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec client.QuotaSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// QuotaSpecList is a collection of QuotaSpecs.
type QuotaSpecList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []QuotaSpec `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// QuotaSpecBinding defines the binding between QuotaSpecs and one or more
// IstioService.
//
// <!-- crd generation tags
// +cue-gen:QuotaSpecBinding:schema:istio.mixer.v1.config.client.QuotaSpecBinding
// +cue-gen:QuotaSpecBinding:groupName:config.istio.io
// +cue-gen:QuotaSpecBinding:version:v1alpha2
// +cue-gen:QuotaSpecBinding:storageVersion
// +cue-gen:QuotaSpecBinding:annotations:helm.sh/resource-policy=keep
// +cue-gen:QuotaSpecBinding:labels:app=istio-mixer,chart=istio,heritage=Tiller,release=istio
// +cue-gen:QuotaSpecBinding:subresource:status
// +cue-gen:QuotaSpecBinding:scope:Namespaced
// +cue-gen:QuotaSpecBinding:resource:categories=istio-io,apim-istio-io
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=config.istio.io/v1alpha2
// +genclient
// +k8s:deepcopy-gen=true
// -->
type QuotaSpecBinding struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec client.QuotaSpecBinding `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// QuotaSpecBindingList is a collection of QuotaSpecBindings.
type QuotaSpecBindingList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []QuotaSpecBinding `json:"items" protobuf:"bytes,2,rep,name=items"`
}

//
// +kubetype-gen:lowerCaseScheme
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// A Rule is a selector and a set of intentions to be executed when the
// selector is `true`
//
// The following example instructs Mixer to invoke `prometheus-handler` handler for all services and pass it the
// instance constructed using the 'RequestCountByService' instance.
//
// ```yaml
// - match: match(destination.service.host, "*")
//   actions:
//   - handler: prometheus-handler
//     instances:
//     - RequestCountByService
// ```
//
// <!-- crd generation tags
// +cue-gen:rule:schema:istio.policy.v1beta1.Rule
// +cue-gen:rule:groupName:config.istio.io
// +cue-gen:rule:version:v1alpha2
// +cue-gen:rule:storageVersion
// +cue-gen:rule:annotations:helm.sh/resource-policy=keep
// +cue-gen:rule:labels:app=mixer,chart=istio,heritage=Tiller,istio=core,package=istio.io.mixer,release=istio
// +cue-gen:rule:subresource:status
// +cue-gen:rule:scope:Namespaced
// +cue-gen:rule:resource:categories=istio-io,policy-istio-io
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=config.istio.io/v1alpha2
// +kubetype-gen:kubeType=Rule
// +kubetype-gen:Rule:tag=kubetype-gen:lowerCaseScheme
// +genclient
// +k8s:deepcopy-gen=true
// -->
type Rule struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec v1beta1.Rule `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RuleList is a collection of Rules.
type RuleList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []Rule `json:"items" protobuf:"bytes,2,rep,name=items"`
}
