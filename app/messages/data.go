package messages

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/maikeulb/friend-meet-friend/app/models"
)

func (m *models.Message) getSentMessages(db *sql.DB) (models.Message, error) {
	query := `
        SELECT m.id, m.body, m.timestamp, u.username, u.id
        FROM messages as m
            INNER JOIN users as u
            ON m.recipient_id = u.id
        WHERE m.sender_id = $1
        ORDER BY m.timestamp;
    `

	rows, err := db.Query(query, m.userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := []models.Message{}

	for rows.Next() {
		var m models.Message
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

func (m *models.Message) getRecievedMessages(db *sql.DB) (models.Message, error) {

	query := `
        SELECT m.id, m.body, m.timestamp, u.username, u.id
        FROM messages
            INNER JOIN users
            ON m.sender_id = u.id
        WHERE m.recipient_id = $1
        ORDER BY m.timestamp
        `

	rows, err := db.Query(query, m.userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := []models.Message{}

	for rows.Next() {
		var m models.Message
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

func (m *models.Message) getMessage(db *sql.DB) error {

	query := `
        SELECT m.id, m.body, m.timestamp, u.username, u.id
        FROM messages
            INNER JOIN users
            ON m.sender_id = u.id
        WHERE m.id = $1`

	return db.QueryRow(query, m.ID).Scan(
		&m.SenderID,
		&m.RecipientID,
		&m.Body)
}

func (m *models.Message) sendMssages(db *sql.DB) error {

	if err != nil {
		return nil, err
	}

	return nil
}

func (m *models.Message) deleteMssages(db *sql.DB) error {

	if err != nil {
		return nil, err
	}

	return nil
}
