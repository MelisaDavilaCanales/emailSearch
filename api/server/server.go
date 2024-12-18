package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"backend/config"
	"backend/emails"
	"backend/persons"
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
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	server := Server{
		Router: router,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the API"))
	})

	emails.AddEmailRoutes(router)
	persons.AddPersonRoutes(router)

	return server
}

func (s *Server) Run() error {
	port := config.API_PORT
	fmt.Println("Server running on port:", port)

	return http.ListenAndServe(":"+port, s.Router)
}
