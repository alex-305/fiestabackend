package db

import (
	"database/sql"

	"github.com/alex-305/fiestabackend/models"
)

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

	return GetFiestaList(query, username, db.DB)
}

func (db *DB) GetLatestFiestas(username string) ([]models.SmallFiesta, error) {
	query := `SELECT f.title, f.username, f.id, i.url
	FROM fiestas f
	JOIN (SELECT fiestaid, url,
	ROW_NUMBER() OVER (PARTITION BY fiestaid ORDER BY url) AS row_num
	FROM images)
	i ON f.id = i.fiestaid AND i.row_num=1
	WHERE username <> $1
	ORDER BY post_date DESC
	LIMIT 20;`

	return GetFiestaList(query, username, db.DB)
}

func (db *DB) GetFollowingFiestas(username string) ([]models.SmallFiesta, error) {
	query := `SELECT f.title, f.username, f.id, i.url
	FROM fiestas f
	JOIN user_follows_user as fl ON f.username = fl.followee
	JOIN (
		SELECT fiestaid, url,
		ROW_NUMBER() OVER (PARTITION BY fiestaid ORDER BY url) AS row_num
		FROM images
	) AS i ON f.id = i.fiestaid AND i.row_num=1
	WHERE fl.follower = $1
	ORDER BY post_date DESC
	LIMIT 20;`

	return GetFiestaList(query, username, db.DB)
}

func GetFiestaList(query, username string, db *sql.DB) ([]models.SmallFiesta, error) {
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
