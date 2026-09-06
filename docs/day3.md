# Day 3 — Production Readiness

Day 3 moves the project from a testable reconciliation engine to a deployable Kubernetes package.

## Added

- Long-running operator runtime with graceful SIGTERM/SIGINT shutdown.
- Kubernetes-compatible `/healthz` and `/readyz` endpoints.
- Unit tests for liveness and readiness behavior.
- Namespace, ServiceAccount, ClusterRoleBinding and Deployment manifests.
- Liveness/readiness probes, resource requests/limits and hardened container security context.
- Kustomize entry point for one-command installation.
- NetworkPolicy baseline.
- Multi-stage distroless container image running as non-root.
- GitHub Actions container-build job after Go CI succeeds.
- Local verification script and production-oriented Make targets.

## Runtime lifecycle

1. Process starts and health server responds on `/healthz`.
2. First reconciliation runs.
3. Readiness changes to healthy after successful reconciliation.
4. Periodic reconciliation keeps desired state converged.
5. SIGTERM/SIGINT marks the process unready and triggers graceful shutdown.

## Important scope note

The reconciliation engine still uses the `WorkloadClient` abstraction with an in-memory implementation. The Kubernetes manifests package the controller process as a real Kubernetes workload, while a concrete Kubernetes API-backed client is intentionally left for the final advanced integration/polish step. This keeps the architecture honest and the Day 2 unit-test boundary intact.
