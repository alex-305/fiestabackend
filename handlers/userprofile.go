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
	User    models.User
	CanEdit bool
}

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]

	log.Printf("handling user: %s", username)
	log.Printf("vars: %s", vars)

	user, err := s.DB.GetUser(username)

	if err != nil {
		http.Error(w, "Could not find user", http.StatusBadRequest)
		return
	}

	token, _ := helpers.GetToken(r)
	canEdit := auth.IsUser(username, token, s.DB)

	var userProf userProfile = userProfile{
		User:    user,
		CanEdit: canEdit,
	}

	userJSON, err := json.Marshal(userProf)

	if err != nil {
		http.Error(w, "Could not parse JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}
