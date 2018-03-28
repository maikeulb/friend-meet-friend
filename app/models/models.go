package models

import (
	"time"
)

type Message struct {
	ID          int
	SenderID    int
	RecipientID int
	Body        string
	timestamp   time.Time
	Sender      MessageUser
	Recipient   MessageUser
}

type MessageUser struct {
	User User
}

type User struct {
	ID       int
	Username string
	// PasswordHash      byte[]
	// PasswordSalt      byte[]
	CreatedOn  time.Time
	LastActive time.Time
	Bio        string
	// Interests         string
	// Neighborhood      string
	MessagesSent     []MessagesSent
	MessagesRecieved []MessagesRecieved
	Followers        []UserFollowers
	Followees        []UserFollowees
}

type MessagesSent struct {
	MessagesSent Message
}

type MessagesRecieved struct {
	MessagesRecieved Message
}

type UserFollowers struct {
	Followers Following
}

type UserFollowees struct {
	Followees Following
}

type Following struct {
	FollowerID int
	FolloweeID int
	Follower   User
	Followee   User
}
