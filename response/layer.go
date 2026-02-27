package response

type Layer struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	ParentLayerId     int    `json:"parentLayerId"`
	DefaultVisibility bool   `json:"defaultVisibility"`
	SubLayerIds       any    `json:"subLayerIds"`
	MinScale          int    `json:"minScale"`
	MaxScale          int    `json:"maxScale"`
	Type              string `json:"type"`
	GeometryType      string `json:"geometryType,omitempty"`
}
