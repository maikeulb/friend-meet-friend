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

func (jm MessageRequest) Message(m Message) Message {
	m.RecipientID = jm.RecipientID
	m.Body = jm.Body
	m.Timestamp = time.Now()

	return m
}

func (jm *MessageRequest) validate() error {
	if jm.RecipientID <= 0 {
		return errors.New("RecipientID must not be empty")
	}
	if jm.Body <= "" {
		return errors.New("Body must not be empty")
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
	jm.Sender.Name = m.Sender.Name
	jm.Recipient.ID = m.Recipient.ID
	jm.Recipient.Name = m.Recipient.Name

	return jm
}

type MessageSenderResponse struct {
	ID       int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type MessageRecipientResponse struct {
	ID       int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
