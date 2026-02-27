package response

import (
	//"encoding/json"
	"fmt"
	"slices"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/rs/zerolog/log"
	"github.com/twpayne/go-proj/v11"
)

// type P2D = [2]float64
// type Ring2D = []P2D
type Polygon2D struct {
	//CurveRings []Multipoint2D `json:"curveRings,omitempty"`
	HasM bool    `json:"hasM"`
	HasZ bool    `json:"hasZ"`
	IDs  [][]int `json:"ids"`
	//Rings            []Ring2D         `json:"rings,omitempty"`
	Rings            []orb.Ring       `json:"rings,omitempty"`
	SpatialReference SpatialReference `json:"spatialReference,omitempty"`
}

func (p Polygon2D) Project(source string, target string) (Geometry, error) {
	pj, err := proj.NewCRSToCRS(source, target, nil)
	if err != nil {
		return GeometryNull{}, fmt.Errorf("project: %w", err)
	}
	rings := make([]orb.Ring, len(p.Rings))
	for i, ring := range p.Rings {
		new_ring := make(orb.Ring, len(ring))
		for j, point := range ring {
			coord := proj.NewCoord(point.X(), point.Y(), 0, 0)
			result, err := pj.Forward(coord)
			if err != nil {
				return GeometryNull{}, fmt.Errorf("forward: %w", err)
			}
			new_ring[j] = orb.Point{result.Y(), result.X()}
		}
		rings[i] = new_ring
	}
	return Polygon2D{
		HasM:             p.HasM,
		HasZ:             p.HasZ,
		IDs:              p.IDs,
		Rings:            rings,
		SpatialReference: SpatialReference{},
	}, nil
}
func (p Polygon2D) String() string { return fmt.Sprintf("poly with %d rings", len(p.Rings)) }
func (p Polygon2D) Type() string   { return "esriGeometryPolygon" }
func (p Polygon2D) ToGeoJSON() (string, error) {
	fc := geojson.NewFeatureCollection()
	rings := make([]orb.Ring, len(p.Rings))
	for i, ring := range p.Rings {
		s := slices.Clone(ring)
		slices.Reverse(s)
		rings[i] = s
	}
	f := geojson.NewFeature(orb.MultiPolygon{rings})
	f.Properties = make(geojson.Properties, 1)
	f.Properties["fake"] = "foo"
	fc.Append(f)
	rawJSON, err := fc.MarshalJSON()
	if err != nil {
		return "", fmt.Errorf("togeojson: %w", err)
	}
	log.Debug().Bytes("geojson", rawJSON).Send()
	return string(rawJSON), nil
}
