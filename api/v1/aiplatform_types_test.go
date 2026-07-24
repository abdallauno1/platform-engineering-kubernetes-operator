package v1

import "testing"

func TestNewDefaultAIPlatform(t *testing.T) {
	platform := NewDefault("demo", "default")
	if platform.Kind != KindAIPlatform {
		t.Fatalf("unexpected kind: %s", platform.Kind)
	}
	if platform.Spec.Image != DefaultImage {
		t.Fatalf("unexpected image: %s", platform.Spec.Image)
	}
	if platform.Metadata.Generation != 1 {
		t.Fatalf("unexpected generation: %d", platform.Metadata.Generation)
	}
	if platform.Status.Phase != PhasePending {
		t.Fatalf("unexpected phase: %s", platform.Status.Phase)
	}
	if err := platform.Validate(); err != nil {
		t.Fatalf("default resource should be valid: %v", err)
	}
}

func TestValidateRejectsInvalidReplicas(t *testing.T) {
	platform := NewDefault("demo", "default")
	platform.Spec.Replicas = 0
	if err := platform.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}
