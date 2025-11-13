package fieldseeker

import (
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"strings"

	"github.com/Gleipnir-Technology/arcgis-go"
)

type FieldSeeker struct {
	arcgis        *arcgis.ArcGIS
	FeatureServer *arcgis.FeatureServer
	ServiceInfo   *arcgis.ServiceInfo
	ServiceName   string
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
		arcgis:        ar,
		FeatureServer: nil,
		ServiceInfo:   nil,
		ServiceName:   "FieldSeekerGIS",
	}
	err = fs.ensureHasFeatureServer()
	if err != nil {
		return nil, fmt.Errorf("Failed to get FieldSeeker service info: %v", err)
	}
	return &fs, nil
}

func (fs *FieldSeeker) DoQuery(layer int, query *arcgis.Query) (*arcgis.QueryResult, error) {
	return fs.arcgis.DoQuery(fs.ServiceName, layer, query)
}
func (fs *FieldSeeker) DoQueryRaw(layer int, query *arcgis.Query) ([]byte, error) {
	return fs.arcgis.DoQueryRaw(fs.ServiceName, layer, query)
}

func (fs *FieldSeeker) FeatureServerLayers() []arcgis.LayerFeature {
	return fs.FeatureServer.Layers
}

func (fs *FieldSeeker) MaxRecordCount() int {
	return fs.FeatureServer.MaxRecordCount
}

func (fs *FieldSeeker) QueryCount(layer_id int) (*arcgis.QueryResultCount, error) {
	return fs.arcgis.QueryCount(fs.ServiceName, layer_id)
}

func (fs *FieldSeeker) WebhookList() ([]arcgis.Webhook, error) {
	return fs.arcgis.WebhookList(fs.ServiceName, arcgis.ServiceTypeFeatureServer)
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
