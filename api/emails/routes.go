package emails

import "github.com/go-chi/chi/v5"

func AddEmailRoutes(router chi.Router) {
	router.Get("/emails", GetEmails)
	router.Get("/emails/{id}", GetEmailHandler)
}
