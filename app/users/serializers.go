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

	var juFollowers []FollowersResponse
	for _, follower := range u.Followers {
		juFollowers = append(juFollowers,
			FollowersResponse{
				ID:       follower.ID,
				Username: follower.Username,
			})
	}
	ju.Followers = juFollowers

	var juFollowees []FolloweesResponse
	for _, followee := range u.Followees {
		juFollowees = append(juFollowees,
			FolloweesResponse{
				ID:       followee.ID,
				Username: followee.Username,
			})
	}
	ju.Followees = juFollowees

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
