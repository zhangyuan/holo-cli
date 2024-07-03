package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var copyFromCsvCmd = &cobra.Command{
	Use:   "copyFromCsv",
	Short: "Copy to Hologres table from csv file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("copyFromCsv called")
	},
}

func init() {
	rootCmd.AddCommand(copyFromCsvCmd)
}
