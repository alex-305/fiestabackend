package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alex-305/fiestabackend/auth"
	"github.com/alex-305/fiestabackend/helpers"
	"github.com/alex-305/fiestabackend/models"
	"github.com/gorilla/mux"
)

type userProfile struct {
	User        models.User
	CanEdit     bool
	IsFollowing bool
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	username := vars["username"]

	log.Printf("handling user: %s", username)
	log.Printf("vars: %s", vars)

	user, err := s.DB.GetUser(username)

	if err != nil {
		http.Error(w, "Could not find user", http.StatusBadRequest)
		return
	}
	token, err := helpers.GetToken(r)

	canEdit := false
	isFollowing := false
	if err == nil {
		log.Printf("Can edit!")
		canEdit = auth.IsUser(username, token, s.DB)
	}

	if !canEdit {
		requestUser, err := auth.ValidateToken(token, s.DB)
		if err == nil {
			isFollowing = s.DB.IsUserFollowing(requestUser.Username, username)
			log.Printf("%t", isFollowing)
		}
	}

	var userProf userProfile = userProfile{
		User:        user,
		CanEdit:     canEdit,
		IsFollowing: isFollowing,
	}

	userJSON, err := json.Marshal(userProf)

	if err != nil {
		http.Error(w, "Could not parse JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

type UserUpdate struct {
	Description string `json:"Description"`
}

func (s *APIServer) handleUserUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]

	log.Printf("handling user: %s", username)
	log.Printf("vars: %s", vars)

	token, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "No JWT token found", http.StatusUnauthorized)
		log.Printf("%s", err)
		return
	}

	user, err := auth.ValidateToken(token, s.DB)

	if err != nil {
		http.Error(w, "Invalid JWT token", http.StatusUnauthorized)
		log.Printf("%s", err)
		return
	}
	var userUpdate UserUpdate

	err = json.NewDecoder(r.Body).Decode(&userUpdate)

	if err != nil {
		http.Error(w, "Could not find new value for description", http.StatusBadRequest)
		log.Printf("%s", err)
		return
	}

	err = s.DB.UpdateDescription(user.Username, userUpdate.Description)

	if err != nil {
		http.Error(w, "Database could not successfully update", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
