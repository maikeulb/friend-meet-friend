package auth

import (
	"encoding/json"
	// "errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID           int
	Username     string
    // UPDATE change username to nickname and make it optionsl
	Email        string
	Token        string
	Password     string
	PasswordHash []byte
	CreatedOn    *time.Time
	LastActive   *time.Time
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

func (u *User) CheckPassword() error {
	bytePassword := []byte(u.Password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func (u *User) SetPassword() error {
	bytePassword := []byte(u.Password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = passwordHash
	return nil
}
