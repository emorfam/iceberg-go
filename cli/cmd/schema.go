/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var schemaCmd = &cobra.Command{
	Use:   "schema table",
	Short: "Get the schema of a table.",
	Long: "Get the schema of a table.",
	Example: `- goiceberg schema nyc.taxis --uri http://localhost:8181`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		out, cat, tableID := parseOutput(), parseCatalog(), args[0]
		tbl := loadTable(out, cat, tableID)
		out.Schema(tbl.Schema())
	},
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}
