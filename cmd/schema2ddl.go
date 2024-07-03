package cmd

import (
	"holo-cli/pkg/schema"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var schema2ddlCmd = &cobra.Command{
	Use:   "schema2ddl",
	Short: "Genenate DDL SQL from schema definition in YAML",
	Run: func(cmd *cobra.Command, args []string) {
		if err := InvokeSchema2ddlCmd(); err != nil {
			log.Fatalln(err)
		}
	},
}

func InvokeSchema2ddlCmd() error {
	schema, err := schema.LoadSchema(schemaPath)
	if err != nil {
		return err
	}

	sql, err := schema.ToSQL()
	if err != nil {
		return err
	}

	outputFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(sql)

	return err
}

func init() {
	rootCmd.AddCommand(schema2ddlCmd)

	schema2ddlCmd.Flags().StringVarP(&schemaPath, "schema-path", "s", "", "Path to schema file in YAML.")
	_ = schema2ddlCmd.MarkFlagRequired("schema-path")
	schema2ddlCmd.Flags().StringVarP(&outputPath, "output-path", "o", "", "Path to output file.")
	_ = schema2ddlCmd.MarkFlagRequired("output-path")
}
