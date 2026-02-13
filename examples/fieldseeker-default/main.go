package main

import (
	"context"
	"os"

	"github.com/Gleipnir-Technology/arcgis-go/fieldseeker"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	ctx := context.TODO()
	ctx = log.With().Str("component", "arcgis").Logger().WithContext(ctx)

	fs, err := fieldseeker.NewFieldSeeker(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create fieldseeker")
		os.Exit(2)
	}
	log.Info().Msg("Getting layers")
	layers, err := fs.FeatureServerLayers(ctx)
	if err != nil {
		log.Error().Err(err).Msg("get layers")
		os.Exit(3)
	}
	for _, l := range layers {
		log.Info().Str("name", l.Name).Msg("layer")
	}
	log.Info().Msg("Done")
}
