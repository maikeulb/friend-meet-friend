package users

import (
	"errors"
	"time"
)

type UserRequest struct {
	ID           int       `json:"id,omitempty"`
	Username     string    `json:"username,omitempty"`
	Email        string    `json:"email,omitempty"`
	Interests    string    `json:"interests,omitempty"`
	Neighborhood string    `json:"neighborhood,omitempty"`
	CreatedOn    time.Time `json:"createdOn,omitempty"`
	LastActive   time.Time `json:"lastActive,omitempty"`
}

func (ju UserRequest) User() User {
	var u User
	u.ID = ju.ID
	u.Username = ju.Username
	u.Email = ju.Email
	u.Interests = ju.Interests
	u.Neighborhood = ju.Neighborhood
	u.CreatedOn = ju.CreatedOn
	u.LastActive = ju.LastActive

	return u
}

func (ju *UserRequest) validate() error {
	if ju.ID <= 0 {
		return errors.New("ID should not be empty")
	}
	if ju.Username <= "" {
		return errors.New("Username should not be empty")
	}
	if ju.Interests <= "" {
		return errors.New("Body should not be empty")
	}

	return nil
}

type UserResponse struct {
	ID           int                 `json:"id,omitempty"`
	Username     string              `json:"username,omitempty"`
	Email        string              `json:"email,omitempty"`
	Interests    string              `json:"interests,omitempty"`
	Neighborhood string              `json:"neighborhood,omitempty"`
	CreatedOn    time.Time           `json:"createdOn,omitempty"`
	LastActive   time.Time           `json:"lastActive,omitempty"`
	Followers    []FollowersResponse `json:"followers,omitempty"`
	Followees    []FolloweesResponse `json:"followees,omitempty"`
}

func Response(u User) UserResponse {
	var ju UserResponse
	ju.ID = u.ID
	ju.Username = u.Username
	ju.Email = u.Email
	ju.Interests = u.Interests
	ju.Neighborhood = u.Neighborhood
	ju.CreatedOn = u.CreatedOn
	ju.LastActive = u.LastActive
	// ju.Followers = u.Followers.ID
	// ju.Followers.Username = u.Followers.Username
	// ju.Followees.ID = u.Followees.ID
	// ju.Followees.Username = u.Followees.Username
	// for x := range p.Fruits {
	// 	if _, err = stmt4.Exec(string(i), x.Type, x.Number); err != nil {
	// 		log.Println("stmt1.Exec: ", err.Error())
	// 		return
	// 	}
	// }

	// for _, follower := range ju.Followers {
	// for _, follower := range ju.Followers {
	// 	// followers = append(followers, follower)
	// 	u.Followers = append(u.Followers, follower)
	// }
	// // u.Followees = append(u.Followees, *u3)
	// // users = append(users, u)
	return ju
}

type FollowersResponse struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

type FolloweesResponse struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}
