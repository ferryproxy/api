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
	// Gateway is the default gateway of the cluster.
	Gateway ClusterInformationSpecGateway `json:"gateway"`
	// Override is will replace the default gateway that will reach the target
	Override map[string]ClusterInformationSpecGateway `json:"override,omitempty"`
}

// ClusterInformationStatus defines the observed state of ClusterInformation
type ClusterInformationStatus struct {
	// ExportedTo is the list of the cluster information exported to.
	ExportedTo []string `json:"exportedTo,omitempty"`
	// ImportedFrom is the list of the cluster information imported from.
	ImportedFrom []string `json:"importedFrom,omitempty"`
	// LastSynchronizationTimestamp is the last time synchronization to the cluster.
	LastSynchronizationTimestamp metav1.Time `json:"lastSynchronizationTimestamp,omitempty"`
	// Phase is the phase of the cluster information.
	Phase string `json:"phase,omitempty"`
	// Conditions current service state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// ClusterInformationSpecGateway defines the desired state of ClusterInformation
type ClusterInformationSpecGateway struct {
	// Reachable indicates that this cluster is reachable, the default unreachable.
	Reachable bool `json:"reachable"`
	// Address is the address of the cluster.
	Address string `json:"address,omitempty"`
	// Navigation is the navigation of the cluster.
	Navigation ClusterInformationSpecGatewayWays `json:"navigation,omitempty"`
	// Reception is the reception of the cluster.
	Reception ClusterInformationSpecGatewayWays `json:"reception,omitempty"`
}

// ClusterInformationSpecGatewayWays defines the desired state of ClusterInformationSpecGateway
type ClusterInformationSpecGatewayWays []ClusterInformationSpecGatewayWay

// ClusterInformationSpecGatewayWay defines the desired state of ClusterInformationSpecGateway
type ClusterInformationSpecGatewayWay struct {
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
//+kubebuilder:printcolumn:name="exported-to",type="string",JSONPath=".status.exportedTo"
//+kubebuilder:printcolumn:name="imported-from",type="string",JSONPath=".status.importedFrom"
//+kubebuilder:printcolumn:name="status",type="string",JSONPath=".status.phase"
//+kubebuilder:printcolumn:name="last-synchronization",type="string",JSONPath=".status.lastSynchronizationTimestamp"
//+kubebuilder:printcolumn:name="age",type="date",JSONPath=".metadata.creationTimestamp"

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
