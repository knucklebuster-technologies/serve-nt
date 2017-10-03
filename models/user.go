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
	ZipCode   int64         `json:"zipCode"  bson:"zip_code"`
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

// FindByEmail query the data for a user by email
func (m *User) FindByEmail() error {
	if m.Email == `` {
		return errors.New("User field Email must be valid")
	}
	err := m.c.Find(bson.M{"email": m.Email}).One(m)
	if err != nil {
		return err
	}
	return nil
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
