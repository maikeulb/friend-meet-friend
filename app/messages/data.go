package messages

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func GetMessage(db *sql.DB, m Message) error {
	return db.QueryRow(
		`SELECT *
        FROM messages
        WHERE id=$1`, m.ID).Scan(
		&m.SenderID,
		&m.RecipientID,
		&m.Body,
		&m.IsRead)
}

func GetUserMssages(db *sql.DB, m Message) (Message, error) {
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
