package followings

type Following struct {
    FollowerID  int
    FolloweeID  int
    Follower    User
    Followee    User
}
