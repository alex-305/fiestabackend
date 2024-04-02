package api

import (
	"github.com/alex-305/fiestabackend/db"
	"github.com/alex-305/fiestabackend/handlers"
)

func CreateServer(listenAddress string, db *db.DB) handlers.APIServer {
	server := handlers.APIServer{
		ListenAddress: listenAddress,
		DB:            db,
	}
	return server
}
