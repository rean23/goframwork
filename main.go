package main

import (
	"goframwork/bootstrap"
	"goframwork/config"
	"net/http"
	"time"
)

func init() {
	config.Initialize()
}
func main() {
	router := bootstrap.SetupRouter()
	bootstrap.ConnectDB()

	srv := &http.Server{
		Handler: router,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
