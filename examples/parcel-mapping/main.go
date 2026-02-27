package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Gleipnir-Technology/arcgis-go"
	"github.com/Gleipnir-Technology/arcgis-go/response"
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
	log.Info().Int("count", len(resp.ObjectIDs)).Msg("prepping jobs")

	to_get := selectRandom(resp.ObjectIDs, 10)

	chanDone := make(chan struct{})

	chanErrors := make(chan error)
	chanResults := make(chan []string)
	chanJobs := make(chan []int)
	go csvWriter(chanResults, "output.csv", chanDone)

	var wg sync.WaitGroup
	num_workers := 2
	for i := 0; i < num_workers; i++ {
		wg.Add(1)
		go worker(fs, uint(*index), *feature_name_apn, *feature_name_desc, chanJobs, chanResults, chanErrors, &wg)
	}
	log.Info().Int("num_workers", num_workers).Msg("started workers")

	max_records := fs.Metadata.MaxRecordCount
	for offset := uint(0); offset < uint(len(to_get)); offset += uint(max_records) {
		end := offset + max_records
		if end > uint(len(to_get)) {
			end = uint(len(to_get))
		}

		batch := to_get[offset:end]
		log.Debug().Int("num", len(batch)).Msg("adding to the jobs queue")
		chanJobs <- batch
	}
	close(chanJobs)

	log.Debug().Msg("Waiting for results")
	go func() {
		wg.Wait()
		close(chanResults)
		close(chanErrors)
	}()

	// Check for errors
	for err := range chanErrors {
		if err != nil {
			log.Error().Err(err).Msg("error in worker")
		}
	}

	// Wait for write to finish
	log.Debug().Msg("Waiting for writes to complete")
	<-chanDone
	log.Info().Msg("Work complete")
}

func worker(fs *arcgis.ServiceFeature, layer_id uint, feature_name_apn string, feature_name_desc string, chanJobs <-chan []int, chanResults chan<- []string, chanErrors chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx := context.Background()
	for to_get := range chanJobs {
		log.Debug().Int("len", len(to_get)).Msg("Working on job")
		id_strs := make([]string, len(to_get))
		for i, idx := range to_get {
			id_strs[i] = strconv.Itoa(idx)
		}
		in_clause := strings.Join(id_strs, ",")
		q := arcgis.Query{
			OutFields: "*",
			Where:     fmt.Sprintf("OBJECTID IN (%s)", in_clause),
		}
		rsp, err := fs.Query(ctx, layer_id, q)
		if err != nil {
			log.Error().Err(err).Msg("query failed")
			chanErrors <- fmt.Errorf("query: %w", err)
			continue
		}
		log.Debug().Int("count", len(rsp.Features)).Msg("processing features")
		for _, feature := range rsp.Features {
			/*
				for k, v := range feature.Attributes {
					log.Info().Str("k", k).Str("v", v.String()).Send()
				}
			*/
			apn := feature.Attributes[feature_name_apn].String()
			desc := feature.Attributes[feature_name_desc].String()
			if desc == "" {
				desc = "\"\""
			}
			geom, err := feature.Geometry.Project("EPSG:2228", "EPSG:4326")
			if err != nil {
				log.Error().Err(err).Msg("project")
				chanErrors <- fmt.Errorf("project: %w", err)
				continue
			}
			geo_json, err := geom.ToGeoJSON()
			chanResults <- []string{
				apn,
				desc,
				geo_json,
			}
		}
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
func csvWriter(chanRows <-chan []string, filename string, chanDone chan<- struct{}) {
	// Create a new file
	file, err := os.Create("output.csv")
	if err != nil {
		log.Error().Err(err).Msg("Error creating file")
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush() // Don't forget to flush!

	// Write header
	header := []string{"APN", "Description", "Geometry"}
	if err := writer.Write(header); err != nil {
		log.Error().Err(err).Msg("Error writing header")
		return
	}
	for row := range chanRows {
		if err := writer.Write(row); err != nil {
			log.Error().Err(err).Msg("Error writing row")
			return
		}
	}
	writer.Flush()
	file.Close()
	chanDone <- struct{}{}
	close(chanDone)
}
func projectGeometry(geom response.Geometry) (string, error) {
	// WKID 2228 is EPSG:2228 (NAD83 / California Zone 3, US Survey Feet)
	/*pj, err := proj.NewCRSToCRS(
		"EPSG:2228",      // Source: California Zone 3
		"EPSG:4326",      // Target: WGS84
		nil,
	)*/
	return "", nil
}
