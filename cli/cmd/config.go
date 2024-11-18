/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

type Config struct {
	Catalog     string
	URI         string
	Output      string
	History     bool
	Cred        string
	Warehouse   string
	Config      string
	Description string
	LocationURI string
}

var cfg Config 