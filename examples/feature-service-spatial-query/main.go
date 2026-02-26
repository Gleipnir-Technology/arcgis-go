package main

import (
	"context"
	//"fmt"
	"flag"
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

	name := flag.String("name", "", "The name of the feature service to use")
	flag.Parse()

	gis, err := arcgis.NewArcGIS(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create arcgis")
		os.Exit(2)
	}

	fs, err := gis.ServiceByName(ctx, *name)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get services")
		os.Exit(3)
	}
	//var feature_service *arcgis.ServiceFeature
	log.Info().Str("name", fs.Name).Str("url", fs.URL.String()).Msg("found map service")
	for _, layer := range fs.Layers {
		count, err := fs.QueryCount(ctx, gis, layer.ID)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get count")
			continue
		}
		log.Info().Str("name", layer.Name).Uint("id", layer.ID).Int("count", count.Count).Msg("found layer")
	}
}
