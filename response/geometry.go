package response

type Geometry interface {
	String() string
	ToGeoJSON() map[string]any
	Type() string
}
