package api

import (
	"os"

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

func NewRouter() *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(corsOptions))
	setRoutes(mux)
	return mux
}

func setRoutes(mux *chi.Mux) {
	authServiceUrl := os.Getenv("AUTH_SERVICE_URL")
	logServiceUrl := os.Getenv("LOGGER_SERVICE_URL")

	mux.Post("/login", ProxyRequestHost(authServiceUrl))

	mux.Post("/users", ProxyRequestHost(authServiceUrl))
	mux.Get("/users", ProxyRequestHost(authServiceUrl))

	mux.Put("/users/:userId", ProxyRequestHost(authServiceUrl))
	mux.Delete("/users/:userId", ProxyRequestHost(authServiceUrl))
	mux.Get("/users/:userId", ProxyRequestHost(authServiceUrl))

	mux.Post("/echo", Echo)

	mux.Post("/logs", ProxyRequestHost(logServiceUrl))
}
