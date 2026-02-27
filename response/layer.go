package response

import (
	"encoding/json"
	"fmt"
)

type Layer struct {
	ID                StringOrNumber `json:"id"`
	Name              string         `json:"name"`
	ParentLayerId     int            `json:"parentLayerId"`
	DefaultVisibility bool           `json:"defaultVisibility"`
	SubLayerIds       any            `json:"subLayerIds"`
	MinScale          int            `json:"minScale"`
	MaxScale          int            `json:"maxScale"`
	Type              string         `json:"type"`
	GeometryType      string         `json:"geometryType,omitempty"`
}

type StringOrNumber string

func (son *StringOrNumber) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a string first
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*son = StringOrNumber(s)
		return nil
	}

	// Try as a number
	var n json.Number
	if err := json.Unmarshal(data, &n); err == nil {
		*son = StringOrNumber(n.String())
		return nil
	}

	return fmt.Errorf("value must be a string or a number")
}
