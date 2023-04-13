package main

import (
	"broker/src/api"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8081"

func main() {
	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: api.NewRouter(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
