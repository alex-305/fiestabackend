package db

import (
	"log"

	"github.com/alex-305/fiestabackend/helpers"
	"github.com/alex-305/fiestabackend/models"
)

func (db *DB) CreateFiesta(fiesta models.Fiesta) (string, error) {
	_, err := db.GetUser(fiesta.Username)
	log.Printf(fiesta.Username + " is attempting to make a fiesta")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	fiestaid := helpers.GenerateRandString(10)

	query := `
	INSERT INTO fiestas(username, title, id)
	VALUES($1, $2, $3);`

	_, err = db.Exec(query, fiesta.Username, fiesta.Title, fiestaid)

	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	for _, url := range fiesta.Images {

		log.Printf("fiestaid:%s", fiestaid)
		log.Printf("url:%s", url)

		_, err = db.Exec(`
		UPDATE images
		SET fiestaid = $1
		WHERE url = $2;`, fiestaid, url)

		if err != nil {
			return "", err
		}
	}

	return fiestaid, nil
}

func (db *DB) GetFiesta(fiestaid string) (models.Fiesta, error) {
	query := `
	SELECT username, title, post_date FROM fiestas
	WHERE id = $1`

	row := db.QueryRow(query, fiestaid)

	fiesta := models.Fiesta{}

	err := row.Scan(&fiesta.Username, &fiesta.Title, &fiesta.Post_date)

	if err != nil {

		return models.Fiesta{}, err
	}

	query = `
	SELECT url FROM images
	WHERE fiestaid = $1`

	rows, err := db.Query(query, fiestaid)

	if err != nil {
		log.Printf("second:%s", err)
		return models.Fiesta{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var image string
		err := rows.Scan(&image)

		fiesta.Images = append(fiesta.Images, image)

		if err != nil {
			return models.Fiesta{}, err
		}
	}

	return fiesta, nil
}

func (db *DB) RemoveFiesta(fiestaid string) error {
	query := `
	DELETE FROM fiestas
	WHERE id = $1`

	_, err := db.Exec(query, fiestaid)

	if err != nil {
		log.Printf("%s", err)
		return err
	}

	return nil
}
