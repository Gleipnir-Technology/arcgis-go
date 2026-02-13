package arcgis

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Gleipnir-Technology/arcgis-go/response"
	"github.com/rs/zerolog"
)

type ServiceFeature struct {
	Name string
	URL  url.URL

	requestor *gisRequestor
}

func newServiceFeature(ctx context.Context, name string, url url.URL, requestor gisRequestor) (*ServiceFeature, error) {
	result := ServiceFeature{
		Name:      name,
		URL:       url,
		requestor: &requestor,
	}
	err := result.init(ctx)
	if err != nil {
		return nil, fmt.Errorf("init: %w", err)
	}
	return &result, nil
}

// Make sure we have the Service IDs we need to use FieldSeeker
func (s *ServiceFeature) init(ctx context.Context) error {
	// POST https://services8.arcgis.com/pV7SH1EgRc6tpxlJ/ArcGIS/rest/services/BorderDistrict/FeatureServer
	logger := zerolog.Ctx(ctx)
	resp, err := reqPostFormToJSONFullURL[response.FeatureService](ctx, *s.requestor, s.URL, map[string]string{})
	if err != nil {
		return fmt.Errorf("get url: %w", err)
	}
	logger.Debug().Str("name", s.Name).Str("description", resp.ServiceDescription).Msg("init ServiceFeature")
	return nil
}
