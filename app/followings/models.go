package followings

import (
	"encoding/json"
	"fmt"
)

type Following struct {
	FollowerID int
	FolloweeID int
}

func (f Following) MarshalJSON() ([]byte, error) {
	fmt.Println("marshalling")
	return json.Marshal(Response(f))
}

func (f *Following) UnmarshalJSON(data []byte) error {
	fmt.Println("unmarshalling")
	var jf FollowingsRequest

	if err := json.Unmarshal(data, &jf); err != nil {
		return err
	}
	if err := jf.validate(); err != nil {
		return err
	}

	jf.Following(f)
	return nil
}
