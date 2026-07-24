package controller

import (
	"context"
	"testing"

	platformv1 "github.com/abdallauno1/platform-engineering-kubernetes-operator/api/v1"
)

func TestReconcileCreatesMissingWorkload(t *testing.T) {
	client := NewMemoryWorkloadClient()
	reconciler := NewAIPlatformReconciler(client)
	resource := platformv1.NewDefault("demo-ai", "platform")
	resource.Spec.Image = "ghcr.io/example/ai-platform:day2"
	resource.Spec.Replicas = 3

	result, err := reconciler.Reconcile(context.Background(), resource)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Action != ActionCreated {
		t.Fatalf("expected Created, got %s", result.Action)
	}
	if result.Status.Phase != platformv1.PhaseReady || result.Status.ReadyReplicas != 3 {
		t.Fatalf("unexpected status: %+v", result.Status)
	}
}

func TestReconcileUpdatesDriftedWorkload(t *testing.T) {
	client := NewMemoryWorkloadClient()
	reconciler := NewAIPlatformReconciler(client)
	resource := platformv1.NewDefault("demo-ai", "platform")

	if _, err := reconciler.Reconcile(context.Background(), resource); err != nil {
		t.Fatalf("initial reconciliation failed: %v", err)
	}
	resource.Metadata.Generation = 2
	resource.Spec.Image = "ghcr.io/example/ai-platform:v2"
	resource.Spec.Replicas = 4

	result, err := reconciler.Reconcile(context.Background(), resource)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Action != ActionUpdated {
		t.Fatalf("expected Updated, got %s", result.Action)
	}
	if result.Status.ObservedGeneration != 2 {
		t.Fatalf("unexpected observed generation: %d", result.Status.ObservedGeneration)
	}
}

func TestReconcileIsIdempotent(t *testing.T) {
	client := NewMemoryWorkloadClient()
	reconciler := NewAIPlatformReconciler(client)
	resource := platformv1.NewDefault("demo-ai", "platform")

	if _, err := reconciler.Reconcile(context.Background(), resource); err != nil {
		t.Fatalf("initial reconciliation failed: %v", err)
	}
	result, err := reconciler.Reconcile(context.Background(), resource)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Action != ActionUnchanged {
		t.Fatalf("expected Unchanged, got %s", result.Action)
	}
}

func TestReconcileSetsFailedStatusForInvalidResource(t *testing.T) {
	client := NewMemoryWorkloadClient()
	reconciler := NewAIPlatformReconciler(client)
	resource := platformv1.NewDefault("demo-ai", "platform")
	resource.Spec.Replicas = 0

	result, err := reconciler.Reconcile(context.Background(), resource)
	if err == nil {
		t.Fatal("expected validation error")
	}
	if result.Action != ActionFailed || result.Status.Phase != platformv1.PhaseFailed {
		t.Fatalf("unexpected result: %+v", result)
	}
}
