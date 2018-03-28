package messages

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func GetSentMessagesForUser(db *sql.DB, m []Message, userID int) ([]Message, error) {
	query := `
        SELECT m.id, m.body, m.timestamp, u.username, u.id
        FROM messages as m
            INNER JOIN users as u
            ON m.recipient_id = u.id
        WHERE m.sender_id = $1
        ORDER BY m.timestamp;
    `

	rows, err := db.Query(query, userID)

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

func GetRecievedMessages(db *sql.DB, m Message) ([]Message, error) {

	query := `
        SELECT m.id, m.body, m.timestamp, u.username, u.id
        FROM messages
            INNER JOIN users
            ON m.sender_id = u.id
        WHERE m.recipient_id = $1
        ORDER BY m.timestamp
        `

	rows, err := db.Query(query, m.ID)

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

// func GetMessage(db *sql.DB, m Message) error {

// 	query := `
//         SELECT m.id, m.body, m.timestamp, u.username, u.id
//         FROM messages
//             INNER JOIN users
//             ON m.sender_id = u.id
//         WHERE m.id = $1`

// 	return db.QueryRow(query, m.ID).Scan(
// 		&m.SenderID,
// 		&m.RecipientID,
// 		&m.Body)
// }

// func SendMssages(db *sql.DB, m Message) error {

// 	if err != nil {
// 		return nil, err
// 	}

// 	return nil
// }

// func DeleteMssages(db *sql.DB, m Message) error {

// 	if err != nil {
// 		return nil, err
// 	}

// 	return nil
// }
