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

func (db *DB) GetFiesta(fiestaDetails models.FiestaDetails) (models.Fiesta, error) {
	query := `
	SELECT username, title, post_date FROM fiestas
	WHERE username = $1
	AND id = $2`

	row := db.QueryRow(query, fiestaDetails.Username, fiestaDetails.FiestaID)

	fiesta := models.Fiesta{}

	err := row.Scan(&fiesta.Username, &fiesta.Title, &fiesta.Post_date)

	if err != nil {
		log.Printf("username: %s, fiestaid: %s", fiestaDetails.Username, fiestaDetails.FiestaID)
		log.Printf("first:%s", err)
		return models.Fiesta{}, err
	}

	query = `
	SELECT url FROM images
	WHERE fiestaid = $1`

	rows, err := db.Query(query, fiestaDetails.FiestaID)

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

func (db *DB) GetUserFiestas(username string) ([]models.SmallFiesta, error) {

	query := `SELECT f.title, f.username, f.id, i.url
	FROM fiestas f
	JOIN (SELECT fiestaid, url,
	ROW_NUMBER() OVER (PARTITION BY fiestaid ORDER BY url) AS row_num
	FROM images)
	i ON f.id = i.fiestaid AND i.row_num=1
	WHERE f.username = $1
	ORDER BY post_date DESC
	LIMIT 20;`

	rows, err := db.Query(query, username)

	if err != nil {
		return []models.SmallFiesta{}, err
	}

	defer rows.Close()

	var fiestas []models.SmallFiesta

	for rows.Next() {
		var fiesta models.SmallFiesta

		err := rows.Scan(&fiesta.Title, &fiesta.Username, &fiesta.ID, &fiesta.CoverImageURL)

		fiesta.Username = username

		fiestas = append(fiestas, fiesta)

		if err != nil {
			return []models.SmallFiesta{}, err
		}
	}

	return fiestas, nil
}

func (db *DB) GetRecentFiestas(username string) ([]models.SmallFiesta, error) {
	query := `SELECT f.title, f.username, f.id, i.url
	FROM fiestas f
	JOIN (SELECT fiestaid, url,
	ROW_NUMBER() OVER (PARTITION BY fiestaid ORDER BY url) AS row_num
	FROM images)
	i ON f.id = i.fiestaid AND i.row_num=1
	WHERE username <> $1
	ORDER BY post_date DESC
	LIMIT 20;`

	rows, err := db.Query(query, username)

	if err != nil {
		return []models.SmallFiesta{}, err
	}

	defer rows.Close()

	var fiestas []models.SmallFiesta

	for rows.Next() {
		var fiesta models.SmallFiesta

		err := rows.Scan(&fiesta.Title, &fiesta.Username, &fiesta.ID, &fiesta.CoverImageURL)

		fiestas = append(fiestas, fiesta)

		if err != nil {
			return []models.SmallFiesta{}, err
		}
	}

	return fiestas, nil
}
