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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HubSpec defines the desired state of Hub
type HubSpec struct {
	// Gateway is the default gateway of this Hub.
	Gateway HubSpecGateway `json:"gateway"`
	// Override will replace the peer default gateway, key is the peer Hub
	Override map[string]HubSpecGateway `json:"override,omitempty"`
}

// HubStatus defines the observed state of Hub
type HubStatus struct {
	// Phase is the phase of the Hub.
	Phase string `json:"phase,omitempty"`
	// LastSynchronizationTimestamp is the last time synchronization
	LastSynchronizationTimestamp metav1.Time `json:"lastSynchronizationTimestamp,omitempty"`
	// Conditions current service state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// HubSpecGateway defines the desired state of Hub
type HubSpecGateway struct {
	// Reachable indicates that this cluster is reachable, the default unreachable.
	Reachable bool `json:"reachable"`
	// Address is the address of the cluster.
	Address string `json:"address,omitempty"`
	// Navigation is the navigation of the Hub.
	Navigation HubSpecGatewayWays `json:"navigation,omitempty"`
	// Reception is the reception of the Hub.
	Reception HubSpecGatewayWays `json:"reception,omitempty"`
}

// HubSpecGatewayWays defines the desired state of HubSpecGateway
type HubSpecGatewayWays []HubSpecGatewayWay

// HubSpecGatewayWay defines the desired state of HubSpecGateway
type HubSpecGatewayWay struct {
	// HubName is the name of Hub to proxy.
	// cannot be specified together with Proxy.
	HubName string `json:"hubName,omitempty"`
	// Proxy is the proxy to use.
	// cannot be specified together with HubName.
	Proxy string `json:"proxy,omitempty"`
}

// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="status",type="string",JSONPath=".status.phase"
//+kubebuilder:printcolumn:name="last-synchronization",type="date",JSONPath=".status.lastSynchronizationTimestamp"
//+kubebuilder:printcolumn:name="age",type="date",JSONPath=".metadata.creationTimestamp"

// Hub is the Schema for the hubs API
type Hub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HubSpec   `json:"spec,omitempty"`
	Status HubStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HubList contains a list of Hub
type HubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hub `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Hub{}, &HubList{})
}
