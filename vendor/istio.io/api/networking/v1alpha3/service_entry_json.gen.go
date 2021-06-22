// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/service_entry.proto

// `ServiceEntry` enables adding additional entries into Istio's
// internal service registry, so that auto-discovered services in the
// mesh can access/route to these manually specified services. A
// service entry describes the properties of a service (DNS name,
// VIPs, ports, protocols, endpoints). These services could be
// external to the mesh (e.g., web APIs) or mesh-internal services
// that are not part of the platform's service registry (e.g., a set
// of VMs talking to services in Kubernetes). In addition, the
// endpoints of a service entry can also be dynamically selected by
// using the `workloadSelector` field. These endpoints can be VM
// workloads declared using the `WorkloadEntry` object or Kubernetes
// pods. The ability to select both pods and VMs under a single
// service allows for migration of services from VMs to Kubernetes
// without having to change the existing DNS names associated with the
// services.
//
// The following example declares a few external APIs accessed by internal
// applications over HTTPS. The sidecar inspects the SNI value in the
// ClientHello message to route to the appropriate external service.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-https
// spec:
//   hosts:
//   - api.dropboxapi.com
//   - www.googleapis.com
//   - api.facebook.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 443
//     name: https
//     protocol: TLS
//   resolution: DNS
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-https
// spec:
//   hosts:
//   - api.dropboxapi.com
//   - www.googleapis.com
//   - api.facebook.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 443
//     name: https
//     protocol: TLS
//   resolution: DNS
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The following configuration adds a set of MongoDB instances running on
// unmanaged VMs to Istio's registry, so that these services can be treated
// as any other service in the mesh. The associated DestinationRule is used
// to initiate mTLS connections to the database instances.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-mongocluster
// spec:
//   hosts:
//   - mymongodb.somedomain # not used
//   addresses:
//   - 192.192.192.192/24 # VIPs
//   ports:
//   - number: 27018
//     name: mongodb
//     protocol: MONGO
//   location: MESH_INTERNAL
//   resolution: STATIC
//   endpoints:
//   - address: 2.2.2.2
//   - address: 3.3.3.3
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-mongocluster
// spec:
//   hosts:
//   - mymongodb.somedomain # not used
//   addresses:
//   - 192.192.192.192/24 # VIPs
//   ports:
//   - number: 27018
//     name: mongodb
//     protocol: MONGO
//   location: MESH_INTERNAL
//   resolution: STATIC
//   endpoints:
//   - address: 2.2.2.2
//   - address: 3.3.3.3
// ```
// {{</tab>}}
// {{</tabset>}}
//
// and the associated DestinationRule
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: mtls-mongocluster
// spec:
//   host: mymongodb.somedomain
//   trafficPolicy:
//     tls:
//       mode: MUTUAL
//       clientCertificate: /etc/certs/myclientcert.pem
//       privateKey: /etc/certs/client_private_key.pem
//       caCertificates: /etc/certs/rootcacerts.pem
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: DestinationRule
// metadata:
//   name: mtls-mongocluster
// spec:
//   host: mymongodb.somedomain
//   trafficPolicy:
//     tls:
//       mode: MUTUAL
//       clientCertificate: /etc/certs/myclientcert.pem
//       privateKey: /etc/certs/client_private_key.pem
//       caCertificates: /etc/certs/rootcacerts.pem
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The following example uses a combination of service entry and TLS
// routing in a virtual service to steer traffic based on the SNI value to
// an internal egress firewall.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-redirect
// spec:
//   hosts:
//   - wikipedia.org
//   - "*.wikipedia.org"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 443
//     name: https
//     protocol: TLS
//   resolution: NONE
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-redirect
// spec:
//   hosts:
//   - wikipedia.org
//   - "*.wikipedia.org"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 443
//     name: https
//     protocol: TLS
//   resolution: NONE
// ```
// {{</tab>}}
// {{</tabset>}}
//
// And the associated VirtualService to route based on the SNI value.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: tls-routing
// spec:
//   hosts:
//   - wikipedia.org
//   - "*.wikipedia.org"
//   tls:
//   - match:
//     - sniHosts:
//       - wikipedia.org
//       - "*.wikipedia.org"
//     route:
//     - destination:
//         host: internal-egress-firewall.ns1.svc.cluster.local
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: VirtualService
// metadata:
//   name: tls-routing
// spec:
//   hosts:
//   - wikipedia.org
//   - "*.wikipedia.org"
//   tls:
//   - match:
//     - sniHosts:
//       - wikipedia.org
//       - "*.wikipedia.org"
//     route:
//     - destination:
//         host: internal-egress-firewall.ns1.svc.cluster.local
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The virtual service with TLS match serves to override the default SNI
// match. In the absence of a virtual service, traffic will be forwarded to
// the wikipedia domains.
//
// The following example demonstrates the use of a dedicated egress gateway
// through which all external service traffic is forwarded.
// The 'exportTo' field allows for control over the visibility of a service
// declaration to other namespaces in the mesh. By default, a service is exported
// to all namespaces. The following example restricts the visibility to the
// current namespace, represented by ".", so that it cannot be used by other
// namespaces.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-httpbin
//   namespace : egress
// spec:
//   hosts:
//   - httpbin.com
//   exportTo:
//   - "."
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-httpbin
//   namespace : egress
// spec:
//   hosts:
//   - httpbin.com
//   exportTo:
//   - "."
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
// ```
// {{</tab>}}
// {{</tabset>}}
//
// Define a gateway to handle all egress traffic.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: Gateway
// metadata:
//  name: istio-egressgateway
//  namespace: istio-system
// spec:
//  selector:
//    istio: egressgateway
//  servers:
//  - port:
//      number: 80
//      name: http
//      protocol: HTTP
//    hosts:
//    - "*"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: Gateway
// metadata:
//  name: istio-egressgateway
//  namespace: istio-system
// spec:
//  selector:
//    istio: egressgateway
//  servers:
//  - port:
//      number: 80
//      name: http
//      protocol: HTTP
//    hosts:
//    - "*"
// ```
// {{</tab>}}
// {{</tabset>}}
//
// And the associated `VirtualService` to route from the sidecar to the
// gateway service (`istio-egressgateway.istio-system.svc.cluster.local`), as
// well as route from the gateway to the external service. Note that the
// virtual service is exported to all namespaces enabling them to route traffic
// through the gateway to the external service. Forcing traffic to go through
// a managed middle proxy like this is a common practice.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: gateway-routing
//   namespace: egress
// spec:
//   hosts:
//   - httpbin.com
//   exportTo:
//   - "*"
//   gateways:
//   - mesh
//   - istio-egressgateway
//   http:
//   - match:
//     - port: 80
//       gateways:
//       - mesh
//     route:
//     - destination:
//         host: istio-egressgateway.istio-system.svc.cluster.local
//   - match:
//     - port: 80
//       gateways:
//       - istio-egressgateway
//     route:
//     - destination:
//         host: httpbin.com
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: VirtualService
// metadata:
//   name: gateway-routing
//   namespace: egress
// spec:
//   hosts:
//   - httpbin.com
//   exportTo:
//   - "*"
//   gateways:
//   - mesh
//   - istio-egressgateway
//   http:
//   - match:
//     - port: 80
//       gateways:
//       - mesh
//     route:
//     - destination:
//         host: istio-egressgateway.istio-system.svc.cluster.local
//   - match:
//     - port: 80
//       gateways:
//       - istio-egressgateway
//     route:
//     - destination:
//         host: httpbin.com
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The following example demonstrates the use of wildcards in the hosts for
// external services. If the connection has to be routed to the IP address
// requested by the application (i.e. application resolves DNS and attempts
// to connect to a specific IP), the discovery mode must be set to `NONE`.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-wildcard-example
// spec:
//   hosts:
//   - "*.bar.com"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: NONE
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-wildcard-example
// spec:
//   hosts:
//   - "*.bar.com"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: NONE
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The following example demonstrates a service that is available via a
// Unix Domain Socket on the host of the client. The resolution must be
// set to STATIC to use Unix address endpoints.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: unix-domain-socket-example
// spec:
//   hosts:
//   - "example.unix.local"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   endpoints:
//   - address: unix:///var/run/example/socket
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: unix-domain-socket-example
// spec:
//   hosts:
//   - "example.unix.local"
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   endpoints:
//   - address: unix:///var/run/example/socket
// ```
// {{</tab>}}
// {{</tabset>}}
//
// For HTTP-based services, it is possible to create a `VirtualService`
// backed by multiple DNS addressable endpoints. In such a scenario, the
// application can use the `HTTP_PROXY` environment variable to transparently
// reroute API calls for the `VirtualService` to a chosen backend. For
// example, the following configuration creates a non-existent external
// service called foo.bar.com backed by three domains: us.foo.bar.com:8080,
// uk.foo.bar.com:9080, and in.foo.bar.com:7080
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: external-svc-dns
// spec:
//   hosts:
//   - foo.bar.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
//   endpoints:
//   - address: us.foo.bar.com
//     ports:
//       http: 8080
//   - address: uk.foo.bar.com
//     ports:
//       http: 9080
//   - address: in.foo.bar.com
//     ports:
//       http: 7080
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: external-svc-dns
// spec:
//   hosts:
//   - foo.bar.com
//   location: MESH_EXTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
//   endpoints:
//   - address: us.foo.bar.com
//     ports:
//       https: 8080
//   - address: uk.foo.bar.com
//     ports:
//       https: 9080
//   - address: in.foo.bar.com
//     ports:
//       https: 7080
// ```
// {{</tab>}}
// {{</tabset>}}
//
// With `HTTP_PROXY=http://localhost/`, calls from the application to
// `http://foo.bar.com` will be load balanced across the three domains
// specified above. In other words, a call to `http://foo.bar.com/baz` would
// be translated to `http://uk.foo.bar.com/baz`.
//
// The following example illustrates the usage of a `ServiceEntry`
// containing a subject alternate name
// whose format conforms to the [SPIFFE standard](https://github.com/spiffe/spiffe/blob/master/standards/SPIFFE-ID.md):
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: httpbin
//   namespace : httpbin-ns
// spec:
//   hosts:
//   - httpbin.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   endpoints:
//   - address: 2.2.2.2
//   - address: 3.3.3.3
//   subjectAltNames:
//   - "spiffe://cluster.local/ns/httpbin-ns/sa/httpbin-service-account"
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: httpbin
//   namespace : httpbin-ns
// spec:
//   hosts:
//   - httpbin.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   endpoints:
//   - address: 2.2.2.2
//   - address: 3.3.3.3
//   subjectAltNames:
//   - "spiffe://cluster.local/ns/httpbin-ns/sa/httpbin-service-account"
// ```
// {{</tab>}}
// {{</tabset>}}
//
// The following example demonstrates the use of `ServiceEntry` with a
// `workloadSelector` to handle the migration of a service
// `details.bookinfo.com` from VMs to Kubernetes. The service has two
// VM-based instances with sidecars as well as a set of Kubernetes
// pods managed by a standard deployment object. Consumers of this
// service in the mesh will be automatically load balanced across the
// VMs and Kubernetes.  VM for the `details.bookinfo.com`
// service. This VM has sidecar installed and bootstrapped using the
// `details-legacy` service account. The sidecar receives HTTP traffic
// on port 80 (wrapped in istio mutual TLS) and forwards it to the
// application on the localhost on the same port.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: WorkloadEntry
// metadata:
//   name: details-vm-1
// spec:
//   serviceAccount: details
//   address: 2.2.2.2
//   labels:
//     app: details
//     instance-id: vm1
// ---
// apiVersion: networking.istio.io/v1alpha3
// kind: WorkloadEntry
// metadata:
//   name: details-vm-2
// spec:
//   serviceAccount: details
//   address: 3.3.3.3
//   labels:
//     app: details
//     instance-id: vm2
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: WorkloadEntry
// metadata:
//   name: details-vm-1
// spec:
//   serviceAccount: details
//   address: 2.2.2.2
//   labels:
//     app: details
//     instance-id: vm1
// ---
// apiVersion: networking.istio.io/v1beta1
// kind: WorkloadEntry
// metadata:
//   name: details-vm-2
// spec:
//   serviceAccount: details
//   address: 3.3.3.3
//   labels:
//     app: details
//     instance-id: vm2
// ```
// {{</tab>}}
// {{</tabset>}}
//
// Assuming there is also a Kubernetes deployment with pod labels
// `app: details` using the same service account `details`, the
// following service entry declares a service spanning both VMs and
// Kubernetes:
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   workloadSelector:
//     labels:
//       app: details
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   workloadSelector:
//     labels:
//       app: details
// ```
// {{</tab>}}
// {{</tabset>}}

package v1alpha3

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for ServiceEntry
func (this *ServiceEntry) MarshalJSON() ([]byte, error) {
	str, err := ServiceEntryMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ServiceEntry
func (this *ServiceEntry) UnmarshalJSON(b []byte) error {
	return ServiceEntryUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ServiceEntryMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	ServiceEntryUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)
