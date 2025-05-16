package arcgis

import (
	"encoding/hex"
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
	LatestWKID int
	WKID       int
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
	Layers                         []Layer
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

var ag *ArcGIS

func DoQuery(service string, layer int, query *Query) (*QueryResult, error) {
	content, err := DoQueryRaw(service, layer, query)
	if err != nil {
		return nil, err
	}
	return parseQueryResult(content)

}

func DoQueryRaw(service string, layer int, query *Query) ([]byte, error) {
	u, err := ag.queryURL(fmt.Sprintf("/services/%s/FeatureServer/%d/query", service, layer), query)
	if err != nil {
		return nil, err
	}

	return requestJSON(u)
}

func GetFeatureServer(service string) (*FeatureServer, error) {
	u, err := ag.serviceUrl(fmt.Sprintf("/services/%s/FeatureServer", service))
	if err != nil {
		return nil, err
	}
	content, err := requestJSON(u)
	if err != nil {
		return nil, err
	}
	return parseFeatureServer(content)
}

func Initialize(service_root string, tenant_id string, token string) {
	ag = new(ArcGIS)
	ag.ServiceRoot = service_root
	ag.TenantId = tenant_id
	ag.Token = token
}

func Services() (*ServiceInfo, error) {
	u, err := ag.serviceUrl("/services")
	if err != nil {
		return nil, err
	}
	content, err := requestJSON(u)
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

type Query struct {
	Limit             int
	ObjectIDs         string
	OutFields         string
	ResultRecordCount int
	ResultOffset      int
	Where             string
}

func NewQuery() *Query {
	q := new(Query)
	return q
}

func (arcgis ArcGIS) queryURL(base string, query *Query) (*url.URL, error) {
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
	if query.Where != "" {
		params["where"] = query.Where
	}
	return arcgis.serviceUrlWithParams(base, params)
}

func QueryCount(service string, layer int) (*QueryResultCount, error) {
	params := make(map[string]string)
	params["returnCountOnly"] = "true"
	params["where"] = "9999=9999"
	u, err := ag.serviceUrlWithParams(fmt.Sprintf("/services/%s/FeatureServer/%d/query", service, layer), params)
	if err != nil {
		return nil, err
	}

	content, err := requestJSON(u)
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
