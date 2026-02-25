package arcgis

import (
	"math"
)

func ExtentSize(e Envelope) Point {
	return Point{
		X: e.XMax - e.XMin,
		Y: e.YMax - e.YMin,
	}
}

const (
	TileSize    = 256       // Standard tile size in pixels
	EarthRadius = 6378137.0 // Earth's radius in meters (Web Mercator)
	OriginShift = math.Pi * EarthRadius
)

// Calculate number of tiles needed to cover the bounds at a given zoom level
func CalculateTileCount(bounds Envelope, zoomLevel int) (tilesX, tilesY, totalTiles int) {
	// Get tile coordinates for the bounds
	minTileX, minTileY := MetersToTile(bounds.XMin, bounds.YMax, zoomLevel)
	maxTileX, maxTileY := MetersToTile(bounds.XMax, bounds.YMin, zoomLevel)

	// Calculate number of tiles in each direction
	tilesX = maxTileX - minTileX + 1
	tilesY = maxTileY - minTileY + 1
	totalTiles = tilesX * tilesY

	return tilesX, tilesY, totalTiles
}

// Convert Web Mercator meters to tile coordinates
func MetersToTile(mx, my float64, zoom int) (tx, ty int) {
	// Convert meters to pixels
	px, py := MetersToPixels(mx, my, zoom)

	// Convert pixels to tile coordinates
	tx = int(math.Floor(px / float64(TileSize)))
	ty = int(math.Floor(py / float64(TileSize)))

	return tx, ty
}

// Convert Web Mercator meters to pixel coordinates
func MetersToPixels(mx, my float64, zoom int) (px, py float64) {
	res := Resolution(zoom)
	px = (mx + OriginShift) / res
	py = (OriginShift - my) / res
	return px, py
}

// Calculate resolution (meters per pixel) at given zoom level
func Resolution(zoom int) float64 {
	return (2 * OriginShift) / (float64(TileSize) * math.Pow(2, float64(zoom)))
}

// If you have scale instead of zoom, convert it
func ScaleToZoom(scale float64, dpi float64) int {
	// Standard conversion: scale = 1 : X means X units per map unit
	// Assuming 96 DPI (common for web)
	if dpi == 0 {
		dpi = 96
	}

	metersPerPixel := scale * 0.0254 / dpi
	zoom := math.Log2((2 * OriginShift) / (metersPerPixel * float64(TileSize)))

	return int(math.Round(zoom))
}
