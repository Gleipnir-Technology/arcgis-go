package response

import (
	"errors"
)

type Polyline struct {
	Paths [][][]float64 `json:"paths"`
}

func (p Polyline) Project(source string, target string) (Geometry, error) {
	return GeometryNull{}, errors.New("not implemented")
}
func (p Polyline) String() string { return "some polyline" }
func (p Polyline) Type() string   { return "LineString" }
func (p Polyline) ToGeoJSON() (string, error) {
	panic("not implemented")
}
