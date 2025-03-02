// Data preprocessing

package processors

import (
	"errors"
	"math"
)

// ValidateInput checks feature constraints
func ValidateInput(features []float64) error {
	if len(features) != 2 {
		return errors.New("input must contain exactly 2 features")
	}

	for _, val := range features {
		if math.IsNaN(val) {
			return errors.New("features contain NaN values")
		}
	}
	return nil
}

// NormalizeFeatures scales features to [0,1]
func NormalizeFeatures(features []float64) []float64 {
	minVal := 1.0
	maxVal := 3.0 // Based on training data stats

	normalized := make([]float64, len(features))
	for i, val := range features {
		normalized[i] = (val - minVal) / (maxVal - minVal)
	}
	return normalized
}
