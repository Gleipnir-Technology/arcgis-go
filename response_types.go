package arcgis

import (
	"context"
	"encoding/json"
)

type AuthInfo struct {
	isTokenBasedSecurity bool
	tokenServiceUrl      string
}
type DefaultValueWrapper string
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
type Envelope struct {
	SpatialReference SpatialReference `json:"spatialReference"`
	XMax             float64          `json:"xmax"`
	YMax             float64          `json:"ymax"`
	XMin             float64          `json:"xmin"`
	YMin             float64          `json:"ymin"`
}
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

func (e ErrorResponse) AsError(ctx context.Context) apiError {
	return newAPIError(ctx, e)
}

type Feature struct {
	Attributes map[string]any
	Geometry   json.RawMessage
}
type Field struct {
	Alias        string
	DefaultValue *DefaultValueWrapper
	Domain       *Domain
	Length       int
	Name         string
	SQLType      string
	Type         string
}
type LayerFeature struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	DefaultVisibility bool    `json:"defaultVisibility"`
	ParentLayerID     int     `json:"parentLayerId"`
	SubLayerIds       []int   `json:"subLayerIds"`
	MinScale          float64 `json:"minScale"`
	MaxScale          float64 `json:"maxScale"`
	Type              string
	GeometryType      string
}
type LOD struct {
	Level      int     `json:"level"`
	Resolution float64 `json:"resolution"`
	Scale      float64 `json:"scale"`
}
type MapServiceMetadata struct {
	CurrentVersion            float64           `json:"currentVersion"`
	ServiceDescription        string            `json:"serviceDescription"`
	MapName                   string            `json:"mapName"`
	Description               string            `json:"description"`
	CopyrightText             string            `json:"copyrightText"`
	SupportsDynamicLayers     bool              `json:"supportsDynamicLayers"`
	Layers                    []LayerFeature    `json:"layers"`
	Tables                    []Table           `json:"tables"`
	SpatialReference          SpatialReference  `json:"spatialReference"`
	SingleFusedMapCache       bool              `json:"singleFusedMapCache"`
	TileInfo                  *TileInfo         `json:"tileInfo,omitempty"`
	StorageInfo               *StorageInfo      `json:"storageInfo,omitempty"`
	InitialExtent             Envelope          `json:"initialExtent"`
	FullExtent                Envelope          `json:"fullExtent"`
	DatesInUnknownTimezone    bool              `json:"datesInUnknownTimezone"`
	TimeInfo                  *TimeInfo         `json:"timeInfo,omitempty"`
	Units                     string            `json:"units"`
	SupportedImageFormatTypes string            `json:"supportedImageFormatTypes"`
	DocumentInfo              map[string]string `json:"documentInfo"`
	Capabilities              string            `json:"capabilities"`
	MaxRecordCount            int               `json:"maxRecordCount"`
	MaxImageHeight            int               `json:"maxImageHeight"`
	MaxImageWidth             int               `json:"maxImageWidth"`
	MinScale                  float64           `json:"minScale"`
	MaxScale                  float64           `json:"maxScale"`
	TileServers               []string          `json:"tileServers"`
	SupportedQueryFormats     string            `json:"supportedQueryFormats"`
	ExportTilesAllowed        bool              `json:"exportTilesAllowed"`
	MaxExportTilesCount       int               `json:"maxExportTilesCount"`
	SupportedExtensions       string            `json:"supportedExtensions"`
	Resampling                bool              `json:"resampling"`
}
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
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
type ResponseServiceInfo struct {
	CurrentVersion float64
	Services       []ServiceListing
}
type RestInfo struct {
	CurrentVersion  float64
	FullVersion     string
	OwningSystemUrl string
	OwningTenant    string
	AuthInfo        AuthInfo
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
type ResponseURLs struct {
	URLs ServerURLCollection `json:"urls"`
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
type ServiceListing struct {
	Name string
	Type string
	URL  string
}
type SpatialReference struct {
	LatestWKID int `json:"latestWkid"`
	WKID       int `json:"wkid"`
}

type StorageInfo struct {
	StorageFormat string `json:"storageFormat"`
	PacketSize    int    `json:"packetSize"`
}
type Table struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	ParentLayerID     int
	DefaultVisibility bool
	SubLayerIDs       *string
	MinScale          int
	MaxScale          int
}
type TileInfo struct {
	Rows               int              `json:"rows"`
	Cols               int              `json:"cols"`
	DPI                int              `json:"dpi"`
	Format             string           `json:"format"`
	CompressionQuality int              `json:"compressionQuality"`
	Origin             Point            `json:"origin"`
	SpatialReference   SpatialReference `json:"spatialReference"`
	LODs               []LOD            `json:"lods"`
}
type TimeInfo struct {
	TimeExtent               []int64        `json:"timeExtent"`
	TimeReference            *TimeReference `json:"timeReference,omitempty"`
	TimeRelation             string         `json:"timeRelation"`
	DefaultTimeInterval      int            `json:"defaultTimeInterval"`
	DefaultTimeIntervalUnits string         `json:"defaultTimeIntervalUnits"`
	DefaultTimeWindow        int            `json:"defaultTimeWindow"`
	HasLiveData              bool           `json:"hasLiveData"`
	LiveModeOffsetDirection  string         `json:"liveModeOffsetDirection"`
}
type TimeReference struct {
	TimeZone               string `json:"timeZone"`
	RespectsDaylightSaving bool   `json:"respectsDaylightSaving"`
}
type UniqueIdField struct {
	Name               string
	IsSystemMaintained bool
}
