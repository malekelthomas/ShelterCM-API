package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/malekelthomas/ShelterCM-API/pkg/config"
	"github.com/malekelthomas/ShelterCM-API/pkg/signature"
	"github.com/malekelthomas/ShelterCM-API/pkg/store/postgres"
	"github.com/malekelthomas/ShelterCM-API/pkg/user"
)

type server struct {
	r *chi.Mux
}

func NewServer() *server {

	us := postgres.NewUserStore()
	ss := postgres.NewSignatureStore()
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Mount("/api/users", user.UserRoutes(us))
	router.Mount("/api/signatures", signature.SignatureRoutes(ss))
	s := server{
		r: router,
	}
	return &s
}

func (s *server) ListenAndServe() error {
	log.Printf("Starting server on port :%d\n", config.Config.Server.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Server.Port), s.r); err != nil {
		return err
	}
	return nil
}
