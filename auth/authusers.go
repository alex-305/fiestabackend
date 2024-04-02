package auth

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/alex-305/fiestabackend/db"
	"github.com/alex-305/fiestabackend/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

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
	creds.Password = hashedPassword
	err = db.CreateUser(creds)

	if err != nil {
		return err
	}
	return nil
}

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	SecretKey := os.Getenv("SECRET_KEY")

	if SecretKey == "" {
		return "", errors.New("secret key could not be retrieved")
	}

	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
