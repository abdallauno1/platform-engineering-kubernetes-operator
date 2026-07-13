# Day 1 Notes — Foundation

## Goal

Create the foundation of a Kubernetes Operator in Go that introduces a custom platform API and a testable reconciliation loop.

## What was built

- Go module structure for an operator project.
- `AIPlatform` custom resource model.
- Kubernetes CRD manifest.
- RBAC scaffold for future in-cluster deployment.
- Reconciler that converts custom resource intent into desired workload state.
- Unit tests for validation and reconciliation.
- GitHub Actions workflow for CI.
- Dockerfile for building the operator binary.

## Run locally

```bash
make test
make build
make run
```

## Apply CRD sample later

```bash
kubectl apply -f config/crd/platform.mady.dev_aiplatforms.yaml
kubectl apply -f config/samples/platform_v1_aiplatform.yaml
```

## Next

Day 2 will evolve the architecture toward a realistic controller: client abstraction, create/update decisions, status phases and stronger reconciliation tests.
