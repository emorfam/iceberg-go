/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/apache/iceberg-go/catalog"
	"github.com/apache/iceberg-go/table"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
)

func parseOutput() Output {
	switch strings.ToLower(cfg.Output) {
	case "text":
		return Text{}
	case "json":
		fallthrough
	default:
		log.Fatalf("Output type %s not implemented.", cfg.Output)
		return nil
	}
}

func parseCatalog() catalog.Catalog {
	var cat catalog.Catalog
	var err error

	switch catalog.CatalogType(cfg.Catalog) {
	case catalog.REST:
		opts := []catalog.Option[catalog.RestCatalog]{}
		if len(cfg.Cred) > 0 {
			opts = append(opts, catalog.WithCredential(cfg.Cred))
		}

		if len(cfg.Warehouse) > 0 {
			opts = append(opts, catalog.WithWarehouseLocation(cfg.Warehouse))
		}


		if cat, err = catalog.NewRestCatalog("rest", cfg.URI, opts...); err != nil {
			log.Fatal(err)
		}
	case catalog.Glue:
		awscfg, err := awsconfig.LoadDefaultConfig(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		opts := []catalog.Option[catalog.GlueCatalog]{
			catalog.WithAwsConfig(awscfg),
		}
		cat = catalog.NewGlueCatalog(opts...)
	default:
		log.Fatalf("Catalog type %s not implemented.", cfg.Catalog)
	}
	return cat
}

func loadTable(out Output, cat catalog.Catalog, id string) *table.Table {
	tbl, err := cat.LoadTable(context.Background(), catalog.ToRestIdentifier(id), nil)
	if err != nil {
		out.Error(err)
		os.Exit(1)
	}

	return tbl
}