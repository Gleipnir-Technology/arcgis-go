package response

type Geometry interface {
	String() string
	ToGeoJSON() (string, error)
	Type() string
}
