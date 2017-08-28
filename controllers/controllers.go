package controllers

import (
	"net/http"
)

// Controller defines a type that controls an http endpoint
type Controller interface {
	Create(http.ResponseWriter, *http.Request)
	Read(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}
