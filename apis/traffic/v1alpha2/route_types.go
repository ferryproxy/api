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

// RouteSpec defines the desired state of Route
type RouteSpec struct {
	// Import is one of import of the Route
	Import RouteSpecRule `json:"import"`

	// Export is one of export of the Route
	Export RouteSpecRule `json:"export"`
}

// RouteStatus defines the observed state of Route
type RouteStatus struct {
	// Way is describe of the way
	Way string `json:"way,omitempty"`

	// Export is describe of the export
	Export string `json:"export,omitempty"`

	// Import is describe of the import
	Import string `json:"import,omitempty"`

	// Phase is the phase of the Route
	Phase string `json:"phase,omitempty"`

	// LastSynchronizationTimestamp is the last time synchronization
	LastSynchronizationTimestamp metav1.Time `json:"lastSynchronizationTimestamp,omitempty"`

	// Conditions current service state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// RouteSpecRule  defines the desired state of RouteSpec
type RouteSpecRule struct {
	// HubName is specifies the name of the Hub
	HubName string `json:"hubName"`

	// Service is the service
	Service RouteSpecRuleService `json:"service"`
}

// RouteSpecRuleService  defines the desired state of RouteSpecRule
type RouteSpecRuleService struct {
	// Name is the service name
	Name string `json:"name"`

	// Namespace is the service namespace
	Namespace string `json:"namespace"`
}

// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="export",type="string",JSONPath=".status.export"
//+kubebuilder:printcolumn:name="way",type="string",JSONPath=".status.way"
//+kubebuilder:printcolumn:name="import",type="string",JSONPath=".status.import"
//+kubebuilder:printcolumn:name="status",type="string",JSONPath=".status.phase"
//+kubebuilder:printcolumn:name="last-synchronization",type="date",JSONPath=".status.lastSynchronizationTimestamp"
//+kubebuilder:printcolumn:name="age",type="date",JSONPath=".metadata.creationTimestamp"

// Route is the Schema for the routes API
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouteSpec   `json:"spec,omitempty"`
	Status RouteStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RouteList contains a list of Route
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
