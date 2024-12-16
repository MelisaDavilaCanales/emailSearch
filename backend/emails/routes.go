package emails

import "github.com/go-chi/chi/v5"

func AddEmailRoutes(router chi.Router) {
	router.Get("/emails", GetAllEmails)
	router.Get("/emails/{id}", GetEmail)
	router.Get("/searchEmails", SearchEmails)
}
