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
	// linfo.Println("Service mongod has been started")

	router := handlers.GetHandler()
	linfo.Println("Routes and their handlers have been created")

	linfo.Println("Web server address " + configuration.Properties.Server.Address)
	startWebserver(configuration.Properties.Server.Address, router)
}
