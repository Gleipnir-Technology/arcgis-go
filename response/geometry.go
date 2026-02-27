package response

import (
	"errors"
)

type Geometry interface {
	Project(string, string) (Geometry, error)
	String() string
	ToGeoJSON() (string, error)
	Type() string
}
type GeometryNull struct{}

func (g GeometryNull) Project(string, string) (Geometry, error) {
	return GeometryNull{}, errors.New("nonsense ")
}
func (g GeometryNull) String() string             { return "" }
func (g GeometryNull) ToGeoJSON() (string, error) { return "", errors.New("nonsense") }
func (g GeometryNull) Type() string               { return "nil" }
