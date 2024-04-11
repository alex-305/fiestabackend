package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alex-305/fiestabackend/auth"
	"github.com/alex-305/fiestabackend/helpers"
	"github.com/alex-305/fiestabackend/models"
	"github.com/gorilla/mux"
)

func (s *APIServer) handleGetComments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	fiestaid := vars["fiestaid"]

	comments, err := s.DB.GetComments(fiestaid)

	if err != nil {
		http.Error(w, "Could not get comments", http.StatusBadRequest)
		return
	}

	commentJSON, err := json.Marshal(comments)

	if err != nil {
		http.Error(w, "Could not parse JSON for response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(commentJSON)

}

func (s *APIServer) handlePostComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not find jwt token", http.StatusUnauthorized)
		return
	}

	user, err := auth.ValidateToken(token, s.DB)

	if err != nil {
		http.Error(w, "Could not authenticate jwt token", http.StatusUnauthorized)
		return
	}

	var comment models.Comment
	comment.Username = user.Username

	err = json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		http.Error(w, "Could not decode json", http.StatusBadRequest)
		return
	}

	commentid, err := s.DB.PostComment(user.Username, comment.Content, comment.Fiestaid)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	comment.ID = commentid

	response, err := json.Marshal(comment)

	if err != nil {
		http.Error(w, "Could not parse JSON for response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
