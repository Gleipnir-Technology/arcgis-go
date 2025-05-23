package arcgis

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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
	ServiceRoot string
	TenantId    string
	Token       string
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
type SpatialReference struct {
	WKID       int
	LatestWKID int
}

type Extent struct {
	XMin             float64
	YMin             float64
	XMax             float64
	YMax             float64
	SpatialReference SpatialReference
}

type Layer struct {
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
	Layers                         []Layer
	// many missing fields
}

// Query endpoint
type UniqueIdField struct {
	Name               string
	IsSystemMaintained bool
}

type Feature struct {
	Attributes map[string]any
}
type QueryResult struct {
	Features          []Feature
	ObjectIdFieldName string
	UniqueIdField     UniqueIdField
}

func parseFeatureServer(data []byte) (*FeatureServer, error) {
	var result FeatureServer
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

func parseServiceInfo(data []byte) (*ServiceInfo, error) {
	var result ServiceInfo
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

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

func logRequestBase(u *url.URL) {
	if u == nil {
		log.Println("URL is nil")
		return
	}

	// Create a copy of the URL to avoid modifying the original
	cleanURL := *u

	// Remove query parameters
	cleanURL.RawQuery = ""

	log.Printf("GET %s", cleanURL.String())

}

func requestJSON(u *url.URL) ([]byte, error) {
	logRequestBase(u)
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	log.Printf("Status %v total bytes %v", resp.StatusCode, resp.ContentLength)
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
	//for n, values := range resp.Header {
	//for _, value := range values {
	//fmt.Println(n, value)
	//}
	//}
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

func (arcgis ArcGIS) Info() (*RestInfo, error) {
	u, err := arcgis.serviceUrl("/info")
	if err != nil {
		return nil, err
	}
	content, err := requestJSON(u)
	if err != nil {
		return nil, err
	}
	return parseRestInfo(content)
}

func (arcgis ArcGIS) FeatureServer(service string) (*FeatureServer, error) {
	u, err := arcgis.serviceUrl(fmt.Sprintf("/services/%s/FeatureServer", service))
	if err != nil {
		return nil, err
	}
	content, err := requestJSON(u)
	if err != nil {
		return nil, err
	}
	return parseFeatureServer(content)
}

func (arcgis ArcGIS) Services() (*ServiceInfo, error) {
	u, err := arcgis.serviceUrl("/services")
	if err != nil {
		return nil, err
	}
	content, err := requestJSON(u)
	if err != nil {
		return nil, err
	}
	return parseServiceInfo(content)
}

type Query struct {
	Limit     int
	OutFields string
	Where     string
}

func NewQuery() *Query {
	q := new(Query)
	return q
}

func (arcgis ArcGIS) Query(service string, layer int, query *Query) (*QueryResult, error) {
	params := make(map[string]string)
	if query.Limit > 0 {
		params["limit"] = strconv.Itoa(query.Limit)
	}
	if query.OutFields != "" {
		params["outFields"] = query.OutFields
	}
	if query.Where != "" {
		params["where"] = query.Where
	}
	u, err := arcgis.serviceUrlWithParams(fmt.Sprintf("/services/%s/FeatureServer/%d/query", service, layer), params)
	if err != nil {
		return nil, err
	}

	content, err := requestJSON(u)
	if err != nil {
		return nil, err
	}
	return parseQueryResult(content)

}
