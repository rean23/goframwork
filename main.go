package main

import (
	"goframwork/bootstrap"
	"goframwork/config"
	c"goframwork/pkg/config"
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
		Addr:    ":" + c.GetString("app.port"),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
