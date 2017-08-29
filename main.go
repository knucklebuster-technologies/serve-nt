package main

import (
	"os"

	"github.com/qawarrior/serve-nt/config"
	"github.com/qawarrior/serve-nt/database"
	"github.com/qawarrior/serve-nt/loggy"
	"github.com/qawarrior/serve-nt/routes"
	"github.com/qawarrior/serve-nt/webserver"
)

func main() {
	loggy.Info("STARTING MAIN")

	loggy.Info("GETTING WORKING DIRECTORY")
	wdir, err := os.Getwd()
	if err != nil {
		loggy.Fatal(err)
	}
	loggy.Info("WORKING DIRECTORY:", wdir)

	loggy.Info("READING CONFIGURATION")
	config := config.Config{}
	err = config.Read(wdir + `\config.json`)
	if err != nil {
		loggy.Fatal(err)
	}

	loggy.Info("STARTING DATABASE SERVER")
	dbsrv := database.NewServer(wdir + `\db`)
	err = dbsrv.Start()
	if err != nil {
		loggy.Fatal(err)
	}
	defer dbsrv.Stop()

	loggy.Info("CREATING DATABASE SESSION")
	err = dbsrv.Connect(config.Data.URI)
	if err != nil {
		loggy.Fatal(err)
	}
	defer dbsrv.Session.Close()

	loggy.Info("SETTING UP ROUTING")
	router, err := routes.Set(config.Data.DbName, dbsrv.Session)
	if err != nil {
		loggy.Fatal(err)
	}

	loggy.Info("STARTING SERVER @ " + config.Server.Address)
	webserver.Start(config.Server.Address, router)
	defer webserver.Stop()
}
