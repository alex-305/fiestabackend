package handlers

import (
	"log"
	"net/http"

	"github.com/alex-305/fiestabackend/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type APIServer struct {
	ListenAddress string
	DB            *db.DB
}

func (s *APIServer) Start() error {
	router := mux.NewRouter()
	s.defineRoutes(router)

	log.Printf("Server is listening on %s", s.ListenAddress)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(router)

	return http.ListenAndServe(s.ListenAddress, corsHandler)
}
