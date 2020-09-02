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
	"encoding/json"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AnsibleJobSpec defines the desired state of AnsibleJob
type AnsibleJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	TowerAuthSecretNamespace string          `json:"tower_auth_secret_namespace,omitempty"`
	TowerAuthSecretName      string          `json:"tower_auth_secret_name,omitempty"`
	JobTemplateName          string          `json:"job_tmeplate_name,omitempty"`
	Inventory                string          `json:"inventory,omitempty"`
	ExtraVars                json.RawMessage `json:"extra_vars,omitempty"`
}

type AnsibleJobResult struct {
	Elapsed  string `json:"elapsed,omitempty"`
	Finished string `json:"finished,omitempty"`
	Started  string `json:"started,omitempty"`
	Status   string `json:"status,omitempty"`
}

//bridging from https://github.com/operator-framework/operator-sdk/blob/master/internal/ansible/controller/status/types.go
// AnsibleResult - encapsulation of the ansible result.
type AnsibleResult struct {
	Ok               int         `json:"ok"`
	Changed          int         `json:"changed"`
	Skipped          int         `json:"skipped"`
	Failures         int         `json:"failures"`
	TimeOfCompletion metav1.Time `json:"completion"`
}

// ConditionType - type of condition
type ConditionType string

const (
	// RunningConditionType - condition type of running.
	RunningConditionType ConditionType = "Running"
	// FailureConditionType - condition type of failure.
	FailureConditionType ConditionType = "Failure"

	JobScussed = "successful"
)

// Condition - the condition for the ansible operator.
type Condition struct {
	Type               ConditionType      `json:"type,omitempty"`
	Status             v1.ConditionStatus `json:"status,omitempty"`
	LastTransitionTime metav1.Time        `json:"lastTransitionTime,omitempty"`
	AnsibleResult      *AnsibleResult     `json:"ansibleResult,omitempty"`
	Reason             string             `json:"reason,omitempty"`
	Message            string             `json:"message,omitempty"`
}

// AnsibleJobStatus defines the observed state of AnsibleJob
type AnsibleJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	AnsibleJobResult `json:"ansiblejobresult,omitempty"`
	Condition        `json:"conditions,omitempty"`
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
