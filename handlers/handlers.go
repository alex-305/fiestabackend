package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alex-305/fiestabackend/auth"
	"github.com/alex-305/fiestabackend/db"
	"github.com/alex-305/fiestabackend/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type APIServer struct {
	ListenAddress string
	DB            *db.DB
}

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
		http.Error(w, "Error generating JWToken", http.StatusInternalServerError)
	}

	w.Header().Set("Authorization", "Bearer"+token)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Token sent successfully"))
}

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
	err = s.DB.CreateUser(creds)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *APIServer) Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/login", s.handleLogin).Methods(http.MethodPost)
	router.HandleFunc("/createAccount", s.handleCreateAccount).Methods(http.MethodPost)
	log.Printf("Server is listening on %s", s.ListenAddress)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(router)

	return http.ListenAndServe(s.ListenAddress, corsHandler)
}
