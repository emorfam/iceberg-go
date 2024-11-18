/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var specCmd = &cobra.Command{
	Use:   "spec table",
	Short: "Return the partition spec of a table.",
	Long: `Return the partition spec of a table.`,
	Example: `goiceberg spec nyc.taxis --uri http://localhost:8181`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		out, cat, tableID := parseOutput(), parseCatalog(), args[0]
		tbl := loadTable(out, cat, tableID)
		out.Spec(tbl.Spec())
	},
}

func init() {
	rootCmd.AddCommand(specCmd)
}

