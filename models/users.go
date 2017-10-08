package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User basic system user
type User struct {
	ID        bson.ObjectId `json:"id"  bson:"_id"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password"  bson:"password"`
	Firstname string        `json:"firstname" bson:"firstname"`
	Lastname  string        `json:"lastname" bson:"lastname"`
	ZipCode   int64         `json:"zipcode"  bson:"zipcode"`
}

// UsersCollection type for managing the events collection
type UsersCollection struct {
	data *mgo.Collection
}

// Insert places types into the collection
func (c UsersCollection) Insert(m *User) (*User, error) {
	if m.ID != `` || bson.IsObjectIdHex(string(m.ID)) {
		return m, nil
	}
	m.ID = bson.NewObjectId()
	err := c.data.Insert(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Find returns a slice of of types found in collection
func (c UsersCollection) Find(q map[string]interface{}) (*[]User, error) {
	m := &[]User{}
	err := c.data.Find(q).All(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// FindOne returns first match in the collectiom
func (c UsersCollection) FindOne(q map[string]interface{}) (*User, error) {
	m := &User{}
	err := c.data.Find(q).One(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete removes items that match query
func (c UsersCollection) Delete(q map[string]interface{}) error {
	return nil
}

// NewUsersCollection returns a ready to use EventsCollection
func NewUsersCollection() *UsersCollection {
	col := mongo.Copy().DB(cfg.Database.Name).C("users")
	idx := newIndex([]string{`email`})
	col.EnsureIndex(idx)
	return &UsersCollection{
		data: col,
	}
}
