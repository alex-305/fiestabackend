package main

func main() {

	/*Open database*/
	db, err := NewDB()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	/*Start the server*/
	server := createAPIServer("localhost:8080", db)
	err = server.Start()

	if err != nil {
		panic(err)
	}

}
