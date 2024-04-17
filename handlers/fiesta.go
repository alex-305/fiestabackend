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

type pathResponse struct {
	Path string `json:"path"`
}

func (s *APIServer) handlePostFiesta(w http.ResponseWriter, r *http.Request) {
	log.Printf("create fiesta running...")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not find JWT token", http.StatusUnauthorized)
		return
	}

	user, err := auth.ValidateToken(token, s.DB)

	if err != nil {
		http.Error(w, "Could not validate token", http.StatusUnauthorized)
		return
	}

	var fiesta models.Fiesta

	err = json.NewDecoder(r.Body).Decode(&fiesta)

	fiesta.Username = user.Username

	if err != nil {
		http.Error(w, "Could not parse JSON", http.StatusInternalServerError)
		return
	}

	fiestaid, err := s.DB.CreateFiesta(fiesta)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	response := pathResponse{
		Path: "/fiesta/" + fiestaid,
	}

	JSONresponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Could not parse JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(JSONresponse)

}

func (s *APIServer) handleGetFiesta(w http.ResponseWriter, r *http.Request) {
	log.Printf("HERE")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	fiestaID := vars["fiestaID"]

	fiesta, err := s.DB.GetFiesta(fiestaID)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	token, _ := helpers.GetToken(r)
	user, err := auth.ValidateToken(token, s.DB)

	if err != nil {
		fiesta.UserLiked = false
		fiesta.IsOwner = false
		fiesta.CanPost = false
	} else {
		fiesta.UserLiked = s.DB.DidUserLike(user.Username, fiestaID)
		fiesta.IsOwner = s.DB.IsOwner(user.Username, fiestaID)
		fiesta.CanPost = s.DB.HasPermission(user.Username, fiestaID)
		log.Printf("canpost:%t isowner: %t", fiesta.CanPost, fiesta.IsOwner)
	}

	fiesta.LikeCount = s.DB.LikeCount(fiestaID)
	jsonResponse, err := json.Marshal(fiesta)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Error parsing json for response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}

func (s *APIServer) handleDeleteFiesta(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	fiestaid := vars["fiestaid"]

	token, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not find JWT token", http.StatusUnauthorized)
		return
	}

	user, err := auth.ValidateToken(token, s.DB)

	if err != nil {
		http.Error(w, "Could not validate token", http.StatusUnauthorized)
		return
	}

	perms := auth.GetPermissions(user.Username, fiestaid, s.DB)

	if !perms.IsOwner {
		http.Error(w, "User does not have access to this resource", http.StatusUnauthorized)
		return
	}

	err = s.DB.RemoveFiesta(fiestaid)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (s *APIServer) handleGetUserFiestas(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)

	username := vars["username"]

	fiestas, err := s.DB.GetUserFiestas(username)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	fiestasJSON, err := json.Marshal(fiestas)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Response JSON could not be parsed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(fiestasJSON)

}

func (s *APIServer) handleGetFiestaList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token, err := helpers.GetToken(r)
	username := ""

	if err == nil {
		user, err := auth.ValidateToken(token, s.DB)
		if err == nil {
			username = user.Username
		}
	}

	vars := mux.Vars(r)
	listType := vars["type"]

	var fiestas []models.SmallFiesta

	switch listType {

	case "following":
		fiestas, err = s.DB.GetFollowingFiestas(username)
	case "latest":
		fiestas, err = s.DB.GetLatestFiestas(username)
	case "popular":
		fiestas, err = s.DB.GetPopularFiestas()

	}

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	fiestasJSON, err := json.Marshal(fiestas)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Response JSON could not be parsed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(fiestasJSON)
}
