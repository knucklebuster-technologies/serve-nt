package main

import (
	"flag"
	"os"

	"github.com/qawarrior/serve-nt/configuration"
)

func main() {
	wd, _ := os.Getwd()
	cpath := flag.String("config", wd+"/config.json", "Path to the config.json file")
	flag.Parse()

	cfg, err := configuration.FromFile(*cpath)
	if err != nil {
		panic(err)
	}

	cfg.Logger.Info.Println("Starting: ", cfg.AppName, " Address: ", cfg.Server.Address)
	startServer(cfg)
}
