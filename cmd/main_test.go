package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersionRoute(t *testing.T) {
	// Create a new test server using our mux
	mux := http.NewServeMux()
	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]string{"version": APP_VERSION})
	})

	// Create a test request
	req := httptest.NewRequest(http.MethodGet, "/version", nil)
	w := httptest.NewRecorder()

	// Serve the request
	mux.ServeHTTP(w, req)

	// Check status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check content type
	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type %q, got %q", "application/json", contentType)
	}

	// Parse response body
	var response map[string]string
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	// Check version
	if version, exists := response["version"]; !exists {
		t.Error("response missing version field")
	} else if version != APP_VERSION {
		t.Errorf("expected version %q, got %q", APP_VERSION, version)
	}
}

func TestHealthRoute(t *testing.T) {
	// Create a new test server using our mux
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// Create a test request
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	// Serve the request
	mux.ServeHTTP(w, req)
}

func TestTemperatureRoute(t *testing.T) {
	// Create a new test server using our mux
	mux := http.NewServeMux()
	mux.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]string{"temperature": "20"})
	})

	// Create a test request
	req := httptest.NewRequest(http.MethodGet, "/temperature", nil)
	w := httptest.NewRecorder()

	// Serve the request
	mux.ServeHTTP(w, req)

	// Check status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check content type
	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type %q, got %q", "application/json", contentType)
	}

	// Parse response body
	var response map[string]string
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	// Check temperature
	if temperature, exists := response["temperature"]; !exists {
		t.Error("response missing temperature field")
	} else if temperature != "20" {
		t.Errorf("expected temperature %q, got %q", "20", temperature)
	}
}
