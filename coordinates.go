package arcgis

import "math"

const R = 6378137.0

type Position struct {
	Latitude  float64
	Longitude float64
}

func Radians(deg float64) float64 {
	return deg * math.Pi / 180
}

func Degrees(rad float64) float64 {
	return rad * 180 / math.Pi
}

func Y2lat(y float64) float64 {
	return Degrees(2*math.Atan(math.Exp(y/R)) - math.Pi/2)
}

func Lat2y(lat float64) float64 {
	return R * math.Log(math.Tan(math.Pi/4+Radians(lat)/2))
}

func X2lat(x float64) float64 {
	return Degrees(x / R)
}

func Lon2x(lon float64) float64 {
	return R * Radians(lon)
}

func GeometryToPosition(g Geometry) Position {
	var result Position
	result.Latitude = X2lat(g.X)
	result.Longitude = Y2lat(g.Y)
	return result
}
