package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type InsertTemplateParameter struct {
	Name string
	Type string
}
type InsertTemplateContext struct {
	Parameters []InsertTemplateParameter
	Schema     string
	Table      string
}

const insertTemplate string = `
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION {{.Schema}}.insert_{{.Table}}(
	p_objectid bigint,
	{{range $i, $p := .Parameters}}
	p_{{ $p.Name }} {{ $p.Type }},{{end}}
	p_geometry jsonb,
	p_geospatial geometry
) RETURNS TABLE(row_inserted boolean, version_num integer) AS $$
DECLARE
	v_next_version integer;
	v_changes_exist boolean;
BEGIN
	-- Check if changes exist
	SELECT NOT EXISTS (
		SELECT 1 FROM {{.Schema}}.{{.Table}} lv 
		WHERE lv.objectid = p_objectid
		{{range .Parameters}}
		AND lv.{{.Name}} IS NOT DISTINCT FROM p_{{.Name}} {{end}}
		AND lv.geometry IS NOT DISTINCT FROM p_geometry
		AND lv.geospatial IS NOT DISTINCT FROM p_geospatial
		ORDER BY VERSION DESC LIMIT 1
	) INTO v_changes_exist;
	
	-- If no changes, return false with current version
	IF NOT v_changes_exist THEN
		RETURN QUERY 
			SELECT 
				FALSE AS row_inserted, 
				(SELECT VERSION FROM {{.Schema}}.{{.Table}} 
				 WHERE objectid = p_objectid ORDER BY VERSION DESC LIMIT 1) AS version_num;
		RETURN;
	END IF;
	
	-- Calculate next version
	SELECT COALESCE(MAX(VERSION) + 1, 1) INTO v_next_version
	FROM {{.Schema}}.{{.Table}}
	WHERE objectid = p_objectid;
	
	-- Insert new version
	INSERT INTO {{.Schema}}.{{.Table}} (
		objectid,
		{{range .Parameters}}
		{{.Name}},{{end}}
		geometry,
		geospatial,
		VERSION
	) VALUES (
		p_objectid,
		{{range .Parameters}}
		p_{{.Name}},{{end}}
		p_geometry,
		p_geospatial,
		v_next_version
	);
	
	-- Return success with new version
	RETURN QUERY SELECT TRUE AS row_inserted, v_next_version AS version_num;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
`

func createInsertFunction(schema, table string, s Schema) (result string, err error) {
	tmpl, err := template.New("insertTemplate").Parse(insertTemplate)
	if err != nil {
		return "", fmt.Errorf("Failed to parse template: %w", err)
	}
	parameters := make([]InsertTemplateParameter, 0)
	for _, field := range s.Fields {
		name := sanitizeSQLName(field.Name)
		// We specifically exclude objectid here because we use it as the key within
		// the table so it needs some special logic
		if name == "objectid" {
			continue
		}
		pgType := mapFieldTypeToPgParamType(field.Type, field.Length)
		parameters = append(parameters, InsertTemplateParameter{
			Name: name,
			Type: pgType,
		})
	}
	context := InsertTemplateContext{
		Parameters: parameters,
		Schema:     schema,
		Table:      table,
	}
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, context)
	if err != nil {
		return "", fmt.Errorf("Failed to execute template: %w", err)
	}
	return buffer.String(), nil
}
