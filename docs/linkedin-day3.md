# LinkedIn — Day 3

Day 3 of my Kubernetes Operator project focused on production readiness rather than adding another feature.

Today I moved the controller from a local reconciliation demo into a Kubernetes-deployable runtime:

- added liveness and readiness endpoints (`/healthz`, `/readyz`)
- implemented graceful shutdown for SIGTERM/SIGINT
- added ServiceAccount, RBAC binding and controller Deployment manifests
- hardened the pod with non-root execution, read-only filesystem and dropped capabilities
- defined CPU/memory requests and limits
- added a NetworkPolicy baseline and a Kustomize installation entry point
- upgraded GitHub Actions so the container image is built only after Go CI passes
- added unit tests for health/readiness behavior

One lesson I keep seeing in Platform Engineering: “it works locally” is only the beginning. A controller also needs lifecycle management, probes, predictable resources, security boundaries and a repeatable delivery path.

Next: Day 4 — advanced integration, metrics, release automation and final portfolio polish.

#Kubernetes #Golang #PlatformEngineering #DevOps #SRE #CloudNative #GitHubActions
