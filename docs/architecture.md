# Day 2 Architecture — Reconciliation with State Management

```mermaid
flowchart LR
    Engineer[Platform Engineer] -->|applies AIPlatform| API[Kubernetes API Server]
    API --> Controller[AIPlatform Controller]
    Controller --> Validate[Validate Custom Resource]
    Validate --> Desired[Build Desired Workload]
    Desired --> Client[Workload Client Abstraction]
    Client -->|not found| Create[Create Workload]
    Client -->|drift detected| Update[Update Workload]
    Client -->|already equal| NoOp[No-op / Idempotent]
    Create --> Status[Status: Ready]
    Update --> Status
    NoOp --> Status
    Validate -->|invalid| Failed[Status: Failed]
```

## Architecture decisions

- `WorkloadClient` isolates reconciliation logic from a concrete Kubernetes client.
- `MemoryWorkloadClient` makes create/update/idempotency behavior testable without a cluster.
- Desired and current state are compared before mutation.
- Reconciliation returns explicit actions: `Created`, `Updated`, `Unchanged`, or `Failed`.
- Status records phase, message, observed generation, and ready replicas.

Day 3 will replace the demonstration runtime with installable operator deployment manifests and production-oriented container/Kubernetes configuration.
