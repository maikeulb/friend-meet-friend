package messages

import (
	"database/sql"
	// "fmt"

	_ "github.com/lib/pq"
)

func GetMessageForUser(db *sql.DB, m *Message, userID int) error {
	query := `
        SELECT m.id,
        m.body,
        m.timestamp,
        u.id,
        u.name,
        u2.id,
        u2.name
        FROM messages as m
        INNER JOIN users as u
        ON m.recipient_id = u.id
        INNER JOIN users as u2
        ON m.sender_id = u2.id
        WHERE m.id = $2 and m.sender_id = $1`

	return db.QueryRow(query, userID, m.ID).Scan(
		&m.ID,
		&m.Body,
		&m.Timestamp,
		&m.Recipient.ID,
		&m.Recipient.Name,
		&m.Sender.ID,
		&m.Sender.Name)
}

func GetSentMessagesForUser(db *sql.DB, m []Message, userID int) ([]Message, error) {
	query := `
        SELECT m.id,
        m.body,
        m.timestamp,
        u.id,
        u.name
        FROM messages as m
        INNER JOIN users as u
        ON m.recipient_id = u.id
        WHERE m.sender_id = $1
        ORDER BY m.timestamp;`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := []Message{}

	for rows.Next() {
		var m = Message{}
		if err := rows.Scan(
			&m.ID,
			&m.Body,
			&m.Timestamp,
			&m.Recipient.ID,
			&m.Recipient.Name); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}
	return messages, nil
}

func GetRecievedMessagesForUser(db *sql.DB, m []Message, userID int) ([]Message, error) {
	query := `
        SELECT m.id,
        m.body,
        m.timestamp,
        u.id,
        u.name
        FROM messages as m
        INNER JOIN users as u
        ON m.sender_id = u.id
        WHERE m.recipient_id = $1
        ORDER BY m.timestamp
        `

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := []Message{}

	for rows.Next() {
		var m = Message{}
		if err := rows.Scan(
			&m.ID,
			&m.Body,
			&m.Timestamp,
			&m.Sender.ID,
			&m.Sender.Name); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}

func SendMessageToUser(db *sql.DB, m *Message) error {
	command := `
        INSERT INTO messages (body, sender_id, recipient_id, timestamp)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	err := db.QueryRow(command, m.Body, m.SenderID, m.RecipientID, m.Timestamp).Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}
