package response

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
