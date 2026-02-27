package response

import (
	"fmt"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/twpayne/go-proj/v11"
)

type Point struct {
	X                float64          `json:"x"`
	Y                float64          `json:"y"`
	Z                float64          `json:"z,omitempty"`
	M                float64          `json:"m,omitempty"`
	SpatialReference SpatialReference `json:"spatialReference,omitempty"`
}

func (p Point) Project(source string, target string) (Geometry, error) {
	pj, err := proj.NewCRSToCRS(
		"EPSG:2228", // Source: California Zone 3
		"EPSG:4326", // Target: WGS84
		nil,
	)
	if err != nil {
		return GeometryNull{}, fmt.Errorf("project: %w", err)
	}

	// Transform coordinates
	coord := proj.NewCoord(p.X, p.Y, 0, 0)
	result, err := pj.Forward(coord)
	if err != nil {
		return GeometryNull{}, fmt.Errorf("forward: %w", err)
	}
	// result.X is longitude, result.Y is latitude
	return Point{X: result.X(), Y: result.Y()}, nil

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
