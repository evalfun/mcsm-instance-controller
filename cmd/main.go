package main

import (
	"flag"

	"github.com/evalfun/mcsm-instance-controller/internal/app"
)

var configPath string

func main() {

	flag.StringVar(&configPath, "config", "config.json", "config file name and path")

	flag.Parse()

	app.StartMCSMController(configPath)

}
