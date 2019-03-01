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
	
	for i, wr := range cfg.Web {
		time 
	}
}


// for 1 s ticker check if this web resource need to be pushed to chan
// 
