package arcgis

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// Root structure for an instance of the ArcGIS API
type ArcGIS struct {
	Authenticator Authenticator
	client        http.Client
	//ServiceRoot   string
	Context *string
	Host    string

	Usage Usage
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

func NewArcGIS(auth Authenticator) *ArcGIS {
	return &ArcGIS{
		Authenticator: auth,
		client:        http.Client{},
		Context:       nil,
		Host:          "https://www.arcgis.com",
		//ServiceRoot:   "https://www.arcgis.com/sharing/rest",
	}
}

func ServiceRootFromTenant(base string, tenantId string) string {
	return fmt.Sprintf("%s/%s", base, tenantId)
}

func (ag *ArcGIS) DoQuery(service string, layer_id uint, query *Query) (*QueryResult, error) {
	content, err := ag.DoQueryRaw(service, layer_id, query)
	if err != nil {
		return nil, err
	}
	return parseQueryResult(content)

}

func (ag *ArcGIS) DoQueryRaw(service string, layer_id uint, query *Query) ([]byte, error) {
	r, err := ag.query(fmt.Sprintf("/services/%s/FeatureServer/%d/query", service, layer_id), query)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %v", err)
	}
	return ag.requestJSON(r)
}

func (ag *ArcGIS) GetFeatureServer(service string) (*FeatureServer, error) {
	req, err := ag.serviceRequest(fmt.Sprintf("/services/%s/FeatureServer", service))
	if err != nil {
		return nil, err
	}
	content, err := ag.requestJSON(req)
	if err != nil {
		return nil, err
	}
	return parseFeatureServer(content)
}

func (ag *ArcGIS) PortalsSelf() (*PortalsResponse, error) {
	// We may need to always direct this request to
	// https://www.arcgis.com/sharing/rest/portals/self?f=json
	// not sure if hosted services are different
	req, err := ag.sharingRequest("/portals/self")
	if err != nil {
		return nil, err
	}
	content, err := ag.requestJSON(req)
	if err != nil {
		return nil, err
	}
	return parsePortalsResponse(content)
}

func (ag *ArcGIS) Search(query string) (*SearchResponse, error) {
	// "https://www.arcgis.com/sharing/rest/search?f=json&q=FieldseekerGIS"
	req, err := ag.sharingRequestWithParams("/search", map[string]string{
		"q": "FieldseekerGIS",
	})
	if err != nil {
		return nil, err
	}
	content, err := ag.requestJSON(req)
	if err != nil {
		return nil, err
	}
	return parseSearchResponse(content)
}
func (ag *ArcGIS) Services() (*ServiceInfo, error) {
	req, err := ag.serviceRequest("/services")
	if err != nil {
		return nil, err
	}
	content, err := ag.requestJSON(req)
	if err != nil {
		return nil, err
	}
	return parseServiceInfo(content)
}

func parseFeatureServer(data []byte) (*FeatureServer, error) {
	var result FeatureServer
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func parsePortalsResponse(data []byte) (*PortalsResponse, error) {
	var result PortalsResponse
	saveResponse(data, "portal.json")
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func parseQueryResult(data []byte) (*QueryResult, error) {
	var result QueryResult
	err := json.Unmarshal(data, &result)
	if err != nil {
		id := uuid.New()
		filename := fmt.Sprintf("debug_response_%s.json", id.String())
		output, err2 := os.Create(filename)
		if err2 != nil {
			log.Warn().Str("filename", filename).Msg("Failed to create debug file, can't dump request.")
			return nil, fmt.Errorf("Failed to parse query result JSON: %w", err)
		}
		defer output.Close()
		output.Write(data)
		log.Info().Str("filename", filename).Msg("Wrote response file for debugging.")
		return nil, fmt.Errorf("Failed to parse query result JSON: %w. Wrote debug file containing the request to %s", err, filename)
	}
	return &result, nil
}

func parseQueryResultCount(data []byte) (*QueryResultCount, error) {
	var result QueryResultCount
	err := json.Unmarshal(data, &result)
	//log.Println("Parsing", string(data))
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

func parseSearchResponse(data []byte) (*SearchResponse, error) {
	saveResponse(data, "search.json")
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

func saveResponse(data []byte, filename string) {
	dest, err := os.Create(filename)
	if err != nil {
		log.Error().Str("filename", filename).Str("err", err.Error()).Msg("Failed to create file")
		return
	}
	_, err = io.Copy(dest, bytes.NewReader(data))
	if err != nil {
		log.Error().Str("filename", filename).Str("err", err.Error()).Msg("Failed to write")
		return
	}
	log.Info().Str("filename", filename).Msg("Wrote response")
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

func (ag *ArcGIS) sharingRequest(endpoint string) (*http.Request, error) {
	u := fmt.Sprintf("%s%s", sharingBaseURL, endpoint)
	fullUrl, err := addParams(u, map[string]string{})
	if err != nil {
		return nil, fmt.Errorf("Failed to add params: %v", err)
	}
	return ag.serviceRequestFromFull(fullUrl)
}

func (ag *ArcGIS) sharingRequestWithParams(endpoint string, params map[string]string) (*http.Request, error) {
	u := fmt.Sprintf("%s%s", sharingBaseURL, endpoint)
	fullUrl, err := addParams(u, params)
	if err != nil {
		return nil, fmt.Errorf("Failed to add params: %v", err)
	}
	return ag.serviceRequestFromFull(fullUrl)
}

func (ag *ArcGIS) serviceUrl(endpoint string) string {
	//u := fmt.Sprintf("%s/%s/arcgis/rest%s", ag.ServiceRoot, ag.TenantId, endpoint)
	//u := fmt.Sprintf("%s%s", ag.ServiceRoot, endpoint)
	if ag.Context != nil {
		return fmt.Sprintf("%s/%s/arcgis/rest%s", ag.Host, *ag.Context, endpoint)
	} else {
		return fmt.Sprintf("%s/arcgis/rest%s", ag.Host, endpoint)
	}
}

func (ag *ArcGIS) serviceRequestFromFull(fullUrl *url.URL) (*http.Request, error) {
	req, err := http.NewRequest("GET", fullUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %v", err)
	}
	return ag.Authenticator.addAuthentication(req)
}

func (ag *ArcGIS) serviceRequest(endpoint string) (*http.Request, error) {
	u := ag.serviceUrl(endpoint)
	fullUrl, err := addParams(u, map[string]string{})
	if err != nil {
		return nil, fmt.Errorf("Failed to add params: %v", err)
	}
	return ag.serviceRequestFromFull(fullUrl)
}

func (ag *ArcGIS) serviceRequestWithParams(endpoint string, params map[string]string) (*http.Request, error) {
	u := ag.serviceUrl(endpoint)
	fullUrl, err := addParams(u, params)
	if err != nil {
		return nil, fmt.Errorf("Failed to add params: %v", err)
	}
	return ag.serviceRequestFromFull(fullUrl)
}

/*
func (arcgis ArcGIS) serviceUrl(endpoint string) (*url.URL, error) {
	u := fmt.Sprintf("%s/%s/arcgis/rest%s", arcgis.ServiceRoot, arcgis.TenantId, endpoint)
	base, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("f", "json")
	params.Add("token", arcgis.Token)
	base.RawQuery = params.Encode()
	return base, nil
}

func (arcgis ArcGIS) serviceUrlWithParams(endpoint string, params map[string]string) (*url.URL, error) {
	u := fmt.Sprintf("%s/%s/arcgis/rest%s", arcgis.ServiceRoot, arcgis.TenantId, endpoint)
	base, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	p := url.Values{}
	p.Add("f", "json")
	p.Add("token", arcgis.Token)
	for k, v := range params {
		p.Add(k, v)
	}
	base.RawQuery = p.Encode()
	return base, nil
}
*/

func logRequestBase(r *http.Request) {
	if r == nil {
		log.Warn().Msg("Can't log request), it's nil")
		return
	}

	// Create a copy of the URL to avoid modifying the original
	cleanURL := *r.URL

	// Remove query parameters
	//cleanURL.RawQuery = ""
	q, _ := url.ParseQuery(cleanURL.RawQuery)
	q.Del("token")
	cleanURL.RawQuery = q.Encode()

	//log.Info().Str("method", r.Method).Str("url", cleanURL.String()).Msg("ArcGIS request")
}

func (ag *ArcGIS) requestJSON(r *http.Request) ([]byte, error) {
	logRequestBase(r)
	resp, err := ag.client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("Failed to make request: %w", err)
	}
	//log.Printf("Status %v total bytes %v", resp.StatusCode, resp.ContentLength)
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return nil, fmt.Errorf("Client request error %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %w", err)
	}
	// During normal operation the ArcGIS server does _not_ respond with the standard
	// 400-level response codes on error. Instead it responds with 200, which is wrong,
	// but we can't force them to be standards-compliant. Instead, we have to attempt
	// to parse an error. If it works we have an error. If not, it must be something else.
	errorFromJSON := tryParseError(body)
	if errorFromJSON != nil {
		return nil, fmt.Errorf("response was an application-level error in JSON: %w", errorFromJSON)
	}
	ag.updateUsage(resp)
	return body, nil
}

func tryParseError(data []byte) error {
	var msg ErrorResponse
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return err
	}
	if msg.Error.Code != 0 || msg.Error.Message != "" || len(msg.Error.Details) > 0 {
		return msg.AsError()
	}
	return nil
}

func (ag *ArcGIS) updateUsage(resp *http.Response) {
	qru := resp.Header["X-Esri-Query-Request-Units"]
	for _, v := range qru {
		n, err := fmt.Sscanf(v, "%d", &ag.Usage.LastRequest)
		if err != nil {
			log.Warn().Str("err", err.Error()).Msg("Failed to parse X-Esri-Query-Request-Units")
		}
		if n < 1 {
			log.Warn().Msg("Parsed no values from X-Esri-Query-Request-Units")
		}
	}
	orupm := resp.Header["X-Esri-Org-Request-Units-Per-Min"]
	for _, v := range orupm {
		// The rupm value is of the form "usage=97;max=10000"
		n, err := fmt.Sscanf(v, "usage=%d;max=%d", &ag.Usage.ThisMinute, &ag.Usage.MaxPerMinute)
		if err != nil {
			log.Warn().Str("err", err.Error()).Msg("Failed to parse X-Esri-Org-Request-Units-Per-Min:")
		}
		if n < 2 {
			log.Warn().Msg("Parsed too few values from X-Esri-Org-Request-Per-Min")
		}
	}
}

func (ag *ArcGIS) Info() (*RestInfo, error) {
	r, err := ag.serviceRequest("/info")
	if err != nil {
		return nil, err
	}
	content, err := ag.requestJSON(r)
	if err != nil {
		return nil, err
	}
	return parseRestInfo(content)
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

func (arcgis ArcGIS) query(base string, query *Query) (*http.Request, error) {
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
	return arcgis.serviceRequestWithParams(base, params)
}

func (ag *ArcGIS) QueryCount(service string, layer_id uint) (*QueryResultCount, error) {
	params := make(map[string]string)
	params["returnCountOnly"] = "true"
	params["where"] = "9999=9999"
	r, err := ag.serviceRequestWithParams(fmt.Sprintf("/services/%s/FeatureServer/%d/query", service, layer_id), params)
	if err != nil {
		return nil, err
	}

	content, err := ag.requestJSON(r)
	if err != nil {
		return nil, err
	}
	return parseQueryResultCount(content)
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

func (ag *ArcGIS) WebhookList(serviceName string, serviceType ServiceType) ([]Webhook, error) {
	result := make([]Webhook, 0)

	u := fmt.Sprintf("%s/%s/ArcGIS/rest/admin/services/%s/%s/webhooks", ag.Host, *ag.Context, serviceName, ServiceTypeNames[serviceType])
	base, err := url.Parse(u)
	if err != nil {
		return result, err
	}
	params := url.Values{}
	params.Add("f", "json")
	base.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", base.String(), nil)
	if err != nil {
		return result, fmt.Errorf("Failed to create request: %v", err)
	}
	req, err = ag.Authenticator.addAuthentication(req)
	if err != nil {
		return result, fmt.Errorf("Failed to add authentication: %v", err)
	}
	content, err := ag.requestJSON(req)
	if err != nil {
		return result, fmt.Errorf("Failed to make request: %v", err)
	}
	resp, err := parseWebhookListResponse(content)
	if err != nil {
		return result, fmt.Errorf("Failed to parse JSON: %v", err)
	}
	return resp.Webhooks, nil
}

type WebhookListResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}

func parseWebhookListResponse(data []byte) (*WebhookListResponse, error) {
	var result WebhookListResponse
	saveResponse(data, "webhook-list.json")
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal JSON: %v", err)
	}
	return &result, nil
}
