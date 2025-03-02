package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/nickemma/mlops-starter/go/processors"
)

type PredictionRequest struct {
	Features []float64 `json:"features"`
}

type PredictionResponse struct {
	Prediction   int       `json:"prediction"`
	Features     []float64 `json:"features"`
	ModelVersion string    `json:"model_version"`
}

func predictHandler(w http.ResponseWriter, r *http.Request) {
	var req PredictionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Validate and preprocess
	if err := processors.ValidateInput(req.Features); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	normalized := processors.NormalizeFeatures(req.Features)

	// Call Python service (TODO: Implement actual call)
	response := PredictionResponse{
		Prediction:   1, // Mock value
		Features:     normalized,
		ModelVersion: "0.1.0",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/predict", predictHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting MLOps API on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
