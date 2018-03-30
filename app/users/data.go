package users

import (
	"database/sql"
	// "fmt"
	_ "github.com/lib/pq"
	// "log"
)

func GetUserProfiles(db *sql.DB, u []*User) ([]*User, error) {

	query := `
    SELECT u.id,
    u.username,
    u.email,
    u.interests,
    u.borough,
    u.created_on,
    u.last_active,
    f.follower_id,
    u2.username,
    f.followee_id,
    u3.username
    FROM users as u
    INNER JOIN followings as f
    ON u.id = f.follower_id
    OR u.id = f.followee_id
    INNER JOIN users as u2
    ON u2.id = f.follower_id
    INNER JOIN users as u3
    ON u3.id = f.followee_id
    ORDER BY u.created_on;`

	rows, err := db.Query(query)

	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*User{}
	ids := []int{}

	for rows.Next() {
		var u = &User{}
		var u2 = &Followers{}
		var u3 = &Followees{}
		if err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.Email,
			&u.Interests,
			&u.Borough,
			&u.CreatedOn,
			&u.LastActive,
			&u2.ID,
			&u2.Username,
			&u3.ID,
			&u3.Username); err != nil {
			return nil, err
		}

		if Contains(ids, u.ID) {
			if IsUnique(u.ID, u2.ID) {
				users[len(users)-1].Followers = append(users[len(users)-1].Followers, *u2)
			}
			if IsUnique(u.ID, u3.ID) {
				users[len(users)-1].Followees = append(users[len(users)-1].Followees, *u3)
			}
		} else {
			if IsUnique(u.ID, u2.ID) {
				u.Followers = append(u.Followers, *u2)
			}
			if IsUnique(u.ID, u3.ID) {
				u.Followees = append(u.Followees, *u3)
			}
			users = append(users, u)
		}
		ids = append(ids, u.ID)
	}
	return users, nil
}

func GetUserProfile(db *sql.DB, u User) (User, error) {

	query := `
        SELECT u.id,
        u.username,
        u.email,
        u.interests,
        u.borough,
        u.created_on,
        u.last_active,
        f.follower_id,
        u2.username,
        f.followee_id,
        u3.username
        FROM users as u
        INNER JOIN followings as f
        ON u.id = f.follower_id
        OR u.id = f.followee_id
        INNER JOIN users as u2
        ON u2.id = f.follower_id
        INNER JOIN users as u3
        ON u3.id = f.followee_id
        WHERE u.id=$1`

	var u2 = &Followers{}
	var u3 = &Followees{}
	err := db.QueryRow(query, u.ID).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.Interests,
		&u.Borough,
		&u.CreatedOn,
		&u.LastActive,
		&u2.ID,
		&u2.Username,
		&u3.ID,
		&u3.Username)

	u.Followers = append(u.Followers, *u2)
	u.Followees = append(u.Followees, *u3)

	if err == sql.ErrNoRows {
		return u, err //check this
	}
	if err != nil {
		return u, err
	}

	return u, nil
}

func UpdateUserProfile(db *sql.DB, u *User) error {
	command := `
        UPDATE users as u
        SET email=case
                when $1='' then u.email
                else $1
            end,
            interests=case
               when $2='' then u.interests
               else $2
            end,
            borough=case
               when $3='' then u.borough
               else $3
            end
        WHERE id = $4;`

	_, err := db.Exec(command, u.Email, u.Interests, u.Borough, u.ID)
	if err != nil {
		return err
	}

	return nil
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func IsUnique(s int, e int) bool {
	if s == e {
		return false
	}
	return true
}
