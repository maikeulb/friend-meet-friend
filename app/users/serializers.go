package users

import (
	// "errors"
	"time"
)

type UserRequest struct {
	Interests string `json:"interests,omitempty"`
	Borough   string `json:"borough,omitempty"`
	Name      string `json:"name,omitempty"`
}

func (ju UserRequest) User(u User) User {
	u.Name = ju.Name
	u.Interests = ju.Interests
	u.Borough = ju.Borough

	return u
}

func (ju *UserRequest) validate() error {

	return nil
}

type UserResponse struct {
	ID         int                 `json:"id,omitempty"`
	Name       string              `json:"name,omitempty"`
	Interests  string              `json:"interests,omitempty"`
	Borough    string              `json:"borough,omitempty"`
	CreatedOn  *time.Time          `json:"createdOn,omitempty"`
	LastActive *time.Time          `json:"lastActive,omitempty"`
	Followers  []FollowersResponse `json:"followers,omitempty"`
	Followees  []FolloweesResponse `json:"followees,omitempty"`
}

func Response(u User) UserResponse {
	var ju UserResponse
	ju.ID = u.ID
	ju.Name = u.Name
	ju.Interests = u.Interests
	ju.Borough = u.Borough
	ju.CreatedOn = u.CreatedOn
	ju.LastActive = u.LastActive

	var juFollowers []FollowersResponse
	for _, follower := range u.Followers {
		juFollowers = append(juFollowers,
			FollowersResponse{
				ID:   follower.ID,
				Name: follower.Name,
			})
	}
	ju.Followers = juFollowers

	var juFollowees []FolloweesResponse
	for _, followee := range u.Followees {
		juFollowees = append(juFollowees,
			FolloweesResponse{
				ID:   followee.ID,
				Name: followee.Name,
			})
	}
	ju.Followees = juFollowees

	return ju
}

type FollowersResponse struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type FolloweesResponse struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
