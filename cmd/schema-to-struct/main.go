package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
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
	outputDir := flag.String("output", "", "Directory where Go files will be written")
	packageName := flag.String("package", "layer", "Package name for generated Go files")
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
	needsTimeImport := false
	needsUUIDImport := false

	// First determine if we need time or UUID imports
	for _, field := range schema.Fields {
		fieldType := mapFieldType(field.Type)
		if fieldType == "" {
			fmt.Printf("No mapping exists for field type %s", field.Type)
			os.Exit(1)
		}
		if fieldType == "time.Time" {
			needsTimeImport = true
		} else if fieldType == "uuid.UUID" {
			needsUUIDImport = true
		}
	}

	// Write package declaration
	code.WriteString(fmt.Sprintf("package %s\n\n", packageName))

	// Write imports
	code.WriteString("import (\n")
	code.WriteString("\t\"encoding/json\"\n")
	if needsTimeImport {
		code.WriteString("\t\"time\"\n")
	}

	if needsUUIDImport {
		if needsTimeImport {
			code.WriteString("\n")
		}
		code.WriteString("\t\"github.com/google/uuid\"\n")
	}

	code.WriteString(")\n\n")

	// Begin struct definition
	code.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	// Keep track of field names to ensure uniqueness
	usedFieldNames := make(map[string]bool)

	// Add fields
	for _, field := range schema.Fields {
		// Use alias if available, otherwise use field name
		displayName := field.Name
		if field.Alias != "" && field.Alias != field.Name {
			// Remove parentheses, question marks, and other non-alphanumeric chars from alias
			cleanAlias := strings.Map(func(r rune) rune {
				if strings.ContainsRune("()[]{}.,;:!@#$%^&*-+?", r) {
					return -1 // Remove the character
				}
				return r
			}, field.Alias)

			if cleanAlias != "" {
				displayName = cleanAlias
			}
		}

		fieldName := toPascalCasePreserveNumbers(displayName)
		fieldName = fixIDSuffix(fieldName)

		// Ensure field name uniqueness
		originalFieldName := fieldName
		suffix := 2
		for usedFieldNames[fieldName] {
			fieldName = fmt.Sprintf("%s%d", originalFieldName, suffix)
			suffix++
		}
		usedFieldNames[fieldName] = true

		// Determine field type
		fieldType := mapFieldType(field.Type)

		code.WriteString(fmt.Sprintf("\t%s %s `field:\"%s\"`\n", fieldName, fieldType, field.Name))
	}
	// Add geometry definition
	code.WriteString("\tGeometry json.RawMessage")

	// Close struct definition
	code.WriteString("}\n")
	code.WriteString(fmt.Sprintf("func (x *%s) GetGeometry() json.RawMessage { return x.Geometry }\n", structName))
	code.WriteString(fmt.Sprintf("func (x *%s) SetGeometry(m json.RawMessage) { x.Geometry = m }\n", structName))

	return code.String()
}

// Create a descriptive enum value name that handles mathematical symbols
func descriptiveEnumValueName(name string) string {
	// Convert numeric values to words for better readability
	re := regexp.MustCompile(`\b(\d+)\b`)
	name = re.ReplaceAllStringFunc(name, func(s string) string {
		n, err := strconv.Atoi(s)
		if err != nil {
			return s
		}
		return numberToWord(n)
	})

	// Replace mathematical symbols with descriptive text
	name = strings.ReplaceAll(name, "~", "About")
	name = strings.ReplaceAll(name, "<", "LessThan")
	name = strings.ReplaceAll(name, ">", "GreaterThan")
	name = strings.ReplaceAll(name, "=", "Equals")
	name = strings.ReplaceAll(name, "≤", "LessThanOrEqual")
	name = strings.ReplaceAll(name, "≥", "GreaterThanOrEqual")

	// Finally clean the string for a Go identifier
	return cleanEnumValueName(name)
}

// Convert a number to its word representation
func numberToWord(n int) string {
	words := map[int]string{
		0:  "Zero",
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		// Add more as needed
	}

	if word, ok := words[n]; ok {
		return word
	}

	// For numbers not in the map, just convert to string
	return strconv.Itoa(n)
}

// Clean a string literal for use in Go code by removing or escaping problematic characters
func cleanStringLiteral(s string) string {
	// Replace quotes with empty string to avoid Go syntax errors
	s = strings.ReplaceAll(s, "\"", "")
	s = strings.ReplaceAll(s, "'", "")

	// Replace other potentially problematic characters
	s = strings.ReplaceAll(s, "\\", "")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\t", " ")

	return s
}

// Clean an identifier for use in Go code
func cleanIdentifier(s string) string {
	// Replace non-alphanumeric characters with empty string
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	cleaned := re.ReplaceAllString(s, "")

	// Ensure it starts with a capital letter
	if len(cleaned) > 0 {
		return strings.ToUpper(cleaned[:1]) + cleaned[1:]
	}

	return "Unknown"
}

// Return the base type for an enum based on field type
func getEnumBaseType(fieldType string) string {
	switch fieldType {
	case "esriFieldTypeSmallInteger":
		return "int16"
	case "esriFieldTypeInteger":
		return "int32"
	case "esriFieldTypeSingle":
		return "float32"
	case "esriFieldTypeDouble":
		return "float64"
	default:
		return "string"
	}
}

func cleanEnumValueName(name string) string {
	// Replace dashes, underscores, spaces with nothing
	re := regexp.MustCompile(`[-_ ]`)
	cleanName := re.ReplaceAllString(name, "")

	// Remove any non-alphanumeric characters
	re = regexp.MustCompile(`[^a-zA-Z0-9]`)
	cleanName = re.ReplaceAllString(cleanName, "")

	// Ensure it starts with a capital letter
	if len(cleanName) > 0 {
		return strings.ToUpper(cleanName[:1]) + cleanName[1:]
	}

	return "Unknown"
}

// Similar to toPascalCase but preserves numeric suffixes
func toPascalCasePreserveNumbers(s string) string {
	// Handle empty strings
	if s == "" {
		return ""
	}

	// Check for numeric suffix
	re := regexp.MustCompile(`^([a-zA-Z_]+)(\d+)$`)
	matches := re.FindStringSubmatch(s)

	var base string
	var numericSuffix string

	if len(matches) == 3 {
		// There is a numeric suffix, separate it
		base = matches[1]
		numericSuffix = matches[2]
	} else {
		base = s
	}

	// Process the base part using our standard Pascal case logic
	// Split on underscores, spaces, or case changes
	var parts []string

	// First, split on underscores and spaces
	for _, part := range strings.FieldsFunc(base, func(r rune) bool {
		return r == '_' || r == ' '
	}) {
		// Check if the part is all uppercase
		if strings.ToUpper(part) == part && len(part) > 1 {
			// Convert all-caps words to lowercase first
			part = strings.ToLower(part)
		}
		parts = append(parts, part)
	}

	// Capitalize the first letter of each part
	for i, part := range parts {
		if len(part) > 0 {
			if i == 0 || len(part) > 1 {
				parts[i] = strings.ToUpper(part[:1]) + part[1:]
			} else {
				parts[i] = strings.ToUpper(part)
			}
		}
	}

	// Rejoin and add numeric suffix if present
	result := strings.Join(parts, "")
	if numericSuffix != "" {
		result += numericSuffix
	}

	return result
}

func fixIDSuffix(n string) string {
	lowered := strings.ToLower(n)
	if strings.HasSuffix(lowered, "id") {
		return n[:len(n)-2] + "ID"
	}
	return n
}

func mapFieldType(fieldType string) string {
	switch fieldType {
	case "esriFieldTypeOID":
		return "uint"
	case "esriFieldTypeSmallInteger":
		return "int16"
	case "esriFieldTypeInteger":
		return "int32"
	case "esriFieldTypeSingle":
		return "float32"
	case "esriFieldTypeDouble":
		return "float64"
	case "esriFieldTypeString":
		return "string"
	case "esriFieldTypeDate":
		return "time.Time"
	case "esriFieldTypeGlobalID":
		return "uuid.UUID"
	case "esriFieldTypeGUID":
		return "uuid.UUID"
	default:
		return ""
	}
}
