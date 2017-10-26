package model

import "time"

// PageData simple structure used in templates
type PageData struct {
	Timestamp time.Time
	AppName   string
}

// ProfileData composite structure used in profile template
type ProfileData struct {
	PageData PageData
	User     User
	Events   *[]Event
}
