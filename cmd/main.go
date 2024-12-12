package main

import (
	"encoding/json"
	"net/http"
	"log"
)

// APP_VERSION is the version of the app
const APP_VERSION = "v0.0.1"

// WriteJSON is a helper function to write a JSON response
// It sets the Content-Type header to application/json and writes the response with the given status code
// It returns an error if the encoding fails
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()
	// Handle the root route
	// This is a simple health check endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]string{"version": APP_VERSION})
	})

	// Handle the health route
	// This is a simple health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// Create a new logger
	log.Println("Starting server on port 8080")
	// Start the server
	http.ListenAndServe(":8080", mux)
}
