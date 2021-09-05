package user

import (
	"encoding/json"
	"net/http"
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
