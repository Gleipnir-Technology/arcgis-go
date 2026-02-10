package arcgis

import (
	"context"
	"fmt"

	"github.com/Gleipnir-Technology/arcgis-go/log"
)

var geocodeURL string = "https://geocode-api.arcgis.com/arcgis/rest/services/World/GeocodeServer/findAddressCandidates"

/*
{
	"spatialReference": {
		"wkid": 4326,
		"latestWkid": 4326
	},
	"candidates": [
		{
			"address": "1 Infinite Loop, Cupertino, California, 95014",
			"location": {
				"x": -122.030230668387,
				"y": 37.331524210671
			},
			"score": 100,
			"attributes": {},
			"extent": {
				"xmin": -122.031230668387,
				"ymin": 37.330524210671,
				"xmax": -122.029230668387,
				"ymax": 37.332524210671
			}
		},
		{
			"address": "Infinite Loop, Cupertino, California, 95014",
			"location": {
				"x": -122.028892485803,
				"y": 37.333228311614
			},
			"score": 98.54,
			"attributes": {},
			"extent": {
				"xmin": -122.029892485803,
				"ymin": 37.332228311614,
				"xmax": -122.027892485803,
				"ymax": 37.334228311614
			}
		}
	]
}
*/
func (ag *ArcGIS) GeocodeFindAddressCandidates(ctx context.Context, address string) error {
	full_url, err := addParams(geocodeURL, map[string]string{
		"f": "json",
		"outFields": "*",
		"SingleLine": address,
	})
	if err != nil {
		return fmt.Errorf("Failed to add params: %w", err)
	}
	r, err := ag.serviceRequestFromFull(full_url)
	if err != nil {
		return fmt.Errorf("Failed to create service request: %w", err)
	}
	body, err := ag.requestJSON(ctx, r)
	if err != nil {
		return fmt.Errorf("Failed to make request: %w", err)
	}
	/*
	var result RestInfo
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
	*/
	log.Info().Str("body", string(body)).Msg("did request")
	return nil
}
