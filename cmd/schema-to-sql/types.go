package main

import (
	"encoding/json"
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
