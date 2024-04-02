package main

import (
	"github.com/alex-305/fiestabackend/api"
	"github.com/alex-305/fiestabackend/db"
)

func main() {

	/*Open database*/
	db, err := db.NewDB()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	/*Start the server*/
	server := api.CreateServer("localhost:8080", db)
	err = server.Start()

	if err != nil {
		panic(err)
	}

}
