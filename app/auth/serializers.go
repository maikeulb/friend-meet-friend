package auth

import (
	"errors"
	"strings"
	// "time"
)

type UserRequest struct {
	Email    string
	Password string
}

func (ju UserRequest) User() User {
	var u User
	u.Email = strings.ToLower(ju.Email)
	u.Password = ju.Password

	return u
}

func (ju *UserRequest) validate() error {
	if ju.Email <= "" {
		return errors.New("Email should not be empty")
	}
	if ju.Password <= "" {
		return errors.New("Password should not be empty")
	}
	// UPDATE Last-active
	return nil
}

type UserResponse struct {
	ID    int    `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Token string `json:"token,omitempty"`
}

func Response(u User) UserResponse {
	var ju UserResponse
	ju.ID = u.ID
	ju.Email = u.Email
	ju.Token = u.Token

	return ju
}
