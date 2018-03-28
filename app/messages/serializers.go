package messages

import (
	"encoding/json"
	"errors"
	"time"
)

type MessageRequest struct {
	SenderID    int       `json:"senderID"`
	RecipientID int       `json:"recipientID"`
	Body        string    `json:"body"`
	Timestamp   time.Time `json:"timestamp"`
}

func (jm MessageRequest) Message() models.Message {
	var m models.Message
	m.SenderID = jm.SenderID
	m.RecipientID = jm.RecipientID
	m.Body = jm.Body
	m.Timestamp = jm.Timestamp

	return m
}

func (jm *MessageRequest) validate() error {
	if jm.SenderID <= "" {
		return errors.New("SenderID should not be empty")
	}
	if jm.RecipientID <= "" {
		return errors.New("RecipientID should not be empty")
	}
	if jm.Body <= "" {
		return errors.New("Body should not be empty")
	}
	if jm.Timestamp <= "" {
		return errors.New("Timestamp should not be empty")
	}

	return nil
}

type MessageResponse struct {
	ID          int       `json:"id"`
	SenderID    int       `json:"senderId"`
	RecipientID int       `json:"recipientId"`
	Body        string    `json:"body"`
	Timestamp   time.time `json:"timstamp"`
	Sender      MessageSenderResponse
	Recipient   MessageRecipientResponse
}

func Response(m models.Message) MessageResponse {
	var jm MessageResponse
	jm.SenderID = m.SenderID
	jm.RecipientID = m.RecipientID
	jm.Body = m.Body
	jm.Timestamp = m.Timestamp
	jm.Sender = m.Sender
	jm.Recipient = m.Recipient

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

func (m models.Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(MessageResonse(m))
}

func (m *models.Message) UnmarshalJSON(data []byte) error {
	var jm MessageRequest

	if err := json.Unmarshal(data, &jm); err != nil {
		return err
	}
	if err := jm.validate(); err != nil {
		return err
	}

	*m = jm.Message()
	return nil
}
