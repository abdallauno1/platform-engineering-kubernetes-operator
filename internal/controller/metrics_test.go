package controller

import (
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMetricsHandler(t *testing.T) {
	m := NewMetrics()
	m.Observe(ActionCreated, nil)
	m.Observe(ActionFailed, errors.New("boom"))
	r := httptest.NewRecorder()
	m.Handler().ServeHTTP(r, httptest.NewRequest("GET", "/metrics", nil))
	body := r.Body.String()
	for _, want := range []string{`aiplatform_reconcile_total{action="Created"} 1`, `aiplatform_reconcile_total{action="Failed"} 1`, `aiplatform_reconcile_errors_total 1`} {
		if !strings.Contains(body, want) {
			t.Fatalf("missing %q in %s", want, body)
		}
	}
}
