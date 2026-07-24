# Day 2 Notes — Architecture Evolution

## Goal

Evolve the Day 1 desired-state calculator into a realistic, idempotent reconciliation flow.

## What changed

- Added a `WorkloadClient` interface to separate controller logic from infrastructure access.
- Added an in-memory implementation for deterministic unit tests and local demonstration.
- Implemented create, update, and no-op reconciliation decisions.
- Added drift detection between desired and current workload state.
- Added status phases: `Pending`, `Ready`, and `Failed`.
- Added `observedGeneration` and `readyReplicas` to status.
- Expanded CRD schema and printer columns.
- Added tests for creation, update, idempotency, and validation failure.
- Added a real `.github/workflows/ci.yml` pipeline.

## Run

```bash
make ci
make run
```

Expected demonstration:

```text
iteration=1 action=Created phase=Ready ...
iteration=2 action=Unchanged phase=Ready ...
iteration=3 action=Updated phase=Ready ...
```

## Why idempotency matters

A Kubernetes controller may reconcile the same resource many times. A correct controller must converge current state toward desired state without producing unnecessary updates each time.
