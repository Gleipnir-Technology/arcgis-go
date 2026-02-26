package response

type Polyline struct {
	Paths [][][]float64 `json:"paths"`
}

func (p Polyline) String() string { return "some polyline" }
func (p Polyline) Type() string   { return "LineString" }
func (p Polyline) ToGeoJSON() map[string]any {
	return map[string]any{
		"type":        "LineString",
		"coordinates": p.Paths,
	}
}
