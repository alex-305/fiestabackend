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
		Path: "/user/" + fiesta.Username + "/fiesta/" + fiestaid,
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
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]
	fiestaID := vars["fiestaID"]

	fiestaDetails := models.FiestaDetails{
		Username: username,
		FiestaID: fiestaID,
	}

	fiesta, err := s.DB.GetFiesta(fiestaDetails)

	log.Printf("title:%s", fiesta.Title)
	log.Printf("post_date:%s", fiesta.Post_date)

	for index := range fiesta.Images {
		log.Printf(fiesta.Images[index])
	}

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	token, _ := helpers.GetToken(r)
	fiesta.CanEdit = auth.IsUser(username, token, s.DB)

	jsonResponse, err := json.Marshal(fiesta)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Error parsing json for response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

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
