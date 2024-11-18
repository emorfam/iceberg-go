/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var propertiesCmd = &cobra.Command{
	Use:   "properties",
	Short: "Properties on tables/namespaces.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("properties called")
	},
}

func init() {
	rootCmd.AddCommand(propertiesCmd)
}
