package users

import (
	"encoding/json"
	// "errors"
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash []byte
	Interests    string
	Borough      string
	CreatedOn    *time.Time
	LastActive   *time.Time
	Followers    []Followers
	Followees    []Followees
}

type Followers struct {
	ID       int
	Username string
}

type Followees struct {
	ID       int
	Username string
}

func (u User) MarshalJSON() ([]byte, error) {
	fmt.Println("marshalling")
	return json.Marshal(Response(u))
}

func (u *User) UnmarshalJSON(data []byte) error {
	fmt.Println("unmarshalling")
	var ju UserRequest

	if err := json.Unmarshal(data, &ju); err != nil {
		return err
	}
	if err := ju.validate(); err != nil {
		return err
	}

	*u = ju.User()
	return nil
}
