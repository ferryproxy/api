/*
Copyright 2022 FerryProxy Authors.

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
	// Gateway is the default gateway of this Hub
	Gateway HubSpecGateway `json:"gateway"`

	// Override will replace the peer default gateway, the key is the name of peer Hub
	Override map[string]HubSpecGateway `json:"override,omitempty"`
}

// HubStatus defines the observed state of Hub
type HubStatus struct {
	// Phase is the phase of the Hub
	Phase string `json:"phase,omitempty"`

	// LastSynchronizationTimestamp is the last time synchronization
	LastSynchronizationTimestamp metav1.Time `json:"lastSynchronizationTimestamp,omitempty"`

	// Conditions current service state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

const (
	// ConnectedCondition means that Controller can connect to the Apiserver
	ConnectedCondition = "Connected"

	// TunnelHealthCondition means the Tunnel is healthy
	TunnelHealthCondition = "TunnelHealth"

	// HubReady means the hub is able to service.
	HubReady = "Ready"
)

// HubSpecGateway defines the desired state of Hub
type HubSpecGateway struct {
	// Reachable indicates that this Hub is reachable
	Reachable bool `json:"reachable"`

	// Address is the address of this Hub,
	// used when Reachable is true
	Address string `json:"address,omitempty"`

	// NavigationWay is a list of Hub names through which this Hub needs to reach other Hubs,
	// used when this Hub reaches other Hubs,
	// used by RoutePolicy to calculate Routes
	NavigationWay []HubSpecGatewayWay `json:"navigationWay,omitempty"`

	// ReceptionWay is a list of Hub names through which other hubs needs to reach this Hub,
	// used when other Hubs reaches this Hub and Reachable is true,
	// used by RoutePolicy to calculate Routes
	ReceptionWay []HubSpecGatewayWay `json:"receptionWay,omitempty"`

	// NavigationProxy is a list of proxies through which this Hub to reach other Hubs must need to go through,
	// used when this Hub reaches other Hubs
	NavigationProxy []HubSpecGatewayProxy `json:"navigationProxy,omitempty"`

	// ReceptionProxy is a list of proxies through which other Hubs to reach this Hub must need to go through,
	// used when other Hubs reaches this Hub and Reachable is true
	ReceptionProxy []HubSpecGatewayProxy `json:"receptionProxy,omitempty"`
}

// HubSpecGatewayProxy defines the desired state of HubSpecGateway
type HubSpecGatewayProxy struct {
	// HubName is the name of Hub to proxy,
	// cannot be specified together with Proxy
	HubName string `json:"hubName,omitempty"`

	// Proxy is the proxy to use,
	// cannot be specified together with HubName
	Proxy string `json:"proxy,omitempty"`
}

// HubSpecGatewayWay defines the desired state of HubSpecGateway
type HubSpecGatewayWay struct {
	// HubName is the name of Hub
	HubName string `json:"hubName"`
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
