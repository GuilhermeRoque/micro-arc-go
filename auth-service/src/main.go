package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5"
)

const webPort = "8082"

func main() {
	dbConn, err := GetConn()
	if err != nil {
		panic(err)
	}

	router := NewRouter(dbConn)

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
