package controller

import (
	"context"
	"fmt"

	platformv1 "github.com/abdallamady/kubernetes-operator-go/api/v1"
)

type Result struct {
	Requeue bool
	Message string
}

type DesiredDeployment struct {
	Name      string
	Namespace string
	Image     string
	Replicas  int
	Labels    map[string]string
}

type AIPlatformReconciler struct{}

func NewAIPlatformReconciler() *AIPlatformReconciler { return &AIPlatformReconciler{} }

func (r *AIPlatformReconciler) Reconcile(ctx context.Context, platform platformv1.AIPlatform) (Result, DesiredDeployment, error) {
	select {
	case <-ctx.Done():
		return Result{}, DesiredDeployment{}, ctx.Err()
	default:
	}

	if err := platform.Validate(); err != nil {
		return Result{Requeue: false, Message: "invalid custom resource"}, DesiredDeployment{}, err
	}

	desired := DesiredDeployment{
		Name:      fmt.Sprintf("%s-workload", platform.Metadata.Name),
		Namespace: platform.Metadata.Namespace,
		Image:     platform.Spec.Image,
		Replicas:  platform.Spec.Replicas,
		Labels: map[string]string{
			"app.kubernetes.io/name":       platform.Metadata.Name,
			"app.kubernetes.io/managed-by": "ai-platform-operator",
			"platform.mady.dev/kind":       platformv1.KindAIPlatform,
		},
	}

	return Result{Requeue: false, Message: "desired deployment calculated"}, desired, nil
}
