package fieldseeker

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type LocationTracking struct {
	ObjectID       uint      `field:"OBJECTID"`
	Accuracy       string    `field:"Accuracy"`
	CreatedUser    string    `field:"created_user"`
	CreatedDate    time.Time `field:"created_date"`
	LastEditedUser string    `field:"last_edited_user"`
	LastEditedDate time.Time `field:"last_edited_date"`
	GlobalID       uuid.UUID `field:"GlobalID"`
	FieldTech      string    `field:"FIELDTECH"`
	CreationDate   time.Time `field:"CreationDate"`
	Creator        string    `field:"Creator"`
	EditDate       time.Time `field:"EditDate"`
	Editor         string    `field:"Editor"`
}

func locationTrackingFromAttributes(attributes map[string]any) (*LocationTracking, error) {
	// Create new LocationTracking instance
	result := &LocationTracking{}

	// Get the reflect.Value and reflect.Type of our struct
	val := reflect.ValueOf(result).Elem()
	typ := val.Type()

	// Iterate over each field in the struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Get the field tag value
		tagValue := fieldType.Tag.Get("field")
		if tagValue == "" {
			continue // Skip fields without a "field" tag
		}

		// Get the attribute value from the map
		attrValue, exists := attributes[tagValue]
		if !exists {
			log.Warn().Str("tag", tagValue).Msg("Missing expected tag for 'LocationTracking'")
			continue // Skip if attribute doesn't exist in the map
		}

		// Skip nil values
		if attrValue == nil {
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
	case reflect.Uint:
		return setUintField(field, value)
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
