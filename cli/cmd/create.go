/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"os"

	"github.com/apache/iceberg-go"
	"github.com/apache/iceberg-go/catalog"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a namespace or table.",
}

var createTableCmd = &cobra.Command{
	Use:   "table",
	Short: "Create a table.",
	Run: func(cmd *cobra.Command, args []string) {
		createTable()
	},
}

var createNamespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Create a namespace.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		createNamespace(args)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createTableCmd)
	createCmd.AddCommand(createNamespaceCmd)
}

func createTable() {
	out := parseOutput()
	out.Error((errors.New("not implemented: create table is WIP")))
}

func createNamespace(args []string) {
	ident := args[0]
	props := iceberg.Properties{}
	if cfg.Description != "" {
		props["Description"] = cfg.Description
	}

	if cfg.LocationURI != "" {
		props["Location"] = cfg.LocationURI
	}

	out, cat := parseOutput(), parseCatalog()
	err := cat.CreateNamespace(context.Background(), catalog.ToRestIdentifier(ident), props)
	if err != nil {
		out.Error(err)
		os.Exit(1)
	}
}
