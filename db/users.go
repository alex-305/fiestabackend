package db

import (
	"database/sql"
	"log"

	"github.com/alex-305/fiestabackend/models"
)

func (db *DB) CreateUser(creds models.Credentials) error {

	_, err := db.GetUser(creds.Username)

	if err == nil {
		return err
	}

	stmt := `
	INSERT INTO users(username, password)
	VALUES($1,$2);`

	_, err = db.Exec(stmt, creds.Username, creds.Password)

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

func (db *DB) UpdateDescription(username, description string) error {
	stmt, err := db.Prepare("UPDATE users SET description = $1 WHERE username = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(description, username)

	if err != nil {
		return err
	}
	return nil
}
