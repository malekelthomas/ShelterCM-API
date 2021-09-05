package signature

import (
	"encoding/json"
	"net/http"
)

func createSignatureHandler(s SignatureService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var newSig Signature
		if err := json.NewDecoder(r.Body).Decode(&newSig); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		createdSig, err := s.Create(&newSig)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if err := json.NewEncoder(w).Encode(createdSig); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
}
