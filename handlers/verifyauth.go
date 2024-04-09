package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alex-305/fiestabackend/auth"
	"github.com/alex-305/fiestabackend/helpers"
)

func (s *APIServer) handleVerifyAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	token, err := helpers.GetToken(r)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Could not find jwt token", http.StatusForbidden)
		return
	}

	user, err := auth.ValidateToken(token, s.DB)

	if err != nil {
		http.Error(w, "No valid token available", http.StatusNotFound)
		return
	}

	userJSON, err := json.Marshal(user)

	if err != nil {
		http.Error(w, "Could not parse user json", http.StatusInternalServerError)
		return
	}
	log.Printf("Successfully verified %s", user.Username)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}
