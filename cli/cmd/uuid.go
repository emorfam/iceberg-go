/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var uuidCmd = &cobra.Command{
	Use:   "uuid table",
	Short: "Return the UUID of a table.",
	Long: "Return the UUID of a table.",
	Example: `- goiceberg uuid nyc.taxis --uri http://localhost:8181`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		out, cat, tableID := parseOutput(), parseCatalog(), args[0]
		tbl := loadTable(out, cat, tableID)
		out.Uuid(tbl.Metadata().TableUUID())
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
}
