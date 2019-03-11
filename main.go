package main

import (
	"flag"
	_ "fmt"
	"time"
)

var (
	confFile = flag.String("config", "./config.yaml", "specify config file")
	cfg *Config
)

func main() {
	flag.Parse()
	cfg = ParseConfig(*confFile)
	
	inputChan := runWorkers(5)
	
	ticker := time.NewTicker(1 * time.Second)
	for range  ticker.C {
		for _, wr := range cfg.Web {
			inputChan <- wr		
		}
	}
}
