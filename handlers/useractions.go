package handlers

import (
	"log"
	"net/http"

	"github.com/alex-305/fiestabackend/auth"
	"github.com/alex-305/fiestabackend/helpers"
	"github.com/gorilla/mux"
)

func (s *APIServer) handlePostFollow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	followee := vars["followee"]

	log.Printf("followee: %s", followee)

	token, err := helpers.GetToken(r)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "JWT token could not be found", http.StatusUnauthorized)
		return
	}

	user, err := auth.ValidateToken(token, s.DB)

	if err != nil {
		http.Error(w, "Could not validate JWT token", http.StatusUnauthorized)
		return
	}

	follower := user.Username
	if s.DB.IsUserFollowing(follower, followee) {
		err = s.DB.UserUnfollowsUser(follower, followee)
	} else {
		err = s.DB.UserFollowsUser(follower, followee)
	}

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
