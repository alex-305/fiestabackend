package helpers

import (
	"database/sql"

	"github.com/alex-305/fiestabackend/models"
)

func GetSmallFiestaList(query, username string, db *sql.DB) ([]models.SmallFiesta, error) {
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
