package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_GetRequestHandler_ReturnsBadRequestForInvalidInput(t *testing.T) {
	// todo research if data provider or similar exists
	req := httptest.NewRequest(http.MethodGet, "/asd", nil)
	rr := httptest.NewRecorder()

	GetRequestHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rr.Code)
	}

	if strings.TrimSpace(rr.Body.String()) != "invalid id" {
		t.Errorf("expected 'invalid id', got '%s'", rr.Body.String())
	}
}

func Test_GetRequestHandler_ReturnsNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/eb43b895f40fbc0f0bdda29d3d52e58a53e2b4b8", nil)
	rr := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{hashedUrl}", GetRequestHandler)

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", rr.Code)
	}
}

func Test_GetRequestHandler_ReturnsRedirect(t *testing.T) {
	urlsMap = map[string]string{"eb43b895f40fbc0f0bdda29d3d52e58a53e2b4b8": "http://www.google.com"}
	req := httptest.NewRequest(http.MethodGet, "/eb43b895f40fbc0f0bdda29d3d52e58a53e2b4b8", nil)
	rr := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{hashedUrl}", GetRequestHandler)

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusMovedPermanently {
		t.Errorf("expected 301, got %d", rr.Code)
	}
}
