package users

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func (model *User) getUsers(db *sql.DB) (User, error) {
	rows, err := db.Query(
		`SELECT u.id, u.username, u.last_active, f.follower_id, f.followee_id
        FROM users
        INNER JOIN followings 
        ON m.sender_id = u.id`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.LastActive); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return messages, nil
}

func (model *User) getProfile(db *sql.DB) (User, error) {
	return db.QueryRow(
		`SELECT u.id, u.username, u.last_active, u.bio, u.created_on, f.followr_id, f.followee_id
                FROM users
                INNER JOIN followings
                WHERE id=$1`, m.ID).Scan(
		&m.ID,
		&m.Username,
		&m.Bio,
		&m.CreatedOn,
		&m.LastActive)
}

func (model *User) editProfile(db *sql.DB) (User, error) {
	return db.QueryRow(
		`SELECT * // join with followings
                FROM messages
                WHERE id=$1`, m.ID).Scan(
		&m.SenderID,
		&m.RecipientID,
		&m.Body,
		&m.IsRead)
}
