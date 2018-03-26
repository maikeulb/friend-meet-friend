package messages

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func (model *Message) getSentMessages(db *sql.DB) (Message, error) {
	rows, err := db.Query(
		`SELECT m.id, m.body, m.timestamp, u.username, u.id
        FROM messages
        INNER JOIN users 
        ON m.sender_id = u.id
        WHERE m.sender_id = $1
        ORDERBY m.timestamp;`,
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

func (model *Message) getRecievedMessages(db *sql.DB) (Message, error) {
	rows, err := db.Query(
		`SELECT m.id, m.body, m.timestamp, u.username, u.id 
		 FROM messages
         INNER JOIN users
         ON m.recipient_id = u.id
	     WHERE m.recipient_id = $1
		 ORDERBY m.timestamp`,
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

func (model *Message) getMessage(db *sql.DB) error {
	return db.QueryRow(
		`SELECT m.id, m.body, m.timestamp, u.username, u.id
        FROM messages
        WHERE id = $1`,
		m.ID).Scan(
		&m.SenderID,
		&m.RecipientID,
		&m.Body)
}

func (model *Message) sendMssages(db *sql.DB) error {

	if err != nil {
		return nil, err
	}

	return nil
}

func (model *Message) deleteMssages(db *sql.DB) error {

	if err != nil {
		return nil, err
	}

	return nil
}
