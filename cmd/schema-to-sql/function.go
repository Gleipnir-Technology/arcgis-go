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
CREATE OR REPLACE FUNCTION {{.Schema}}.insert_{{.Table}}(
	p_objectid bigint,
	{{range $i, $p := .Parameters}}{{if $i}},{{end}}p_{{ $p.Name }} {{ $p.Type }}{{end}}
) RETURNS TABLE(row_inserted boolean, version_num integer) AS $$
DECLARE
	v_next_version integer;
	v_changes_exist boolean;
BEGIN
	-- Check if changes exist
	SELECT NOT EXISTS (
		SELECT 1 FROM {{.Schema}}.{{.Table}} lv 
		WHERE lv.objectid = p_objectid
		{{range .Parameters}}AND lv.{{.Name}} IS NOT DISTINCT FROM p_{{.Name}} {{end}}
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
		{{range .Parameters}}{{.Name}}, {{end}}
		VERSION
	) VALUES (
		p_objectid,
		{{range .Parameters}}p_{{.Name}}, {{end}}
		v_next_version
	);
	
	-- Return success with new version
	RETURN QUERY SELECT TRUE AS row_inserted, v_next_version AS version_num;
END;
$$ LANGUAGE plpgsql;
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
		pgType := toPGType(schema, table, field)
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

func toPGType(schema, table string, field Field) (pgType string) {
	if field.Domain != nil && field.Domain.Type == "codedValue" && len(field.Domain.CodedValues) > 0 {
		// For enum fields, use the fully qualified enum type
		pgType = fmt.Sprintf("%s.%s_%s_enum", schema, table, sanitizeSQLName(field.Domain.Name))
	} else {
		// Map field type to PostgreSQL parameter type
		pgType = mapFieldTypeToPgParamType(field.Type, field.Length)
	}
	return pgType
}
