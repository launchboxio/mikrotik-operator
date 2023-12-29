/*
Copyright 2023.

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

// AddressSpec defines the desired state of Address
type AddressSpec struct {
	Address   string `json:"address"`
	Interface string `json:"interface"`
	Disabled  bool   `json:"disabled,omitempty"`
	Comment   string `json:"comment,omitempty"`
}

// AddressStatus defines the observed state of Address
type AddressStatus struct {
	Identifier string `json:"identifier,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Address is the Schema for the addresses API
type Address struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AddressSpec   `json:"spec,omitempty"`
	Status AddressStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AddressList contains a list of Address
type AddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Address `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Address{}, &AddressList{})
}
