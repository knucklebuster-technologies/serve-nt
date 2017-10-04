package handlers

import (
	"time"

	"github.com/qawarrior/serve-nt/models"
)

type pagedata struct {
	Timestamp time.Time
	AppName   string
}

type profiledata struct {
	pagedata
	models.User
}
