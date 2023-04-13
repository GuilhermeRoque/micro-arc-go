package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8085"

func main() {
	router := NewRouter()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
	log.Printf("Started on port %s\n", webPort)

}
