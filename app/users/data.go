package users

import (
	"database/sql"
	// "fmt"
	_ "github.com/lib/pq"
)

func GetUserProfiles(db *sql.DB, u []*User) ([]*User, error) {
	query := `
    SELECT u.id,
    u.username,
    u.email,
    u.interests,
    u.neighborhood,
    u.created_on,
    u.last_active,
    f.follower_id,
    u2.username,
    f.followed_id
    u3.username
    FROM users as u
    INNER JOIN followings as f
    ON u.id = f.follower_id
    OR u.id = f.followed_id
    INNER JOIN users as u2
    ON u2.id = f.follower_id
    INNER JOIN users as u3
    ON u3.id = f.followed_id
    ORDER BY m.timestamp;`

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		var u = &User{}
		var u2 = &Followers{}
		var u3 = &Followees{}
		if err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.Email,
			&u.Interests,
			&u.Neighborhood,
			&u.CreatedOn,
			&u.LastActive,
			&u2.ID,
			&u2.Username,
			&u3.ID,
			&u3.Username); err != nil {
			return nil, err
		}
		u.Followers = append(u.Followers, *u2)
		u.Followees = append(u.Followees, *u3)
		users = append(users, u)
	}
	return users, nil
}

// func (model *User) getProfile(db *sql.DB) (User, error) {

// 	query := `
//         SELECT u.id,
//         u.username,
//         u.last_active,
//         u.bio,
//         u.created_on,
//         f.follower_id,
//         f.followed_id,
//         u2.username,
//         u3.username
//         FROM users u
//         INNER JOIN followings as f
//         ON u.id = f.follower_id
//         OR u.id = f.followed_id
//         INNER JOIN users as u2
//         ON u2.id = f.follower_id
//         INNER JOIN users as u3
//         ON u3.id = f.followed_id
//         WHERE id=$1`

// 	rows, err := db.QueryRow(query, m.ID).Scan(
// 		&m.ID,
// 		&m.Username,
// 		&m.Bio,
// 		&m.CreatedOn,
// 		&m.LastActive)

// 	if err == sql.ErrNoRows {
// 		log.Printf("No users")
// 	}

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return rows
// }

// func (model *User) editProfile(db *sql.DB) (User, error) {

// 	query := `

//             SELECT u.id,
//             u.username,
//             u.last_active,
//             u.bio,
//             u.created_on,
//             f.follower_id,
//             f.followed_id,
//             u2.username,
//             u3.username
//             FROM users u
//             INNER JOIN followings as f
//             ON u.id = f.follower_id
//             OR u.id = f.followed_id
//             INNER JOIN users as u2
//             ON u2.id = f.follower_id
//             INNER JOIN users as u3
//             ON u3.id = f.followed_id
//             WHERE id=$1`

// 	return db.QueryRow(query, m.ID).Scan(
// 		&m.SenderID,
// 		&m.RecipientID,
// 		&m.Body,
// 		&m.IsRead)

// 	if err == sql.ErrNoRows {
// 		log.Printf("No users")
// 	}

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
