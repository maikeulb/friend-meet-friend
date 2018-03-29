package messages

import (
	"errors"
	"time"
)

type MessageRequest struct {
	SenderID    int       `json:"senderID,omitempty"`
	RecipientID int       `json:"recipientID,omitempty"`
	Body        string    `json:"body,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
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
	ID          int                      `json:"id,omitempty"`
	SenderID    int                      `json:"senderId,omitempty"`
	RecipientID int                      `json:"recipientId,omitempty"`
	Body        string                   `json:"body,omitempty"`
	Timestamp   time.Time                `json:"timestamp,omitempty"`
	Sender      MessageSenderResponse    `json:"sender,omitempty"`
	Recipient   MessageRecipientResponse `json:"recipient,omitempty"`
}

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
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

type MessageRecipientResponse struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}
