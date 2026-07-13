package controller

import (
	"context"
	"testing"

	platformv1 "github.com/abdallamady/kubernetes-operator-go/api/v1"
)

func TestReconcileBuildsDesiredDeployment(t *testing.T) {
	resource := platformv1.NewDefault("demo-ai", "platform")
	resource.Spec.Image = "ghcr.io/example/ai-platform:day1"
	resource.Spec.Replicas = 3

	result, desired, err := NewAIPlatformReconciler().Reconcile(context.Background(), resource)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Requeue {
		t.Fatalf("expected no requeue")
	}
	if desired.Name != "demo-ai-workload" {
		t.Fatalf("unexpected deployment name: %s", desired.Name)
	}
	if desired.Image != resource.Spec.Image || desired.Replicas != 3 {
		t.Fatalf("desired deployment mismatch: %+v", desired)
	}
}

func TestReconcileRejectsInvalidResource(t *testing.T) {
	resource := platformv1.NewDefault("demo-ai", "platform")
	resource.Spec.Replicas = 0

	_, _, err := NewAIPlatformReconciler().Reconcile(context.Background(), resource)
	if err == nil {
		t.Fatalf("expected validation error")
	}
}
