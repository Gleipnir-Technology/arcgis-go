package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Gleipnir-Technology/arcgis-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var BaseURL, ClientID, ClientSecret, Environment, FieldseekerSchemaDirectory, MapboxToken string

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	ctx := context.TODO()
	ctx = log.With().Str("component", "arcgis").Logger().WithContext(ctx)

	url_str := flag.String("feature-server-url", "", "The URL of the feature server we'll extract the data from")
	index := flag.Int("layer-index", -1, "The layer index within the feature server to extract from")
	feature_name_apn := flag.String("feature-name-apn", "", "The name of the feature to pull APN values from")
	feature_name_desc := flag.String("feature-name-description", "", "The name of the feature to pull description values from")
	flag.Parse()

	if *feature_name_apn == "" {
		log.Error().Msg("You must specify -feature-name-apn")
		os.Exit(1)
	}
	if *feature_name_desc == "" {
		log.Error().Msg("You must specify -feature-name-description")
		os.Exit(1)
	}
	if *url_str == "" {
		log.Error().Msg("You must specify -feature-server-url")
		os.Exit(1)
	}
	if *index == -1 {
		log.Error().Msg("You must specify -index")
		os.Exit(1)
	}
	gis, err := arcgis.NewArcGIS(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create arcgis")
		os.Exit(2)
	}

	u, err := url.Parse(*url_str)
	if err != nil {
		log.Error().Err(err).Msg("url parse failed")
		os.Exit(3)
	}

	fs, err := gis.ServiceByURL(ctx, *u)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get services")
		os.Exit(3)
	}

	log.Info().Str("name", fs.Name).Str("url", fs.URL.String()).Str("item id", fs.Metadata.ServiceItemId).Int("len layers", len(fs.Metadata.Layers)).Msg("found feature service")
	layer := fs.Metadata.Layers[*index]
	log.Info().Str("name", layer.Name).Msg("layer")
	resp, err := fs.QueryIDs(ctx, uint(*index), arcgis.Query{
		Where: "1=1",
	})
	if err != nil {
		log.Error().Err(err).Msg("failed query")
	}
	log.Info().Int("count", len(resp.ObjectIDs)).Msg("done")
	to_get := selectRandom(resp.ObjectIDs, 10)
	id_strs := make([]string, len(to_get))
	for i, idx := range to_get {
		id_strs[i] = strconv.Itoa(idx)
	}
	in_clause := strings.Join(id_strs, ",")
	q := arcgis.Query{
		OutFields: "*",
		Where:     fmt.Sprintf("OBJECTID IN (%s)", in_clause),
	}
	rsp, err := fs.Query(ctx, uint(*index), q)
	if err != nil {
		log.Error().Err(err).Msg("failed query")
	}
	for _, feature := range rsp.Features {
		/*
			for k, v := range feature.Attributes {
				log.Info().Str("k", k).Str("v", v.String()).Send()
			}
		*/
		v := feature.Attributes[*feature_name_apn]
		if v == nil {
			log.Error().Str("apn-name", *feature_name_apn).Msg("nil v")
			continue
		}
		apn := v.String()
		desc := feature.Attributes[*feature_name_desc].String()
		geom := feature.Geometry.String()
		log.Info().Str("apn", apn).Str("desc", desc).Str("geom", geom).Send()
	}
}
func selectRandom(cohort []int, num int) []int {
	result := make([]int, num)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	for i := range num {
		ig := r.Intn(len(cohort))
		result[i] = cohort[ig]
	}
	return result
}
