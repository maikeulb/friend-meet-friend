package messages

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	ID          int
	Body        string
	Timestamp   time.Time
	SenderID    int
	RecipientID int
	Sender      MessageUser
	Recipient   MessageUser
}

type MessageUser struct {
	ID   int
	Name string
}

type MessagesSent struct {
	ID          int
	RecipientID int
}

type MessagesRecieved struct {
	ID       int
	SenderID int
}

func (m Message) MarshalJSON() ([]byte, error) {
	fmt.Println("marshalling")
	return json.Marshal(Response(m))
}

func (m *Message) UnmarshalJSON(data []byte) error {
	fmt.Println("unmarshalling")
	var jm MessageRequest

	if err := json.Unmarshal(data, &jm); err != nil {
		return err
	}
	if err := jm.validate(); err != nil {
		return err
	}

	*m = jm.Message(*m)
	return nil
}
