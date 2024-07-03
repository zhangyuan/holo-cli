package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var copyToCsvCmd = &cobra.Command{
	Use:   "copyToCsv",
	Short: "Copy Hologres table to local csv file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("copyToCsv called")
	},
}

func init() {
	rootCmd.AddCommand(copyToCsvCmd)
}
