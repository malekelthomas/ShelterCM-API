package user

import "github.com/go-chi/chi/v5"

func UserRoutes(s UserService) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/create", createUserHandler(s))
	return router
}
