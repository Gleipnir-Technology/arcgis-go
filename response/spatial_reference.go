package response

// SpatialReference defines the spatial reference system
type SpatialReference struct {
	WKID          *int     `json:"wkid,omitempty"`
	LatestWKID    *int     `json:"latestWkid,omitempty"`
	VCSWKID       *int     `json:"vcsWkid,omitempty"`
	LatestVCSWKID *int     `json:"latestVcsWkid,omitempty"`
	XYTolerance   *float64 `json:"xyTolerance,omitempty"`
	ZTolerance    *float64 `json:"zTolerance,omitempty"`
	MTolerance    *float64 `json:"mTolerance,omitempty"`
	FalseX        *float64 `json:"falseX,omitempty"`
	FalseY        *float64 `json:"falseY,omitempty"`
	XYUnits       *float64 `json:"xyUnits,omitempty"`
	FalseZ        *float64 `json:"falseZ,omitempty"`
	ZUnits        *float64 `json:"zUnits,omitempty"`
	FalseM        *float64 `json:"falseM,omitempty"`
	MUnits        *float64 `json:"mUnits,omitempty"`
}
