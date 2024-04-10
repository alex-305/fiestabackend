package db

func (db *DB) DidUserLike(username, fiestaid string) bool {
	query := `
	SELECT username
	FROM user_likes_fiesta
	WHERE username = $1
	AND fiestaid = $2`
	var follow string
	row := db.QueryRow(query, username, fiestaid).Scan(&follow)

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
	REMOVE FROM user_likes_fiesta
	WHERE username = $1
	AND fiestaid = $2;`

	_, err := db.Exec(query, username, fiestaid)

	if err != nil {
		return err
	}

	return nil
}
