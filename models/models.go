package models

import (
	"errors"
	"log"

	"github.com/qawarrior/serve-nt/configuration"
	"gopkg.in/mgo.v2"
)

var (
	mongoURI       string
	mongoDB        string
	mongoDBSession *mgo.Session
)

// Collector any type that can return its mongo db collection
type Collector interface {
	Collection() *mgo.Collection
}

func init() {
	err := UpdateDbSession(configuration.Properties.Data.URI, configuration.Properties.Data.DbName)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateDbSession sets internal state used by all models
func UpdateDbSession(dburi, dbname string) error {
	mongoURI = dburi
	mongoDB = dbname
	s, err := mgo.Dial(mongoURI)
	if err != nil {
		return err
	}
	mongoDBSession = s
	return nil
}

// Close cleanup internal state of package
func Close() {
	mongoDBSession.Close()
}

func checkDBVars() error {
	if mongoDB == `` || mongoURI == `` || mongoDBSession == nil {
		return errors.New("UpdateDbSession must be called before models can function")
	}
	return nil
}
