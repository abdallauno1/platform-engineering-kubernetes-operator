# LinkedIn Post — Day 1

Today I started Project 1 of my AI Platform Engineer roadmap: **Kubernetes Operator in Go**.

Day 1 was all about the foundation:

- Designed the first `AIPlatform` Custom Resource.
- Added a Kubernetes CRD manifest.
- Built a small reconciliation loop in Go.
- Added unit tests for validation and desired-state calculation.
- Prepared Docker support and GitHub Actions CI from the first commit.
- Documented the architecture and reconciliation sequence using Mermaid diagrams.

The main idea: a Platform Engineer should not only deploy workloads, but also build APIs that let teams describe intent while automation handles the operational details.

Tomorrow I will evolve the controller logic with a more realistic reconciliation flow, status transitions and stronger architecture boundaries.

#Kubernetes #Golang #PlatformEngineering #DevOps #CloudNative #KubernetesOperator
