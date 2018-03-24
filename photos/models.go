package photos

import (
	"time"
)

type PhotoModel struct {
	ID          int    `json:"id,omitempty"`
	UserID      string `json:"userId,omitempty"`
	Url         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	IsMain      int    `json:"isMain,omitempty"`
	User        UserModel
}

type UserModel struct {
	UserModel users.UserModel
}
