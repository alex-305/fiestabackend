package main

type APIServer struct {
	listenAddress string
}

func createAPIServer(listenAddress string) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
	}
}
