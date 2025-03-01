package ml_service

import (
	"encoding/json"
	"net/http"
)

// Go HTTP handler calling Python ML service
func predictHandler(w http.ResponseWriter, r *http.Request) {
	// Preprocess input in Go
	processed := preprocess(r.Body)

	// Call Python ML service
	resp, _ := http.Post(
		"http://ml-service:8000/predict",
		"application/json",
		processed,
	)

	// Post-process in Go
	result := postprocess(resp.Body)
	json.NewEncoder(w).Encode(result)
}
