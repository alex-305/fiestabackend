package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alex-305/fiestabackend/auth"
	"github.com/alex-305/fiestabackend/models"
)

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {
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
	err = auth.LoginUser(creds, s.DB)

	if err != nil {
		http.Error(w, "Invalid username or password: ", http.StatusUnauthorized)
		return
	}

	log.Printf("Successful login: %s", creds.Username)

	token, err := auth.GenerateJWT(creds.Username)

	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error generating JWToken", http.StatusInternalServerError)
	}

	response := map[string]string{"token": token}
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err)
		http.Error(w, "Could not send JWToken", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
