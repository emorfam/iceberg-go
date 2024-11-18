/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/apache/iceberg-go/catalog"
	"github.com/apache/iceberg-go/cli/output"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [namespace]",
	Short: "Lists namespaces or tables.",
	Long: `Lists namespaces or tables. If a namespace is provided, tables in the namespace are listed.`,
	Example: `- goiceberg list --uri http://localhost:8181
- goiceberg list nyc --uri http://localhost:8181`,
	Run: func(cmd *cobra.Command, args []string) {
		parent := ""
		if len(args) > 0 {
			parent = args[0]
		}
		if cfg.URI == "" {
			fmt.Println("URI missing, please provide using --uri, the config or environment variable PYICEBERG_CATALOG__DEFAULT__URI")
			return
		}
		out, cat := parseOutput(), parseCatalog()
		list(out, cat, parent)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(out output.Output, cat catalog.Catalog, parent string) {
	prnt := catalog.ToRestIdentifier(parent)

	ids, err := cat.ListNamespaces(context.Background(), prnt)
	if err != nil {
		out.Error(err)
		os.Exit(1)
	}

	if len(ids) == 0 && parent != "" {
		ids, err = cat.ListTables(context.Background(), prnt)
		if err != nil {
			out.Error(err)
			os.Exit(1)
		}
	}
	out.Identifiers(ids)
}