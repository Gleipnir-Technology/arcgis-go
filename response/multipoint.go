package response

type Multipoint2D struct {
	HasZ             bool             `json:"hasZ"`
	HasM             bool             `json:"hasM"`
	IDs              []int            `json:"ids"`
	Points           [][2]float64     `json:"points"`
	SpatialReference SpatialReference `json:"spatialReference"`
}

func (p Multipoint2D) Type() string { return "esriGeometryMultipoint" }
func (p Multipoint2D) ToGeoJSON() map[string]any {
	return map[string]any{
		"type": "LineString",
		//"coordinates": p.Paths,
	}
}
