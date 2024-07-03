package cmd

import (
	"fmt"
	"holo-cli/pkg/loader"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var copyFromCsvCmd = &cobra.Command{
	Use:   "copyFromCsv",
	Short: "Copy to Hologres table from csv file",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(".env"); os.IsNotExist(err) {
		} else {
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}
		}

		dsn := os.Getenv("DSN")
		loader := loader.NewLoader(dsn)
		rowsAffected, err := loader.CopyFromCsv(csvPath, schemaName, tableName, options)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(rowsAffected)
	},
}

func init() {
	rootCmd.AddCommand(copyFromCsvCmd)

	copyFromCsvCmd.Flags().StringVarP(&csvPath, "csv-path", "c", "", "Path to csv file")
	_ = copyFromCsvCmd.MarkFlagRequired("csv-path")

	copyFromCsvCmd.Flags().StringVarP(&schemaName, "schema-name", "s", "", "Schema name")
	_ = copyFromCsvCmd.MarkFlagRequired("schema-name")

	copyFromCsvCmd.Flags().StringVarP(&tableName, "table-name", "t", "", "Table name")
	_ = copyFromCsvCmd.MarkFlagRequired("table-name")

	copyFromCsvCmd.Flags().StringVar(&options, "options", "", "COPY options, seperated by ,")
}
