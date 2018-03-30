package followings

import (
	"errors"
)

type FollowingsRequest struct {
	FolloweeID int `json:"followeeID,omitempty"`
}

func (jf FollowingsRequest) Following(f *Following) {
	f.FolloweeID = jf.FolloweeID

	return
}

func (jf *FollowingsRequest) validate() error {
	if jf.FolloweeID <= 0 {
		return errors.New("FolloweeID must not be empty")
	}

	return nil
}

type FollowingsResponse struct {
	FollowerID int
	FolloweeID int
}

func Response(m Following) FollowingsResponse {
	var jf FollowingsResponse
	jf.FollowerID = m.FollowerID
	jf.FolloweeID = m.FolloweeID

	return jf
}
