package arcgis

import (
	"context"
	"fmt"
	"math"
	"net/url"
)

type MapService struct {
	ID    string
	Name  string
	Title string
	URL   url.URL

	meta *MapServiceMetadata
}

func (ms MapService) Metadata(ctx context.Context, ag *ArcGIS) (*MapServiceMetadata, error) {
	if ms.meta != nil {
		return ms.meta, nil
	}
	return reqGetJSONFullURL[MapServiceMetadata](ctx, ag.requestor, ms.URL)
}
func (ms MapService) Tile(ctx context.Context, ag *ArcGIS, level, row, column int) ([]byte, error) {
	// From https://developers.arcgis.com/documentation/portal-and-data-services/data-services/map-tile-services/introduction/
	// GET https://{host}/{organizationId}/arcgis/rest/services/{serviceName}/MapServer/tile/{z}/{y}/{x}
	// But the URL value we popluate above is already most of it.
	// It should look like
	// https://tiles.arcgis.com/tiles/{organizationId}/arcgis/rest/services/{serviceName}/MapServer
	u, err := url.Parse(fmt.Sprintf("%s/tile/%d/%d/%d", ms.URL, level, row, column))
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}
	return reqGetParamsHeadersFullURL(ctx, ag.requestor, *u, map[string]string{}, map[string]string{})
}
func (ms MapService) TileGPS(ctx context.Context, ag *ArcGIS, level int, lat, lng float64) ([]byte, error) {
	row, col := LatLngToTile(level, lat, lng)
	return ms.Tile(ctx, ag, level, row, col)
}

// LatLngToTile converts GPS coordinates to ArcGIS tile coordinates
func LatLngToTile(level int, lat, lng float64) (row, column int) {
	// Get number of tiles per dimension at this zoom level
	numTiles := math.Pow(2, float64(level))

	// Convert longitude to tile column
	// Range: -180 to 180 degrees maps to 0 to numTiles
	column = int(math.Floor((lng + 180.0) / 360.0 * numTiles))

	// Convert latitude to tile row using Mercator projection
	// First convert lat to radians
	latRad := lat * math.Pi / 180.0

	// Apply Mercator projection formula
	// This maps latitude from -85.0511 to 85.0511 degrees to 0 to numTiles
	mercatorY := 0.5 - math.Log(math.Tan(latRad)+1/math.Cos(latRad))/(2*math.Pi)
	row = int(math.Floor(mercatorY * numTiles))

	// Ensure values are within valid range
	if column < 0 {
		column = 0
	} else if column >= int(numTiles) {
		column = int(numTiles) - 1
	}

	if row < 0 {
		row = 0
	} else if row >= int(numTiles) {
		row = int(numTiles) - 1
	}

	return row, column
}
