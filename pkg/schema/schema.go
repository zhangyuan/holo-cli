package schema

import (
	_ "embed"
	"holo-cli/pkg/render"
	"os"

	"gopkg.in/yaml.v3"
)

//go:embed templates/create_table.tmpl
var createTableSQL string

const ProcTimeColumnName = "proc_time"

type Column struct {
	Name     string `yaml:"name"`
	DataType string `yaml:"data_type"`
	Comment  string `yaml:"comment"`
}

type Schema struct {
	SchemaName       string   `yaml:"schema_name,omitempty"`
	TableName        string   `yaml:"table_name"`
	Description      string   `yaml:"description"`
	Columns          []Column `yaml:"columns"`
	PrimaryKeys      []string `yaml:"primary_keys"`
	DistributionKeys []string `yaml:"distribution_keys"`
}

func LoadSchema(path string) (*Schema, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var schema Schema
	if err := yaml.Unmarshal(bytes, &schema); err != nil {
		return nil, err
	}
	return &schema, nil
}

func (s *Schema) ToSQL() (string, error) {
	r, err := render.NewRender("CreateTable", createTableSQL)
	if err != nil {
		return "", nil
	}

	return r.Render(s)
}
