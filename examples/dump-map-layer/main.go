package main

import (
	"context"
	//"fmt"
	//"net/url"
	"os"

	"github.com/Gleipnir-Technology/arcgis-go"
	//"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var BaseURL, ClientID, ClientSecret, Environment, FieldseekerSchemaDirectory, MapboxToken string

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	ctx := context.TODO()
	ctx = log.With().Str("component", "arcgis").Logger().WithContext(ctx)

	gis, err := arcgis.NewArcGIS(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create arcgis")
		os.Exit(2)
	}

	map_services, err := gis.MapServices(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get map service")
		os.Exit(3)
	}

	if err != nil {
		log.Error().Err(err).Msg("Failed to get map services")
		os.Exit(4)
	}
	var map_service arcgis.MapService
	for _, ms := range map_services {
		log.Info().Str("name", ms.Name).Str("id", ms.ID).Str("title", ms.Title).Str("url", ms.URL.String()).Msg("found map service")
		map_service = ms
	}
	meta, err := map_service.Metadata(ctx, gis)
	log.Info().
		Str("description", meta.ServiceDescription).
		Float64("min-scale", meta.MinScale).
		Float64("max-scale", meta.MaxScale).
		Msg("got meta")
	extent_size := arcgis.ExtentSize(meta.FullExtent)
	log.Info().Float64("width", extent_size.X).Float64("height", extent_size.Y).Msg("full extent")
	for _, l := range meta.Layers {
		log.Info().Str("name", l.Name).Msg("got layer")
	}
	if meta.TileInfo != nil {
		t := meta.TileInfo
		log.Info().Int("rows", t.Rows).Int("cols", t.Cols).Int("dpi", t.DPI).Str("format", t.Format).Msg("tile info")
		for _, lod := range t.LODs {
			tiles_x, tiles_y, total := arcgis.CalculateTileCount(meta.FullExtent, lod.Level)
			log.Info().Int("level_", lod.Level).Float64("resolution", lod.Resolution).Float64("scale", lod.Scale).Int("tiles_x", tiles_x).Int("tiles_y", tiles_y).Int("total tiles", total).Msg("lod")
		}
	} else {
		log.Info().Msg("no tile info")
	}
}
