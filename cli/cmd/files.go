/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "List all the files of the table.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("files called")
	},
}

func init() {
	rootCmd.AddCommand(filesCmd)
}
