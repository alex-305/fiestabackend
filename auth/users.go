package auth

import (
	"log"
	"strings"

	"github.com/alex-305/fiestabackend/db"
	"github.com/alex-305/fiestabackend/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(creds models.Credentials, db *db.DB) error {
	dbPass, err := db.GetPassword(creds.Username)
	if err != nil {
		log.Printf("Error getting password: %v", err)
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(creds.Password))
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil
}

func CreateUser(creds models.Credentials, db *db.DB) error {
	hashedPassword, err := hashPassword(creds.Password)

	if err != nil {
		return err
	}
	newUser := models.Credentials{
		Username: strings.ToLower(creds.Username),
		Password: hashedPassword,
	}

	err = db.CreateUser(newUser)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func IsUser(username, token string, db *db.DB) bool {
	user, err := ValidateToken(token, db)

	if err != nil || username != user.Username {
		return false
	}

	return true
}
