/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var schemaPath string
var outputPath string
var csvPath string
var schemaName string
var tableName string
var options string
var sqlQuery string

var rootCmd = &cobra.Command{
	Use:   "holo-cli",
	Short: "Hologres Cli",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
