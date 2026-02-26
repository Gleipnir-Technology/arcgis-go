package response

import (
	"fmt"
)

type Point struct {
	X                float64          `json:"x"`
	Y                float64          `json:"y"`
	Z                float64          `json:"z,omitempty"`
	M                float64          `json:"m,omitempty"`
	SpatialReference SpatialReference `json:"spatialReference,omitempty"`
}

func (p Point) String() string { return fmt.Sprintf("%f,%f", p.X, p.Y) }
func (p Point) Type() string   { return "esriGeometryPoint" }
func (p Point) ToGeoJSON() map[string]any {
	return map[string]any{
		"type":        "Point",
		"coordinates": []float64{p.X, p.Y},
	}
}
