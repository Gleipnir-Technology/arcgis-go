package response

type Extent struct {
	MMin             *float64 `json:"mmin,omitempty"`
	MMax             *float64 `json:"mmax,omitempty"`
	XMin             float64  `json:"xmin"`
	YMin             float64  `json:"ymin"`
	XMax             float64  `json:"xmax"`
	YMax             float64  `json:"ymax"`
	ZMin             *float64 `json:"zmin,omitempty"`
	ZMax             *float64 `json:"zmax,omitempty"`
	SpatialReference SpatialReference
}
