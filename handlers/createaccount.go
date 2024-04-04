package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/fiestabackend/auth"
	"github.com/alex-305/fiestabackend/models"
)

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	err = auth.CreateUser(creds, s.DB)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	token, err := auth.GenerateJWT(creds.Username)

	if err != nil {
		http.Error(w, "Error generating JWToken", http.StatusInternalServerError)
	}

	response := map[string]string{"token": token}
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Could not send JWToken", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
