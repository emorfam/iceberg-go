/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/apache/iceberg-go/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goiceberg",
	Short: "CLI for Apache Iceberg written in Go",
	Long: `CLI for Apache Iceberg written in Go`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	rootCmd.PersistentFlags().StringVarP(&cfg.Catalog, "catalog", "c", "rest", "specify the catalog name")
	rootCmd.PersistentFlags().StringVarP(&cfg.URI, "uri", "u", "", "specify the catalog URI")
	rootCmd.PersistentFlags().StringVarP(&cfg.Output, "output", "o", "text", "output type (json/text)")
	rootCmd.PersistentFlags().StringVarP(&cfg.Cred, "credential", "r", "", "specify credential for the catalog")
	rootCmd.PersistentFlags().StringVarP(&cfg.Warehouse, "warehouse", "w", "", "specify the warehouse to use")
	rootCmd.PersistentFlags().StringVarP(&cfg.Config, "config", "f", "", "specify the path to the configuration to use")
	rootCmd.PersistentFlags().StringVarP(&cfg.Description, "description", "d", "", "specify a description for the namespace")
	rootCmd.PersistentFlags().StringVarP(&cfg.LocationURI, "location_uri", "l", "", "specify a location URI for the namespace")

    cobra.OnFinalize(buildConf)
}

func buildConf() {
    fileCfg := config.ParseConfig(config.LoadConfig(cfg.Config), "default")
	if fileCfg != nil {
		mergeConf(fileCfg, &cfg)
	}	
}

func mergeConf(fileConf *config.CatalogConfig, resConfig *Config) {
	if len(resConfig.Catalog) == 0 {
		resConfig.Catalog = fileConf.Catalog
	}
	if len(resConfig.URI) == 0 {
		resConfig.URI = fileConf.URI
	}
	if len(resConfig.Output) == 0 {
		resConfig.Output = fileConf.Output
	}
	if len(resConfig.Cred) == 0 {
		resConfig.Cred = fileConf.Credential
	}
	if len(resConfig.Warehouse) == 0 {
		resConfig.Warehouse = fileConf.Warehouse
	}
}
