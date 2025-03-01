// Go: main.go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type PredictionRequest struct {
	Features []float64 `json:"features"`
}

func main() {
	http.HandleFunc("/predict", func(w http.ResponseWriter, r *http.Request) {
		var req PredictionRequest
		json.NewDecoder(r.Body).Decode(&req)

		// Call Python service
		jsonData, err := json.Marshal(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := http.Post("http://localhost:5000/invoke", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8080", nil)
}
