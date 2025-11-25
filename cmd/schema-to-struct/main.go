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
	Name string `json:"name"`
	Code string `json:"code"`
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
	var enums []string

	// Extract domains that need to be converted to enums
	domains := make(map[string]*Domain)
	for _, field := range schema.Fields {
		if field.Domain != nil && field.Domain.Type == "codedValue" && len(field.Domain.CodedValues) > 0 {
			domains[field.Domain.Name] = field.Domain
		}
	}

	// Write package declaration and imports
	code.WriteString(fmt.Sprintf("package %s\n\n", packageName))
	code.WriteString("import (\n")
	code.WriteString("\t\"time\"\n\n")
	code.WriteString("\t\"github.com/google/uuid\"\n")
	code.WriteString(")\n\n")

	// Generate enum types
	for domainName, domain := range domains {
		enumName := domainName + "Type"
		enumPrefix := domainName

		// Begin enum type definition
		code.WriteString(fmt.Sprintf("type %s string\n\n", enumName))
		code.WriteString("const (\n")

		// Add enum values
		for _, value := range domain.CodedValues {
			valueName := cleanEnumValueName(value.Name)
			constName := fmt.Sprintf("%s%s", enumPrefix, valueName)
			code.WriteString(fmt.Sprintf("\t%s %s = \"%s\"\n", constName, enumName, value.Code))
		}

		// Close enum definition
		code.WriteString(")\n\n")

		// Keep track of generated enums
		enums = append(enums, enumName)
	}

	// Begin struct definition
	code.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	// Add fields
	for _, field := range schema.Fields {
		// Use alias if available, otherwise use field name
		displayName := field.Name
		if field.Alias != "" && field.Alias != field.Name {
			// Remove parentheses and other non-alphanumeric chars from alias if present
			cleanAlias := strings.Map(func(r rune) rune {
				if strings.ContainsRune("()[]{}.,;:!@#$%^&*-+", r) {
					return -1 // Remove the character
				}
				return r
			}, field.Alias)

			if cleanAlias != "" {
				displayName = cleanAlias
			}
		}

		fieldName := toPascalCase(displayName)

		// Determine field type
		var fieldType string
		if field.Domain != nil && field.Domain.Type == "codedValue" && len(field.Domain.CodedValues) > 0 {
			fieldType = field.Domain.Name + "Type"
		} else {
			fieldType = mapFieldType(field.Type)
		}

		code.WriteString(fmt.Sprintf("\t%s %s `field:\"%s\"`\n", fieldName, fieldType, field.Name))
	}

	// Close struct definition
	code.WriteString("}\n")

	return code.String()
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

func toPascalCase(s string) string {
	// Handle empty strings
	if s == "" {
		return ""
	}

	// Split on underscores, spaces, or case changes
	var parts []string

	// First, split on underscores and spaces
	for _, part := range strings.FieldsFunc(s, func(r rune) bool {
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

	return strings.Join(parts, "")
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
