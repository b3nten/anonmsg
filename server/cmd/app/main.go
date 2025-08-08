package main

import (
	"flag"
	"fmt"
	"os"

	"benton.codes/anonmsg/cfg"
	"benton.codes/anonmsg/internal/app"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "config_dev.yaml", "Path to the configuration file")
}

func main() {
	c, err := cfg.Init(configFile)
	if err != nil {
		fmt.Printf("Error initializing configuration: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(c)
	app.Run(c)
}
