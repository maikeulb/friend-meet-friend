package messages

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
)

type MessageRequest struct {
	SenderID    int       `json:"senderID"`
	RecipientID int       `json:"recipientID"`
	Body        string    `json:"body"`
	Timestamp   time.time `json:"timestamp"`
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
}

func Response(m Message) MessageResponse {
	var jm MessageResponse
	jm.SenderID = m.SenderID
	jm.RecipientID = m.RecipientID
	jm.Body = m.Body
	jm.Timestamp = m.Timestamp

	return jm
}

func (m Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(MessageResonse(m))
}

func (m *Message) UnmarshalJSON(data []byte) error {
	var jm MessageRequest

	if err := json.Unmarshal(data, &jm); err != nil {
		return err // panic?
	}
	if err := jm.validate(); err != nil {
		return err
	}

	*m = jm.Message()
	return nil
}
