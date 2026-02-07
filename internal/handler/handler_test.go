package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/reneroboter/urlshortener/internal/store"
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
	store.UrlsMap = sync.Map{}
	store.UrlsMap.Store("eb43b895f40fbc0f0bdda29d3d52e58a53e2b4b8", "http://www.google.com")
	req := httptest.NewRequest(http.MethodGet, "/eb43b895f40fbc0f0bdda29d3d52e58a53e2b4b8", nil)
	rr := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{hashedUrl}", GetRequestHandler)

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusMovedPermanently {
		t.Errorf("expected 301, got %d", rr.Code)
	}
}

func Test_PostRequestHandler_ReturnsId(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /shorten", PostRequestHandler)

	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer([]byte(`{"url":"http://www.google.com"}`)))
	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}

	if !strings.ContainsAny("738ddf35b3a85a7a6ba7b232bd3d5f1e4d284ad1", rr.Body.String()) {
		t.Errorf("expected '738ddf35b3a85a7a6ba7b232bd3d5f1e4d284ad1', got '%s'", rr.Body.String())
	}
}

func Test_PostRequestHandler_BadRequest(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /shorten", PostRequestHandler)

	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer([]byte(`{"url":"www.google.com"}`)))
	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rr.Code)
	}

	if strings.TrimSpace(rr.Body.String()) != "invalid URL format" {
		t.Errorf("expected 'invalid URL format', got '%s'", rr.Body.String())
	}
}

func Test_PostRequestHandler_ReturnBadRequestIfRecordExists(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /shorten", PostRequestHandler)

	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer([]byte(`{"url":"http://www.github.com"}`)))
	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}

	if !strings.ContainsAny("2041ee3c75afc15ce115e52bca5cfe48c7abbc96", rr.Body.String()) {
		t.Errorf("expected '2041ee3c75afc15ce115e52bca5cfe48c7abbc96', got '%s'", rr.Body.String())
	}

	req = httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer([]byte(`{"url":"http://www.github.com"}`)))
	rr = httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rr.Code)
	}

	if strings.TrimSpace(rr.Body.String()) != "url already exists" {
		t.Errorf("expected 'url already exists', got '%s'", rr.Body.String())
	}
}
