package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo(ctx context.Context) (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))
	clientOptions.SetAuth(options.Credential{

		Username:      os.Getenv("DB_USER"),
		Password:      os.Getenv("DB_PASSWORD"),
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    "admin",
	})

	log.Printf("Connecting in database with auth options %+v", clientOptions.Auth)

	// connect
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	log.Println("Connected to mongo!")

	return client, nil

}
