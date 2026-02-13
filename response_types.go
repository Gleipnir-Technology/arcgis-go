package arcgis

import (
	"context"
	"encoding/json"
)

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

// Listing of available services
type ServiceListing struct {
	Name string
	Type string
	URL  string
}

type ResponseServiceInfo struct {
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

// Basic information about the REST API itself
type AuthInfo struct {
	isTokenBasedSecurity bool
	tokenServiceUrl      string
}

// Errors at the API level
type ErrorFromAPI struct {
	Code        int      `json:"code"`
	Details     []string `json:"details"`
	Error       string   `json:"error"`
	Description string   `json:"error_description"`
	Message     string   `json:"message"`
}

type ErrorResponse struct {
	Error ErrorFromAPI `json:"error"`
}

type ResourceInfo struct {
	CurrentVersion        float32 `json:"currentVersion"`
	MapName               string  `json:"mapName"`
	SupportsDynamicLayers bool    `json:"supportsDynamicLayers"`
	//Layers []LayerResource
	//Tables TableResource
	//SpatialReference SpatialReference `json:"spatialReference"`
	SingleFusedMapCache bool `json:"singleFusedMapCache"`
	//TileInfo TileInfo `json:"tileInfo"`
	//InitialExtent Extent `json:"initialExtent"`
	//FullExtent Extent `json:"fullExtent"`
}

type SpatialReference struct {
	LatestWKID int `json:"wkid"`
	WKID       int `json:"latestWkid"`
}

type RestInfo struct {
	CurrentVersion  float64
	FullVersion     string
	OwningSystemUrl string
	OwningTenant    string
	AuthInfo        AuthInfo
}

type SearchResponse struct {
	Total     int            `json:"total"`
	Start     int            `json:"start"`
	Num       int            `json:"num"`
	NextStart int            `json:"nextStart"`
	Results   []SearchResult `json:"results"`
}
type SearchResult struct {
	ID                      string                  `json:"id"`
	Owner                   string                  `json:"owner"`
	OrgID                   string                  `json:"orgId"`
	Created                 int64                   `json:"created"`
	IsOrgItem               bool                    `json:"isOrgItem"`
	Modified                int64                   `json:"modified"`
	Guid                    *string                 `json:"guid"`
	Name                    string                  `json:"name"`
	Title                   string                  `json:"title"`
	Type                    string                  `json:"type"`
	TypeKeywords            []string                `json:"typeKeywords"`
	Description             string                  `json:"description"`
	Tags                    []string                `json:"tags"`
	Snippet                 string                  `json:"snippet"`
	Thumbnail               string                  `json:"thumbnail"`
	Documentation           *string                 `json:"documentation"`
	Extent                  [][]float32             `json:"extent"`
	Categories              []string                `json:"categories"`
	SpatialReference        string                  `json:"spatialReference"`
	AccessInformation       string                  `json:"accessInformation"`
	Classification          *string                 `json:"classification"`
	LicenseInfo             string                  `json:"licenseInfo"`
	Culture                 string                  `json:"culture"`
	Properties              *map[string]interface{} `json:"properties"`
	AdvancedSettings        *string                 `json:"advancedSettings"`
	URL                     string                  `json:"url"`
	ProxyFilter             *string                 `json:"proxyFilter"`
	Access                  string                  `json:"access"`
	Size                    int                     `json:"size"`
	SubInfo                 int                     `json:"subInfo"`
	AppCategories           []string                `json:"appCategories"`
	Industries              []string                `json:"industries"`
	Languages               []string                `json:"languages"`
	LargeThumbnail          *string                 `json:"largeThumbnail"`
	Banner                  *string                 `json:"banner"`
	Screenshots             []string                `json:"screenshots"`
	Listed                  bool                    `json:"listed"`
	NumComments             int                     `json:"numComments"`
	NumRatings              int                     `json:"numRatings"`
	AvgRating               int                     `json:"avgRating"`
	NumViews                int                     `json:"numViews"`
	ScoreCompleteness       int                     `json:"scoreCompleteness"`
	GroupDesignations       *string                 `json:"groupDesignations"`
	ApiToken1ExpirationDate int                     `json:"apiToken1ExpirationDate"`
	ApiToken2ExpirationDate int                     `json:"apiToken2ExpirationDate"`
	LastViewed              int64                   `json:"lastViewed"`
}

type ServerURL struct {
	HTTPS []string `json:"https"`
}
type ServerURLCollection struct {
	Features  ServerURL `json:"features"`
	Insights  ServerURL `json:"insights"`
	Notebooks ServerURL `json:"notebooks"`
	Tiles     ServerURL `json:"tiles"`
}
type ResponseURLs struct {
	URLs ServerURLCollection `json:"urls"`
}

func (e ErrorResponse) AsError(ctx context.Context) apiError {
	return newAPIError(ctx, e)
}
