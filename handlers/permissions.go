package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alex-305/fiestabackend/auth"
	"github.com/alex-305/fiestabackend/helpers"
	"github.com/gorilla/mux"
)

func (s *APIServer) handlePostPermissions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not available", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	fiestaid := vars["fiestaid"]

	token, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not find JWT token", http.StatusMethodNotAllowed)
		return
	}

	user, err := auth.ValidateToken(token, s.DB)

	if err != nil || !auth.GetPermissions(user.Username, fiestaid, s.DB).IsOwner {
		log.Printf("%s", err)
		http.Error(w, "User is not authorized to add permissions to this fiesta", http.StatusUnauthorized)
		return
	}

	var userToAdd struct {
		Username string `json:"username"`
	}

	err = json.NewDecoder(r.Body).Decode(&userToAdd)

	log.Printf("user:%s", userToAdd)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Error decoding request JSON", http.StatusBadRequest)
		return
	}

	if user.Username == userToAdd.Username {
		log.Printf("%s", err)
		http.Error(w, "User already owns fiesta", http.StatusBadRequest)
		return
	}

	err = s.DB.AddPermission(userToAdd.Username, fiestaid)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (s *APIServer) handleGetPermissions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not available", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	fiestaid := vars["fiestaid"]

	permissionList, err := s.DB.GetPermissions(fiestaid)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	type permission struct {
		Usernames []string `json:"usernames"`
	}

	permissions := permission{
		Usernames: permissionList,
	}

	permissionsJSON, err := json.Marshal(permissions)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Could not parse JSON for response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(permissionsJSON)
}

func (s *APIServer) handleDeletePermissions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not available", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	fiestaid := vars["fiestaid"]
	userToRemove := vars["username"]

	token, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not find JWT token", http.StatusMethodNotAllowed)
		return
	}

	user, err := auth.ValidateToken(token, s.DB)

	if err != nil || !auth.GetPermissions(user.Username, fiestaid, s.DB).IsOwner {
		http.Error(w, "User is not authorized to add permissions to this fiesta", http.StatusUnauthorized)
		return
	}

	err = s.DB.RevokePermission(userToRemove, fiestaid)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
