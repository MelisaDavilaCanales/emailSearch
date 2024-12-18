package persons

import "github.com/go-chi/chi/v5"

func AddPersonRoutes(router chi.Router) {
	router.Get("/persons", GetPersonsHandler)
}
