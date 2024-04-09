package db

import "github.com/alex-305/fiestabackend/models"

func (db *DB) AddImage(image models.Image) error {

	_, err := db.GetUser(image.Username)

	if err != nil {
		return err
	}

	query := `
	INSERT INTO images(username, url)
	VALUES($1, $2);
	`
	_, err = db.Exec(query, image.Username, image.Url)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) DeleteImage(image models.Image) error {
	_, err := db.GetUser(image.Username)

	if err != nil {
		return err
	}

	query := `
	DELETE FROM images
	WHERE username = $1
	AND url = $2;
	`

	_, err = db.Exec(query, image.Username, image.Url)

	if err != nil {
		return err
	}

	return nil
}
