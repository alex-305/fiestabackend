package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	username := creds.Username
	password := creds.Password

	log.Printf("user: %s, password: %s", username, password)

	// user, err := s.getUser(username, password)

	// if err != nil {
	// 	http.Error(w, "Invalid username or password.", http.StatusUnauthorized)
	// 	return
	// }
	// hashedPassword := hashPassword(password)

	// if user.password != hashedPassword {
	// 	http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	// 	return
	// }
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	db.CreateUser(creds)

}

func (s *APIServer) Start() error {
	router := mux.NewRouter()
	router.HandleFunc("/login", s.handleLogin).Methods(http.MethodPost)
	log.Printf("Server is listening on %s", s.listenAddress)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(router)

	return http.ListenAndServe(s.listenAddress, corsHandler)
}
