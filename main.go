package main

import (
	"github.com/qawarrior/serve-nt/configuration"
	"github.com/qawarrior/serve-nt/handlers"
)

func main() {

	// m, err := startMongod(configuration.Properties.Data.DbPath)
	// if err != nil {
	// 	lerror.Fatal(err)
	// }
	// defer m.Process.Kill()
	// configuration.Linfo.Println("Service mongod has been started")

	router := handlers.GetHandler()
	configuration.Linfo.Println("Routes and their handlers have been created")

	configuration.Linfo.Println("Web server address " + configuration.Properties.Server.Address)
	startWebserver(configuration.Properties.Server.Address, router)
}
