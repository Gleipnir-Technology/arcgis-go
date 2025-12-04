package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

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
		log.Error().Err(err).Msg("Error creating output directory")
		os.Exit(1)
	}

	// Process all JSON files in the input directory
	files, err := filepath.Glob(filepath.Join(*inputDir, "*.json"))
	if err != nil {
		log.Error().Err(err).Msg("Error reading input directory")
		os.Exit(1)
	}

	processed := 0
	skipped := 0

	for _, file := range files {
		did_process, err := processFile(file, *outputDir, *dbSchema)
		if err != nil {
			log.Error().Err(err).Str("file", file).Msg("Failed to process file")
			os.Exit(2)
		}
		if did_process {
			processed++
		} else {
			skipped++
		}
	}

	log.Info().Int("processed", processed).Int("skipped", skipped).Msg("Successfully schema files")
}

func readJSON(filePath string) (baseName string, schema Schema, err error) {
	baseName = filepath.Base(filePath)

	// Read and parse the JSON file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return baseName, schema, errors.New(fmt.Sprintf("Error reading file %s: %v\n", filePath, err))
	}

	if err := json.Unmarshal(data, &schema); err != nil {
		return baseName, schema, errors.New(fmt.Sprintf("Error parsing JSON in file %s: %v\n", filePath, err))
	}
	return baseName, schema, nil
}

func processFile(filePath, outputDir, dbSchema string) (processed bool, err error) {
	baseName, schema, err := readJSON(filePath)
	if err != nil {
		return false, fmt.Errorf("Failed to read JSON: %w", err)
	}
	sanitizedTableName := sanitizeSQLName(strings.TrimSuffix(baseName, ".json"))

	// Check if the file is effectively empty (no fields or only has features array that's empty)
	if len(schema.Fields) == 0 {
		// This is an empty schema or a features file with no schema
		if len(schema.Features) == 0 {
			log.Warn().Str("file", filePath).Msg("File appears to be empty (no fields and empty or missing features array)")
		} else {
			log.Warn().Str("file", filePath).Msg("File contains features but no schema fields")
		}
		return false, nil
	}

	err = sanityCheckSchema(schema)
	if err != nil {
		return false, fmt.Errorf("Failed sanity check: %w", err)
	}
	// Generate the SQL code for table definition
	tableSqlCode, err := createTableFunction(dbSchema, sanitizedTableName, schema)
	if err != nil {
		return false, fmt.Errorf("Failed to generate table SQL: %w", err)
	}

	// Write the table SQL code to a file
	outputFileName := strings.ToLower(sanitizedTableName) + ".sql"
	outputPath := filepath.Join(outputDir, "table", outputFileName)

	if err := ioutil.WriteFile(outputPath, []byte(tableSqlCode), 0644); err != nil {
		return false, fmt.Errorf("Error writing SQL file %s: %w", outputPath, err)
	}

	insertSql, err := createInsertFunction(dbSchema, sanitizedTableName, schema)
	if err != nil {
		return false, fmt.Errorf("Failed to create insert query: %w", err)
	}
	// Generate the SQL code for prepared insert statement
	//insertSqlCode := generatePreparedInsertSQL(sanitizedTableName, schema, dbSchema)

	// Write the insert SQL code to a file in the insert subdirectory
	insertFileName := strings.ToLower(fmt.Sprintf("insert_%s_versioned.sql", strings.ToLower(sanitizedTableName)))
	insertPath := filepath.Join(outputDir, "insert", insertFileName)

	if err := ioutil.WriteFile(insertPath, []byte(insertSql), 0644); err != nil {
		return false, fmt.Errorf("Error writing SQL file %s: %w\n", insertPath, err)
	}

	log.Info().Str("input", filePath).Str("output", outputPath).Str("insert", insertPath).Msg("Generated SQL")
	return true, nil
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
		log.Warn().Str("type", fieldType).Msg("No mapping exists for field type, using TEXT")
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

func sanityCheckSchema(schema Schema) error {
	if !(schema.GeometryType == "esriGeometryPoint" || schema.GeometryType == "esriGeometryPolyline" || schema.GeometryType == "esriGeometryPolygon") {
		return fmt.Errorf("Unrecognized geometry type '%s'", schema.GeometryType)
	}
	if schema.SpatialReference.LatestWKID != 3857 {
		return fmt.Errorf("The spatial reference is '%d' rather than 3857, which is baked in to a bunch of assumptions, so we'll just fail here.", schema.SpatialReference.LatestWKID)
	}
	return nil
}

func toGeospatial(t string) (string, error) {
	switch t {
	case "esriGeometryPoint":
		return "POINT", nil
	case "esriGeometryPolyline":
		return "LINESTRING", nil
	case "esriGeometryPolygon":
		return "POLYGON", nil
	default:
		return "", fmt.Errorf("Unrecognized esri geometry type '%s", t)
	}
}
