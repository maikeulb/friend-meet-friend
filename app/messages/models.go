package messages

import (
	"encoding/json"
	"time"
)

type Message struct {
	ID          int
	SenderID    int
	RecipientID int
	Body        string
	Timestamp   time.Time
	Sender      MessageUser
	Recipient   MessageUser
}

type MessageUser struct {
	ID       int
	Username string
}

type MessagesSent struct {
	ID          int
	RecipientID int
}

type MessagesRecieved struct {
	ID       int
	SenderID int
}

// func (m Message) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(MessageResponse(m))
// }

func (m *Message) UnmarshalJSON(data []byte) error {
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
