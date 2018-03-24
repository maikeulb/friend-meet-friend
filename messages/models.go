package messages

type MessageModel struct {
	ID          int
	SenderID    int
	RecipientID int
	Body        string
	IsRead      bool // can be improved
	Sender      MessageUserModel
	Recipient   MessageUserModel
}

type MessageUserModel struct {
	UserModel users.UserModel
}
