package main

import (
	"net/http"
	"time"

	"github.com/qawarrior/serve-nt/configuration"
	"github.com/qawarrior/serve-nt/handlers"
)

func startServer(c *configuration.Config) error {
	wsrvr := &http.Server{
		Handler:      handlers.New(c),
		Addr:         c.Server.Address,
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
