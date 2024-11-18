/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "List all the files of the table.",
	Run: func(cmd *cobra.Command, args []string) {
		out, cat, tableID := parseOutput(), parseCatalog(), args[0]
		tbl := loadTable(out, cat, tableID)
		out.Files(tbl, cfg.History)
	},
}

func init() {
	rootCmd.AddCommand(filesCmd)
}
