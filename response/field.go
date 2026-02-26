package response

import (
	"encoding/json"
	"fmt"
)

// https://developers.arcgis.com/net/api-reference/api/uwp/Esri.ArcGISRuntime/Esri.ArcGISRuntime.Data.FieldType.html
type FieldType int

const (
	FieldTypeBlob FieldType = iota
	FieldTypeDate
	FieldTypeDateOnly
	FieldTypeFloat32
	FieldTypeFloat64
	FieldTypeGeometry
	FieldTypeGlobalID
	FieldTypeGuid
	FieldTypeIntegerSmall
	FieldTypeInteger
	FieldTypeInt64
	FieldTypeOID
	FieldTypeRaster
	FieldTypeString
	FieldTypeText
	FieldTypeTimeOnly
	FieldTypeTimestampOffset
	FieldTypeXml
)

// Field represents a field in the layer or table
type Field struct {
	Name         *string   `json:"name,omitempty"`
	Type         FieldType `json:"type,omitempty"`
	Alias        *string   `json:"alias,omitempty"`
	Domain       any       `json:"domain,omitempty"`
	Editable     *bool     `json:"editable,omitempty"`
	Nullable     *bool     `json:"nullable,omitempty"`
	Length       *int      `json:"length,omitempty"`
	DefaultValue any       `json:"defaultValue,omitempty"`
	ModelName    *string   `json:"modelName,omitempty"`
}

func (f *Field) UnmarshalJSON(data []byte) error {
	type Alias Field
	aux := &struct {
		Type *string `json:"type,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	if aux.Type == nil {
		return fmt.Errorf("nil field type")
	}
	switch *aux.Type {
	case "esriFieldTypeBlob":
		f.Type = FieldTypeBlob
	case "esriFieldTypeDate":
		f.Type = FieldTypeDate
	case "esriFieldTypeDateOnly":
		f.Type = FieldTypeDateOnly
	case "esriFieldTypeDouble":
		f.Type = FieldTypeFloat64
	case "esriFieldTypeFloat32":
		f.Type = FieldTypeFloat32
	case "esriFieldTypeFloat64":
		f.Type = FieldTypeFloat64
	case "esriFieldTypeGeometry":
		f.Type = FieldTypeGeometry
	case "esriFieldTypeGlobalID":
		f.Type = FieldTypeGlobalID
	case "esriFieldTypeGuid":
		f.Type = FieldTypeGuid
	case "esriFieldTypeInteger":
		f.Type = FieldTypeInteger
	case "esriFieldTypeInt64":
		f.Type = FieldTypeInt64
	case "esriFieldTypeOID":
		f.Type = FieldTypeOID
	case "esriFieldTypeRaster":
		f.Type = FieldTypeRaster
	case "esriFieldTypeSingle":
		f.Type = FieldTypeFloat32
	case "esriFieldTypeSmallInteger":
		f.Type = FieldTypeIntegerSmall
	case "esriFieldTypeString":
		f.Type = FieldTypeString
	case "esriFieldTypeText":
		f.Type = FieldTypeText
	case "esriFieldTypeTimeOnly":
		f.Type = FieldTypeTimeOnly
	case "esriFieldTypeTimestampOffset":
		f.Type = FieldTypeTimestampOffset
	case "esriFieldTypeXml":
		f.Type = FieldTypeXml
	default:
		return fmt.Errorf("Unrecognized esri field type '%s'", *aux.Type)
	}
	return nil
}
