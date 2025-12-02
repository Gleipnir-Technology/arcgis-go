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
	ObjectIDFieldName string            `json:"objectIdFieldName"`
	GeometryType      string            `json:"geometryType"`
	Fields            []Field           `json:"fields"`
	Features          []json.RawMessage `json:"features"` // Added to detect empty files
}

func main() {
	// Parse command line arguments
	inputDir := flag.String("input", "", "Directory containing JSON schema files")
	outputDir := flag.String("output", "", "Directory where SQL files will be written")
	dbSchema := flag.String("schema", "public", "PostgreSQL schema name for the tables")
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

	processed := 0
	skipped := 0

	for _, file := range files {
		if processFile(file, *outputDir, *dbSchema) {
			processed++
		} else {
			skipped++
		}
	}

	fmt.Printf("Successfully processed %d schema files, skipped %d empty files\n", processed, skipped)
}

func processFile(filePath, outputDir, dbSchema string) bool {
	// Read and parse the JSON file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return false
	}

	var schema Schema
	if err := json.Unmarshal(data, &schema); err != nil {
		fmt.Printf("Error parsing JSON in file %s: %v\n", filePath, err)
		return false
	}

	// Generate table name from the filename
	baseName := filepath.Base(filePath)
	tableName := strings.TrimSuffix(baseName, ".json")

	// Check if the file is effectively empty (no fields or only has features array that's empty)
	if len(schema.Fields) == 0 {
		// This is an empty schema or a features file with no schema
		if len(schema.Features) == 0 {
			fmt.Printf("Warning: Skipping %s - file appears to be empty (no fields and empty or missing features array)\n", filePath)
		} else {
			fmt.Printf("Warning: Skipping %s - file contains features but no schema fields\n", filePath)
		}
		return false
	}

	// Generate the SQL code
	sqlCode := generateSQLCode(tableName, schema, dbSchema)

	// Write the SQL code to a file
	outputFileName := strings.ToLower(tableName) + ".sql"
	outputPath := filepath.Join(outputDir, outputFileName)

	if err := ioutil.WriteFile(outputPath, []byte(sqlCode), 0644); err != nil {
		fmt.Printf("Error writing SQL file %s: %v\n", outputPath, err)
		return false
	}

	fmt.Printf("Generated %s from %s\n", outputPath, filePath)
	return true
}

func generateSQLCode(tableName string, schema Schema, dbSchema string) string {
	var code strings.Builder
	domainTypes := make(map[string][]CodedValue)
	schemaName := sanitizeSQLName(dbSchema)
	sanitizedTableName := sanitizeSQLName(tableName)

	// Store fields with enum defaults for later comment generation
	enumDefaultComments := make(map[string]string)

	// Collect all domains for potential enum types
	for _, field := range schema.Fields {
		if field.Domain != nil && field.Domain.Type == "codedValue" && len(field.Domain.CodedValues) > 0 {
			domainTypes[field.Domain.Name] = field.Domain.CodedValues
		}
	}

	// Add header comment
	code.WriteString(fmt.Sprintf("-- Table definition for %s.%s\n", schemaName, tableName))
	code.WriteString("-- Includes versioning for tracking changes\n\n")

	// Create schema if not exists
	code.WriteString(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;\n\n", schemaName))

	// Create enum types for domains
	for domainName, codedValues := range domainTypes {
		enumTypeName := fmt.Sprintf("%s.%s_%s_enum", schemaName, sanitizedTableName, sanitizeSQLName(domainName))
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

	// Begin table definition with schema qualification
	code.WriteString(fmt.Sprintf("CREATE TABLE %s.%s (\n", schemaName, sanitizedTableName))

	// Process fields
	var primaryKeyField string
	for _, field := range schema.Fields {
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

		// Flag to track if this field uses an enum type
		isEnumField := false

		// Handle domains (enum types) - use schema qualified type name
		if field.Domain != nil && field.Domain.Type == "codedValue" && len(field.Domain.CodedValues) > 0 {
			enumTypeName := fmt.Sprintf("%s.%s_%s_enum", schemaName, sanitizedTableName, sanitizeSQLName(field.Domain.Name))
			fieldType = enumTypeName
			isEnumField = true
		}

		// Build column definition
		columnDef := fmt.Sprintf("  %s %s", fieldName, fieldType)

		// Add NOT NULL for ObjectID field
		if field.Name == schema.ObjectIDFieldName {
			columnDef += " NOT NULL"
		}

		// Handle default values - only add DEFAULT clause for non-enum fields
		if field.Default != nil {
			if isEnumField {
				// For enum fields, store the default value as a comment for later
				var defaultComment string
				switch v := field.Default.(type) {
				case string:
					defaultComment = fmt.Sprintf("'%s'", escapeSQLString(v))
				default:
					defaultComment = fmt.Sprintf("%v", v)
				}
				enumDefaultComments[fieldName] = defaultComment
			} else {
				// For non-enum fields, add the DEFAULT clause
				switch v := field.Default.(type) {
				case string:
					columnDef += fmt.Sprintf(" DEFAULT '%s'", escapeSQLString(v))
				case float64, int, int64:
					columnDef += fmt.Sprintf(" DEFAULT %v", v)
				case bool:
					columnDef += fmt.Sprintf(" DEFAULT %t", v)
				}
			}
		}

		// Always add a comma since we'll add VERSION column later
		columnDef += ","
		code.WriteString(columnDef + "\n")
	}

	// Add VERSION column for tracking changes
	code.WriteString("  VERSION INTEGER NOT NULL DEFAULT 1,\n")

	// Add primary key constraint including VERSION
	if primaryKeyField != "" {
		code.WriteString(fmt.Sprintf("  PRIMARY KEY (%s, VERSION)\n", primaryKeyField))
	} else {
		// If no ObjectID field exists, warn but still create the VERSION column
		code.WriteString("  -- Warning: No ObjectID field found, VERSION column added but not in primary key\n")
		code.WriteString("  PRIMARY KEY (VERSION)\n")
	}

	// Close table definition
	code.WriteString(");\n\n")

	// Add comment for VERSION column
	code.WriteString(fmt.Sprintf("COMMENT ON COLUMN %s.%s.VERSION IS 'Tracks version changes to the row. Increases when data is modified.';\n\n",
		schemaName, sanitizedTableName))

	// Add comments for fields with aliases - use schema qualified table name
	for _, field := range schema.Fields {
		if field.Alias != "" && field.Alias != field.Name {
			code.WriteString(fmt.Sprintf("COMMENT ON COLUMN %s.%s.%s IS '%s';\n",
				schemaName,
				sanitizedTableName,
				sanitizeSQLName(field.Name),
				escapeSQLString(field.Alias)))
		}
	}

	// Add comments for enum fields with default values
	for fieldName, defaultValue := range enumDefaultComments {
		code.WriteString(fmt.Sprintf("\n-- Field %s has default value: %s\n", fieldName, defaultValue))
	}

	// Generate PREPARE statement for conditional insert with versioning
	generatePreparedStatement(&code, schema, tableName, schemaName, domainTypes)

	return code.String()
}

func generatePreparedStatement(code *strings.Builder, schema Schema, tableName, schemaName string, domainTypes map[string][]CodedValue) {
	sanitizedTableName := sanitizeSQLName(tableName)

	// Generate unique prepared statement name based on table name
	preparedStatementName := fmt.Sprintf("insert_%s_versioned", sanitizedTableName)

	// Add header comment
	code.WriteString(fmt.Sprintf("\n-- Prepared statement for conditional insert with versioning\n"))
	code.WriteString(fmt.Sprintf("-- Only inserts a new version if data has changed\n"))

	// Start preparing parameter type list
	var paramTypes []string
	var columnNames []string
	var paramPlaceholders []string
	var conditionClauses []string

	paramCounter := 1

	// Process all fields to build parameter lists
	for _, field := range schema.Fields {
		fieldName := sanitizeSQLName(field.Name)
		columnNames = append(columnNames, fieldName)

		// Get PostgreSQL type for this field
		var pgType string

		if field.Domain != nil && field.Domain.Type == "codedValue" && len(field.Domain.CodedValues) > 0 {
			// For enum fields, use the fully qualified enum type
			pgType = fmt.Sprintf("%s.%s_%s_enum", schemaName, sanitizedTableName, sanitizeSQLName(field.Domain.Name))
		} else {
			// Map field type to PostgreSQL parameter type
			pgType = mapFieldTypeToPgParamType(field.Type, field.Length)
		}

		// Add to parameter types list
		paramTypes = append(paramTypes, pgType)

		// Add placeholder for parameter in INSERT
		paramPlaceholders = append(paramPlaceholders, fmt.Sprintf("$%d", paramCounter))

		// Add condition for checking if data has changed
		conditionClauses = append(conditionClauses, fmt.Sprintf("    lv.%s IS NOT DISTINCT FROM $%d", fieldName, paramCounter))

		paramCounter++
	}

	// Generate the PREPARE statement with type declarations
	code.WriteString(fmt.Sprintf("PREPARE %s(%s) AS\n",
		preparedStatementName, strings.Join(paramTypes, ", ")))

	// Add the conditional insert query with CTEs
	code.WriteString("WITH\n")
	code.WriteString("-- Get the current latest version of this record\n")
	code.WriteString(fmt.Sprintf("latest_version AS (\n  SELECT * FROM %s.%s\n", schemaName, sanitizedTableName))
	code.WriteString(fmt.Sprintf("  WHERE %s = $1\n", sanitizeSQLName(schema.ObjectIDFieldName)))
	code.WriteString("  ORDER BY VERSION DESC\n")
	code.WriteString("  LIMIT 1\n")
	code.WriteString("),\n")

	code.WriteString("-- Calculate the next version number\n")
	code.WriteString("next_version AS (\n")
	code.WriteString("  SELECT COALESCE(MAX(VERSION) + 1, 1) as version_num\n")
	code.WriteString(fmt.Sprintf("  FROM %s.%s\n", schemaName, sanitizedTableName))
	code.WriteString(fmt.Sprintf("  WHERE %s = $1\n", sanitizeSQLName(schema.ObjectIDFieldName)))
	code.WriteString(")\n")

	// Start INSERT statement
	code.WriteString("-- Perform conditional insert\n")
	code.WriteString(fmt.Sprintf("INSERT INTO %s.%s (\n", schemaName, sanitizedTableName))
	code.WriteString(fmt.Sprintf("  %s,\n", strings.Join(columnNames, ", ")))
	code.WriteString("  VERSION\n")
	code.WriteString(")\n")

	// Select clause
	code.WriteString("SELECT\n")
	code.WriteString(fmt.Sprintf("  %s,\n", strings.Join(paramPlaceholders, ", ")))
	code.WriteString("  v.version_num\n")
	code.WriteString("FROM next_version v\n")

	// Where clause for conditional insert
	code.WriteString("WHERE\n")
	code.WriteString("  -- Only insert if no record exists yet OR data has changed\n")
	code.WriteString("  NOT EXISTS (SELECT 1 FROM latest_version lv WHERE\n")
	code.WriteString(strings.Join(conditionClauses, " AND\n"))
	code.WriteString("\n  )\n")

	// Return the inserted row
	code.WriteString("RETURNING *;\n")

	// Add execution example
	code.WriteString(fmt.Sprintf("\n-- Example usage: EXECUTE %s(id, value1, value2, ...);\n", preparedStatementName))
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

// Map field type to PostgreSQL parameter type for PREPARE statements
func mapFieldTypeToPgParamType(fieldType string, length int) string {
	switch fieldType {
	case "esriFieldTypeOID":
		return "bigint"
	case "esriFieldTypeSmallInteger":
		return "smallint"
	case "esriFieldTypeInteger":
		return "integer"
	case "esriFieldTypeSingle":
		return "real"
	case "esriFieldTypeDouble":
		return "double precision"
	case "esriFieldTypeString":
		if length > 0 {
			return "varchar"
		}
		return "text"
	case "esriFieldTypeDate":
		return "timestamp"
	case "esriFieldTypeGlobalID", "esriFieldTypeGUID":
		return "uuid"
	default:
		return "text"
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
