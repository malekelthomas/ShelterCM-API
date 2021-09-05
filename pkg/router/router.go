package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/malekelthomas/ShelterCM-API/pkg/store/postgres"
	"github.com/malekelthomas/ShelterCM-API/pkg/user"
)

type server struct {
	r *chi.Mux
}

func NewServer() *server {

	us := postgres.NewUserStore()
	router := chi.NewRouter()
	router.Mount("/api/users", user.UserRoutes(us))
	s := server{
		r: router,
	}
	return &s
}

func (s *server) ListenAndServe() error {
	if err := http.ListenAndServe("", s.r); err != nil {
		return err
	}
	return nil
}
