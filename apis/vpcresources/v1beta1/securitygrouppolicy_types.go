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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Important: Run "make" to regenerate code after modifying this file

// SecurityGroupPolicySpec defines the desired state of SecurityGroupPolicy
type SecurityGroupPolicySpec struct {
	PodSelector            *metav1.LabelSelector  `json:"podSelector,omitempty"`
	ServiceAccountSelector ServiceAccountSelector `json:"serviceAccountSelector,omitempty"`
	SecurityGroups         GroupIds               `json:"securityGroups,omitempty"`
}

// GroupIds has security groups carried in the CRD SecurityGroupPolicy.
type GroupIds struct {
	// Groups is the list of EC2 Security Groups Ids that need to be applied to the ENI of a Pod.
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=5
	Groups []string `json:"groupIds,omitempty"`
}

// ServiceAccountSelector provides the selector for service account label matches and its name matches.
type ServiceAccountSelector struct {
	*metav1.LabelSelector `json:",omitempty"`
	// matchNames is the list of service account names. The requirements are ANDed
	// +kubebuilder:validation:MinItems=1
	MatchNames []string `json:"matchNames,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="Security-Group-Ids",type=string,JSONPath=`.spec.securityGroups.groupIds`,description="The security group IDs to apply to the elastic network interface of pods that match this policy"
// +kubebuilder:resource:shortName=sgp

// Custom Resource Definition for applying security groups to pods
type SecurityGroupPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SecurityGroupPolicySpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupPolicyList contains a list of SecurityGroupPolicy
type SecurityGroupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroupPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecurityGroupPolicy{}, &SecurityGroupPolicyList{})
}