package db

func (db *DB) IsUserFollowing(follower, followee string) bool {
	query := `
	SELECT follower
	FROM user_follows_user
	WHERE follower = $1
	AND followee = $2`
	var follow string
	row := db.QueryRow(query, follower, followee).Scan(&follow)

	return row == nil

}

func (db *DB) FollowUser(follower, followee string) error {
	query := `
	INSERT into user_follows_user(follower, followee)
	VALUES($1,$2);
	`
	_, err := db.Exec(query, follower, followee)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UnfollowUser(follower, followee string) error {
	query := `
	DELETE FROM user_follows_user
	WHERE follower = $1
	AND followee = $2;
	`
	_, err := db.Exec(query, follower, followee)

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) FollowerCount(user string) int {
	query := `
	SELECT COUNT(*)
	FROM user_follows_user
	WHERE followee = $1`

	var count int
	_ = db.QueryRow(query, user).Scan(&count)
	return count
}

func (db *DB) FollowingCount(user string) int {
	query := `
	SELECT COUNT(*)
	FROM user_follows_user
	WHERE follower = $1`

	var count int
	_ = db.QueryRow(query, user).Scan(&count)
	return count
}
