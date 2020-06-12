package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// imageJob describes a imageJob.
type ImageJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ImageJobSpec `json:"spec"`
}

// imagejobSpec is the spec for a jobImage resource
type ImageJobSpec struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Encoding string `json:"encoding,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// imageJobList is a list of imageJob resources
type ImageJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ImageJob `json:"items"`
}
