package response

type LOD struct {
	Level      int     `json:"level"`
	Resolution float64 `json:"resolution"`
	Scale      float64 `json:"scale"`
}
