package models

import "io"
import "gopkg.in/mgo.v2"

var (
	dbURI     string
	dbnName   string
	dbSession *mgo.Session
)

// Modeler interface for models to marshal and decode JSON
type Modeler interface {
	Encode(io.Writer) error
	Decode(io.Reader) error
}

// DataDetailer returns db specific data
type DataDetailer interface {
	CollectionName() string
	DBSession() *mgo.Session
}

// Init sets internal state of package to all the structs to operate correctly
func Init(dburi, dbname string) error {
	dbURI = dburi
	dbnName = dbname
	s, err := mgo.Dial(dbURI)
	if err != nil {
		return err
	}
	dbSession = s
	return nil
}

// Close cleanup internal state of package
func Close() error {
	dbSession.Close()
	return nil
}
