package controller

import (
	"fmt"
	"net/http"
	"sync"
)

type Metrics struct {
	mu     sync.Mutex
	total  map[Action]uint64
	errors uint64
}

func NewMetrics() *Metrics { return &Metrics{total: make(map[Action]uint64)} }
func (m *Metrics) Observe(action Action, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.total[action]++
	if err != nil {
		m.errors++
	}
}
func (m *Metrics) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		m.mu.Lock()
		defer m.mu.Unlock()
		w.Header().Set("Content-Type", "text/plain; version=0.0.4")
		fmt.Fprintln(w, "# HELP aiplatform_reconcile_total Total reconciliations by action.")
		fmt.Fprintln(w, "# TYPE aiplatform_reconcile_total counter")
		for _, a := range []Action{ActionCreated, ActionUpdated, ActionUnchanged, ActionFailed} {
			fmt.Fprintf(w, "aiplatform_reconcile_total{action=%q} %d\n", a, m.total[a])
		}
		fmt.Fprintln(w, "# HELP aiplatform_reconcile_errors_total Total reconciliation errors.")
		fmt.Fprintln(w, "# TYPE aiplatform_reconcile_errors_total counter")
		fmt.Fprintf(w, "aiplatform_reconcile_errors_total %d\n", m.errors)
	})
}
