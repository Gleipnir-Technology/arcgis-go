package arcgis

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/Gleipnir-Technology/arcgis-go/response"
	"github.com/rs/zerolog"
)

type ServiceFeature struct {
	Layers         []response.Layer
	Name           string
	MaxRecordCount uint
	URL            url.URL

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
		Layers:         info.Layers,
		MaxRecordCount: info.MaxRecordCount,
		Name:           name,
		URL:            url,
		info:           *info,
		requestor:      &requestor,
	}
	return &result, nil
}
func (sf *ServiceFeature) QueryCount(ctx context.Context, ag *ArcGIS, layer_id uint) (*QueryResultCount, error) {
	params := make(map[string]string)
	params["returnCountOnly"] = "true"
	params["where"] = "9999=9999"
	url := sf.URL.JoinPath(strconv.Itoa(int(layer_id)), "query")
	return reqGetJSONParamsHeadersFullURL[QueryResultCount](ctx, ag.requestor, *url, params, map[string]string{})
}

func (sf *ServiceFeature) QueryPoint(ctx context.Context, ag *ArcGIS, layer_id uint) (*QueryResult, error) {
	params := make(map[string]string)
	params["where"] = "9999=9999"
	params["geometry"] = "-119.12,36.36"
	params["geometryType"] = "esriGeometryPoint"
	url := sf.URL.JoinPath(strconv.Itoa(int(layer_id)), "query")
	return reqGetJSONParamsHeadersFullURL[QueryResult](ctx, ag.requestor, *url, params, map[string]string{})

}
