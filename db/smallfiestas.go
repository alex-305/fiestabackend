package db

import (
	"database/sql"

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
	LIMIT 50;`

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
	LIMIT 50;`

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
	LIMIT 50;`

	return db.GetFiestaList(query, username)
}

func (db *DB) GetPopularFiestas() ([]models.SmallFiesta, error) {
	query := `
	SELECT f.title, f.username, f.id, f.post_date, i.url
	FROM fiestas f
	
	JOIN (
		SELECT fiestaid, MIN(url) AS url
		FROM images
		GROUP BY fiestaid
	) 
	AS i ON f.id = i.fiestaid
	JOIN comments c ON f.id = c.fiestaid
	JOIN user_likes_fiesta l ON f.id = l.fiestaid
	GROUP BY f.title, f.username, f.id, f.post_date, i.url
	ORDER BY COUNT(l.username) + COUNT(c.id) DESC
	LIMIT 50;`

	return db.GetFiestaList(query, "")

}

func (db *DB) GetFiestaList(query, username string) ([]models.SmallFiesta, error) {
	var rows *sql.Rows
	var err error
	if username != "" {
		rows, err = db.Query(query, username)
	} else {
		rows, err = db.Query(query)
	}

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
