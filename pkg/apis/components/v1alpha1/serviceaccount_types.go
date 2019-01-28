// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceAccount is the Schema for the serviceaccounts API
// +k8s:openapi-gen=true
type ServiceAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceAccountSpec   `json:"spec,omitempty"`
	Status ServiceAccountStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceAccountList contains a list of ServiceAccount
type ServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccount `json:"items"`
}

// ServiceAccountSpec defines the desired state of Secret
type ServiceAccountSpec struct {
	corev1.ServiceAccount `json:",inline"`
	Cluster               string `json:"cluster"`
	Name                  string `json:"name"`
	Namespace             string `json:"namespace"`
}

// ServiceAccountStatus defines the observed state of ServiceAccount
type ServiceAccountStatus struct {
	Status string `json:"status,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ServiceAccount{}, &ServiceAccountList{})
}
