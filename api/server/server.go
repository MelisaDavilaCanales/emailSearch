package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"github.com/MelisaDavilaCanales/emailSearch/api/config"
)

type Server struct {
	Router *chi.Mux
}

func NewServer() Server {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(middleware.CleanPath)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{config.CLIENT_API_URL},
		AllowedMethods:   []string{http.MethodGet, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	AddRoutes(router)

	server := Server{
		Router: router,
	}

	return server
}

func AddRoutes(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Welcome to the API")) //nolint:errcheck
	})

	router.NotFound(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("404 Not Found")) //nolint:errcheck
	})

	AddEmailRoutes(router)
	AddPersonRoutes(router)
}

func (s *Server) Run() error {
	port := config.API_PORT
	fmt.Println("Server running on port:", port)

	return http.ListenAndServe(":"+port, s.Router)
}
