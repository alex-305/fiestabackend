package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/alex-305/fiestabackend/auth"
	"github.com/alex-305/fiestabackend/helpers"
	"github.com/alex-305/fiestabackend/models"
	"github.com/gorilla/mux"
)

const IMAGE_DIR_PATH = "./uploads/images/"

func (s *APIServer) handleGetImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	imageURL := vars["imageURL"]

	imagePath := filepath.Join(IMAGE_DIR_PATH, imageURL)

	imageFile, err := os.Open(imagePath)

	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	defer imageFile.Close()

	w.Header().Set("Content-Type", "image/png")

	_, err = io.Copy(w, imageFile)

	if err != nil {
		http.Error(w, "Error serving the image", http.StatusInternalServerError)
		return
	}

}

func (s *APIServer) handleDeleteImage(w http.ResponseWriter, r *http.Request) {
	log.Printf("We're here!")
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	imageURL := vars["imageURL"]

	token, err := helpers.GetToken(r)

	if err != nil {
		http.Error(w, "Could not find JWT Token", http.StatusUnauthorized)
		log.Printf("%s", err)
		return
	}

	user, err := auth.ValidateToken(token, s.DB)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Could not validate JWT Token", http.StatusUnauthorized)
		return
	}

	image := models.Image{
		Username: user.Username,
		Url:      imageURL,
	}

	err = os.Remove(IMAGE_DIR_PATH + image.Url)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Could not find image", http.StatusNotFound)
		return
	}

	err = s.DB.DeleteImage(image)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Could not delete image from database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)

}

type imageJSON struct {
	Url string `json:"Url"`
}

func (s *APIServer) handlePostImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token, err := helpers.GetToken(r)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Token not found", http.StatusUnauthorized)
		return
	}

	user, err := auth.ValidateToken(token, s.DB)

	fiestaid := r.FormValue("fiestaid")
	if fiestaid != "" {
		perms := auth.GetPermissions(user.Username, fiestaid, s.DB)
		if !perms.IsOwner && !perms.CanPost {
			http.Error(w, "User is not authorized", http.StatusUnauthorized)
		}
	}

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Could not validate token", http.StatusUnauthorized)
		return
	}

	err = r.ParseMultipartForm(6 << 20) //this makes file max 6mb

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	userImg, handler, err := r.FormFile("image")

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}

	defer userImg.Close()

	filename := user.Username + "_" + helpers.GenerateRandString(40) + handler.Filename

	file, err := os.OpenFile(IMAGE_DIR_PATH+filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Error opening image", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, userImg)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	imgFile := models.Image{
		Username: user.Username,
		Url:      filename,
	}

	err = s.DB.AddImage(imgFile, fiestaid)

	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	imgJSON := imageJSON{
		Url: filename,
	}

	jsonResponse, err := json.Marshal(imgJSON)

	if err != nil {
		log.Printf("%s", err)
		http.Error(w, "Error parsing json for response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
