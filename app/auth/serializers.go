package auth

import (
// "errors"
// "time"
"strings"
)

type UserRequest struct {
	Username string
	Email    string
	Password string
}

func (ju UserRequest) User() User {
	var u User
	u.Username = strings.ToLower(ju.Username)
	u.Email = strings.ToLower(ju.Email)
	u.SetPassword(ju.Password)

	return u
}

func (ju *UserRequest) validate() error {
	// if ju.Username <= "" {
	// return errors.New("Username should not be empty")
	// }
	// if ju.Email <= "" {
	// return errors.New("Body should not be empty")
	// }
	// if ju.Password <= "" {
	// return errors.New("Body should not be empty")
	// }
	// if len(password) == 0 {
	// return errors.New("password should not be empty!")
	// }

	return nil
}

type UserResponse struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Token    string `json:"token,omitempty"`
}

func Response(u User) UserResponse {
	var ju UserResponse
	ju.ID = u.ID
	ju.Username = u.Username
	ju.Email = u.Email
	ju.Token = u.Token

	return ju
}
