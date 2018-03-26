package users

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func (model *User) getUsers(db *sql.DB) (User, error) {
	query := `
        SELECT u.id,
            u.username,
            u.last_active,
            f.follower_id,
            f.followed_id
            u2.username,
            u3.username
        FROM users as u
            INNER JOIN followings as f
            ON u.id = f.follower_id
            OR u.id = f.followed_id
            INNER JOIN users as u2
            ON u2.id = f.follower_id
            INNER JOIN users as u3
            ON u3.id = f.followed_id`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.LastActive); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	return messages, nil
}

func (model *User) getProfile(db *sql.DB) (User, error) {

	query := `
        SELECT u.id,
            u.username,
            u.last_active,
            u.bio,
            u.created_on,
            f.follower_id,
            f.followed_id,
            u2.username,
            u3.username
        FROM users u
            INNER JOIN followings as f
            ON u.id = f.follower_id
            OR u.id = f.followed_id
            INNER JOIN users as u2
            ON u2.id = f.follower_id
            INNER JOIN users as u3
            ON u3.id = f.followed_id
        WHERE id=$1`

	rows, err := db.QueryRow(query, m.ID).Scan(
		&m.ID,
		&m.Username,
		&m.Bio,
		&m.CreatedOn,
		&m.LastActive)

	if err == sql.ErrNoRows {
		log.Printf("No users")
	}

	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func (model *User) editProfile(db *sql.DB) (User, error) {

	query := `

        SELECT u.id,
            u.username,
            u.last_active,
            u.bio,
            u.created_on,
            f.follower_id,
            f.followed_id,
            u2.username,
            u3.username
        FROM users u
            INNER JOIN followings as f
            ON u.id = f.follower_id
            OR u.id = f.followed_id
            INNER JOIN users as u2
            ON u2.id = f.follower_id
            INNER JOIN users as u3
            ON u3.id = f.followed_id
        WHERE id=$1`

	return db.QueryRow(query, m.ID).Scan(
		&m.SenderID,
		&m.RecipientID,
		&m.Body,
		&m.IsRead)

	if err == sql.ErrNoRows {
		log.Printf("No users")
	}

	if err != nil {
		log.Fatal(err)
	}
}
