package main

type APIServer struct {
	listenAddress string
	db            *DB
}

func createAPIServer(listenAddress string, db *DB) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		db:            db,
	}
}
