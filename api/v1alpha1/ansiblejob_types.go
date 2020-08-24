/*


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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AnsibleJobSpec defines the desired state of AnsibleJob
type AnsibleJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	TowerAuthSecret string `json:"tower_auth_secret,omitempty"`
	JobTemplateName string `json:"job_tmeplate_name,omitempty"`
	Inventory       string `json:"inventory,omitempty"`
	ExtraVars       string `json:"extra_vars,omitempty"`
}

type AnsibleJobResult struct {
	Elapsed  string      `json:"elapsed,omitempty"`
	Finished metav1.Time `json:"finished,omitempty"`
	Started  metav1.Time `json:"started,omitempty"`
	Status   string      `json:"status"`
}

type Conditions struct {
	AnsibleResults     []AnsibleResult `json:"ansibleresult,omitempty"`
	LastTransitionTime metav1.Time     `json:"finished,omitempty"`
	Message            string          `json:"message,omitempty"`
	Reason             string          `json:"reason,omitempty"`
	Status             string          `json:"status,omitempty"`
	Type               string          `json:"type,omitempty"`
}

type AnsibleResult struct {
	Changed    int         `json:"changed,omitempty"`
	Completion metav1.Time `json:"completion,omitempty"`
	Failures   int         `json:"failurs,omitempty"`
	Ok         int         `json:"ok,omitempty"`
	Skipped    int         `json:"skipped,omitempty"`
}

// AnsibleJobStatus defines the observed state of AnsibleJob
type AnsibleJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	AnsibleJobResult `json:"ansiblejobresult,omitempty"`
	Conditions       `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// AnsibleJob is the Schema for the ansiblejobs API
type AnsibleJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AnsibleJobSpec   `json:"spec,omitempty"`
	Status AnsibleJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnsibleJobList contains a list of AnsibleJob
type AnsibleJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnsibleJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AnsibleJob{}, &AnsibleJobList{})
}
