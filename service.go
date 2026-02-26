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
	Name     string
	Metadata *response.FeatureService
	URL      url.URL

	requestor *gisRequestor
}

func newServiceFeature(ctx context.Context, name string, url url.URL, requestor gisRequestor) *ServiceFeature {
	result := ServiceFeature{
		Metadata: nil,
		Name:     name,
		URL:      url,

		requestor: &requestor,
	}
	return &result
}
func (sf *ServiceFeature) LayerMetadata(ctx context.Context, layer_id uint) (*response.LayerMetadata, error) {
	url := sf.URL.JoinPath(strconv.Itoa(int(layer_id)))
	return reqGetJSONParamsHeadersFullURL[response.LayerMetadata](ctx, *sf.requestor, *url, map[string]string{}, map[string]string{})
}

func (sf *ServiceFeature) Query(ctx context.Context, layer_id uint, query Query) (*response.QueryResult, error) {
	params := query.toParams()
	url := sf.URL.JoinPath(strconv.Itoa(int(layer_id)), "query")
	return reqGetJSONParamsHeadersFullURL[response.QueryResult](ctx, *sf.requestor, *url, params, map[string]string{})
}
func (sf *ServiceFeature) QueryRaw(ctx context.Context, service string, layer_id uint, query Query) ([]byte, error) {
	// path := fmt.Sprintf("/services/%s/FeatureServer/%d/query?f=json", service, layer_id)
	url := sf.URL.JoinPath(strconv.Itoa(int(layer_id)), "query?f=json")
	return reqGetParamsHeadersFullURL(ctx, *sf.requestor, *url, map[string]string{}, map[string]string{})
}

func (sf *ServiceFeature) QueryCount(ctx context.Context, layer_id uint) (*response.QueryResultCount, error) {
	params := make(map[string]string)
	params["returnCountOnly"] = "true"
	params["where"] = "9999=9999"
	url := sf.URL.JoinPath(strconv.Itoa(int(layer_id)), "query")
	return reqGetJSONParamsHeadersFullURL[response.QueryResultCount](ctx, *sf.requestor, *url, params, map[string]string{})
}

func (sf *ServiceFeature) QueryEnvelope(ctx context.Context, layer_id uint, point Point) (*response.QueryResult, error) {
	params := make(map[string]string)
	//params["where"] = "9999=9999"
	xmin := point.X - 0.01
	ymin := point.Y - 0.01
	xmax := point.X + 0.01
	ymax := point.Y + 0.01
	params["geometry"] = fmt.Sprintf("%f,%f,%f,%f", xmin, ymin, xmax, ymax)
	params["geometryType"] = "esriGeometryEnvelope"
	params["inSR"] = point.SpatialReference
	params["outFields"] = "*"
	url := sf.URL.JoinPath(strconv.Itoa(int(layer_id)), "query")
	return reqGetJSONParamsHeadersFullURL[response.QueryResult](ctx, *sf.requestor, *url, params, map[string]string{})

}
func (sf *ServiceFeature) QueryWithin(ctx context.Context, layer_id uint, point Point) (*response.QueryResult, error) {
	p_j, err := point.asJSON()
	if err != nil {
		return nil, fmt.Errorf("marshal point: %w", err)
	}
	params := make(map[string]string)
	//params["where"] = "9999=9999"
	params["f"] = "json"
	params["geometry"] = p_j
	params["geometryType"] = "esriGeometryPoint"
	params["inSR"] = point.SpatialReference
	params["outFields"] = "*"
	params["returnAllRecords"] = "true"
	params["returnGeometry"] = "true"
	params["spatialRel"] = "esriSpatialRelWithin"
	params["where"] = "1=1"
	url := sf.URL.JoinPath(strconv.Itoa(int(layer_id)), "query")
	return reqPostFormToJSONFullURL[response.QueryResult](ctx, *sf.requestor, *url, params)

}
func (sf *ServiceFeature) PopulateMetadata(ctx context.Context) (*response.FeatureService, error) {
	logger := zerolog.Ctx(ctx)
	// POST https://services8.arcgis.com/<context>/ArcGIS/rest/services/BorderDistrict/FeatureServer
	meta, err := reqPostFormToJSONFullURL[response.FeatureService](ctx, *sf.requestor, sf.URL, map[string]string{})
	if err != nil || meta == nil {
		return nil, fmt.Errorf("get url: %w", err)
	}
	logger.Debug().Str("name", sf.Name).Str("description", meta.ServiceDescription).Msg("populate ServiceFeature")

	sf.Metadata = meta
	return meta, nil
}
func (sf *ServiceFeature) Layers(ctx context.Context) ([]response.Layer, error) {
	if sf.Metadata != nil {
		return sf.Metadata.Layers, nil
	}
	meta, err := sf.PopulateMetadata(ctx)
	if err != nil {
		return make([]response.Layer, 0), err
	}
	return meta.Layers, nil
}
