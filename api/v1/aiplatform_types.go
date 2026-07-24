package v1

import (
	"fmt"
	"strings"
)

const (
	KindAIPlatform      = "AIPlatform"
	DefaultImage        = "nginx:1.27-alpine"
	DefaultReplicas int = 1

	PhasePending = "Pending"
	PhaseReady   = "Ready"
	PhaseFailed  = "Failed"
)

type AIPlatformSpec struct {
	Image    string            `json:"image"`
	Replicas int               `json:"replicas"`
	Env      map[string]string `json:"env,omitempty"`
}

type AIPlatformStatus struct {
	Phase              string `json:"phase"`
	Message            string `json:"message,omitempty"`
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	ReadyReplicas      int    `json:"readyReplicas,omitempty"`
}

type AIPlatform struct {
	APIVersion string           `json:"apiVersion"`
	Kind       string           `json:"kind"`
	Metadata   ObjectMeta       `json:"metadata"`
	Spec       AIPlatformSpec   `json:"spec"`
	Status     AIPlatformStatus `json:"status,omitempty"`
}

type ObjectMeta struct {
	Name       string            `json:"name"`
	Namespace  string            `json:"namespace,omitempty"`
	Generation int64             `json:"generation,omitempty"`
	Labels     map[string]string `json:"labels,omitempty"`
}

func (p AIPlatform) Validate() error {
	if strings.TrimSpace(p.Metadata.Name) == "" {
		return fmt.Errorf("metadata.name is required")
	}
	if strings.TrimSpace(p.Spec.Image) == "" {
		return fmt.Errorf("spec.image is required")
	}
	if p.Spec.Replicas < 1 {
		return fmt.Errorf("spec.replicas must be >= 1")
	}
	return nil
}

func NewDefault(name, namespace string) AIPlatform {
	return AIPlatform{
		APIVersion: "platform.mady.dev/v1",
		Kind:       KindAIPlatform,
		Metadata: ObjectMeta{
			Name:       name,
			Namespace:  namespace,
			Generation: 1,
			Labels: map[string]string{
				"app.kubernetes.io/name":       name,
				"app.kubernetes.io/managed-by": "ai-platform-operator",
			},
		},
		Spec: AIPlatformSpec{
			Image:    DefaultImage,
			Replicas: DefaultReplicas,
			Env:      map[string]string{},
		},
		Status: AIPlatformStatus{
			Phase:   PhasePending,
			Message: "Waiting for first reconciliation",
		},
	}
}
