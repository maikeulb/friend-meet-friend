package users

import (
	// "errors"
	"time"
)

type UserRequest struct {
	Interests string `json:"interests,omitempty"`
	Borough   string `json:"borough,omitempty"`
}

func (ju UserRequest) User() User {
	var u User
	u.Interests = ju.Interests
	u.Borough = ju.Borough

	return u
}

func (ju *UserRequest) validate() error {
	// if ju.Interests <= "" {
	// return errors.New("Username should not be empty")
	// }
	// if ju.Boroughs <= "" {
	// return errors.New("Body should not be empty")
	// }

	return nil
}

type UserResponse struct {
	ID         int                 `json:"id,omitempty"`
	Username   string              `json:"username,omitempty"`
	Email      string              `json:"email,omitempty"`
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
	ju.Username = u.Username
	ju.Email = u.Email
	ju.Interests = u.Interests
	ju.Borough = u.Borough
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
