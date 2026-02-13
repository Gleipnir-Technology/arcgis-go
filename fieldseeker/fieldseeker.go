package fieldseeker

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/Gleipnir-Technology/arcgis-go"
	"github.com/Gleipnir-Technology/arcgis-go/response"
	"github.com/rs/zerolog"
)

type LayerType uint

const (
	LayerUnknown LayerType = iota
	LayerAerialSpraySession
	LayerAerialSprayLine
	LayerBarrierSpray
	LayerBarrierSprayRoute
	LayerContainerRelate
	LayerFieldScoutingLog
	LayerHabitatRelate
	LayerInspectionSample
	LayerInspectionSampleDetail
	LayerLandingCount
	LayerLandingCountLocation
	LayerLineLocation
	LayerLocationTracking
	LayerMosquitoInspection
	LayerOfflineMapAreas
	LayerProposedTreatmentArea
	LayerPointLocation
	LayerPolygonLocation
	LayerPoolDetail
	LayerPool
	LayerPoolBuffer
	LayerQALarvCount
	LayerQAMosquitoInspection
	LayerQAProductObservation
	LayerRestrictedArea
	LayerRodentInspection
	LayerRodentLocation
	LayerSampleCollection
	LayerSampleLocation
	LayerServiceRequest
	LayerSpeciesAbundance
	LayerStormDrain
	LayerTracklog
	LayerTrapLocation
	LayerTrapData
	LayerTimeCard
	LayerTreatment
	LayerTreatmentArea
	LayerULVSprayRoute
	LayerZones
	LayerZones2
)

type FieldSeeker struct {
	Arcgis         *arcgis.ArcGIS
	ServiceFeature *arcgis.ServiceFeature
	ServiceName    string

	layerToID map[LayerType]uint
}

var fieldseekerFeatureName string = "FieldSeekerGIS"

func NewFieldSeeker(ctx context.Context) (*FieldSeeker, error) {
	ag, err := arcgis.NewArcGIS(ctx)
	if err != nil {
		return nil, fmt.Errorf("new arcgis: %w", err)
	}
	return NewFieldSeekerFromAG(ctx, *ag)
}
func NewFieldSeekerFromAG(ctx context.Context, ag arcgis.ArcGIS) (*FieldSeeker, error) {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("Created arcgis client, searching for FieldSeekerGIS FeatureServer")
	resp, err := ag.Search(ctx, fmt.Sprintf("name:\"%s\"", fieldseekerFeatureName))
	if err != nil {
		return nil, fmt.Errorf("search %s: %w", fieldseekerFeatureName, err)
	}
	logger.Debug().Int("total", resp.Total).Int("num", resp.Num).Msg("fieldseeker search results")
	for _, r := range resp.Results {
		logger.Debug().Str("name", r.Name).Str("type", r.Type).Str("url", r.URL).Msg("Search Result")
		if r.Name == fieldseekerFeatureName && r.Type == "Feature Service" {
			return NewFieldSeekerFromURL(ctx, ag, r.URL)
		}
	}
	return nil, fmt.Errorf("Failed to find a Feature Service named '%s'", fieldseekerFeatureName)
}
func NewFieldSeekerFromURL(ctx context.Context, ag arcgis.ArcGIS, url_str string) (*FieldSeeker, error) {
	logger := zerolog.Ctx(ctx)
	u, err := url.Parse(url_str)
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}
	service, err := ag.NewServiceFeature(ctx, fieldseekerFeatureName, *u)
	if err != nil {
		return nil, fmt.Errorf("new service feature: %w", err)
	}
	layer_to_id := make(map[LayerType]uint, 0)
	for _, layer := range service.Layers {
		t, err := NameToLayerType(layer.Name)
		if err != nil {
			logger.Warn().Err(err).Msg("Failed to handle layer")
			continue
		}
		layer_to_id[t] = layer.ID
	}

	result := FieldSeeker{
		Arcgis:         &ag,
		ServiceFeature: service,
		//ServiceInfo:   nil,
		ServiceName: "FieldSeekerGIS",
		layerToID:   layer_to_id,
	}
	return &result, nil
}

func (fs *FieldSeeker) AdminInfo(ctx context.Context) (*arcgis.AdminInfo, error) {
	return fs.Arcgis.AdminInfo(ctx, fs.ServiceName, arcgis.ServiceTypeFeatureServer)
}

func (fs *FieldSeeker) Layers() []response.Layer {
	return fs.ServiceFeature.Layers
}

func (fs *FieldSeeker) MaxRecordCount(ctx context.Context) (uint, error) {
	return fs.ServiceFeature.MaxRecordCount, nil
}

func (fs *FieldSeeker) PermissionList(ctx context.Context) (*response.PermissionSlice, error) {
	return fs.Arcgis.PermissionList(ctx, fs.ServiceName, arcgis.ServiceTypeFeatureServer)
}
func (fs *FieldSeeker) QueryCount(ctx context.Context, layer_id uint) (*arcgis.QueryResultCount, error) {
	return fs.Arcgis.QueryCount(ctx, fs.ServiceName, layer_id)
}

func (fs *FieldSeeker) SchemaRaw(ctx context.Context, layer_id uint) ([]byte, error) {
	query := arcgis.NewQuery()
	query.ResultRecordCount = 1
	query.ResultOffset = 0
	query.OutFields = "*"
	query.Where = "1=1"
	return fs.Arcgis.QueryRaw(ctx, fs.ServiceName, layer_id, query)
}
func (fs *FieldSeeker) Schema(ctx context.Context, layer_id uint) (*arcgis.QueryResult, error) {
	query := arcgis.NewQuery()
	query.ResultRecordCount = 1
	query.ResultOffset = 0
	query.OutFields = "*"
	query.Where = "1=1"
	return fs.Arcgis.Query(ctx, fs.ServiceName, layer_id, query)
}

func (fs *FieldSeeker) WebhookList(ctx context.Context) (*arcgis.WebhookSlice, error) {
	return fs.Arcgis.WebhookList(ctx, fs.ServiceName, arcgis.ServiceTypeFeatureServer)
}

func (fs *FieldSeeker) doQueryAll(ctx context.Context, layer_id uint, offset uint) (*arcgis.QueryResult, error) {
	q := arcgis.NewQuery()
	count, err := fs.MaxRecordCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to query for max count: %w", err)
	}
	q.ResultRecordCount = count
	q.ResultOffset = offset
	q.SpatialReference = "4326"
	q.OutFields = "*"
	q.Where = "1=1"
	return fs.Arcgis.Query(ctx, fs.ServiceName, layer_id, q)
}

func NameToLayerType(n string) (LayerType, error) {
	switch n {
	case "LocationTracking":
		return LayerLocationTracking, nil
	case "Tracklog":
		return LayerTracklog, nil
	case "ServiceRequest":
		return LayerServiceRequest, nil
	case "TrapLocation":
		return LayerTrapLocation, nil
	case "LandingCountLocation":
		return LayerLandingCountLocation, nil
	case "SampleLocation":
		return LayerSampleLocation, nil
	case "ContainerRelate":
		return LayerContainerRelate, nil
	case "HabitatRelate":
		return LayerHabitatRelate, nil
	case "PoolDetail":
		return LayerPoolDetail, nil
	case "Pool":
		return LayerPool, nil
	case "SpeciesAbundance":
		return LayerSpeciesAbundance, nil
	case "PointLocation":
		return LayerPointLocation, nil
	case "InspectionSample":
		return LayerInspectionSample, nil
	case "InspectionSampleDetail":
		return LayerInspectionSampleDetail, nil
	case "MosquitoInspection":
		return LayerMosquitoInspection, nil
	case "TrapData":
		return LayerTrapData, nil
	case "LandingCount":
		return LayerLandingCount, nil
	case "TimeCard":
		return LayerTimeCard, nil
	case "Treatment":
		return LayerTreatment, nil
	case "SampleCollection":
		return LayerSampleCollection, nil
	case "StormDrain":
		return LayerStormDrain, nil
	case "QAProductObservation":
		return LayerQAProductObservation, nil
	case "QALarvCount":
		return LayerQALarvCount, nil
	case "QAMosquitoInspection":
		return LayerQAMosquitoInspection, nil
	case "FieldScoutingLog":
		return LayerFieldScoutingLog, nil
	case "BarrierSpray":
		return LayerBarrierSpray, nil
	case "BarrierSprayRoute":
		return LayerBarrierSprayRoute, nil
	case "LineLocation":
		return LayerLineLocation, nil
	case "ULVSprayRoute":
		return LayerULVSprayRoute, nil
	case "OfflineMapAreas":
		return LayerOfflineMapAreas, nil
	case "TreatmentArea":
		return LayerTreatmentArea, nil
	case "RestrictedArea":
		return LayerRestrictedArea, nil
	case "ProposedTreatmentArea":
		return LayerProposedTreatmentArea, nil
	case "PolygonLocation":
		return LayerPolygonLocation, nil
	case "Zones2":
		return LayerZones2, nil
	case "Zones":
		return LayerZones, nil
	case "AerialSpraySession":
		return LayerAerialSpraySession, nil
	case "AerialSprayLine":
		return LayerAerialSprayLine, nil
	case "PoolBuffer":
		return LayerPoolBuffer, nil
	case "RodentInspection":
		return LayerRodentInspection, nil
	case "RodentLocation":
		return LayerRodentLocation, nil
	default:
		return LayerUnknown, errors.New(fmt.Sprintf("'%s' is not a recognized layer name", n))
	}
}

type Geometric interface {
	SetGeometry(json.RawMessage)
	GetGeometry() json.RawMessage
}

func featureToStruct[T any, PT interface {
	*T
	Geometric
}](ctx context.Context, fs *FieldSeeker, layer LayerType, offset uint) ([]PT, error) {
	var results []PT

	layer_id, ok := fs.layerToID[layer]
	if !ok {
		return results, fmt.Errorf("Cannot get layer %s", layer)
	}
	qr, err := fs.doQueryAll(ctx, layer_id, offset)
	if err != nil {
		return results, fmt.Errorf("Failed to query %s (layer %d): %w", layer, layer_id, err)
	}

	for _, feature := range qr.Features {
		//logFeature(feature)
		s, err := structFromFeature[T, PT](ctx, &feature)
		if err != nil {
			return results, fmt.Errorf("Failed to get %s from query result: %w", layer, err)
		}
		results = append(results, s)
	}
	return results, nil
}

func logFeature(f arcgis.Feature) {
	/*
		kv := make(map[string]string, 0)
		l := zerolog.Dict()
		for k, v := range f.Attributes {
				s, ok := v.(string)
				if ok {
					kv[k] = s
					continue
				}
				i, ok := v.(int)
				if ok {
					kv[k] = string(i)
					continue
				}
				kv[k] = "*unknown*"
			l.Interface(k, v)
		}
		log.Debug().Dict("feature", l).Msg("Handling feature")
	*/
}
func stringOrEmpty(data map[string]any, key string) string {
	source, ok := data[key].(string)
	if ok {
		return source
	}
	return ""
}

func truncateURL(u string, marker string) (string, error) {
	// Parse the URL
	parsedURL, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	path := parsedURL.Path
	idx := strings.Index(path, marker)
	if idx == -1 {
		return "", fmt.Errorf("%s not found in URL path", marker)
	}

	parsedURL.Path = path[:idx+len(marker)]

	parsedURL.RawQuery = ""
	parsedURL.Fragment = ""

	// Return the truncated URL
	return parsedURL.String(), nil
}
