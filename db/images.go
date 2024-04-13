package db

import "github.com/alex-305/fiestabackend/models"

func (db *DB) AddImage(image models.Image, fiestaid string) error {
	var err error
	if fiestaid == "" {
		_, err = db.Exec(`INSERT INTO images(username, url) 
		VALUES($1, $2)`, image.Username, image.Url)
	} else {
		_, err = db.Exec(`INSERT INTO images(username, url, fiestaid)
		VALUES($1, $2, $3)`, image.Username, image.Url, fiestaid)
	}

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) DeleteImage(image models.Image) error {

	query := `
	DELETE FROM images
	WHERE username = $1
	AND url = $2;
	`

	_, err := db.Exec(query, image.Username, image.Url)

	if err != nil {
		return err
	}

	return nil
}
