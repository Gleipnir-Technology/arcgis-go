package response

type Table struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	ParentLayerID     int
	DefaultVisibility bool
	SubLayerIDs       *string
	MinScale          int
	MaxScale          int
}
