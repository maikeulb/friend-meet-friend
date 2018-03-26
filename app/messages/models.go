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

type User struct {
    ID                int
    Username          string
    PasswordHash      byte[]
    PasswordSalt      byte[]
    CreatedOn         time.time
    LastActive        time.time
    Bio               string
    // Interests         string
    // Neighborhood      string
    MessagesSent      []MessagesSent
    MessagesRecieved  []MessagesRecieved
    Followers         []Followings
    Followees         []Followings
}

