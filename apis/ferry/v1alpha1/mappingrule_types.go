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

// MappingRuleSpec defines the desired state of MappingRule
type MappingRuleSpec struct {
	// Import is one of import of the  MappingRule.
	Import MappingRuleSpecPorts `json:"import"`
	// Export is one of export of the  MappingRule.
	Export MappingRuleSpecPorts `json:"export"`
}

// MappingRuleSpecPorts  defines the desired state of MappingRuleSpec
type MappingRuleSpecPorts struct {
	// ClusterName is specifies the name of the cluster
	ClusterName string `json:"clusterName"`
	// Service is the service
	Service MappingRuleSpecPortsService `json:"service"`
}

// MappingRuleSpecPortsService  defines the desired state of MappingRuleSpecPorts
type MappingRuleSpecPortsService struct {
	// Name is the service name.
	Name string `json:"name"`
	// Namespace is the service namespace.
	Namespace string `json:"namespace"`
}

// MappingRuleStatus defines the observed state of MappingRule
type MappingRuleStatus struct {
	// Export is describe of the export
	Export string `json:"export,omitempty"`
	// Import is describe of the import
	Import string `json:"import,omitempty"`
	// Phase is the phase of the mapping rule.
	Phase string `json:"phase,omitempty"`
	// LastSynchronizationTimestamp is the last time synchronization to the cluster.
	LastSynchronizationTimestamp metav1.Time `json:"lastSynchronizationTimestamp,omitempty"`
	// Conditions current service state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="export",type="string",JSONPath=".status.export"
//+kubebuilder:printcolumn:name="import",type="string",JSONPath=".status.import"
//+kubebuilder:printcolumn:name="status",type="string",JSONPath=".status.phase"
//+kubebuilder:printcolumn:name="last-synchronization",type="date",JSONPath=".status.lastSynchronizationTimestamp"
//+kubebuilder:printcolumn:name="age",type="date",JSONPath=".metadata.creationTimestamp"

// MappingRule is the Schema for the mappingrules API
type MappingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MappingRuleSpec   `json:"spec,omitempty"`
	Status MappingRuleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MappingRuleList contains a list of MappingRule
type MappingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MappingRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MappingRule{}, &MappingRuleList{})
}
