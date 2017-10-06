package models

import (
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User basic system user
type User struct {
	c         *mgo.Collection
	ID        bson.ObjectId `json:"id"  bson:"_id"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password"  bson:"password"`
	Firstname string        `json:"firstname" bson:"firstname"`
	Lastname  string        `json:"lastname" bson:"lastname"`
	ZipCode   int64         `json:"zipcode"  bson:"zipcode"`
}

// NewUser returns an initialized User type
func NewUser() *User {
	c := mongo.Copy().DB(cfg.Database.Name).C("users")
	i := newIndex([]string{`email`})
	c.EnsureIndex(i)
	return &User{
		c: c,
	}
}

// Find takes a standard mongo map for a query
func (m *User) Find(q map[string]interface{}) error {
	return m.c.Find(q).One(m)
}

// Insert adds the data for a user to the collection
func (m *User) Insert() error {
	if m.ID != `` || bson.IsObjectIdHex(string(m.ID)) {
		return errors.New("User is already in the collection")
	}
	m.ID = bson.NewObjectId()
	err := m.c.Insert(m)
	if err != nil {
		return errors.Wrap(err, "User data insert failed")
	}
	return nil
}
