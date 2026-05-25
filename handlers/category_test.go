package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethodCheck_OK(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/categories/scenes",
		nil,
	)

	rr := httptest.NewRecorder()

	result := methodCheck(rr, req, http.MethodGet)

	if !result {
		t.Error("expected true, got false")
	}
}

func TestMethodCheck_NG(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/categories/scenes",
		nil,
	)

	rr := httptest.NewRecorder()

	result := methodCheck(rr, req, http.MethodGet)

	if result {
		t.Error("expected false, got true")
	}

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rr.Code)
	}
}

