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
	loggy.Info.Println("STARTING MAIN")

	loggy.Info.Println("GETTING WORKING DIRECTORY")
	wdir, err := os.Getwd()
	if err != nil {
		loggy.Fatal(err)
	}
	loggy.Info.Println("WORKING DIRECTORY:", wdir)

	loggy.Info.Println("READING CONFIGURATION")
	config := config.Config{}
	err = config.Read(wdir + `\config.json`)
	if err != nil {
		loggy.Fatal(err)
	}

	loggy.Info.Println("STARTING DATABASE SERVER")
	dbsrv := database.NewServer(wdir + `\db`)
	err = dbsrv.Start()
	if err != nil {
		loggy.Fatal(err)
	}
	defer dbsrv.Stop()

	loggy.Info.Println("CREATING DATABASE SESSION")
	err = dbsrv.Connect(config.Data.URI)
	if err != nil {
		loggy.Fatal(err)
	}
	defer dbsrv.Session.Close()

	loggy.Info.Println("SETTING UP ROUTING")
	router, err := routes.Set(config.Data.DbName, dbsrv.Session)
	if err != nil {
		loggy.Fatal(err)
	}

	loggy.Info.Println("STARTING SERVER @ " + config.Server.Address)
	webserver.Start(config.Server.Address, router)
	defer webserver.Stop()
}
