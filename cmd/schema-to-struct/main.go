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

	// Write imports only if needed
	if needsTimeImport || needsUUIDImport {
		code.WriteString("import (\n")

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
	}

	// Generate enum types for fields with domains
	domainFields := make(map[string]Field)
	domainTypeNames := make(map[string]string) // Maps domain names to their clean Go type names

	// First pass: collect all fields with domains
	for _, field := range schema.Fields {
		if field.Domain != nil && field.Domain.Type == "codedValue" && len(field.Domain.CodedValues) > 0 {
			domainFields[field.Domain.Name] = field

			// Create a clean domain name for Go
			cleanDomainName := cleanIdentifier(field.Domain.Name)
			domainTypeNames[field.Domain.Name] = structName + cleanDomainName + "Type"
		}
	}

	// Second pass: generate enums
	for domainName, field := range domainFields {
		// Use the clean type name we created in the first pass
		enumName := domainTypeNames[domainName]
		enumPrefix := structName + cleanIdentifier(domainName)

		// Determine enum base type based on field type
		enumBaseType := getEnumBaseType(field.Type)

		// Begin enum type definition
		code.WriteString(fmt.Sprintf("type %s %s\n\n", enumName, enumBaseType))
		code.WriteString("const (\n")

		// Keep track of used constant names to ensure uniqueness
		usedConstNames := make(map[string]bool)

		// Add enum values
		for _, value := range field.Domain.CodedValues {
			valueName := descriptiveEnumValueName(value.Name)
			constName := fmt.Sprintf("%s%s", enumPrefix, valueName)

			// Ensure uniqueness of const names
			originalConstName := constName
			suffix := 1
			for usedConstNames[constName] {
				constName = fmt.Sprintf("%s%d", originalConstName, suffix)
				suffix++
			}
			usedConstNames[constName] = true

			// Format the code value based on its type
			var codeValue string

			// Detect if it's a string or numeric value by checking first character
			if len(value.Code) > 0 && value.Code[0] == '"' {
				// It's a string value
				var strVal string
				if err := json.Unmarshal(value.Code, &strVal); err == nil {
					// Clean up the string value to remove or escape problematic characters
					cleanedStr := cleanStringLiteral(strVal)
					codeValue = fmt.Sprintf("\"%s\"", cleanedStr)
				} else {
					codeValue = "\"\" // Error parsing code value"
				}
			} else {
				// It's a numeric value, use as is
				codeValue = string(value.Code)
			}

			code.WriteString(fmt.Sprintf("\t%s %s = %s\n", constName, enumName, codeValue))
		}

		// Close enum definition
		code.WriteString(")\n\n")
	}

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
		var fieldType string
		if field.Domain != nil && field.Domain.Type == "codedValue" && len(field.Domain.CodedValues) > 0 {
			// Use our clean domain type name from the map
			fieldType = domainTypeNames[field.Domain.Name]
		} else {
			fieldType = mapFieldType(field.Type)
		}

		code.WriteString(fmt.Sprintf("\t%s %s `field:\"%s\"`\n", fieldName, fieldType, field.Name))
	}

	// Close struct definition
	code.WriteString("}\n")

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
