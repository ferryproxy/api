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

// FerryPolicySpec defines the desired state of FerryPolicy
type FerryPolicySpec struct {
	Rules []FerryPolicySpecRule `json:"rules"`
}

// FerryPolicySpecRule defines the desired rule of FerryPolicyRule
// Service mapping from Exports cluster to Imports cluster
type FerryPolicySpecRule struct {
	// Exports is a list of exports of the FerryPolicy.
	Exports []FerryPolicySpecRuleExport `json:"exports"`
	// Imports is a list of imports of the FerryPolicy.
	Imports []FerryPolicySpecRuleImport `json:"imports"`
}

// FerryPolicySpecRuleExport defines the desired export of FerryPolicyRule
type FerryPolicySpecRuleExport struct {
	// ClusterName is specifies the name of the cluster
	ClusterName string `json:"clusterName"`
	// Match is specifies the service of matched
	Match Match `json:"match,omitempty"`
}

// FerryPolicySpecRuleImport defines the desired import of FerryPolicyRule
type FerryPolicySpecRuleImport struct {
	// ClusterName is specifies the name of the cluster
	ClusterName string `json:"clusterName"`
	// Match is specifies the service of matched
	Match Match `json:"match,omitempty"`
}

// Match defines the desired match of FerryPolicyRule
type Match struct {
	// Labels is specifies the labels of matched
	// cannot be specified together with Name
	Labels map[string]string `json:"labels,omitempty"`
	// Namespace is specifies the namespace of matched
	Namespace string `json:"namespace,omitempty"`
	// Name is specifies the name of matched
	// cannot be specified together with Labels
	Name string `json:"name,omitempty"`
}

// FerryPolicyStatus defines the observed state of FerryPolicy
type FerryPolicyStatus struct {
	// LastSynchronizationTimestamp is the last time synchronization to the cluster.
	LastSynchronizationTimestamp metav1.Time `json:"lastSynchronizationTimestamp,omitempty"`
	// Conditions current service state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="last-synchronization",type="string",JSONPath=".status.lastSynchronizationTimestamp"
//+kubebuilder:printcolumn:name="age",type="date",JSONPath=".metadata.creationTimestamp"

// FerryPolicy is the Schema for the ferrypolicies API
type FerryPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FerryPolicySpec   `json:"spec,omitempty"`
	Status FerryPolicyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FerryPolicyList contains a list of FerryPolicy
type FerryPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FerryPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FerryPolicy{}, &FerryPolicyList{})
}
