package main

import (
	"context"
	"fmt"
	"log"

	platformv1 "github.com/abdallauno1/platform-engineering-kubernetes-operator/api/v1"
	"github.com/abdallauno1/platform-engineering-kubernetes-operator/internal/controller"
)

func main() {
	ctx := context.Background()
	client := controller.NewMemoryWorkloadClient()
	reconciler := controller.NewAIPlatformReconciler(client)
	resource := platformv1.NewDefault("demo-ai-platform", "default")
	resource.Spec.Image = "ghcr.io/example/ai-platform:day2"
	resource.Spec.Replicas = 2

	for iteration := 1; iteration <= 2; iteration++ {
		result, err := reconciler.Reconcile(ctx, resource)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("iteration=%d action=%s phase=%s message=%q readyReplicas=%d\n",
			iteration, result.Action, result.Status.Phase, result.Message, result.Status.ReadyReplicas)
	}

	resource.Metadata.Generation++
	resource.Spec.Replicas = 3
	result, err := reconciler.Reconcile(ctx, resource)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("iteration=3 action=%s phase=%s message=%q readyReplicas=%d\n",
		result.Action, result.Status.Phase, result.Message, result.Status.ReadyReplicas)
}
