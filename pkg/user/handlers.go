package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func createUserHandler(s UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		createdUser, err := s.Create(&newUser)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if err := json.NewEncoder(w).Encode(createdUser); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
}

func getUsersHandler(s UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		role := r.URL.Query().Get("role")
		users, err := s.GetAll(Role(role))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if err := json.NewEncoder(w).Encode(users); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

	}
}

func getUserHandler(s UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		user, err := s.GetOneByID(id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

	}
}
