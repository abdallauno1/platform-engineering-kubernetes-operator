package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	platformv1 "github.com/abdallauno1/platform-engineering-kubernetes-operator/api/v1"
	"github.com/abdallauno1/platform-engineering-kubernetes-operator/internal/controller"
	"github.com/abdallauno1/platform-engineering-kubernetes-operator/internal/health"
)

func main() {
	mode := flag.String("mode", "operator", "runtime mode: operator or demo")
	healthAddr := flag.String("health-addr", ":8081", "health probe listen address")
	metricsAddr := flag.String("metrics-addr", ":8080", "metrics listen address")
	reconcileInterval := flag.Duration("reconcile-interval", 30*time.Second, "reconciliation interval in operator mode")
	flag.Parse()

	switch *mode {
	case "demo":
		runDemo()
	case "operator":
		runOperator(*healthAddr, *metricsAddr, *reconcileInterval)
	default:
		log.Fatalf("unsupported mode %q", *mode)
	}
}

func runOperator(healthAddr, metricsAddr string, reconcileInterval time.Duration) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	probeServer := health.NewServer(healthAddr)
	go health.LogRunError(probeServer.Run(ctx))

	metrics := controller.NewMetrics()
	metricsServer := &http.Server{Addr: metricsAddr, Handler: metrics.Handler(), ReadHeaderTimeout: 5 * time.Second}
	go func() {
		if err := metricsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("metrics server failed: %v", err)
			stop()
		}
	}()
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = metricsServer.Shutdown(shutdownCtx)
	}()

	client := controller.NewMemoryWorkloadClient()
	reconciler := controller.NewAIPlatformReconciler(client)
	resource := platformv1.NewDefault("demo-ai-platform", "default")
	resource.Spec.Image = "ghcr.io/example/ai-platform:v1.0.0"
	resource.Spec.Replicas = 2

	if err := reconcileOnce(ctx, reconciler, metrics, resource); err != nil {
		log.Printf("initial reconciliation failed: %v", err)
	} else {
		probeServer.SetReady(true)
	}

	ticker := time.NewTicker(reconcileInterval)
	defer ticker.Stop()

	log.Printf("operator runtime started healthAddr=%s metricsAddr=%s reconcileInterval=%s", healthAddr, metricsAddr, reconcileInterval)
	for {
		select {
		case <-ctx.Done():
			probeServer.SetReady(false)
			log.Printf("shutdown signal received")
			return
		case <-ticker.C:
			if err := reconcileOnce(ctx, reconciler, metrics, resource); err != nil {
				probeServer.SetReady(false)
				log.Printf("reconciliation failed: %v", err)
				continue
			}
			probeServer.SetReady(true)
		}
	}
}

func runDemo() {
	ctx := context.Background()
	client := controller.NewMemoryWorkloadClient()
	reconciler := controller.NewAIPlatformReconciler(client)
	resource := platformv1.NewDefault("demo-ai-platform", "default")
	resource.Spec.Image = "ghcr.io/example/ai-platform:v1.0.0"
	resource.Spec.Replicas = 2

	for iteration := 1; iteration <= 2; iteration++ {
		result, err := reconciler.Reconcile(ctx, resource)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("iteration=%d action=%s phase=%s message=%q readyReplicas=%d\n", iteration, result.Action, result.Status.Phase, result.Message, result.Status.ReadyReplicas)
	}
	resource.Metadata.Generation++
	resource.Spec.Replicas = 3
	result, err := reconciler.Reconcile(ctx, resource)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("iteration=3 action=%s phase=%s message=%q readyReplicas=%d\n", result.Action, result.Status.Phase, result.Message, result.Status.ReadyReplicas)
}

func reconcileOnce(ctx context.Context, reconciler *controller.AIPlatformReconciler, metrics *controller.Metrics, resource platformv1.AIPlatform) error {
	result, err := reconciler.Reconcile(ctx, resource)
	metrics.Observe(result.Action, err)
	if err != nil {
		return err
	}
	log.Printf("reconcile action=%s phase=%s readyReplicas=%d message=%q",
		result.Action, result.Status.Phase, result.Status.ReadyReplicas, result.Message)
	return nil
}
