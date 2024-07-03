package cmd

import (
	"fmt"
	"holo-cli/pkg/loader"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var copyToCsvCmd = &cobra.Command{
	Use:   "copyToLocal",
	Short: "Copy Hologres query to local file",
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
		rowsAffected, err := loader.CopyToLocal(sqlQuery, outputPath, options)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(rowsAffected)
	},
}

func init() {
	rootCmd.AddCommand(copyToCsvCmd)

	copyToCsvCmd.Flags().StringVarP(&sqlQuery, "query", "q", "", "SQL query")
	_ = copyToCsvCmd.MarkFlagRequired("query")

	copyToCsvCmd.Flags().StringVarP(&outputPath, "output-path", "o", "", "Path to output file")
	_ = copyToCsvCmd.MarkFlagRequired("output-path")

	copyToCsvCmd.Flags().StringVar(&options, "options", "", "COPY options, seperated by ,")
}
