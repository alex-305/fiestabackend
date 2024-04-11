package db

import (
	"github.com/alex-305/fiestabackend/helpers"
	"github.com/alex-305/fiestabackend/models"
)

func (db *DB) GetComments(fiestaid string) ([]models.Comment, error) {
	query := `
	SELECT * FROM comments
	WHERE fiestaid=$1
	ORDER BY post_date DESC;`

	rows, err := db.Query(query, fiestaid)

	if err != nil {
		return []models.Comment{}, err
	}

	var comments []models.Comment

	for rows.Next() {
		var comment models.Comment

		err := rows.Scan(&comment.ID, &comment.Username, &comment.Content, &comment.Fiestaid, &comment.Post_date)

		comments = append(comments, comment)

		if err != nil {
			return []models.Comment{}, err
		}
	}

	return comments, nil

}

func (db *DB) PostComment(username, content, fiestaid string) (string, error) {

	id := helpers.GenerateRandString(24)

	query := `
	INSERT into comments(username,content,fiestaid,id)
	VALUES($1, $2, $3, $4)`

	_, err := db.Exec(query, username, content, fiestaid, id)

	if err != nil {
		return "", err
	}

	return id, nil
}
