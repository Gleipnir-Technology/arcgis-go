package response

type Extent struct {
	XMin             float64 `json:"xmin"`
	YMin             float64 `json:"ymin"`
	XMax             float64 `json:"xmax"`
	YMax             float64 `json:"ymax"`
	SpatialReference SpatialReference
}
