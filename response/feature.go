package response

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

type AttributeValue interface {
	String() string
}

// Concrete types for each field type
type TextValue struct{ V string }

func (t TextValue) String() string { return t.V }
func (t TextValue) Value() any     { return t.V }

type Int16Value struct{ V int16 }

func (i Int16Value) String() string { return strconv.FormatInt(int64(i.V), 10) }
func (i Int16Value) Value() any     { return i.V }

type Int32Value struct{ V int32 }

func (i Int32Value) String() string { return strconv.FormatInt(int64(i.V), 10) }
func (i Int32Value) Value() any     { return i.V }

type Float32Value struct{ V float32 }

func (f Float32Value) String() string { return strconv.FormatFloat(float64(f.V), 'f', -1, 32) }
func (f Float32Value) Value() any     { return f.V }

type Float64Value struct{ V float64 }

func (f Float64Value) String() string { return strconv.FormatFloat(f.V, 'f', -1, 64) }
func (f Float64Value) Value() any     { return f.V }

type DateValue struct{ V time.Time }

func (d DateValue) String() string { return d.V.Format(time.RFC3339) }
func (d DateValue) Value() any     { return d.V }

type Feature struct {
	Attributes map[string]AttributeValue `json:"-"` // Handle manually
	Geometry   Geometry                  `json:"-"` // Handle manually
}

func (f *Feature) UnmarshalWithType(data []byte, geometryType string, fieldNameToType map[string]FieldType) error {
	aux := &struct {
		Attributes map[string]any  `json:"attributes"`
		Geometry   json.RawMessage `json:"geometry"`
	}{}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// Unmarshal geometry based on type
	var err error
	switch geometryType {
	case "esriGeometryPoint":
		var geom Point
		err = json.Unmarshal(aux.Geometry, &geom)
		f.Geometry = geom
	case "esriGeometryPolyline":
		var geom Polyline
		err = json.Unmarshal(aux.Geometry, &geom)
		f.Geometry = geom
	case "esriGeometryPolygon":
		var geom Polygon2D
		err = json.Unmarshal(aux.Geometry, &geom)
		f.Geometry = geom
	// ... other cases
	default:
		return fmt.Errorf("unknown geometry type: %s", geometryType)
	}

	// Unmarshal attributes based on type
	f.Attributes = make(map[string]AttributeValue, len(aux.Attributes))
	for attrName, rawValue := range aux.Attributes {
		typ, ok := fieldNameToType[attrName]
		if !ok {
			return fmt.Errorf("lookup field name to type")
		}

		typedValue, err := convertToTypedValue(rawValue, typ)
		if err != nil {
			return fmt.Errorf("converting %s: %w", attrName, err)
		}
		f.Attributes[attrName] = typedValue
		if typedValue == nil {
			log.Warn().Str("name", attrName).Msg("set to nil")
		}
	}

	return err
}

func convertToTypedValue(raw any, fieldType FieldType) (AttributeValue, error) {
	if raw == nil {
		return NullValue{}, nil
	}

	switch fieldType {
	case FieldTypeText:
		if v, ok := raw.(string); ok {
			return TextValue{V: v}, nil
		}
		return nil, fmt.Errorf("expected string, got %T", raw)

	case FieldTypeIntegerSmall:
		// JSON numbers come as float64
		if v, ok := raw.(float64); ok {
			return Int16Value{V: int16(v)}, nil
		}
		return nil, fmt.Errorf("expected number, got %T", raw)

	case FieldTypeInteger:
		// JSON numbers come as float64
		if v, ok := raw.(float64); ok {
			return Int32Value{V: int32(v)}, nil
		}
		return nil, fmt.Errorf("expected number, got %T", raw)

	case FieldTypeFloat32:
		if v, ok := raw.(float32); ok {
			return Float32Value{V: v}, nil
		}
		return nil, fmt.Errorf("expected number, got %T", raw)
	case FieldTypeFloat64:
		if v, ok := raw.(float64); ok {
			return Float64Value{V: v}, nil
		}
		return nil, fmt.Errorf("expected number, got %T", raw)

	case FieldTypeDate:
		// Assuming epoch milliseconds or ISO string
		switch v := raw.(type) {
		case float64:
			return DateValue{V: time.UnixMilli(int64(v))}, nil
		case string:
			t, err := time.Parse(time.RFC3339, v)
			if err != nil {
				return nil, err
			}
			return DateValue{V: t}, nil
		}
		return nil, fmt.Errorf("expected date, got %T", raw)

	// ... handle other types

	default:
		return GenericValue{V: raw}, nil
	}
}

// NullValue for handling nulls
type NullValue struct{}

func (NullValue) String() string { return "" }
func (NullValue) Value() any     { return nil }

// GenericValue as fallback
type GenericValue struct{ V any }

func (g GenericValue) String() string { return fmt.Sprintf("%v", g.V) }
func (g GenericValue) Value() any     { return g.V }
