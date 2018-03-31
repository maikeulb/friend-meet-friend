package auth

import (
	"database/sql"
	// "fmt"
	_ "github.com/lib/pq"
)

func GetUser(db *sql.DB, u *User) error {
	query := `
        SELECT u.id,
        u.email,
        u.password_hash
        FROM users as u
        WHERE u.email=$1;`

	return db.QueryRow(query, u.Email).Scan(
		&u.ID,
		&u.Email,
		&u.PasswordHash)
}

func IsEmailExists(db *sql.DB, u User) (bool, error) {
	query := `
        SELECT u.id,
        u.email
        FROM users as u
        WHERE u.email=$1;`
	err := db.QueryRow(query, u.Email).Scan(
		&u.ID,
		&u.Email)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func SaveUser(db *sql.DB, u *User) error {
	command := `
        INSERT INTO users(email, password_hash)
        VALUES($1, $2)
        RETURNING id;`

	err := db.QueryRow(command, u.Email, u.PasswordHash).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}
