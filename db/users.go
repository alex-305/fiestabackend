package db

import (
	"database/sql"
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

func (db *DB) GetUser(username string) (models.User, error) {

	row := db.QueryRow("SELECT username, description, join_date FROM users WHERE username = $1", username)

	var user models.User
	var description sql.NullString
	err := row.Scan(&user.Username, &description, &user.Join_date)

	if description.Valid {
		user.Description = description.String
	} else {
		user.Description = ""
	}

	if err != nil {
		log.Printf("%s", err)
		return models.User{}, err
	}

	return user, nil
}
