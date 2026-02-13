package arcgis

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Gleipnir-Technology/arcgis-go/response"
	"github.com/rs/zerolog"
)

type ServiceFeature struct {
	Layers []response.Layer
	Name   string
	URL    url.URL

	info      response.FeatureService
	requestor *gisRequestor
}

func newServiceFeature(ctx context.Context, name string, url url.URL, requestor gisRequestor) (*ServiceFeature, error) {
	logger := zerolog.Ctx(ctx)
	// POST https://services8.arcgis.com/<context>/ArcGIS/rest/services/BorderDistrict/FeatureServer
	info, err := reqPostFormToJSONFullURL[response.FeatureService](ctx, requestor, url, map[string]string{})
	if err != nil || info == nil {
		return nil, fmt.Errorf("get url: %w", err)
	}
	logger.Debug().Str("name", name).Str("description", info.ServiceDescription).Msg("init ServiceFeature")

	result := ServiceFeature{
		Layers:    info.Layers,
		Name:      name,
		URL:       url,
		info:      *info,
		requestor: &requestor,
	}
	return &result, nil
}
