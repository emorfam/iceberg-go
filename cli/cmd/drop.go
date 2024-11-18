/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Operations to drop a namespace or table.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("drop called")
	},
}

func init() {
	rootCmd.AddCommand(dropCmd)
}
