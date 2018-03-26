package users

import (
    "time"
)

type User struct {
    ID                int
    Username          string
    PasswordHash      byte[]
    PasswordSalt      byte[]
    CreatedOn         time.time
    LastActive        time.time
    Bio               string
    Interests         string
    Neighborhood      string
    MessagesSent      []MessagesSent
    MessagesRecieved  []MessagesRecieved
    Followers         []Followings
    Followees         []Followings
}

type MessagesSent struct {
    MessagesSent messages.MessagesSent
}

type MessagesRecieved struct {
    MessagesRecieved messages.MessagesSent
}

