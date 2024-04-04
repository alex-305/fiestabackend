package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/fiestabackend/auth"
	"github.com/alex-305/fiestabackend/helpers"
)

func (s *APIServer) handleVerifyAuth(w http.ResponseWriter, r *http.Request) {
	token, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not find jwt token", http.StatusForbidden)
	}

	user, err := auth.ValidateToken(token, s.DB)

	if err != nil {
		http.Error(w, "No valid token available", http.StatusNotFound)
	}

	userJSON, err := json.Marshal(user)

	if err != nil {
		http.Error(w, "Could not parse user json", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}
