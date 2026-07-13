package main

import (
	"context"
	"fmt"
	"log"

	platformv1 "github.com/abdallamady/kubernetes-operator-go/api/v1"
	"github.com/abdallamady/kubernetes-operator-go/internal/controller"
)

func main() {
	resource := platformv1.NewDefault("demo-ai-platform", "default")
	reconciler := controller.NewAIPlatformReconciler()

	result, desired, err := reconciler.Reconcile(context.Background(), resource)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result=%s deployment=%s image=%s replicas=%d\n", result.Message, desired.Name, desired.Image, desired.Replicas)
}
