package handlers

import (
	"fmt"
	"net/http"
)

func profileidGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to your profile")
}
