package fieldseeker

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/Gleipnir-Technology/arcgis-go"
	"github.com/google/uuid"
)

type ServiceRequestSourceType int

const (
	SourceTypeUnknown ServiceRequestSourceType = iota
	SourceTypePhone
	SourceTypeEmail
	SourceTypeWebsite
	SourceTypeDropin
)

type ServiceRequestPriorityType int

const (
	ServiceRequestPriorityUnknown ServiceRequestPriorityType = iota
	ServiceRequestPriorityLow
	ServiceRequestPriorityMedium
	ServiceRequestPriorityHigh
	ServiceRequestPriorityFollowupVisit
	ServiceRequestPriorityHTCResponse
	ServiceRequestPriorityDiseaseActivityResponse
)

type ServiceRequest struct {
	ObjectID        uint                       `field:"OBJECTID"`
	Received        time.Time                  `field:"RECDATETIME"`
	Source          ServiceRequestSourceType   `field:"SOURCE"`
	EnteredBy       string                     `field:"ENTRYTECH"`
	Priority        ServiceRequestPriorityType `field:"PRIORITY"`
	Supervisor      string                     `field:"SUPERVISOR"`
	AssignedTech    string                     `field:"ASSIGNEDTECH"`
	Status          string                     `field:"STATUS"`
	AnonymousCaller bool                       `field:"CLRANON"`
	CallerName      string                     `field:"CLRFNAME"`
	CallerPhone     string                     `field:"CLRPHONE"`
}

type Tracklog struct {
	ObjectID uint      `field:"OBJECTID"`
	GlobalID uuid.UUID `field:"GlobalID"`
}

func structFromFeature[T any, PT interface {
	*T
	Geometric
}](ctx context.Context, feature *arcgis.Feature) (PT, error) {
	logger := arcgis.LoggerFromContext(ctx)
	// Create new LocationTracking instance
	result := PT(new(T))

	// Get the reflect.Value and reflect.Type of our struct
	val := reflect.ValueOf(result).Elem()
	typ := val.Type()

	// Iterate over each field in the struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// If it's the Geometry field, pull it from a special place
		if fieldType.Name == "Geometry" {
			result.SetGeometry(feature.Geometry)
			continue
		}
		// Get the field tag value
		tagValue := fieldType.Tag.Get("field")
		if tagValue == "" {
			logger.Warn().Str("field", fieldType.Name).Msg("No field tag")
			continue // Skip fields without a "field" tag
		}

		// Get the attribute value from the map
		attrValue, exists := feature.Attributes[tagValue]
		if !exists {
			logger.Warn().Str("tag", tagValue).Str("type", typ.Name()).Msg("Missing expected tag")
			continue // Skip if attribute doesn't exist in the map
		}

		// Skip nil values
		if attrValue == nil {
			//logger.Warn().Str("field", fieldType.Name).Msg("nil value")
			continue
		}

		// Set the field based on its type
		if err := setFieldValue(field, attrValue); err != nil {
			return nil, fmt.Errorf("error setting field %s: %w", fieldType.Name, err)
		}
	}

	return result, nil
}

func setFieldValue(field reflect.Value, value any) error {
	// Check if the field can be set
	if !field.CanSet() {
		return fmt.Errorf("field cannot be set")
	}

	switch field.Kind() {
	case reflect.Float32:
		return setFloat32Field(field, value)
	case reflect.Float64:
		return setFloat64Field(field, value)
	case reflect.Int16:
		return setInt16Field(field, value)
	case reflect.Int32:
		return setInt32Field(field, value)
	case reflect.Uint:
		return setUintField(field, value)
	// json.RawMessage is a slice
	case reflect.Slice:
		return fmt.Errorf("not sure what to do with slice of '%s'", field.Elem().Kind())
	case reflect.String:
		return setStringField(field, value)
	case reflect.Struct:
		return setStructField(field, value)
	// UUID is an array of bytes
	case reflect.Array:
		if field.Len() == 16 && field.Type().Elem().Kind() == reflect.Uint8 {
			return setUUIDField(field, value)
		} else {
			return fmt.Errorf("not sure what to do with array of %d '%s'", field.Len(), field.Elem().Kind())
		}
	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}
}

func setFloat32Field(field reflect.Value, value any) error {
	var floatVal float32

	// Handle different input types
	switch v := value.(type) {
	case float32:
		floatVal = v
	case float64:
		floatVal = float32(v)
	case int:
		floatVal = float32(v)
	case int64:
		floatVal = float32(v)
	case uint:
		floatVal = float32(v)
	case int16:
		floatVal = float32(v)
	case string:
		parsedVal, err := strconv.ParseFloat(v, 10)
		if err != nil {
			return err
		}
		floatVal = float32(parsedVal)
	default:
		return fmt.Errorf("cannot convert %T to uint", value)
	}

	field.SetFloat(float64(floatVal))
	return nil
}

func setFloat64Field(field reflect.Value, value any) error {
	var floatVal float64

	// Handle different input types
	switch v := value.(type) {
	case float32:
		floatVal = float64(v)
	case float64:
		floatVal = v
	case int:
		floatVal = float64(v)
	case int64:
		floatVal = float64(v)
	case uint:
		floatVal = float64(v)
	case int16:
		floatVal = float64(v)
	case string:
		parsedVal, err := strconv.ParseFloat(v, 10)
		if err != nil {
			return err
		}
		floatVal = float64(parsedVal)
	default:
		return fmt.Errorf("cannot convert %T to uint", value)
	}

	field.SetFloat(floatVal)
	return nil
}

func setInt16Field(field reflect.Value, value any) error {
	var intVal int16

	// Handle different input types
	switch v := value.(type) {
	case float64:
		intVal = int16(v)
	case int:
		intVal = int16(v)
	case int64:
		intVal = int16(v)
	case uint:
		intVal = int16(v)
	case int16:
		intVal = v
	case string:
		parsedVal, err := strconv.ParseInt(v, 10, 16)
		if err != nil {
			return err
		}
		intVal = int16(parsedVal)
	default:
		return fmt.Errorf("cannot convert %T to uint", value)
	}

	field.SetInt(int64(intVal))
	return nil
}

func setInt32Field(field reflect.Value, value any) error {
	var intVal int32

	// Handle different input types
	switch v := value.(type) {
	case float64:
		intVal = int32(v)
	case int:
		intVal = int32(v)
	case int64:
		intVal = int32(v)
	case uint:
		intVal = int32(v)
	case int16:
		intVal = int32(v)
	case int32:
		intVal = v
	case string:
		parsedVal, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return err
		}
		intVal = int32(parsedVal)
	default:
		return fmt.Errorf("cannot convert %T to uint", value)
	}

	field.SetInt(int64(intVal))
	return nil
}

func setUintField(field reflect.Value, value any) error {
	var uintVal uint64

	// Handle different input types
	switch v := value.(type) {
	case float64:
		uintVal = uint64(v)
	case int:
		uintVal = uint64(v)
	case int64:
		uintVal = uint64(v)
	case uint:
		uintVal = uint64(v)
	case uint64:
		uintVal = v
	case string:
		parsedVal, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return err
		}
		uintVal = parsedVal
	default:
		return fmt.Errorf("cannot convert %T to uint", value)
	}

	field.SetUint(uintVal)
	return nil
}

func setStringField(field reflect.Value, value any) error {
	// Convert to string
	strValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot convert %T to string", value)
	}

	//logger.Debug().Str("field", field.Type().Name()).Str("value", strValue).Msg("Set field")
	field.SetString(strValue)
	return nil
}

func setStructField(field reflect.Value, value any) error {
	// Handle specific struct types
	switch field.Type() {
	case reflect.TypeOf(time.Time{}):
		return setTimeField(field, value)
	case reflect.TypeOf(uuid.UUID{}):
		return setUUIDField(field, value)
	default:
		return fmt.Errorf("unsupported struct type: %s", field.Type().Name())
	}
}

func setTimeField(field reflect.Value, value any) error {
	var timeVal time.Time
	var err error

	// Handle different time formats
	switch v := value.(type) {
	case string:
		// Try common time formats
		layouts := []string{
			time.RFC3339,
			"2006-01-02T15:04:05Z",
			"2006-01-02 15:04:05",
			"2006-01-02",
		}

		for _, layout := range layouts {
			timeVal, err = time.Parse(layout, v)
			if err == nil {
				break
			}
		}
		if err != nil {
			return fmt.Errorf("cannot parse time string: %s", v)
		}
	case float64:
		// Assume it's a Unix timestamp in milliseconds
		sec := int64(v / 1000)
		msec := int64(v) % 1000
		timeVal = time.Unix(sec, msec*1000000)
	case int64:
		// Assume it's a Unix timestamp in milliseconds
		sec := v / 1000
		msec := v % 1000
		timeVal = time.Unix(sec, msec*1000000)
	case time.Time:
		timeVal = v
	default:
		return fmt.Errorf("cannot convert %T to time.Time", value)
	}

	field.Set(reflect.ValueOf(timeVal))
	return nil
}

func setUUIDField(field reflect.Value, value any) error {
	var uuidVal uuid.UUID
	var err error

	switch v := value.(type) {
	case string:
		uuidVal, err = uuid.Parse(v)
		if err != nil {
			return fmt.Errorf("invalid UUID string: %s", v)
		}
	case []byte:
		uuidVal, err = uuid.ParseBytes(v)
		if err != nil {
			return fmt.Errorf("invalid UUID bytes")
		}
	case uuid.UUID:
		uuidVal = v
	default:
		return fmt.Errorf("cannot convert %T to uuid.UUID", value)
	}

	field.Set(reflect.ValueOf(uuidVal))
	return nil
}
