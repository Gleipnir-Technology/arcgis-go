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

	ctx := log.With().Str("component", "arcgis").Logger().WithContext(context.TODO())

	batch_size := flag.Int("batch-size", 50, "The number of IDs to request in a single batch")
	index := flag.Int("layer-index", -1, "The layer index within the feature server to extract from")
	output := flag.String("output", "output.csv", "The name of the CSV file to output")
	url_str := flag.String("feature-server-url", "", "The URL of the feature server we'll extract the data from")
	workers := flag.Int("workers", 2, "The number of downloader workers to run in parallel")
	flag.Parse()

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
	//layer := fs.Metadata.Layers[*index]
	log.Info().Str("name", fs.Metadata.Layers[*index].Name).Msg("layer")
	layer, err := fs.LayerMetadata(ctx, uint(*index))
	if err != nil {
		log.Error().Err(err).Msg("failed layer metadata")
		os.Exit(4)
	}
	field_names := make([]string, len(layer.Fields)+1)
	for i, field := range layer.Fields {
		field_names[i] = *field.Name
	}
	field_names[len(layer.Fields)] = "geometry"
	log.Info().Strs("field names", field_names).Send()
	resp, err := fs.QueryIDs(ctx, uint(*index), arcgis.Query{
		Where: "1=1",
	})
	if err != nil {
		log.Error().Err(err).Msg("failed query")
		os.Exit(5)
	}
	log.Info().Int("count", len(resp.ObjectIDs)).Msg("prepping jobs")

	//to_get := selectRandom(resp.ObjectIDs, 10)
	to_get := resp.ObjectIDs

	chanDone := make(chan struct{})

	chanErrors := make(chan error)
	chanResults := make(chan []string)
	chanJobs := make(chan []int)
	go csvWriter(chanResults, *output, field_names, chanDone)

	var wg sync.WaitGroup
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go worker(fs, uint(*index), field_names, chanJobs, chanResults, chanErrors, &wg)
	}
	log.Info().Int("num_workers", *workers).Msg("started workers")

	//max_records := fs.Metadata.MaxRecordCount
	max_records := uint(*batch_size)
	for offset := uint(0); offset < uint(len(to_get)); offset += uint(max_records) {
		end := offset + max_records
		if end > uint(len(to_get)) {
			end = uint(len(to_get))
		}

		batch := to_get[offset:end]
		//log.Debug().Int("num", len(batch)).Msg("adding to the jobs queue")
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

func worker(fs *arcgis.ServiceFeature, layer_id uint, field_names []string, chanJobs <-chan []int, chanResults chan<- []string, chanErrors chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx := context.Background()
	name_to_value := make(map[string]string, len(field_names))
	for to_get := range chanJobs {
		//log.Debug().Int("len", len(to_get)).Msg("Working on job")
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
		//log.Debug().Int("count", len(rsp.Features)).Msg("processing features")
		for _, feature := range rsp.Features {
			for k, v := range feature.Attributes {
				name_to_value[k] = v.String()
			}
			geom, err := feature.Geometry.Project("EPSG:2228", "EPSG:4326")
			if err != nil {
				log.Error().Err(err).Msg("project")
				chanErrors <- fmt.Errorf("project: %w", err)
				continue
			}
			geo_json, err := geom.ToGeoJSON()
			name_to_value["geometry"] = geo_json
			row := make([]string, len(field_names))
			for i, f := range field_names {
				row[i] = name_to_value[f]
			}
			chanResults <- row
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
func csvWriter(chanRows <-chan []string, filename string, field_names []string, chanDone chan<- struct{}) {
	// Create a new file
	file, err := os.Create(filename)
	if err != nil {
		log.Error().Err(err).Msg("Error creating file")
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush() // Don't forget to flush!

	// Write header
	//header := []string{"APN", "Description", "Geometry"}
	if err := writer.Write(field_names); err != nil {
		log.Error().Err(err).Msg("Error writing header")
		return
	}
	row_count := 0
	for row := range chanRows {
		if err := writer.Write(row); err != nil {
			log.Error().Err(err).Msg("Error writing row")
			return
		}
		row_count++
		if row_count%1000 == 0 {
			log.Info().Int("rowcount", row_count).Send()
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
