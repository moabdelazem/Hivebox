package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// APP_VERSION is the version of the app
const APP_VERSION = "v0.0.2"
const OPEN_SENSE_API_ID = "5eba5fbad46fb8001b799786"
const API_URL = "https://api.opensensemap.org/boxes/%s?format=json"

// WriteJSON is a helper function to write a JSON response
// It sets the Content-Type header to application/json and writes the response with the given status code
// It returns an error if the encoding fails
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func GetSenseBoxData(apiId string, from time.Time) (map[string]interface{}, error) {
	// Create API Call to OpenSense API
	// https://api.opensensemap.org/boxes/:senseBoxId?format=:format
	// Get the data from the API

	// Create the API URL
	apiUrl := fmt.Sprintf(API_URL, apiId)

	// Make the API call
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// Read and parse the JSON response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, nil
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Handle the version route
	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]string{"version": APP_VERSION})
	})

	// Handle the health route
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// Handle the temperature route
	// This route will return the current average temperature based on all senseBox data.
	// The data is no older than 1 hour.
	mux.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		// Get the current time
		now := time.Now()

		// Get the current time minus 1 hour
		oneHourAgo := now.Add(-1 * time.Hour)

		// Get the data from the senseBox
		data, err := GetSenseBoxData(OPEN_SENSE_API_ID, oneHourAgo)
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to get data from senseBox"})
			return
		}

		// Extract sensors array from data
		sensors, ok := data["sensors"].([]interface{})
		if !ok {
			WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Invalid sensor data format"})
			return
		}

		// Find temperature sensor and get latest measurement
		var temperature float64
		var count int
		for _, sensor := range sensors {
			sensorMap := sensor.(map[string]interface{})
			if sensorMap["title"] == "Temperatur" {
				if lastMeasurement, ok := sensorMap["lastMeasurement"].(map[string]interface{}); ok {
					if value, ok := lastMeasurement["value"].(string); ok {
						if val, err := strconv.ParseFloat(value, 64); err == nil {
							temperature += val
							count++
						}
					}
				}
			}
		}

		// Check if there is any temperature data
		if count == 0 {
			WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "No temperature data found"})
			return
		}

		// Calculate average temperature
		avgTemperature := temperature / float64(count)

		// Write the average temperature to the response
		WriteJSON(w, http.StatusOK, map[string]interface{}{
			"average_temperature": avgTemperature,
			"unit":                "°C",
			"sensor_count":        count,
		})
	})

	// Create a new logger
	log.Println("Starting server on port 8080")
	// Start the server
	http.ListenAndServe(":8080", mux)
}
