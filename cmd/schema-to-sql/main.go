package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// CodedValue represents a value in a coded domain
type CodedValue struct {
	Name string          `json:"name"`
	Code json.RawMessage `json:"code"`
}

// Domain represents a domain definition
type Domain struct {
	Type        string       `json:"type"`
	Name        string       `json:"name"`
	MergePolicy string       `json:"mergePolicy"`
	SplitPolicy string       `json:"splitPolicy"`
	CodedValues []CodedValue `json:"codedValues"`
}

// Field represents a field in the schema
type Field struct {
	Name    string      `json:"name"`
	Type    string      `json:"type"`
	Alias   string      `json:"alias"`
	SQLType string      `json:"sqlType"`
	Length  int         `json:"length,omitempty"`
	Domain  *Domain     `json:"domain"`
	Default interface{} `json:"defaultValue"`
}

// Schema represents the schema definition from the JSON file
type Schema struct {
	ObjectIDFieldName string  `json:"objectIdFieldName"`
	GeometryType      string  `json:"geometryType"`
	Fields            []Field `json:"fields"`
}

func main() {
	// Parse command line arguments
	inputDir := flag.String("input", "", "Directory containing JSON schema files")
	outputDir := flag.String("output", "", "Directory where SQL files will be written")
	flag.Parse()

	// Validate input
	if *inputDir == "" || *outputDir == "" {
		fmt.Println("Error: Both input and output directories must be specified")
		flag.Usage()
		os.Exit(1)
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Process all JSON files in the input directory
	files, err := filepath.Glob(filepath.Join(*inputDir, "*.json"))
	if err != nil {
		fmt.Printf("Error reading input directory: %v\n", err)
		os.Exit(1)
	}

	for _, file := range files {
		processFile(file, *outputDir)
	}

	fmt.Printf("Successfully processed %d schema files\n", len(files))
}

func processFile(filePath, outputDir string) {
	// Read and parse the JSON file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return
	}

	var schema Schema
	if err := json.Unmarshal(data, &schema); err != nil {
		fmt.Printf("Error parsing JSON in file %s: %v\n", filePath, err)
		return
	}

	// Generate table name from the filename
	baseName := filepath.Base(filePath)
	tableName := strings.TrimSuffix(baseName, ".json")

	// Generate the SQL code
	sqlCode := generateSQLCode(tableName, schema)

	// Write the SQL code to a file
	outputFileName := strings.ToLower(tableName) + ".sql"
	outputPath := filepath.Join(outputDir, outputFileName)

	if err := ioutil.WriteFile(outputPath, []byte(sqlCode), 0644); err != nil {
		fmt.Printf("Error writing SQL file %s: %v\n", outputPath, err)
		return
	}

	fmt.Printf("Generated %s from %s\n", outputPath, filePath)
}

func generateSQLCode(tableName string, schema Schema) string {
	var code strings.Builder
	domainTypes := make(map[string][]CodedValue)

	// Collect all domains for potential enum types
	for _, field := range schema.Fields {
		if field.Domain != nil && field.Domain.Type == "codedValue" && len(field.Domain.CodedValues) > 0 {
			domainTypes[field.Domain.Name] = field.Domain.CodedValues
		}
	}

	// Add header comment
	code.WriteString(fmt.Sprintf("-- Table definition for %s\n\n", tableName))

	// Create enum types for domains
	for domainName, codedValues := range domainTypes {
		enumTypeName := fmt.Sprintf("%s_%s_enum", tableName, sanitizeSQLName(domainName))
		code.WriteString(fmt.Sprintf("CREATE TYPE %s AS ENUM (\n", enumTypeName))

		for i, value := range codedValues {
			var enumValue string

			// If the Code field is a string (starts with a quote)
			if len(value.Code) > 0 && value.Code[0] == '"' {
				var strVal string
				if err := json.Unmarshal(value.Code, &strVal); err == nil {
					enumValue = fmt.Sprintf("'%s'", escapeSQLString(strVal))
				} else {
					enumValue = "'' -- Error parsing code value"
				}
			} else {
				// For non-string values, convert to string representation
				enumValue = fmt.Sprintf("'%s'", string(value.Code))
			}

			// Add comma for all but the last item
			if i < len(codedValues)-1 {
				code.WriteString(fmt.Sprintf("  %s,\n", enumValue))
			} else {
				code.WriteString(fmt.Sprintf("  %s\n", enumValue))
			}
		}
		code.WriteString(");\n\n")
	}

	// Begin table definition
	code.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", sanitizeSQLName(tableName)))

	// Process fields
	var primaryKeyField string
	for i, field := range schema.Fields {
		fieldName := sanitizeSQLName(field.Name)
		fieldType := mapFieldTypeToSQL(field.Type, field.Length)

		// Check if this is the ObjectID field (primary key)
		if field.Name == schema.ObjectIDFieldName {
			primaryKeyField = fieldName
			// For the primary key field, use SERIAL or BIGSERIAL
			if fieldType == "INTEGER" {
				fieldType = "SERIAL"
			} else if fieldType == "BIGINT" {
				fieldType = "BIGSERIAL"
			}
		}

		// Handle domains (enum types)
		if field.Domain != nil && field.Domain.Type == "codedValue" && len(field.Domain.CodedValues) > 0 {
			enumTypeName := fmt.Sprintf("%s_%s_enum", tableName, sanitizeSQLName(field.Domain.Name))
			fieldType = enumTypeName
		}

		// Build column definition
		columnDef := fmt.Sprintf("  %s %s", fieldName, fieldType)

		// Add NOT NULL for ObjectID field
		if field.Name == schema.ObjectIDFieldName {
			columnDef += " NOT NULL"
		}

		// Handle default values
		if field.Default != nil {
			// Format default value based on its type
			switch v := field.Default.(type) {
			case string:
				columnDef += fmt.Sprintf(" DEFAULT '%s'", escapeSQLString(v))
			case float64, int, int64:
				columnDef += fmt.Sprintf(" DEFAULT %v", v)
			case bool:
				columnDef += fmt.Sprintf(" DEFAULT %t", v)
			}
		}

		// Add comma if not the last field
		if i < len(schema.Fields)-1 {
			columnDef += ","
		}

		code.WriteString(columnDef + "\n")
	}

	// Add primary key constraint if we have an ObjectID field
	if primaryKeyField != "" {
		code.WriteString(fmt.Sprintf("  , PRIMARY KEY (%s)\n", primaryKeyField))
	}

	// Close table definition
	code.WriteString(");\n\n")

	// Add comments for fields with aliases
	for _, field := range schema.Fields {
		if field.Alias != "" && field.Alias != field.Name {
			code.WriteString(fmt.Sprintf("COMMENT ON COLUMN %s.%s IS '%s';\n",
				sanitizeSQLName(tableName),
				sanitizeSQLName(field.Name),
				escapeSQLString(field.Alias)))
		}
	}

	return code.String()
}

func mapFieldTypeToSQL(fieldType string, length int) string {
	switch fieldType {
	case "esriFieldTypeOID":
		return "BIGINT"
	case "esriFieldTypeSmallInteger":
		return "SMALLINT"
	case "esriFieldTypeInteger":
		return "INTEGER"
	case "esriFieldTypeSingle":
		return "REAL"
	case "esriFieldTypeDouble":
		return "DOUBLE PRECISION"
	case "esriFieldTypeString":
		if length > 0 {
			return fmt.Sprintf("VARCHAR(%d)", length)
		}
		return "TEXT"
	case "esriFieldTypeDate":
		return "TIMESTAMP"
	case "esriFieldTypeGlobalID", "esriFieldTypeGUID":
		return "UUID"
	default:
		fmt.Printf("Warning: No mapping exists for field type %s, using TEXT\n", fieldType)
		return "TEXT"
	}
}

// Sanitize SQL identifiers to prevent SQL injection and syntax errors
func sanitizeSQLName(name string) string {
	// Remove or replace invalid characters
	re := regexp.MustCompile(`[^a-zA-Z0-9_]`)
	sanitized := re.ReplaceAllString(name, "_")

	// Ensure it doesn't start with a number
	if len(sanitized) > 0 && sanitized[0] >= '0' && sanitized[0] <= '9' {
		sanitized = "_" + sanitized
	}

	return strings.ToLower(sanitized)
}

// Escape single quotes and other special characters in SQL strings
func escapeSQLString(s string) string {
	// Replace single quotes with two single quotes (SQL standard)
	return strings.ReplaceAll(s, "'", "''")
}
