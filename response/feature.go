package response

import (
	"encoding/json"
)

type Feature struct {
	Attributes map[string]any  `json:"attributes"`
	Geometry   json.RawMessage `json:"geometry"`
}
