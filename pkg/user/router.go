package user

import "github.com/go-chi/chi/v5"

func UserRoutes(s UserService) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", createUserHandler(s))
	router.Get("/", getUsersHandler(s))
	router.Get("/{id}", getUserHandler(s))
	return router
}
