package signature

import (
	"github.com/go-chi/chi/v5"
)

func SignatureRoutes(s SignatureService) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", createSignatureHandler(s))
	return router
}
