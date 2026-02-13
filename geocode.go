package arcgis

import (
	"context"

	"github.com/Gleipnir-Technology/arcgis-go/response"
)

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

var geocodeURL string = "https://geocode-api.arcgis.com/arcgis/rest/services/World/GeocodeServer/findAddressCandidates"

func (ag *ArcGIS) GeocodeFindAddressCandidates(ctx context.Context, address string) (*response.GeocodeFindAddressCandidates, error) {
	sub := ag.requestor.withHost("https://geocode-api.arcgis.com")
	path := "/arcgis/rest/services/World/GeocodeServer/findAddressCandidates"
	params := map[string]string{
		"f":          "json",
		"outFields":  "*",
		"SingleLine": address,
	}
	return reqGetJSONParams[response.GeocodeFindAddressCandidates](ctx, sub, path, params)
}
