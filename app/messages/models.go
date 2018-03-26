package messages

type Message struct {
	ID          int
	SenderID    int
	RecipientID int
	Body        string
	timestamp   time.time
	Sender      MessageUser
	Recipient   MessageUser
}

type MessageUser struct {
	User users.User
}
