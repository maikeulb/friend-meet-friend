package messages

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func GetSentMessagesForUser(db *sql.DB, m []*Message, userID int) ([]*Message, error) {
	query := `
        SELECT m.id, 
            m.body, 
            m.timestamp, 
            u.id, 
            u.username
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

	messages := []*Message{}

	for rows.Next() {
		var m = &Message{}
		if err := rows.Scan(
			&m.ID,
			&m.Body,
			&m.Timestamp,
			&m.Recipient.ID,
			&m.Recipient.Username); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}

func GetRecievedMessagesForUser(db *sql.DB, m []*Message, userID int) ([]*Message, error) {

	query := `
        SELECT m.id, 
            m.body, 
            m.timestamp, 
            u.id, 
            u.username
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

	messages := []*Message{}

	for rows.Next() {
		var m = &Message{}
		if err := rows.Scan(
			&m.ID,
			&m.Body,
			&m.Timestamp,
			&m.Sender.ID,
			&m.Sender.Username); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}

func GetMessageForUser(db *sql.DB, m *Message, userID int) error {

	query := `
        SELECT m.id, 
            m.body, 
            m.timestamp,
            u.id,
            u.username,
            u2.id,
            u2.username
        FROM messages as m
            INNER JOIN users as u
            ON m.recipient_id = u.id
            INNER JOIN users as u2
            ON m.sender_id = u2.id
        WHERE m.id = $2 and m.sender_id = $1`

	return db.QueryRow(query, userID, m.ID).Scan( // should I check error?
		&m.ID,
		&m.Body,
		&m.Timestamp,
		&m.Recipient.ID,
		&m.Recipient.Username,
		&m.Sender.ID,
		&m.Sender.Username)
}

// func SendMssages(db *sql.DB, m Message) error {

//  if err != nil {
//      return nil, err
//  }

//  return nil
// }

// func DeleteMssages(db *sql.DB, m Message) error {

//  if err != nil {
//      return nil, err
//  }

//  return nil
// }
