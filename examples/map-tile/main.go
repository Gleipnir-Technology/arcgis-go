package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Gleipnir-Technology/arcgis-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	var lat = flag.Float64("lat", 36.3440579, "The latitude to pull the tile from")
	var lng = flag.Float64("lng", -119.2941189, "The longitude to pull the tile from")
	var level = flag.Uint("level", 22, "The level of detail to use")
	flag.Parse()

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
	for _, s := range map_services {
		log.Info().Str("id", s.ID).Str("name", s.Name).Str("url", s.URL.String()).Str("title", s.Title).Msg("Extracting tiles from map service")
		metadata, err := s.PopulateMetadata(ctx)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get metadata")
			os.Exit(4)
		}
		log.Info().Int("layers", len(metadata.Layers)).Msg("got map metadata")
		parent_dir := fmt.Sprintf("tiles/%s", s.Name)
		err = os.MkdirAll(parent_dir, 0750)
		if err != nil {
			log.Error().Err(err).Str("parent_dir", parent_dir).Msg("Failed to make parent dir")
			os.Exit(4)
		}
		for {
			filename := fmt.Sprintf("%s/%d.jpg", parent_dir, level)
			img, err := s.TileGPS(ctx, *level, *lat, *lng)
			if err != nil {
				log.Error().Err(err).Msg("tile failure")
				os.Exit(4)
			}
			//img, err := s.Tile(ctx, 22, 1642076, 707271)
			if err != nil {
				log.Error().Err(err).Msg("failed to get tile")
				os.Exit(4)
			}
			// Create file in configured directory
			dst, err := os.Create(filename)
			if err != nil {
				log.Error().Err(err).Msg("file create failure")
				os.Exit(4)
			}
			defer dst.Close()
			// Copy rest of request body to file
			_, err = io.Copy(dst, bytes.NewReader(img))
			if err != nil {
				log.Error().Err(err).Msg("file copy failure")
				os.Exit(4)
			}
			log.Info().Str("filename", filename).Msg("Wrote file")
			*level += 1
			if *level > 0 {
				break
			}
		}
	}
}
