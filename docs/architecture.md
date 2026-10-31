# Architecture

```mermaid
flowchart TD
  CR[AIPlatform CR] --> REC[Reconciler]
  REC --> DES[Desired Workload]
  DES --> CLIENT[WorkloadClient]
  CLIENT --> MEM[Memory adapter - shipped]
  CLIENT -. extension .-> K8S[Kubernetes API adapter]
  REC --> STATUS[Status result]
  STATUS --> METRICS[Metrics]
  STATUS --> EVENTS[Events]
  RUNTIME[Operator runtime] --> REC
  RUNTIME --> PROBES[Health / readiness]
```

The controller owns policy: validation, desired-state derivation, drift detection and result/status calculation. Infrastructure access is behind `WorkloadClient`, making policy independently testable.
