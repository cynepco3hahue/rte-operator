/*
Copyright 2021.

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

// This is borrowed from the kubernetes source, because controller-gen
// complains about the kube native type:
// encountered struct field "Namespace" without JSON tag in type "NamespacedName"
// at least until kube catches up, we just inline this simple struct here.

// NamespacedName comprises a resource name, with a mandatory namespace,
// rendered as "<namespace>/<name>".
type NamespacedName struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
}

const (
	Separator = '/'
)

// String returns the general purpose string representation
func (n NamespacedName) String() string {
	return n.Namespace + string(Separator) + n.Name
}

// ResourceTopologyExporterSpec defines the desired state of ResourceTopologyExporter
type ResourceTopologyExporterSpec struct {
}

// ResourceTopologyExporterStatus defines the observed state of ResourceTopologyExporter
type ResourceTopologyExporterStatus struct {
	DaemonSet *NamespacedName `json:"daemonset,omitempty"`

	// Conditions show the current state of the ResourceTopologyExporter Operator
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName=rte,path=resourcetopologyexporters

// ResourceTopologyExporter is the Schema for the resourcetopologyexporters API
type ResourceTopologyExporter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceTopologyExporterSpec   `json:"spec,omitempty"`
	Status ResourceTopologyExporterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ResourceTopologyExporterList contains a list of ResourceTopologyExporter
type ResourceTopologyExporterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceTopologyExporter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResourceTopologyExporter{}, &ResourceTopologyExporterList{})
}
