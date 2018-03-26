package messages

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func (model *User) getUsers(db *sql.DB) (User, error) {
	rows, err := db.Query(
		`SELECT *
        FROM messages
        WHERE user_id = $1`,
		m.userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := []Message{}

	for rows.Next() {
		var m Message
		if err := rows.Scan(
			&m.ID,
			&m.Body,
			&m.RecipientID,
			&m.SenderID); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}

func (model *User) getProfile(db *sql.DB) (User, error) {
	return db.QueryRow(
		`SELECT * // join with followings
                FROM messages
                WHERE id=$1`, m.ID).Scan(
		&m.SenderID,
		&m.RecipientID,
		&m.Body,
		&m.IsRead)
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

func AddFollowing(db *sql.DB, m Message) error {

	if err != nil {
		return nil, err
	}

	return nil
}

func RemoveFollowing(db *sql.DB, m Message) error {

	if err != nil {
		return nil, err
	}

	return nil
}
