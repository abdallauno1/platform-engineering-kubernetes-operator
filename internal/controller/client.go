package controller

import (
	"context"
	"errors"
	"sync"
)

var ErrNotFound = errors.New("workload not found")

type WorkloadClient interface {
	Get(ctx context.Context, namespace, name string) (Workload, error)
	Create(ctx context.Context, workload Workload) error
	Update(ctx context.Context, workload Workload) error
}

type MemoryWorkloadClient struct {
	mu        sync.RWMutex
	workloads map[string]Workload
}

func NewMemoryWorkloadClient() *MemoryWorkloadClient {
	return &MemoryWorkloadClient{workloads: make(map[string]Workload)}
}

func (c *MemoryWorkloadClient) Get(ctx context.Context, namespace, name string) (Workload, error) {
	if err := ctx.Err(); err != nil {
		return Workload{}, err
	}
	c.mu.RLock()
	defer c.mu.RUnlock()
	workload, ok := c.workloads[namespace+"/"+name]
	if !ok {
		return Workload{}, ErrNotFound
	}
	return workload, nil
}

func (c *MemoryWorkloadClient) Create(ctx context.Context, workload Workload) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.workloads[workload.Key()] = workload
	return nil
}

func (c *MemoryWorkloadClient) Update(ctx context.Context, workload Workload) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.workloads[workload.Key()]; !ok {
		return ErrNotFound
	}
	c.workloads[workload.Key()] = workload
	return nil
}
