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
	router.HandleFunc("/fiesta/{fiestaID}", s.handleGetFiesta).Methods(http.MethodGet)
	router.HandleFunc("/fiesta", s.handlePostFiesta).Methods(http.MethodPost)
	router.HandleFunc("/fiesta/{fiestaid}", s.handleDeleteFiesta).Methods(http.MethodDelete)
	//Fiesta List Endpoint
	router.HandleFunc("/fiestas/{type}", s.handleGetFiestaList).Methods(http.MethodGet)
	router.HandleFunc("/user/{username}/fiesta", s.handleGetUserFiestas).Methods(http.MethodGet)
	//Image Endpoints
	router.HandleFunc("/image", s.handlePostImage).Methods(http.MethodPost)
	router.HandleFunc("/image/{imageURL}", s.handleDeleteImage).Methods(http.MethodDelete)
	router.HandleFunc("/image/{imageURL}", s.handleGetImage).Methods(http.MethodGet)
	//Follow Endpoints
	router.HandleFunc("/user/follows/{followee}", s.handlePostFollow).Methods(http.MethodPost)
	//Like Endpoints
	router.HandleFunc("/fiesta/{fiestaid}/like", s.handlePostLike).Methods(http.MethodPost)
	//Post Permissions
	router.HandleFunc("/fiesta/{fiestaid}/permissions", s.handlePostPermissions).Methods(http.MethodPost)
	router.HandleFunc("/fiesta/{fiestaid}/permissions", s.handleGetPermissions).Methods(http.MethodGet)
	router.HandleFunc("/fiesta/{fiestaid}/permissions/{username}", s.handleDeletePermissions).Methods(http.MethodDelete)
	//Comment Endpoints
	router.HandleFunc("/fiesta/{fiestaid}/comments", s.handleGetComments).Methods(http.MethodGet)
	router.HandleFunc("/fiesta/{fiestaID}/comments", s.handlePostComment).Methods(http.MethodPost)
}
