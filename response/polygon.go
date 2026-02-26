package response

import (
	"fmt"
)

type P2D = [2]float64
type Ring2D = []P2D
type Polygon2D struct {
	//CurveRings []Multipoint2D `json:"curveRings,omitempty"`
	HasZ             bool             `json:"hasZ"`
	HasM             bool             `json:"hasM"`
	IDs              [][]int          `json:"ids"`
	Rings            []Ring2D         `json:"rings,omitempty"`
	SpatialReference SpatialReference `json:"spatialReference,omitempty"`
}

func (p Polygon2D) String() string { return fmt.Sprintf("poly with %d rings", len(p.Rings)) }
func (p Polygon2D) Type() string   { return "esriGeometryPolygon" }
func (p Polygon2D) ToGeoJSON() map[string]any {
	return map[string]any{
		"type": "Polygon",
		//"coordinates": p.Paths,
	}
}
