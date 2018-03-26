package users

type UserModel struct {
    ID                int
    Username          string
    Hash              byte[]
    Salt              byte[]
    Gender            string
    DateOfBirth       time.time
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
    followerID  UserModel
    FolloweeID  UserModel
    Follower    UserModel
    Followee    UserModel
}
