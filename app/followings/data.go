package followings

import (
	"database/sql"
	// "fmt"

	_ "github.com/lib/pq"
)

func AddFollowing(db *sql.DB, f *Following) error {
	command := `
        INSERT into followings (follower_id, followee_id)
        VALUES ($1, $2)
        RETURNING follower_id, followee_id`

	err := db.QueryRow(command, f.FollowerID, f.FolloweeID).Scan(
		&f.FollowerID, &f.FolloweeID)
	if err != nil {
		return err
	}

	return nil
}

func RemoveFollowing(db *sql.DB, f *Following) error {
	command := `
        DELETE from followings 
        WHERE follower_id=$1 and followee_id=$2
        RETURNING follower_id, followee_id`

	err := db.QueryRow(command, f.FollowerID, f.FolloweeID).Scan(
		&f.FollowerID, &f.FolloweeID)
	if err != nil {
		return err
	}

	return nil
}
