package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8084"

func main() {
	ctx := context.Background()
	// connect to mongo
	mongoClient, err := ConnectToMongo(ctx)
	if err != nil {
		log.Panic(err)
	}

	// close connection
	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	router := NewRouter(mongoClient)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: router,
	}

	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
	log.Printf("Started on port %s\n", webPort)

}
