package db

import (
	"github.com/alex-305/fiestabackend/models"
)

func (db *DB) GetUserFiestas(username string) ([]models.SmallFiesta, error) {

	query := `
	SELECT f.title, f.username, f.id, f.post_date, i.url
	FROM fiestas f
	JOIN (
		SELECT fiestaid, MIN(url) AS url
		FROM images
		GROUP BY fiestaid
	)
	i ON f.id = i.fiestaid
	WHERE f.username = $1
	ORDER BY f.post_date DESC
	LIMIT 20;`

	return db.GetFiestaList(query, username)
}

func (db *DB) GetLatestFiestas(username string) ([]models.SmallFiesta, error) {
	query := `
	SELECT f.title, f.username, f.id, f.post_date, i.url
	FROM fiestas f
	JOIN (
		SELECT fiestaid, MIN(url) AS url
		FROM images
		GROUP BY fiestaid
	)
	i ON f.id = i.fiestaid
	WHERE username <> $1
	ORDER BY f.post_date DESC
	LIMIT 20;`

	return db.GetFiestaList(query, username)
}

func (db *DB) GetFollowingFiestas(username string) ([]models.SmallFiesta, error) {
	query := `
	SELECT f.title, f.username, f.id, f.post_date, i.url
	FROM fiestas f
	JOIN user_follows_user as fl ON f.username = fl.followee
	JOIN (
		SELECT fiestaid, MIN(url) AS url
		FROM images
		GROUP BY fiestaid
	)
	AS i ON f.id = i.fiestaid
	WHERE fl.follower = $1
	ORDER BY f.post_date DESC
	LIMIT 20;`

	return db.GetFiestaList(query, username)
}

func (db *DB) GetFiestaList(query, username string) ([]models.SmallFiesta, error) {
	rows, err := db.Query(query, username)

	if err != nil {
		return []models.SmallFiesta{}, err
	}

	defer rows.Close()

	var fiestas []models.SmallFiesta

	for rows.Next() {
		var fiesta models.SmallFiesta

		err := rows.Scan(&fiesta.Title, &fiesta.Username, &fiesta.ID, &fiesta.PostDate, &fiesta.CoverImageURL)

		fiestas = append(fiestas, fiesta)

		if err != nil {
			return []models.SmallFiesta{}, err
		}
	}

	return fiestas, nil
}
