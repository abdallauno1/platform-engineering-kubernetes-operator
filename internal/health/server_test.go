package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthzAlwaysHealthy(t *testing.T) {
	server := NewServer(":0")
	request := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	response := httptest.NewRecorder()

	server.healthz(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", response.Code)
	}
}

func TestReadyzReflectsReadiness(t *testing.T) {
	server := NewServer(":0")
	request := httptest.NewRequest(http.MethodGet, "/readyz", nil)

	response := httptest.NewRecorder()
	server.readyz(response, request)
	if response.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 before readiness, got %d", response.Code)
	}

	server.SetReady(true)
	response = httptest.NewRecorder()
	server.readyz(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("expected 200 after readiness, got %d", response.Code)
	}
}
