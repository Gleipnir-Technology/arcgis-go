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

	"github.com/Gleipnir-Technology/arcgis-go/response"
	"github.com/rs/zerolog"
)

// Root structure for an instance of the ArcGIS API
type ArcGIS struct {
	AccountID string

	requestor gisRequestor

	urlFeatures  string
	urlInsights  string
	urlNotebooks string
	urlTiles     string

	usage Usage
}

type Usage struct {
	// Units used on the last request
	LastRequest int
	// The maximum units we can use in a minute
	MaxPerMinute int
	// The amount used in the current minute
	ThisMinute int
}

func NewArcGIS(ctx context.Context) (*ArcGIS, error) {
	var username = os.Getenv("ARCGIS_USERNAME")
	if username == "" {
		return nil, fmt.Errorf("no username")
	}
	var password = os.Getenv("ARCGIS_PASSWORD")
	if password == "" {
		return nil, fmt.Errorf("no password")
	}
	auth := AuthenticatorUsernamePassword{
		Password: password,
		Username: username,
	}
	return NewArcGISAuth(ctx, &auth)
}

func NewArcGISAuth(ctx context.Context, auth Authenticator) (*ArcGIS, error) {
	var err error
	logger := zerolog.Ctx(ctx)
	var mitm_proxy = os.Getenv("MITM_PROXY")
	transport := &http.Transport{}
	if mitm_proxy != "" {
		transport, err = MITMProxyTransport()
		if err != nil {
			return nil, fmt.Errorf("create mitm proxy: %w", err)
		}
		logger.Warn().Msg("Using MITM proxy for debugging ArcGIS")
	}
	return NewArcGISTransport(ctx, nil, auth, transport)
}
func NewArcGISTransport(ctx context.Context, host *string, auth Authenticator, transport *http.Transport) (*ArcGIS, error) {
	if host == nil {
		h := os.Getenv("ARCGIS_BASE")
		if h == "" {
			h = "https://www.arcgis.com"
		}
		host = &h
	}
	h := "https://www.arcgis.com"
	if host != nil {
		h = *host
	}
	requestor, err := newGisRequestor(ctx, auth, h, transport)
	if err != nil {
		return nil, fmt.Errorf("create requestor: %w", err)
	}
	result := &ArcGIS{
		requestor: requestor,
		//ServiceRoot:   "https://www.arcgis.com/sharing/rest",
	}
	err = result.switchHostByPortal(ctx)
	if err != nil {
		return nil, fmt.Errorf("switch portal: %w", err)
	}
	err = result.populateURLs(ctx)
	if err != nil {
		return nil, fmt.Errorf("populate urls: %w", err)
	}
	return result, nil
}
func ServiceRootFromTenant(base string, tenantId string) string {
	return fmt.Sprintf("%s/%s", base, tenantId)
}

func (ag *ArcGIS) Query(ctx context.Context, service string, layer_id uint, query *Query) (*QueryResult, error) {
	path := fmt.Sprintf("/arcgis/rest/services/%s/FeatureServer/%d/query", service, layer_id)
	url, err := ag.urlFeature(path)
	if err != nil {
		return nil, fmt.Errorf("make url: %w", err)
	}
	params := query.toParams()
	return reqGetJSONParamsHeadersFullURL[QueryResult](ctx, ag.requestor, *url, params, map[string]string{})
}
func (ag *ArcGIS) QueryRaw(ctx context.Context, service string, layer_id uint, query *Query) ([]byte, error) {
	// path := fmt.Sprintf("/services/%s/FeatureServer/%d/query?f=json", service, layer_id)
	path := fmt.Sprintf("/arcgis/rest/services/%s/FeatureServer/%d/query?f=json", service, layer_id)
	url, err := ag.urlFeature(path)
	if err != nil {
		return nil, fmt.Errorf("make url: %w", err)
	}
	return ag.requestor.doGetParamsHeadersFullURL(ctx, *url, map[string]string{}, map[string]string{})
}

type AdminInfo struct {
	CurrentVersion string `json:"currentversion"`
}

func (ag *ArcGIS) AdminInfo(ctx context.Context, serviceName string, serviceType ServiceType) (*AdminInfo, error) {
	// We may need to always direct this request to
	path := fmt.Sprintf("/arcgis/rest/admin/services/%s/%s/permissions", serviceName, ServiceTypeNames[serviceType])
	url, err := ag.urlFeature(path)
	if err != nil {
		return nil, fmt.Errorf("make url: %w", err)
	}
	return reqGetJSONFullURL[AdminInfo](ctx, ag.requestor, *url)
}

func (ag *ArcGIS) NewServiceFeature(ctx context.Context, name string, url url.URL) (*ServiceFeature, error) {
	return newServiceFeature(ctx, name, url, ag.requestor)
}
func (ag *ArcGIS) MapServices(ctx context.Context) ([]MapService, error) {
	logger := zerolog.Ctx(ctx)
	resp, err := ag.SearchInAccount(ctx, "type:\"Map Service\"")
	if err != nil {
		return nil, fmt.Errorf("search err: %w", err)
	}
	logger.Debug().Int("total", resp.Total).Msg("got results")
	results := make([]MapService, 0)
	for _, r := range resp.Results {
		if r.Type != "Map Service" {
			logger.Warn().Str("type", r.Type).Msg("Got the wrong type for a map service")
			continue
		}
		u, err := url.Parse(r.URL)
		if err != nil {
			logger.Error().Err(err).Msg("map service url parse error")
			continue
		}
		m := MapService{
			ID:    r.ID,
			Name:  r.Name,
			Title: r.Title,
			URL:   *u,

			meta: nil,
		}
		results = append(results, m)
	}
	return results, nil
}

var globalBaseURL string = "https://www.arcgis.com/"

func (ag *ArcGIS) PortalsGlobal(ctx context.Context) (*response.Portal, error) {
	// So, this is a bit nuts. Bear with me.
	// There is a special endpoint at GET https://www.arcgis.com/sharing/rest/portals/self?f=json
	req_url, err := url.Parse(globalBaseURL + "/sharing/rest/portals")
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}
	return reqGetJSONParamsHeadersFullURL[response.Portal](ctx, ag.requestor, *req_url, map[string]string{}, map[string]string{})
}
func (ag *ArcGIS) PortalsSelf(ctx context.Context) (*response.Portal, error) {
	// We may need to always direct this request to
	//
	// not sure if hosted services are different
	//
	// GET https://<urlkey>.maps.arcgis.com/sharing/rest/portals/self/urls?f=json
	// seems to also work, ond may give different data.
	return reqGetJSON[response.Portal](ctx, ag.requestor, "/sharing/rest/portals/self")
}
func (ag *ArcGIS) Search(ctx context.Context, query string) (*SearchResponse, error) {
	return reqPostFormToJSON[SearchResponse](ctx, ag.requestor, "/sharing/rest/search", map[string]string{
		"f":     "json",
		"q":     query,
		"start": "1",
		"num":   "100",
	})
}
func (ag *ArcGIS) SearchInAccount(ctx context.Context, query string) (*SearchResponse, error) {
	params := map[string]string{
		"f":         "json",
		"q":         fmt.Sprintf("%s accountid:%s", query, ag.AccountID),
		"start":     "1",
		"num":       "100",
		"sortField": "avgRating",
		"sortOrder": "desc",
	}
	return reqPostFormToJSON[SearchResponse](ctx, ag.requestor, "/sharing/rest/search", params)
}
func (ag *ArcGIS) Services(ctx context.Context) ([]*ServiceFeature, error) {
	//return reqGetJSON[ServiceInfo](ctx, ag.requestor, org_path)
	u, err := ag.urlFeature("/arcgis/rest/services")
	if err != nil {
		return nil, fmt.Errorf("make url: %w", err)
	}
	resp, err := reqGetJSONParamsHeadersFullURL[ResponseServiceInfo](ctx, ag.requestor, *u, map[string]string{}, map[string]string{})
	if err != nil {
		return nil, err
	}

	logger := zerolog.Ctx(ctx)
	results := make([]*ServiceFeature, 0)
	for _, s := range resp.Services {
		logger.Info().Str("name", s.Name).Str("type", s.Type).Str("url", s.URL).Msg("service")

		u, err := url.Parse(s.URL)
		if err != nil {
			return results, fmt.Errorf("parse url: %w", err)
		}
		sf, err := newServiceFeature(ctx, s.Name, *u, ag.requestor)
		if err != nil {
			return results, fmt.Errorf("new service feature: %w", err)
		}
		results = append(results, sf)
	}
	return results, nil
}
func (ag *ArcGIS) ServiceByName(ctx context.Context, name string) (*ServiceFeature, error) {
	logger := zerolog.Ctx(ctx)
	query := fmt.Sprintf("type:\"Feature Service\" name:\"%s\"", name)
	resp, err := ag.SearchInAccount(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("search err: %w", err)
	}
	logger.Debug().Int("total", resp.Total).Str("query", query).Msg("got results")
	for _, result := range resp.Results {
		if result.Type == "Feature Service" && result.Name == name {
			u, err := url.Parse(result.URL)
			if err != nil {
				return nil, fmt.Errorf("parse url: %w", err)
			}
			r, err := newServiceFeature(ctx, name, *u, ag.requestor)
			if err != nil {
				return nil, fmt.Errorf("creating service: %w", err)
			}
			return r, nil
		}
	}
	return nil, fmt.Errorf("no matching results")
}
func (ag *ArcGIS) populateURLs(ctx context.Context) error {
	logger := zerolog.Ctx(ctx)
	path := "/sharing/rest/portals/self/urls"
	resp, err := reqGetJSON[ResponseURLs](ctx, ag.requestor, path)
	if err != nil {
		return fmt.Errorf("get urls: %w", err)
	}
	ag.urlFeatures = resp.URLs.Features.HTTPS[0]
	ag.urlInsights = resp.URLs.Insights.HTTPS[0]
	ag.urlNotebooks = resp.URLs.Notebooks.HTTPS[0]
	ag.urlTiles = resp.URLs.Tiles.HTTPS[0]
	logger.Info().Str("feature-url", ag.urlFeatures).Msg("Populated URLs")
	return nil
}
func (ag *ArcGIS) switchHostByPortal(ctx context.Context) error {
	logger := zerolog.Ctx(ctx)
	portals, err := ag.PortalsSelf(ctx)
	if err != nil {
		return fmt.Errorf("Failed to get portals: %w", err)
	} else if portals == nil {
		return errors.New("Returned portals was nil")
	}
	ag.AccountID = portals.ID
	ag.requestor.host = fmt.Sprintf("https://%s.maps.arcgis.com", portals.UrlKey)
	logger.Debug().Str("id", portals.ID).Str("name", portals.PortalName).Str("urlkey", portals.UrlKey).Str("host", ag.requestor.host).Msg("Switched host by portal")
	return nil
}
func (ag *ArcGIS) urlFeature(path string) (*url.URL, error) {
	return url.Parse(fmt.Sprintf("https://%s/%s%s", ag.urlFeatures, ag.AccountID, path))
}

func saveResponse(ctx context.Context, data []byte, filename string) {
	logger := zerolog.Ctx(ctx)
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
func (ag *ArcGIS) updateUsage(ctx context.Context, resp *http.Response) {
	logger := zerolog.Ctx(ctx)
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
	return reqGetJSON[RestInfo](ctx, ag.requestor, "/sharing/rest/info")
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

func (ag *ArcGIS) PermissionList(ctx context.Context, serviceName string, serviceType ServiceType) (*response.PermissionSlice, error) {
	path := fmt.Sprintf("/arcgis/rest/admin/services/%s/%s/permissions", serviceName, ServiceTypeNames[serviceType])
	url, err := ag.urlFeature(path)
	if err != nil {
		return nil, fmt.Errorf("make url: %w", err)
	}
	return reqGetJSONParamsHeadersFullURL[response.PermissionSlice](ctx, ag.requestor, *url, map[string]string{}, map[string]string{})
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
	path := fmt.Sprintf("/arcgis/rest/admin/services/%s/%s/webhooks", serviceName, ServiceTypeNames[serviceType])
	return reqGetJSON[WebhookSlice](ctx, ag.requestor, path)
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
