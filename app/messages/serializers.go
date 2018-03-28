package messages

import (
	"errors"
	"time"
)

type MessageRequest struct {
	SenderID    int       `json:"senderID"`
	RecipientID int       `json:"recipientID"`
	Body        string    `json:"body"`
	Timestamp   time.Time `json:"timestamp"`
}

func (jm MessageRequest) Message() Message {
	var m Message
	m.SenderID = jm.SenderID
	m.RecipientID = jm.RecipientID
	m.Body = jm.Body
	m.Timestamp = jm.Timestamp

	return m
}

func (jm *MessageRequest) validate() error {
	if jm.SenderID <= 0 {
		return errors.New("SenderID should not be empty")
	}
	if jm.RecipientID <= 0 {
		return errors.New("RecipientID should not be empty")
	}
	if jm.Body <= "" {
		return errors.New("Body should not be empty")
	}

	return nil
}

type MessageResponse struct {
	ID          int       `json:"id"`
	SenderID    int       `json:"senderId"`
	RecipientID int       `json:"recipientId"`
	Body        string    `json:"body"`
	Timestamp   time.Time `json:"timstamp"`
	Sender      MessageSenderResponse
	Recipient   MessageRecipientResponse
}

// func (m Message) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(MessageResponse(m))
// }

func Response(m Message) MessageResponse {
	var jm MessageResponse
	jm.SenderID = m.SenderID
	jm.RecipientID = m.RecipientID
	jm.Body = m.Body
	jm.Timestamp = m.Timestamp
	jm.Sender.ID = m.Sender.ID
	jm.Sender.Username = m.Sender.Username
	jm.Recipient.ID = m.Recipient.ID
	jm.Recipient.Username = m.Recipient.Username

	return jm
}

type MessageSenderResponse struct {
	ID       int    `json:"id"`
	Username string `json:username"`
}

type MessageRecipientResponse struct {
	ID       int    `json:"id"`
	Username string `json:username"`
}
