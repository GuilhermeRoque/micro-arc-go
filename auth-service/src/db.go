package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func GetConn() (*sql.DB, error) {
	dsn := flag.String("dsn", os.Getenv("DSN"), "connection data source name")
	return OpenDB(*dsn)
}

func OpenDB(dsn string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	for i := 0; i < 3; i++ {
		db, err = sql.Open("pgx", dsn)
		if err == nil {
			break
		}
		log.Printf("Failed connecting to database %d, trying again in 3 seconds...\n", i+1)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
