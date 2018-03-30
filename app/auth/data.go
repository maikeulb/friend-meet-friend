package users

import (
	"database/sql"
	// "fmt"
	_ "github.com/lib/pq"
	"github.com/maikeulb/friend-meet-friend/app/users"
	"log"
)

func GetUser(db *sql.DB, u users.User) (users.User, error) {

	query := `
        SELECT u.id,
        u.username,
        u.email,
        u.password_hash
        FROM users as u
        WHERE u.username=$1;`

	err := db.QueryRow(query, u.username).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.PasswordHash)

	if err == sql.ErrNoRows {
		log.Printf("No user")
	}
	if err != nil {
		log.Fatal(err)
	}

	return u, nil
}

func SaveUser(db *sql.DB, u users.User) error {

	command := `
            INSERT INTO users(username, email, password_hash)
            VALUES($1, $2, $3);`

	_, err := db.Exec(command, u.Username, u.Email, u.PasswordHash)

	return err
}
