package models

import (
	"encoding/json"
	"time"
)

type User struct {
	ID               int
	Username         string
	Email            string
	// PasswordHash     []byte
	// PasswordSalt     []byte
	Bio              string
	Neighborhood     string
	CreatedOn        time.Time
	LastActive       time.Time
	MessagesSent     []MessagesSent
	MessagesRecieved []MessagesRecieved
	Followers        []UserFollowers
	Followees        []UserFollowees
}

type MessagesSent struct {
	ID          int
	RecipientID int
}

type MessagesRecieved struct {
	ID       int
	SenderID int
}

type UserFollowers struct {
	FollowerID int
	Username   string
}

type UserFollowees struct {
	FolloweeID int
	Username   string
}

type Following struct {
	FollowerID int
	FolloweeID int
	Follower   User
	Followee   User
}
