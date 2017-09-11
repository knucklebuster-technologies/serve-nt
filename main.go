package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/qawarrior/loggy"
	"github.com/qawarrior/serve-nt/routes"
	"github.com/qawarrior/srvcontrol"
	mgo "gopkg.in/mgo.v2"
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
	config, err := readConfiguration(wdir + `\config.json`)
	if err != nil {
		loggy.Fatal(err)
	}

	loggy.Info("STARTING DATABASE SERVER")
	err = srvcontrol.StartDataSrvr(wdir + `\datastore`)
	if err != nil {
		loggy.Fatal(err)
	}
	defer srvcontrol.StopDataSrvr()

	loggy.Info("CREATING DATABASE SESSION")
	dbsession, err := mgo.Dial(config.Data.URI)
	if err != nil {
		loggy.Fatal(err)
	}
	defer dbsession.Close()

	loggy.Info("SETTING UP ROUTING")
	router, err := routes.Set(config.Data.DbName, dbsession)
	if err != nil {
		loggy.Fatal(err)
	}

	loggy.Info("STARTING SERVER @ " + config.Server.Address)
	srvcontrol.StartWebSrvr(config.Server.Address, router)
}

type configuration struct {
	ProjectEnv string `json:"projectEnv"`
	AppName    string `json:"appName"`
	AppURI     string `json:"appURI"`
	Version    string `json:"version"`
	Server     struct {
		Address string `json:"address"`
	} `json:"server"`
	Data struct {
		URI    string `json:"uri"`
		DbName string `json:"dbName"`
	} `json:"data"`
}

func readConfiguration(path string) (*configuration, error) {
	c := &configuration{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(file, c)
	if err != nil {
		return c, err
	}
	return c, nil
}
