package models

import (
	"encoding/json"
	"io"
)

// Servent the person that can take on an event to complete
type Servent struct {
	User
}

// Encode writes the structs value to a stream
func (a *Servent) Encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(a)
}

// Decode reads a stream and assigns values to the structs properties
func (a *Servent) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(a)
}
