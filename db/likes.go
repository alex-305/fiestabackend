package db

import (
	"log"
)

func (db *DB) DidUserLike(liker, fiestaid string) bool {
	query := `
	SELECT username
	FROM user_likes_fiesta
	WHERE username = $1
	AND fiestaid = $2`

	var username string

	log.Printf("liker:%sfiestaid:%s", liker, fiestaid)

	row := db.QueryRow(query, liker, fiestaid).Scan(&username)

	return row == nil
}

func (db *DB) LikeFiesta(username, fiestaid string) error {
	query := `
	INSERT INTO user_likes_fiesta(username, fiestaid)
	VALUES($1, $2);`

	_, err := db.Exec(query, username, fiestaid)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UnLikeFiesta(username, fiestaid string) error {
	query := `
	DELETE FROM user_likes_fiesta
	WHERE username = $1
	AND fiestaid = $2;`

	_, err := db.Exec(query, username, fiestaid)

	if err != nil {
		log.Printf("%s", err)
		return err
	}

	return nil
}

func (db *DB) LikeCount(fiestaid string) int {
	query := `
	SELECT COUNT(*) FROM user_likes_fiesta
	WHERE fiestaid = $1;`

	var count int
	_ = db.QueryRow(query, fiestaid).Scan(&count)
	log.Printf("count:%d", count)
	return count
}
