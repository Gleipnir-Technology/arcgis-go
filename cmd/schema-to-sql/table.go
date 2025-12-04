package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type TableTemplateParameter struct {
	Name      string
	SourceKey string
	Type      string
}
type TableTemplateContext struct {
	Geospatial string
	Parameters []TableTemplateParameter
	Schema     string
	Table      string
	WKID       int
}

const tableTemplate string = `
-- Table definition for {{.Schema}}.{{.Table}}
-- Includes versioning for tracking changes

-- When inserting a new row, VERSION defaults to 1
-- When updating a row, insert a new row with the same ID but incremented VERSION
-- The most recent version of a row has the highest VERSION value
CREATE SCHEMA IF NOT EXISTS {{.Schema}};

CREATE TABLE {{.Schema}}.{{.Table}} (
  objectid BIGSERIAL NOT NULL,
  {{ range .Parameters }}
  {{.Name}} {{.Type}},{{ end }}
  geometry JSONB NOT NULL,
  geospatial GEOMETRY({{.Geospatial}}, {{.WKID}}),
  VERSION INTEGER NOT NULL DEFAULT 1,
  PRIMARY KEY (objectid, VERSION)
);

{{ range .Parameters }}
COMMENT ON COLUMN {{$.Schema}}.{{$.Table}}.{{.Name}} IS 'Original attribute from ArcGIS API is {{.SourceKey}}';{{ end }}
`

func createTableFunction(schema, table string, s Schema) (result string, err error) {
	tmpl, err := template.New("tableTemplate").Parse(tableTemplate)
	if err != nil {
		return "", fmt.Errorf("Failed to parse table template: %w", err)
	}
	parameters := make([]TableTemplateParameter, 0)
	for _, field := range s.Fields {
		name := sanitizeSQLName(field.Name)
		// We specifically exclude objectid here because we use it as the key within
		// the table so it needs some special logic
		if name == "objectid" {
			continue
		}
		pgType := mapFieldTypeToSQL(field.Type, field.Length)
		parameters = append(parameters, TableTemplateParameter{
			Name:      name,
			SourceKey: field.Name,
			Type:      pgType,
		})
	}
	geospatial, err := toGeospatial(s.GeometryType)
	if err != nil {
		return "", fmt.Errorf("failed to figure out geometry type: %w", err)
	}
	context := TableTemplateContext{
		Geospatial: geospatial,
		Parameters: parameters,
		Schema:     schema,
		Table:      table,
		WKID:       s.SpatialReference.LatestWKID,
	}
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, context)
	if err != nil {
		return "", fmt.Errorf("Failed to execute table template: %w", err)
	}
	return buffer.String(), nil
}
