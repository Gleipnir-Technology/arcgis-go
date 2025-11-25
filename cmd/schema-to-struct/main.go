package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Field represents a field in the schema
type Field struct {
	Name    string      `json:"name"`
	Type    string      `json:"type"`
	Alias   string      `json:"alias"`
	SQLType string      `json:"sqlType"`
	Length  int         `json:"length,omitempty"`
	Domain  interface{} `json:"domain"`
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
	outputDir := flag.String("output", "", "Directory where Go files will be written")
	packageName := flag.String("package", "layers", "Package name for generated Go files")
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
		processFile(file, *outputDir, *packageName)
	}

	fmt.Printf("Successfully processed %d schema files\n", len(files))
}

func processFile(filePath, outputDir, packageName string) {
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

	// Generate Go struct name from the filename
	baseName := filepath.Base(filePath)
	structName := strings.TrimSuffix(baseName, ".json")

	// Generate the Go code
	goCode := generateGoCode(structName, schema, packageName)

	// Write the Go code to a file
	outputFileName := strings.ToLower(structName) + ".go"
	outputPath := filepath.Join(outputDir, outputFileName)

	if err := ioutil.WriteFile(outputPath, []byte(goCode), 0644); err != nil {
		fmt.Printf("Error writing Go file %s: %v\n", outputPath, err)
		return
	}

	fmt.Printf("Generated %s from %s\n", outputPath, filePath)
}

func generateGoCode(structName string, schema Schema, packageName string) string {
	var code strings.Builder

	// Write package declaration and imports
	code.WriteString(fmt.Sprintf("package %s\n\n", packageName))
	code.WriteString("import (\n")
	code.WriteString("\t\"time\"\n\n")
	code.WriteString("\t\"github.com/google/uuid\"\n")
	code.WriteString(")\n\n")

	// Begin struct definition
	code.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	// Add fields
	for _, field := range schema.Fields {
		fieldName := toPascalCase(field.Name)
		fieldType := mapFieldType(field.Type)
		code.WriteString(fmt.Sprintf("\t%s %s `field:\"%s\"`\n", fieldName, fieldType, field.Name))
	}

	// Close struct definition
	code.WriteString("}\n")

	return code.String()
}

func toPascalCase(s string) string {
	// Simple implementation - capitalize first letter
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func mapFieldType(fieldType string) string {
	switch fieldType {
	case "esriFieldTypeOID":
		return "uint"
	case "esriFieldTypeDouble":
		return "float64"
	case "esriFieldTypeString":
		return "string"
	case "esriFieldTypeDate":
		return "time.Time"
	case "esriFieldTypeGlobalID":
		return "uuid.UUID"
	default:
		return "interface{}"
	}
}
