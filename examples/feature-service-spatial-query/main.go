package main

import (
	"context"
	"flag"
	"fmt"
	//"net/url"
	"os"
	"reflect"

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
	meta, err := fs.PopulateMetadata(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to populate meta")
		os.Exit(4)
	}

	//var feature_service *arcgis.ServiceFeature
	log.Info().Str("name", fs.Name).Str("url", fs.URL.String()).Str("item id", meta.ServiceItemId).Int("len layers", len(meta.Layers)).Msg("found feature service")
	for _, layer := range meta.Layers {
		p := arcgis.Point{
			SpatialReference: "4326",
			X:                -119.321129,
			Y:                36.303404,
		}
		result, err := fs.QueryWithin(ctx, layer.ID, p)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get count")
			continue
		}
		log.Info().Str("gid", result.GlobalIDFieldName).Send()
		for _, feature := range result.Features {
			log.Info().Int("len attributes", len(feature.Attributes)).Msg("feature")
			for k, v := range feature.Attributes {
				var v_ string
				switch v.(type) {
				case int:
					v_ = fmt.Sprintf("%d", v)
				case float32:
					v_ = fmt.Sprintf("%f", v)
				case float64:
					v_ = fmt.Sprintf("%f", v)
				case string:
					v_ = v.(string)
				default:
					log.Warn().Str("type", reflect.TypeOf(v).Name()).Msg("Not sure what to do with type '%s'")
				}
				log.Info().Str("v", v_).Str("k", k).Msg("attribute")
			}
		}
		/*
			meta, err := fs.LayerMetadata(ctx, gis, layer.ID)
			if err != nil {
				log.Error().Err(err).Msg("Failed to get count")
				continue
			}
			log.Info().Str("name", layer.Name).Uint("id", layer.ID).Str("type", *meta.Type).Msg("found layer")
		*/
	}
}
