package main

import (
	"flag"
	_ "fmt"
)

var (
	confFile = flag.String("config", "./config.yaml", "specify config file")
	rc       *RuntimeConfig
)

func main() {
	flag.Parse()
	cfg := ParseConfig(*confFile)
	rc = CreateRuntimeConfig(cfg)
}

