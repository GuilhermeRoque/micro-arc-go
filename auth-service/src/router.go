package main

import (
	"auth-service/src/users"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

var corsOptions = cors.Options{
	AllowedOrigins:   []string{"https://*", "http://*"},
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	ExposedHeaders:   []string{"Link"},
	AllowCredentials: true,
	MaxAge:           300,
}

func GetUserIdFromURL(url string) (int, error) {
	userIdString := strings.TrimPrefix(url, fmt.Sprintf("%s/", users.UsersEndpoint))
	userIdInt, err := strconv.Atoi(userIdString)
	return userIdInt, err
}

func NewRouter(conn *sql.DB) *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(corsOptions))
	userController := users.NewUserController(conn)
	users.SetUsersRoutes(mux, userController)

	return mux
}
