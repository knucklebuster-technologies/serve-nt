package models

import (
	"github.com/qawarrior/serve-nt/configuration"
	"gopkg.in/mgo.v2"
)

var cfg *configuration.Config
var mongo *mgo.Session

// Init sets up the models package for data processing
func Init(c *configuration.Config) error {
	m, err := mgo.Dial(c.Database.URI)
	if err != nil {
		return err
	}
	mongo = m
	cfg = c
	return nil
}

// INTERNAL FUNCTIONS

func newIndex(k []string) mgo.Index {
	i := mgo.Index{
		Key:        k,
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	return i
}
