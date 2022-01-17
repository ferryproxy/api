/*
Copyright 2022 Shiming Zhang.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterInformationSpec defines the desired state of ClusterInformation
type ClusterInformationSpec struct {
	// Kubeconfig is the kubeconfig of the cluster,
	// cannot be specified together with InCluster.
	Kubeconfig []byte `json:"kubeconfig,omitempty"`
	// InCluster indicates whether the cluster is itself cluster,
	// cannot be specified together with Kubeconfig.
	InCluster bool `json:"inCluster,omitempty"`
	// Ingress is the ingress of the cluster.
	Ingress *ClusterInformationSpecRoute `json:"ingress,omitempty"`
	// Egress is the egress of the cluster.
	Egress *ClusterInformationSpecRoute `json:"egress,omitempty"`
}

// ClusterInformationStatus defines the observed state of ClusterInformation
type ClusterInformationStatus struct {
	// ExportedTo is the list of the cluster information exported to.
	ExportedTo []string `json:"exportedTo,omitempty"`
	// ImportedFrom is the list of the cluster information imported from.
	ImportedFrom []string `json:"importedFrom,omitempty"`
	// Conditions current service state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// ClusterInformationSpecRoute defines the desired state of ClusterInformation
type ClusterInformationSpecRoute struct {
	// Port is the port to expose.
	Port int32 `json:"port"`
	// IP is the IP address to expose.
	// cannot be specified together with ServiceName and ServiceNamespace.
	IP string `json:"ip,omitempty"`
	// ServiceName is the name of the service to expose.
	// cannot be specified together with IP.
	ServiceName string `json:"serviceName,omitempty"`
	// ServiceNamespace is the namespace of the service to expose.
	// cannot be specified together with IP.
	ServiceNamespace string `json:"serviceNamespace,omitempty"`
	// Proxies is a list of proxies to use for the route
	Proxies map[string]Proxies `json:"proxies,omitempty"`
	// DefaultProxies is a default list of proxies to use for the route
	DefaultProxies Proxies `json:"defaultProxies,omitempty"`
}

// Proxies defines the desired state of ClusterInformationSpecRoute
type Proxies []Proxy

// Proxy defines the desired state of ClusterInformationSpecRoute
type Proxy struct {
	// ClusterName is the name of the cluster to proxy to.
	// cannot be specified together with Proxy.
	ClusterName string `json:"clusterName,omitempty"`
	// Proxy is the proxy to use.
	// cannot be specified together with ClusterName.
	Proxy string `json:"proxy,omitempty"`
}

// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ClusterInformation is the Schema for the clusterinformations API
type ClusterInformation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterInformationSpec   `json:"spec,omitempty"`
	Status ClusterInformationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterInformationList contains a list of ClusterInformation
type ClusterInformationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterInformation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterInformation{}, &ClusterInformationList{})
}
