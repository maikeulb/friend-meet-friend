package users

import (
    "time"
)

type UserModel struct {
    ID                int
    Username          string
    PasswordHash      byte[]
    PasswordSalt      byte[]
    CreatedOn         time.time
    LastActive        time.time
    Bio               string
    Interests         string
    Neighborhood      string
    MessagesSent      []MessagesSentModel
    MessagesRecieved  []MessagesRecievedModel
    Followers         []FollowingsModel
    Followees         []FollowingsModel
}

type MessagesSentModel struct {
    MessagesSent messages.MessagesSentModel
}

type MessagesRecievedModel struct {
    MessagesRecievedModel messages.MessagesSentModel
}

type FollowersModel struct {
    FollowerID  UserModel
    FolloweeID  UserModel
    Follower    UserModel
    Followee    UserModel
}
