package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *APIServer) defineRoutes(router *mux.Router) {
	//Auth Endpoints
	router.HandleFunc("/login", s.handleLogin).Methods(http.MethodPost)
	router.HandleFunc("/createAccount", s.handleCreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/user/{username}", s.handleGetUser).Methods(http.MethodGet)
	router.HandleFunc("/auth/verify", s.handleVerifyAuth).Methods(http.MethodGet)
	router.HandleFunc("/user/{username}/update", s.handleUserUpdate).Methods(http.MethodPost)
	//Fiesta Endpoints
	router.HandleFunc("/fiesta", s.handlePostFiesta).Methods(http.MethodPost)
	router.HandleFunc("/fiesta", s.handleRecentFiestas).Methods(http.MethodGet)
	router.HandleFunc("/user/{username}/fiesta/{fiestaID}", s.handleGetFiesta).Methods(http.MethodGet)
	router.HandleFunc("/fiesta/{fiestaID}/comment", s.handleFiestaComment).Methods(http.MethodPost)
	router.HandleFunc("/user/{username}/fiesta", s.handleGetUserFiestas).Methods(http.MethodGet)
	//image endpoints
	router.HandleFunc("/image", s.handlePostImage).Methods(http.MethodPost)
	router.HandleFunc("/image/{imageURL}", s.handleRemoveImage).Methods(http.MethodDelete)
	router.HandleFunc("/image/{imageURL}", s.handleGetImage).Methods(http.MethodGet)
	//User action endpoints
	router.HandleFunc("/user/follows/{followee}", s.handlePostFollow).Methods(http.MethodPost)
}
