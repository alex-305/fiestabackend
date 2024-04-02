package db

import (
	"log"

	"github.com/alex-305/fiestabackend/models"
)

func (db *DB) CreateUser(cred models.Credentials) error {

	stmt := `
	INSERT INTO users(username, password)
	VALUES($1,$2);`

	_, err := db.Exec(stmt, cred.Username, cred.Password)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetPassword(username string) (string, error) {
	var password string
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&password)
	log.Printf("dbpass: %s", password)
	if err != nil {
		return "", err
	}

	return password, nil
}
