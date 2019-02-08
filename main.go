package main

import (
	"flag"
	"time"
)

var (
	cfg    *Config
	keeper *StateKeeper
)

func main() {
	confFile := flag.String("config", "./config.yaml", "Please specify YAML config file location")
	flag.Parse()
	cfg = ParseConfig(*confFile)

	i, o := DispatchWorkers(cfg.Global.Workers)
	DispatchNotificators(1, o)

	keeper = CreateKeeper(cfg.Web...)

	ticker := time.NewTicker(cfg.data.interval)
	for _ = range ticker.C {
		for _, wr := range cfg.Web {
			i <- wr
		}
	}
}
