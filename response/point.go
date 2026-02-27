package response

import (
	"fmt"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
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
func (p Point) ToGeoJSON() (string, error) {
	fc := geojson.NewFeatureCollection()
	fc.Append(geojson.NewFeature(orb.Point{
		p.X,
		p.Y,
	}))
	rawJSON, err := fc.MarshalJSON()
	if err != nil {
		return "", fmt.Errorf("togeojson: %w", err)
	}
	return string(rawJSON), nil
}
