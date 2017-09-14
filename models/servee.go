package models

import (
	"encoding/json"
	"io"
)

// Servee the one posting an event that needs to be filled
type Servee struct {
	User
}

// NewServee initialize a Servee value
func NewServee() {}

// Encode writes the structs value to a stream
func (a *Servee) Encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(a)
}

// Decode reads a stream and assigns values to the structs properties
func (a *Servee) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(a)
}

// Create will create a DBRecord for this Model
func (a *Servee) Create(v ...interface{}) (Servee, error) {
	return *a, nil
}
