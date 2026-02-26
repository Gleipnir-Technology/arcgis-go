package arcgis

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	SpatialReference string
	X                float64 `json:"x"`
	Y                float64 `json:"y"`
}

func (p Point) asJSON() (string, error) {
	result, err := json.Marshal(p)
	if err != nil {
		return "", fmt.Errorf("asJSON: %w", err)
	}
	return string(result), nil
}
