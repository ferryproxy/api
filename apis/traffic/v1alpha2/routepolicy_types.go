/*
Copyright 2022 FerryProxy.

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

// RoutePolicySpec defines the desired state of RoutePolicy
type RoutePolicySpec struct {
	// Exports is a list of exports of the RoutePolicy
	Exports []RoutePolicySpecRule `json:"exports"`

	// Imports is a list of imports of the RoutePolicy
	Imports []RoutePolicySpecRule `json:"imports"`
}

// RoutePolicyStatus defines the observed state of RoutePolicy
type RoutePolicyStatus struct {
	// RouteCount is the number of Route in the RoutePolicy
	RouteCount int `json:"routeCount,omitempty"`

	// Phase is the phase of the RoutePolicy
	Phase string `json:"phase,omitempty"`

	// LastSynchronizationTimestamp is the last time synchronization
	LastSynchronizationTimestamp metav1.Time `json:"lastSynchronizationTimestamp,omitempty"`

	// Conditions current service state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// RoutePolicySpecRule defines the desired import of RoutePolicySpec
type RoutePolicySpecRule struct {
	// HubName is specifies the name of the Hub
	HubName string `json:"hubName"`

	// Service is specifies the service of matched
	Service RoutePolicySpecRuleService `json:"service,omitempty"`
}

// RoutePolicySpecRuleService defines the desired match of RoutePolicySpecRule
type RoutePolicySpecRuleService struct {
	// Labels is specifies the labels of matched
	Labels map[string]string `json:"labels,omitempty"`

	// Namespace is specifies the namespace of matched
	Namespace string `json:"namespace,omitempty"`

	// Name is specifies the name of matched
	Name string `json:"name,omitempty"`
}

// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="route-count",type="integer",JSONPath=".status.routeCount"
//+kubebuilder:printcolumn:name="status",type="string",JSONPath=".status.phase"
//+kubebuilder:printcolumn:name="last-synchronization",type="date",JSONPath=".status.lastSynchronizationTimestamp"
//+kubebuilder:printcolumn:name="age",type="date",JSONPath=".metadata.creationTimestamp"

// RoutePolicy is the Schema for the routepolicies API
type RoutePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RoutePolicySpec   `json:"spec,omitempty"`
	Status RoutePolicyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RoutePolicyList contains a list of RoutePolicy
type RoutePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoutePolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RoutePolicy{}, &RoutePolicyList{})
}
