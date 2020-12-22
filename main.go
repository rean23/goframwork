package main

import (
	"fmt"
	"goframwork/pkg/config"
)

func main() {
	//router := bootstrap.SetupRouter();
	fmt.Println(config.Get("app.a"));
	/*srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe();*/
}
