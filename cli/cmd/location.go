/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var locationCmd = &cobra.Command{
	Use:   "location table",
	Short: "Return the location of a table.",
	Long: `Return the location of a table.`,
	Example: `- goiceberg location nyc.taxis --uri http://localhost:8181`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		out, cat, tableID := parseOutput(), parseCatalog(), args[0]
		tbl := loadTable(out, cat, tableID)
		out.Text(tbl.Location())
	},
}

func init() {
	rootCmd.AddCommand(locationCmd)
}
