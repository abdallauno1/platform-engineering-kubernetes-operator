package controller

import (
	"context"
	"errors"
	"fmt"

	platformv1 "github.com/abdallauno1/platform-engineering-kubernetes-operator/api/v1"
)

type Action string

const (
	ActionCreated   Action = "Created"
	ActionUpdated   Action = "Updated"
	ActionUnchanged Action = "Unchanged"
	ActionFailed    Action = "Failed"
)

type Result struct {
	Action  Action
	Requeue bool
	Message string
	Status  platformv1.AIPlatformStatus
}

type AIPlatformReconciler struct {
	client WorkloadClient
}

func NewAIPlatformReconciler(client WorkloadClient) *AIPlatformReconciler {
	return &AIPlatformReconciler{client: client}
}

func (r *AIPlatformReconciler) Reconcile(ctx context.Context, platform platformv1.AIPlatform) (Result, error) {
	if err := ctx.Err(); err != nil {
		return failedResult(platform, err), err
	}
	if err := platform.Validate(); err != nil {
		return failedResult(platform, err), err
	}

	desired := desiredWorkload(platform)
	current, err := r.client.Get(ctx, desired.Namespace, desired.Name)

	switch {
	case errors.Is(err, ErrNotFound):
		if err := r.client.Create(ctx, desired); err != nil {
			return failedResult(platform, err), err
		}
		return readyResult(platform, ActionCreated, "managed workload created"), nil
	case err != nil:
		return failedResult(platform, err), err
	case !current.Equal(desired):
		if err := r.client.Update(ctx, desired); err != nil {
			return failedResult(platform, err), err
		}
		return readyResult(platform, ActionUpdated, "managed workload updated after drift detection"), nil
	default:
		return readyResult(platform, ActionUnchanged, "managed workload already matches desired state"), nil
	}
}

func desiredWorkload(platform platformv1.AIPlatform) Workload {
	return Workload{
		Name:      fmt.Sprintf("%s-workload", platform.Metadata.Name),
		Namespace: platform.Metadata.Namespace,
		Image:     platform.Spec.Image,
		Replicas:  platform.Spec.Replicas,
		Env:       cloneMap(platform.Spec.Env),
		Labels: map[string]string{
			"app.kubernetes.io/name":       platform.Metadata.Name,
			"app.kubernetes.io/managed-by": "ai-platform-operator",
			"platform.mady.dev/kind":       platformv1.KindAIPlatform,
		},
	}
}

func readyResult(platform platformv1.AIPlatform, action Action, message string) Result {
	return Result{
		Action:  action,
		Requeue: false,
		Message: message,
		Status: platformv1.AIPlatformStatus{
			Phase:              platformv1.PhaseReady,
			Message:            message,
			ObservedGeneration: platform.Metadata.Generation,
			ReadyReplicas:      platform.Spec.Replicas,
		},
	}
}

func failedResult(platform platformv1.AIPlatform, err error) Result {
	return Result{
		Action:  ActionFailed,
		Requeue: false,
		Message: err.Error(),
		Status: platformv1.AIPlatformStatus{
			Phase:              platformv1.PhaseFailed,
			Message:            err.Error(),
			ObservedGeneration: platform.Metadata.Generation,
		},
	}
}

func cloneMap(input map[string]string) map[string]string {
	output := make(map[string]string, len(input))
	for key, value := range input {
		output[key] = value
	}
	return output
}
