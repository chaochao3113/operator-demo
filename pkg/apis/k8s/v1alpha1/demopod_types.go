package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DemoPodSpec defines the desired state of DemoPod
//期望状态
type DemoPodSpec struct {
	Replicas int `json:"replicas"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// DemoPodStatus defines the observed state of DemoPod
//运行状态
type DemoPodStatus struct {
	Replicas int `json:"replicas"`
	PodNames []string `json:"podNames"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DemoPod is the Schema for the demopods API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=demopods,scope=Namespaced
type DemoPod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DemoPodSpec   `json:"spec,omitempty"`
	Status DemoPodStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DemoPodList contains a list of DemoPod
type DemoPodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DemoPod `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DemoPod{}, &DemoPodList{})
}
