package models

import "gopkg.in/mgo.v2/bson"
import "gopkg.in/mgo.v2"

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

// NewUser return an initialized User type
func NewUser() *User {
	c := mongoDBSession.Copy().DB(mongoDB).C(`Users`)
	return &User{
		c: c,
	}
}

// FindByEmail query the data for a user by email
func (m *User) FindByEmail() error {
	err := m.c.Find(bson.M{"email": m.Email}).One(m)
	if err != nil {
		return err
	}
	return nil
}
