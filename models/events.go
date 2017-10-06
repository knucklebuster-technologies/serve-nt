package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Event data type for events returned by an EventsCollection
type Event struct {
	ID          bson.ObjectId `json:"id"  bson:"_id"`
	Title       string        `json:"title"  bson:"title"`
	Description string        `json:"description"  bson:"description"`
	ServeeID    bson.ObjectId `json:"serveeid" bson:"serveeid"`
	ServentID   bson.ObjectId `json:"serventid, omitempty" bson:"serventid, omitempty"`
}

// EventsCollection type for managing the events collection
type EventsCollection struct {
	data *mgo.Collection
}

// Insert places types into the collection
func (c EventsCollection) Insert(m *Event) (*Event, error) {
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

// Find returns a slice of *Event
func (c EventsCollection) Find(q map[string]interface{}) (*[]Event, error) {
	m := &[]Event{}
	err := c.data.Find(q).All(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Delete removes items that match query
func (c EventsCollection) Delete(q map[string]interface{}) error {
	return nil
}

// NewEventsCollection returns a ready to use EventsCollection
func NewEventsCollection() *EventsCollection {
	col := mongo.Copy().DB(cfg.Database.Name).C("events")
	idx := newIndex([]string{`title`, `serveeid`})
	col.EnsureIndex(idx)
	return &EventsCollection{
		data: col,
	}
}
