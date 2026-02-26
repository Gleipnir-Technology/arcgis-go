package response

import (
	"encoding/json"
)

type QueryResult struct {
	Features          []Feature        `json:"-"` // Handle manually
	Fields            []Field          `json:"-"` // Handle manually
	GeometryType      string           `json:"geometryType"`
	GlobalIDFieldName string           `json:"globalIdFieldName"`
	ObjectIDFieldName string           `json:"objectIdFieldName"`
	SpatialReference  SpatialReference `json:"spatialReference"`
	UniqueIDField     UniqueIdField    `json:"uniqueIdField"`
}

func (qr *QueryResult) UnmarshalJSON(data []byte) error {
	// First, unmarshal into a temporary struct to get the geometry type and field types
	type Alias QueryResult
	aux := &struct {
		Features []json.RawMessage `json:"features"`
		Fields   []Field           `json:"fields"`
		*Alias
	}{
		Alias: (*Alias)(qr),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	field_name_to_type := make(map[string]FieldType)
	for _, f := range aux.Fields {
		field_name_to_type[*f.Name] = f.Type
	}
	// Now unmarshal each feature with the known geometry type
	qr.Features = make([]Feature, len(aux.Features))
	for i, rawFeature := range aux.Features {
		if err := qr.Features[i].UnmarshalWithType(rawFeature, qr.GeometryType, field_name_to_type); err != nil {
			return err
		}
	}

	return nil
}
