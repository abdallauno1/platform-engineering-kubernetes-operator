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
	if err := platform.Validate(); err != nil {
		t.Fatalf("default resource should be valid: %v", err)
	}
}
