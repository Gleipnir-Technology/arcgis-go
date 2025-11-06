package arcgis

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// Errors at the API level
type ErrorFromAPI struct {
	Code    int
	Details []string
	Message string
}

type ErrorMessage struct {
	Error ErrorFromAPI
}

// Root structure for an instance of the ArcGIS API
type ArcGIS struct {
	Authenticator Authenticator
	client http.Client
	ServiceRoot string
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
	ID                int
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
	MaxRecordCount                 int
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

type Geometry struct {
	X float64
	Y float64
}

type Feature struct {
	Attributes map[string]any
	Geometry   Geometry
}

type CodeWrapper string

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
		client: http.Client{},
		ServiceRoot: "https://www.arcgis.com/sharing/rest",
	}
}

func ServiceRootFromTenant(base string, tenantId string) string {
	return fmt.Sprintf("%s/%s", base, tenantId)
}

func (ag *ArcGIS) DoQuery(service string, layer int, query *Query) (*QueryResult, error) {
	content, err := ag.DoQueryRaw(service, layer, query)
	if err != nil {
		return nil, err
	}
	return parseQueryResult(content)

}

func (ag *ArcGIS) DoQueryRaw(service string, layer int, query *Query) ([]byte, error) {
	r, err := ag.query(fmt.Sprintf("/services/%s/FeatureServer/%d/query", service, layer), query)
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
	req, err := ag.serviceRequest("/portals/self")
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
	baseURL := "https://www.arcgis.com/sharing/rest/search?q=FieldseekerGIS&f=pjson"
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		slog.Error("Failed to make request", slog.String("err", err.Error()))
		return nil, err
	}
	auth, ok := ag.Authenticator.(AuthenticatorOAuth)
	if !ok {
		slog.Error("Couldn't munch auth")
		return nil, errors.New("Bad auth munge")
	}
	req.Header.Add("X-ESRI-Authorization", "Bearer "+auth.AccessToken)
	resp, err := ag.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to do request: %v", err)
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	slog.Info("Search response", slog.Int("status", resp.StatusCode))
	/*
	params := make(map[string]string)
	params["q"] = "FieldseekerGIS"
	req, err := ag.serviceRequestWithParams("/search", params)
	if err != nil {
		return nil, err
	}
	content, err := ag.requestJSON(req)
	if err != nil {
		return nil, err
	}
	*/
	dest, err := os.Create("search.json")
	if err != nil {
		slog.Error("Failed to create search.json", slog.String("err", err.Error()))
		return nil, err
	}
	_, err = io.Copy(dest, bytes.NewReader(content))
	if err != nil {
		slog.Error("Failed to write search.json", slog.String("err", err.Error()))
		return nil, err
	}
	slog.Info("Wrote search.json")

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
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func parseQueryResult(data []byte) (*QueryResult, error) {
	var result QueryResult
	err := json.Unmarshal(data, &result)
	//log.Println("Parsing", string(data))
	if err != nil {
		return nil, err
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

func (ag *ArcGIS) serviceRequest(endpoint string) (*http.Request, error) {
	//u := fmt.Sprintf("%s/%s/arcgis/rest%s", ag.ServiceRoot, ag.TenantId, endpoint)
	u := fmt.Sprintf("%s%s", ag.ServiceRoot, endpoint)
	base, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("f", "json")
	base.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", base.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %v", err)
	}
	return ag.Authenticator.addAuthentication(req)
}

func (ag *ArcGIS) serviceRequestWithParams(endpoint string, params map[string]string) (*http.Request, error) {
	u := fmt.Sprintf("%s%s", ag.ServiceRoot, endpoint)
	base, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	p := url.Values{}
	p.Add("f", "json")
	for k, v := range params {
		p.Add(k, v)
	}
	base.RawQuery = p.Encode()
	req, err := http.NewRequest("GET", base.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %v", err)
	}
	return ag.Authenticator.addAuthentication(req)
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
		slog.Warn("Can't log request, it's nil")
		return
	}

	// Create a copy of the URL to avoid modifying the original
	cleanURL := *r.URL

	// Remove query parameters
	//cleanURL.RawQuery = ""
	q, _ := url.ParseQuery(cleanURL.RawQuery)
	q.Del("token")
	cleanURL.RawQuery = q.Encode()

	slog.Info("ArcGIS request", slog.String("method", r.Method), slog.String("url", cleanURL.String()))
}

func (ag *ArcGIS) requestJSON(r *http.Request) ([]byte, error) {
	logRequestBase(r)
	resp, err := ag.client.Do(r)
	if err != nil {
		return nil, err
	}
	//log.Printf("Status %v total bytes %v", resp.StatusCode, resp.ContentLength)
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return nil, fmt.Errorf("Client request error %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// During normal operation the ArcGIS server does _not_ respond with the standard
	// 400-level response codes on error. Instead it responds with 200, which is wrong,
	// but we can't force them to be standards-compliant. Instead, we have to attempt
	// to parse an error. If it works we have an error. If not, it must be something else.
	errorFromJSON := tryParseError(body)
	if errorFromJSON != nil {
		return nil, errorFromJSON
	}
	ag.updateUsage(resp)
	return body, nil
}

func tryParseError(data []byte) error {
	var msg ErrorMessage
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return err
	}
	if msg.Error.Code != 0 || msg.Error.Message != "" || len(msg.Error.Details) > 0 {
		return fmt.Errorf("ArcGIS API Error: %v", msg)
	}
	return nil
}

func (ag *ArcGIS) updateUsage(resp *http.Response) {
	qru := resp.Header["X-Esri-Query-Request-Units"]
	for _, v := range qru {
		n, err := fmt.Sscanf(v, "%d", &ag.Usage.LastRequest)
		if err != nil {
			slog.Warn("Failed to parse X-Esri-Query-Request-Units", slog.String("err", err.Error()))
		}
		if n < 1 {
			slog.Warn("Parsed no values from X-Esri-Query-Request-Units")
		}
	}
	orupm := resp.Header["X-Esri-Org-Request-Units-Per-Min"]
	for _, v := range orupm {
		// The rupm value is of the form "usage=97;max=10000"
		n, err := fmt.Sscanf(v, "usage=%d;max=%d", &ag.Usage.ThisMinute, &ag.Usage.MaxPerMinute)
		if err != nil {
			slog.Warn("Failed to parse X-Esri-Org-Request-Units-Per-Min:", slog.String("err", err.Error()))
		}
		if n < 2 {
			slog.Warn("Parsed too few values from X-Esri-Org-Request-Per-Min")
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
	ResultRecordCount int
	ResultOffset      int
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
		params["resultOffset"] = strconv.Itoa(query.ResultOffset)
	}
	if query.Where != "" {
		params["where"] = query.Where
	}
	if len(query.SpatialReference) > 0 {
		params["outSR"] = query.SpatialReference
	}
	return arcgis.serviceRequestWithParams(base, params)
}

func (ag *ArcGIS) QueryCount(service string, layer int) (*QueryResultCount, error) {
	params := make(map[string]string)
	params["returnCountOnly"] = "true"
	params["where"] = "9999=9999"
	r, err := ag.serviceRequestWithParams(fmt.Sprintf("/services/%s/FeatureServer/%d/query", service, layer), params)
	if err != nil {
		return nil, err
	}

	content, err := ag.requestJSON(r)
	if err != nil {
		return nil, err
	}
	return parseQueryResultCount(content)
}

func (c *CodeWrapper) UnmarshalJSON(data []byte) (err error) {
	// Does it look like a string?
	var content string
	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		content = string(data[1 : len(data)-1])
	} else {
		if data[0] == 0 {
			content = "0"
		} else {
			content = "1"
		}
	}

	*c = CodeWrapper(string(content))
	return nil
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
