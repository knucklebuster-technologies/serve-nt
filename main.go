package main

import (
	"github.com/qawarrior/serve-nt/configuration"
	"github.com/qawarrior/serve-nt/handlers"
)

func main() {
	configuration.Linfo.Println("Starting - ", configuration.Properties.AppName)
	configuration.Linfo.Println("Server address " + configuration.Properties.Server.Address)
	startWebserver(configuration.Properties.Server.Address, handlers.GetHandler())
}
