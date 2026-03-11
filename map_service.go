package arcgis

import (
	"context"
	"math"
	"net/url"
	"strconv"

	"github.com/Gleipnir-Technology/arcgis-go/response"
	"github.com/rs/zerolog/log"
)

type MapService struct {
	ID       string
	Metadata *response.MapServiceMetadata
	Name     string
	Title    string
	URL      url.URL

	requestor *gisRequestor
}

func (ms MapService) PopulateMetadata(ctx context.Context) (*response.MapServiceMetadata, error) {
	if ms.Metadata != nil {
		return ms.Metadata, nil
	}
	return reqGetJSONFullURL[response.MapServiceMetadata](ctx, *ms.requestor, ms.URL)
}
func (ms MapService) Tile(ctx context.Context, level, row, column uint) ([]byte, *ErrorWithStatus) {
	// From https://developers.arcgis.com/documentation/portal-and-data-services/data-services/map-tile-services/introduction/
	// GET https://{host}/{organizationId}/arcgis/rest/services/{serviceName}/MapServer/tile/{z}/{y}/{x}
	log.Info().Str("url", ms.URL.String()).Uint("lvl", level).Uint("row", row).Uint("col", column).Msg("creating tile URL")
	u := ms.URL.JoinPath("tile", strconv.Itoa(int(level)), strconv.Itoa(int(row)), strconv.Itoa(int(column)))
	result, err := reqGetParamsHeadersFullURL(ctx, *ms.requestor, *u, map[string]string{}, map[string]string{})
	if err != nil {
		if err.Status == 404 {
			log.Info().Msg("404, assuming empty tile data")
			return []byte{}, nil
		}
		return nil, err
	}
	return result, nil
}
func (ms MapService) TileGPS(ctx context.Context, level uint, lat, lng float64) ([]byte, error) {
	row, col := LatLngToTile(level, lat, lng)
	log.Info().Float64("lat", lat).Float64("lng", lng).Uint("row", row).Uint("col", col).Msg("GPS to tile")
	return ms.Tile(ctx, level, row, col)
}

// LatLngToTile converts GPS coordinates to ArcGIS tile coordinates
func LatLngToTile(level uint, lat, lng float64) (row, column uint) {
	// Get number of tiles per dimension at this zoom level
	numTiles := math.Pow(2, float64(level))

	// Convert longitude to tile column
	// Range: -180 to 180 degrees maps to 0 to numTiles
	column = uint(math.Floor((lng + 180.0) / 360.0 * numTiles))

	// Convert latitude to tile row using Mercator projection
	// First convert lat to radians
	latRad := lat * math.Pi / 180.0

	// Apply Mercator projection formula
	// This maps latitude from -85.0511 to 85.0511 degrees to 0 to numTiles
	mercatorY := 0.5 - math.Log(math.Tan(latRad)+1/math.Cos(latRad))/(2*math.Pi)
	row = uint(math.Floor(mercatorY * numTiles))

	// Ensure values are within valid range
	if column < 0 {
		column = 0
	} else if column >= uint(numTiles) {
		column = uint(numTiles) - 1
	}

	if row < 0 {
		row = 0
	} else if row >= uint(numTiles) {
		row = uint(numTiles) - 1
	}

	return row, column
}
