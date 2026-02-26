package response

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
