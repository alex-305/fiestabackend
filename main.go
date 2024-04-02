package main

func main() {

	/*Start the server*/
	server := createAPIServer("localhost:8080")
	err := server.Start()

	if err != nil {
		panic(err)
	}

	/*Open database*/

	db, err := NewDB()

	if err != nil {
		panic(err)
	}

	defer db.Close()

}
