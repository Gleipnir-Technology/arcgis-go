package main

import (
	"context"
	"os"

	"github.com/Gleipnir-Technology/arcgis-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

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
	resp, err := gis.Search(ctx, "FieldseekerGIS")
	if err != nil {
		log.Error().Err(err).Msg("search fieldseeker")
		os.Exit(3)
	}
	for _, result := range resp.Results {
		log.Info().Str("url", result.URL).Str("name", result.Name).Msg("Found")
	}
}
