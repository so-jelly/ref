package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SubSpec defines the desired state of Sub
type SubSpec MainSpec

// SubStatus defines the observed state of Sub
type SubStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Sub is the Schema for the subs API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=subs,scope=Namespaced
type Sub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SubSpec   `json:"spec,omitempty"`
	Status SubStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SubList contains a list of Sub
type SubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Sub `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Sub{}, &SubList{})
}
