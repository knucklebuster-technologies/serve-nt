package main

import (
	"net/http"
	"time"
)

func startWebserver(addr string, hndlr http.Handler) error {
	wsrvr := &http.Server{
		Handler:      hndlr,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	defer wsrvr.Shutdown(nil)
	err := wsrvr.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
