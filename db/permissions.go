package db

import "log"

func (db *DB) HasPermission(username, fiestaid string) bool {
	query := `
	SELECT username FROM fiesta_post_permissions
	WHERE username = $1
	AND fiestaid = $2;`

	var perms string
	row := db.QueryRow(query, username, fiestaid).Scan(&perms)

	return row == nil
}

func (db *DB) IsOwner(user, fiestaid string) bool {
	query := `
	SELECT username FROM fiestas
	WHERE id = $1`

	var username string
	err := db.QueryRow(query, fiestaid).Scan(&username)

	if err != nil {
		log.Printf("%s", err)
		return false
	}

	return user == username
}

func (db *DB) AddPermission(username, fiestaid string) error {
	query := `
	INSERT INTO fiesta_post_permissions(username, fiestaid)
	VALUES($1, $2);`

	_, err := db.Exec(query, username, fiestaid)

	return err
}

func (db *DB) GetPermissions(fiestaid string) ([]string, error) {

	var permissions []string

	query := `
	SELECT username
	FROM fiestas
	WHERE id = $1;`

	var username string

	err := db.QueryRow(query, fiestaid).Scan(&username)

	permissions = append(permissions, username)

	if err != nil {
		log.Printf("%s", err)
		return []string{}, err
	}

	query = `
	SELECT username
	FROM fiesta_post_permissions
	WHERE fiestaid = $1;`

	permissionsList, err := db.GetPermissionList(query, fiestaid)

	if err != nil {
		return []string{}, err
	}

	log.Printf("owner is %s", username)

	permissions = append(permissions, permissionsList...)

	return permissions, nil
}

func (db *DB) RevokePermission(username, fiestaid string) error {
	query := `
	DELETE FROM fiesta_post_permissions
	WHERE username = $1 AND fiestaid = $2;`

	_, err := db.Exec(query, username, fiestaid)

	return err
}

func (db *DB) GetPermissionList(query, fiestaid string) ([]string, error) {
	rows, err := db.Query(query, fiestaid)

	if err != nil {
		return []string{}, err
	}

	defer rows.Close()

	var permissions []string

	for rows.Next() {
		var permission string

		err := rows.Scan(&permission)

		permissions = append(permissions, permission)

		if err != nil {
			return []string{}, err
		}
	}

	return permissions, nil
}
