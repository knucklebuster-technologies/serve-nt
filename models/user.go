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

func (m *User) FindByLogin(e, p string) {
	m.collection.Find(bson.M{"email": e, "password": p}).One(&m)
}
