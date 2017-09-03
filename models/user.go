package models

import "gopkg.in/mgo.v2/bson"

type user struct {
	ID            bson.ObjectId `json:"id"  bson:"_id"`
	Email         string        `json:"email" bson:"email"`
	Username      string        `json:"usernamre"  bson:"username"`
	Password      string        `json:"password"  bson:"password"`
	Firstname     string        `json:"firstname" bson:"firstname"`
	Lastname      string        `json:"lastname" bson:"lastname"`
	SessionID     string        `json:"sessionid"  bson:"session_id"`
	StreetAddress string        `json:"streetAddress"  bson:"street_address"`
	City          string        `json:"city"  bson:"city"`
	State         string        `json:"state"  bson:"state"`
	ZipCode       int64         `json:"zipCode"  bson:"zip_code"`
}
