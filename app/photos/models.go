package photos

import (
	"time"
)

type PhotoModel struct {
	ID          int
	UserID      string
	Url         string
	Description string
	IsMain      int
	User        UserModel
}

type UserModel struct {
	UserModel users.UserModel
}
