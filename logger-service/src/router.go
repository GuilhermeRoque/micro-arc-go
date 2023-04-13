package main

import (
	"fmt"
	"logger-service/src/logs"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

var corsOptions = cors.Options{
	AllowedOrigins:   []string{"https://*", "http://*"},
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	ExposedHeaders:   []string{"Link"},
	AllowCredentials: true,
	MaxAge:           300,
}

func GetIdFromURL(url string) string {
	logIdString := strings.TrimPrefix(url, fmt.Sprintf("%s/", logs.LogsEndpoint))
	return logIdString
}

func NewRouter(mongoClient *mongo.Client) *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(corsOptions))
	logController := logs.NewLogController(mongoClient)
	logs.SetRoutes(mux, logController)

	return mux
}
