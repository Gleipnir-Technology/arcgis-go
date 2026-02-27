package response

import (
	//"encoding/json"
	"fmt"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/rs/zerolog/log"
)

type P2D = [2]float64
type Ring2D = []P2D
type Polygon2D struct {
	//CurveRings []Multipoint2D `json:"curveRings,omitempty"`
	HasZ bool    `json:"hasZ"`
	HasM bool    `json:"hasM"`
	IDs  [][]int `json:"ids"`
	//Rings            []Ring2D         `json:"rings,omitempty"`
	Rings            []orb.Ring       `json:"rings,omitempty"`
	SpatialReference SpatialReference `json:"spatialReference,omitempty"`
}

func (p Polygon2D) String() string { return fmt.Sprintf("poly with %d rings", len(p.Rings)) }
func (p Polygon2D) Type() string   { return "esriGeometryPolygon" }
func (p Polygon2D) ToGeoJSON() (string, error) {
	fc := geojson.NewFeatureCollection()
	fc.Append(geojson.NewFeature(orb.MultiPolygon{p.Rings}))
	rawJSON, err := fc.MarshalJSON()
	if err != nil {
		return "", fmt.Errorf("togeojson: %w", err)
	}
	log.Debug().Bytes("geojson", rawJSON).Send()
	return string(rawJSON), nil
}
