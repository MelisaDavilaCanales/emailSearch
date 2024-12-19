package server

import (
	"github.com/go-chi/chi/v5"

	"github.com/MelisaDavilaCanales/emailSearch/api/controllers"
)

func AddEmailRoutes(router chi.Router) {
	router.Get("/emails", controllers.GetEmails)
	router.Get("/emails/{id}", controllers.GetEmail)
}

func AddPersonRoutes(router chi.Router) {
	router.Get("/persons", controllers.GetPersons)
}
