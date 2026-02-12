package arcgis

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/google/uuid"
)

// Root structure for an instance of the ArcGIS API
type ArcGIS struct {
	//ServiceRoot   string
	requestor gisRequestor

	usage Usage
}

// Basic information about the REST API itself
type AuthInfo struct {
	isTokenBasedSecurity bool
	tokenServiceUrl      string
}
type RestInfo struct {
	CurrentVersion  float64
	FullVersion     string
	OwningSystemUrl string
	OwningTenant    string
	AuthInfo        AuthInfo
}

// Listing of available services
type ServiceListing struct {
	Name string
	Type string
	URL  string
}

type ServiceInfo struct {
	CurrentVersion float64
	Services       []ServiceListing
}

// Feature Server details
type LayerFeature struct {
	ID                uint
	Name              string
	ParentLayerID     int
	DefaultVisibility bool
	SubLayerIDs       *string
	MinScale          int
	MaxScale          int
	Type              string
	GeometryType      string
}

type Table struct {
	ID                int
	Name              string
	ParentLayerID     int
	DefaultVisibility bool
	SubLayerIDs       *string
	MinScale          int
	MaxScale          int
}

type FeatureServer struct {
	CurrentVersion                 float64
	ServiceItemId                  string
	ServiceDescription             string
	HasVersionedData               bool
	HasSharedDomains               bool
	MaxRecordCount                 uint
	SupportedQueryFormats          string
	SupportsVCSProjection          bool
	SupportedExportFormats         string
	SupportedConvertFileFormats    string
	SupportedConvertContentFormats string
	SupportedFullTextLocales       []string
	Capabilities                   string
	Description                    string
	CopyrightText                  string
	SpatialReference               SpatialReference
	InitialExtent                  Extent
	FullExtent                     Extent
	AllowGeometryUpdates           bool
	SupportsTrueCurve              bool
	SupportedCurveTypes            []string
	AllowTrueCurvesUpdates         bool
	Layers                         []LayerFeature
	Tables                         []Table
	// many missing fields
}

// Query endpoint
type UniqueIdField struct {
	Name               string
	IsSystemMaintained bool
}

type Feature struct {
	Attributes map[string]any
	Geometry   json.RawMessage
}

type CodedValue struct {
	Code CodeWrapper
	Name string
}

type Domain struct {
	CodedValues []CodedValue
	MergePolicy string
	Name        string
	SplitPolicy string
	Type        string
}

type DefaultValueWrapper string

type Field struct {
	Alias        string
	DefaultValue *DefaultValueWrapper
	Domain       *Domain
	Length       int
	Name         string
	SQLType      string
	Type         string
}

type QueryResult struct {
	Features          []Feature
	Fields            []Field
	GeometryType      string
	GlobalIDFieldName string
	ObjectIdFieldName string
	SpatialReference  SpatialReference
	UniqueIdField     UniqueIdField
}

type QueryResultCount struct {
	Count int
}

type Usage struct {
	// Units used on the last request
	LastRequest int
	// The maximum units we can use in a minute
	MaxPerMinute int
	// The amount used in the current minute
	ThisMinute int
}

func NewArcGIS(auth Authenticator, host *string) *ArcGIS {
	h := "https://www.arcgis.com"
	if host != nil {
		h = *host
	}
	return &ArcGIS{
		requestor: newGisRequestor(auth, h),
		//ServiceRoot:   "https://www.arcgis.com/sharing/rest",
	}
}

func ServiceRootFromTenant(base string, tenantId string) string {
	return fmt.Sprintf("%s/%s", base, tenantId)
}

func (ag *ArcGIS) Query(ctx context.Context, service string, layer_id uint, query *Query) (*QueryResult, error) {
	path := fmt.Sprintf("/services/%s/FeatureServer/%d/query", service, layer_id)
	return doJSONGet[QueryResult](ctx, ag.requestor, path)
}
func (ag *ArcGIS) QueryRaw(ctx context.Context, service string, layer_id uint, query *Query) ([]byte, error) {
	path := fmt.Sprintf("/services/%s/FeatureServer/%d/query", service, layer_id)
	return doGet(ctx, ag.requestor, path)
}

type AdminInfo struct {
	CurrentVersion string `json:"currentversion"`
}

func (ag *ArcGIS) AdminInfo(ctx context.Context, serviceName string, serviceType ServiceType) (*AdminInfo, error) {
	// We may need to always direct this request to
	path := fmt.Sprintf("/ArcGIS/rest/admin/services/%s/%s/permissions", serviceName, ServiceTypeNames[serviceType])
	return doJSONGet[AdminInfo](ctx, ag.requestor, path)
}

func (ag *ArcGIS) GetFeatureServer(ctx context.Context, service string) (*FeatureServer, error) {
	path := fmt.Sprintf("/services/%s/FeatureServer", service)
	return doJSONGet[FeatureServer](ctx, ag.requestor, path)
}

/*
not valid
func (ag *ArcGIS) Root(ctx context.Context) (string, error) {
	req, err := ag.serviceRequest("/")
	if err != nil {
		return "", fmt.Errorf("Failed to create root request: %w", err)
	}
	content, err := ag.requestJSON(ctx, req)
	if err != nil {
		return "", fmt.Errorf("Failed to request JSON: %w", err)
	}
	return string(content), err
}
*/

func (ag *ArcGIS) Search(ctx context.Context, query string) (*SearchResponse, error) {
	return doJSONPost[SearchResponse](ctx, ag.requestor, "/sharing/rest/search", map[string]string{
		"q": query,
	})
}
func (ag *ArcGIS) Services(ctx context.Context) (*ServiceInfo, error) {
	return doJSONGet[ServiceInfo](ctx, ag.requestor, "/sharing/rest/services")
}
func (ag *ArcGIS) SwitchHostByPortal(ctx context.Context) error {
	logger := LoggerFromContext(ctx)
	portals, err := ag.PortalsSelf(ctx)
	if err != nil {
		return fmt.Errorf("Failed to get portals: %w", err)
	} else if portals == nil {
		return errors.New("Returned portals was nil")
	}
	logger.Debug().Str("id", portals.ID).Str("name", portals.PortalName).Str("urlkey", portals.UrlKey).Msg("Found a portal")
	ag.requestor.host = fmt.Sprintf("https://%s.maps.arcgis.com", portals.UrlKey)
	return nil
}
func parseFeatureServer(data []byte) (*FeatureServer, error) {
	var result FeatureServer
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func parseQueryResult(ctx context.Context, data []byte) (*QueryResult, error) {
	var result QueryResult
	logger := LoggerFromContext(ctx)
	err := json.Unmarshal(data, &result)
	if err != nil {
		id := uuid.New()
		filename := fmt.Sprintf("debug_response_%s.json", id.String())
		output, err2 := os.Create(filename)
		if err2 != nil {
			logger.Warn().Str("filename", filename).Msg("Failed to create debug file, can't dump request.")
			return nil, fmt.Errorf("Failed to parse query result JSON: %w", err)
		}
		defer output.Close()
		output.Write(data)
		logger.Info().Str("filename", filename).Msg("Wrote response file for debugging.")
		return nil, fmt.Errorf("Failed to parse query result JSON: %w. Wrote debug file containing the request to %s", err, filename)
	}
	return &result, nil
}

func parseQueryResultCount(data []byte) (*QueryResultCount, error) {
	var result QueryResultCount
	err := json.Unmarshal(data, &result)
	//Println("Parsing", string(data))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func parseRestInfo(data []byte) (*RestInfo, error) {
	var result RestInfo
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func parseSearchResponse(ctx context.Context, data []byte) (*SearchResponse, error) {
	saveResponse(ctx, data, "search.json")
	var result SearchResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func parseServiceInfo(data []byte) (*ServiceInfo, error) {
	var result ServiceInfo
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func saveResponse(ctx context.Context, data []byte, filename string) {
	logger := LoggerFromContext(ctx)
	dest, err := os.Create(filename)
	if err != nil {
		logger.Error().Str("filename", filename).Str("err", err.Error()).Msg("Failed to create file")
		return
	}
	_, err = io.Copy(dest, bytes.NewReader(data))
	if err != nil {
		logger.Error().Str("filename", filename).Str("err", err.Error()).Msg("Failed to write")
		return
	}
	logger.Info().Str("filename", filename).Msg("Wrote response")
}

var sharingBaseURL string = "https://www.arcgis.com/sharing/rest"

func addParams(u string, params map[string]string) (*url.URL, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	_, ok := params["f"]
	if !ok {
		params["f"] = "json"
	}
	vals := url.Values{}
	for k, v := range params {
		vals.Add(k, v)
	}
	parsed.RawQuery = vals.Encode()
	return parsed, nil
}

func (ag *ArcGIS) updateUsage(ctx context.Context, resp *http.Response) {
	logger := LoggerFromContext(ctx)
	qru := resp.Header["X-Esri-Query-Request-Units"]
	for _, v := range qru {
		n, err := fmt.Sscanf(v, "%d", &ag.usage.LastRequest)
		if err != nil {
			logger.Warn().Str("err", err.Error()).Msg("Failed to parse X-Esri-Query-Request-Units")
		}
		if n < 1 {
			logger.Warn().Msg("Parsed no values from X-Esri-Query-Request-Units")
		}
	}
	orupm := resp.Header["X-Esri-Org-Request-Units-Per-Min"]
	for _, v := range orupm {
		// The rupm value is of the form "usage=97;max=10000"
		n, err := fmt.Sscanf(v, "usage=%d;max=%d", &ag.usage.ThisMinute, &ag.usage.MaxPerMinute)
		if err != nil {
			logger.Warn().Str("err", err.Error()).Msg("Failed to parse X-Esri-Org-Request-Units-Per-Min:")
		}
		if n < 2 {
			logger.Warn().Msg("Parsed too few values from X-Esri-Org-Request-Per-Min")
		}
	}
}

func (ag *ArcGIS) Info(ctx context.Context) (*RestInfo, error) {
	return doJSONGet[RestInfo](ctx, ag.requestor, "/sharing/rest/info")
}

type Query struct {
	Limit             int
	ObjectIDs         string
	OutFields         string
	ResultRecordCount uint
	ResultOffset      uint
	SpatialReference  string // Should eventually make an enum, probably
	Where             string
}

func NewQuery() *Query {
	q := new(Query)
	return q
}

func (query Query) toParams() map[string]string {
	params := make(map[string]string)
	if query.Limit > 0 {
		params["limit"] = strconv.Itoa(query.Limit)
	}
	if query.ObjectIDs != "" {
		params["objectIds"] = query.ObjectIDs
	}
	if query.OutFields != "" {
		params["outFields"] = query.OutFields
	}
	if query.ResultOffset > 0 {
		params["resultOffset"] = strconv.Itoa(int(query.ResultOffset))
	}
	if query.Where != "" {
		params["where"] = query.Where
	}
	if len(query.SpatialReference) > 0 {
		params["outSR"] = query.SpatialReference
	}
	return params
}

type PermissionSlice = []Permission

func (ag *ArcGIS) PermissionList(ctx context.Context, serviceName string, serviceType ServiceType) (*PermissionSlice, error) {
	path := fmt.Sprintf("/ArcGIS/rest/admin/services/%s/%s/permissions", serviceName, ServiceTypeNames[serviceType])
	return doJSONGet[PermissionSlice](ctx, ag.requestor, path)
}

type PermissionEntry struct {
	Constraint string `json:"constraint"`
	IsAllowed  bool   `json:"isAllowed"`
}
type Permission struct {
	ChildURL   string          `json:"childURL"`
	Operation  string          `json:"operation"`
	Permission PermissionEntry `json:"permission"`
	Principal  string          `json:"principal"`
}
type PermissionListResponse struct {
	Permissions []Permission `json:"permissions"`
}

func parsePermissionListResponse(ctx context.Context, data []byte) (*PermissionListResponse, error) {
	var result PermissionListResponse
	saveResponse(ctx, data, "permission-list.json")
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal JSON: %w", err)
	}
	return &result, nil
}
func (ag *ArcGIS) QueryCount(ctx context.Context, service string, layer_id uint) (*QueryResultCount, error) {
	params := make(map[string]string)
	params["returnCountOnly"] = "true"
	params["where"] = "9999=9999"
	path := fmt.Sprintf("/services/%s/FeatureServer/%d/query", service, layer_id)
	return doJSONGetParams[QueryResultCount](ctx, ag.requestor, path, params)
}

func (d *DefaultValueWrapper) UnmarshalJSON(data []byte) (err error) {
	// Does it look like a string?
	var content string
	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		content = string(data[1 : len(data)-1])
	} else {
		content = hex.EncodeToString(data)
	}

	*d = DefaultValueWrapper(string(content))
	return nil
}

type ServiceType int

const (
	ServiceTypeGPServer ServiceType = iota
	ServiceTypeFeatureServer
)

var ServiceTypeNames = map[ServiceType]string{
	ServiceTypeGPServer:      "GPServer",
	ServiceTypeFeatureServer: "FeatureServer",
}

type Webhook struct {
	Name string
}
type WebhookSlice = []Webhook

func (ag *ArcGIS) WebhookList(ctx context.Context, serviceName string, serviceType ServiceType) (*WebhookSlice, error) {
	path := fmt.Sprintf("/ArcGIS/rest/admin/services/%s/%s/webhooks", serviceName, ServiceTypeNames[serviceType])
	return doJSONGet[WebhookSlice](ctx, ag.requestor, path)
}

type WebhookListResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}

func parseWebhookListResponse(ctx context.Context, data []byte) (*WebhookListResponse, error) {
	var result WebhookListResponse
	saveResponse(ctx, data, "webhook-list.json")
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal JSON: %w", err)
	}
	return &result, nil
}
