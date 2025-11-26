package fieldseeker

import (
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"strings"

	"github.com/Gleipnir-Technology/arcgis-go"
	"github.com/rs/zerolog/log"
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
	FeatureServer *arcgis.FeatureServer
	ServiceInfo   *arcgis.ServiceInfo
	ServiceName   string

	arcgis    *arcgis.ArcGIS
	layerToID map[LayerType]uint
}

func extractURLParts(urlString string) (string, []string, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", nil, err
	}

	host := parsedURL.Scheme + "://" + parsedURL.Host

	// Split the path and filter empty parts
	var pathParts []string
	for _, part := range strings.Split(parsedURL.Path, "/") {
		if part != "" {
			pathParts = append(pathParts, part)
		}
	}

	return host, pathParts, nil
}

func NewFieldSeeker(ar *arcgis.ArcGIS, fieldseeker_url string) (*FieldSeeker, error) {
	// The URL for fieldseeker should be something like
	// https://foo.arcgis.com/123abc/arcgis/rest/services/FieldSeekerGIS/FeatureServer
	// We need to break it up
	host, pathParts, err := extractURLParts(fieldseeker_url)
	if err != nil {
		return nil, fmt.Errorf("Failed to break up provided url: %v", err)
	}
	if len(pathParts) < 1 {
		return nil, errors.New("Didn't get enough path parts")
	}
	context := pathParts[0]
	slog.Info("Using base fieldseeker URL", slog.String("host", host), slog.String("context", context))
	ar.Context = &context
	ar.Host = host
	fs := FieldSeeker{
		FeatureServer: nil,
		ServiceInfo:   nil,
		ServiceName:   "FieldSeekerGIS",
		arcgis:        ar,
		layerToID:     make(map[LayerType]uint, 0),
	}
	err = fs.ensureHasFeatureServer()
	if err != nil {
		return nil, fmt.Errorf("Failed to get FieldSeeker service info: %v", err)
	}
	return &fs, nil
}

func (fs *FieldSeeker) FeatureServerLayers() []arcgis.LayerFeature {
	return fs.FeatureServer.Layers
}

func (fs *FieldSeeker) MaxRecordCount() uint {
	return fs.FeatureServer.MaxRecordCount
}

func (fs *FieldSeeker) QueryCount(layer_id uint) (*arcgis.QueryResultCount, error) {
	return fs.arcgis.QueryCount(fs.ServiceName, layer_id)
}

func (fs *FieldSeeker) Schema(layer_id uint) ([]byte, error) {
	query := arcgis.NewQuery()
	query.ResultRecordCount = 1
	query.ResultOffset = 0
	query.OutFields = "*"
	query.Where = "1=1"
	return fs.arcgis.DoQueryRaw(fs.ServiceName, layer_id, query)
}

func (fs *FieldSeeker) WebhookList() ([]arcgis.Webhook, error) {
	return fs.arcgis.WebhookList(fs.ServiceName, arcgis.ServiceTypeFeatureServer)
}

func (fs *FieldSeeker) doQueryAll(layer_id uint, offset uint) (*arcgis.QueryResult, error) {
	q := arcgis.NewQuery()
	q.ResultRecordCount = fs.MaxRecordCount()
	q.ResultOffset = offset
	q.SpatialReference = "4326"
	q.OutFields = "*"
	q.Where = "1=1"
	qr, err := fs.arcgis.DoQuery(fs.ServiceName, layer_id, q)
	return qr, err
}

// Make sure we have the Layer IDs we need to perform queries
func (fs *FieldSeeker) ensureHasFeatureServer() error {
	err := fs.ensureHasServices()
	if err != nil {
		return fmt.Errorf("Failed to ensure has services: %v", err)
	}
	if fs.FeatureServer != nil {
		slog.Info("already has feature server")
		return nil
	}
	s, err := fs.arcgis.GetFeatureServer(fs.ServiceName)
	if err != nil {
		return fmt.Errorf("Failed to get feature server: %v", err)
	}
	if s == nil {
		return errors.New("Got a null feature server")
	}
	slog.Info("Add feature server", slog.String("item id", s.ServiceItemId))
	fs.FeatureServer = s
	for _, layer := range fs.FeatureServerLayers() {
		t, err := NameToLayerType(layer.Name)
		if err != nil {
			log.Warn().Err(err).Msg("Failed to handle layer")
			continue
		}
		fs.layerToID[t] = layer.ID
	}
	return nil
}

// Make sure we have the Service IDs we need to use FieldSeeker
func (fs *FieldSeeker) ensureHasServices() error {
	if fs.ServiceInfo != nil {
		slog.Info("already has services")
		return nil
	}
	s, err := fs.arcgis.Services()
	if err != nil {
		return fmt.Errorf("Failed to query services: %v", err)
	}
	if s == nil {
		return errors.New("Got a null service info")
	}
	fs.ServiceInfo = s
	slog.Info("Add service info", slog.Float64("version", s.CurrentVersion), slog.Int("services", len(s.Services)))
	return nil
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
func featureToStruct[S any](fs *FieldSeeker, layer LayerType, offset uint) ([]*S, error) {
	var results []*S

	layer_id, ok := fs.layerToID[layer]
	if !ok {
		return results, fmt.Errorf("Cannot get layer %s", layer)
	}
	qr, err := fs.doQueryAll(layer_id, offset)
	if err != nil {
		return results, fmt.Errorf("Failed to query %s (layer %d): %w", layer, layer_id, err)
	}

	for _, feature := range qr.Features {
		s, err := structFromAttributes[S](feature.Attributes)
		if err != nil {
			return results, fmt.Errorf("Failed to get %s from query result: %w", layer, err)
		}
		results = append(results, s)
	}
	return results, nil
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
