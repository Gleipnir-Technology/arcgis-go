package arcgis

import (
	"testing"
)

func TestMercatorToGPS(t *testing.T) {
	g := Geometry{-1.32815821632e+07, 4.3482645502e+06}
	p := GeometryToPosition(g)

	if p.Latitude != -119.31048254491778 {
		t.Error("Bad Latitude", p.Latitude)
	}
	if p.Longitude != 36.34548680219802 {
		t.Error("Bad Longitude", p.Longitude)
	}
}
