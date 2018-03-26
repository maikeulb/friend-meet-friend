package messages

type MessageModel struct {
	ID          int
	SenderID    int
	RecipientID int
	Body        string
	IsRead      bool
	Sender      MessageUserModel
	Recipient   MessageUserModel
}

type MessageUserModel struct {
	UserModel users.UserModel
}
