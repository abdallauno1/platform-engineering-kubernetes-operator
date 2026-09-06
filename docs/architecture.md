# Day 3 Architecture — Production Runtime and Kubernetes Packaging

```mermaid
flowchart TB
    Engineer[Platform Engineer] -->|kubectl apply -k config| API[Kubernetes API Server]
    API --> CRD[AIPlatform CRD]
    API --> Deployment[Controller Manager Deployment]

    subgraph Pod[Operator Pod]
      Runtime[Long-running Operator Runtime]
      Probe[Health Server :8081]
      Reconciler[AIPlatform Reconciler]
      Client[WorkloadClient Abstraction]
      Runtime --> Reconciler
      Runtime --> Probe
      Reconciler --> Client
    end

    Deployment --> Pod
    Kubelet[Kubelet] -->|GET /healthz| Probe
    Kubelet -->|GET /readyz| Probe
    RBAC[ServiceAccount + ClusterRole + Binding] --> Deployment
    NetPol[NetworkPolicy] --> Pod
```

## Production-readiness decisions

- Distroless, non-root image minimizes runtime surface area.
- Read-only root filesystem and dropped Linux capabilities reduce container privileges.
- Liveness and readiness endpoints are independent so Kubernetes can distinguish process health from controller readiness.
- Resource requests and limits make scheduling behavior explicit.
- Kustomize provides a single installation entry point without duplicating manifests.
- CI gates the container build on successful formatting, vetting, race-enabled tests and Go build.
