package auth

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	"net/http"
	// "strings"
	// "strconv"

	// "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func LoginUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	email := "mickjag@email.com"
	password := "P@ssw3rd!"

	u := &User{Email: email}
	err := GetUser(db, u)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusForbidden, err.Error())
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := u.CheckPassword(password); err != nil {
		respondWithError(w, http.StatusForbidden, err.Error())
		return
	}
	if err := GenerateToken(u); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, u)
}

func RegisterUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	email := "mickjag@email.com"
	username := "mickjag"
	password := "P@ssw3rd!"

	u := User{Email: email, Username: username}

	var err error
	exists, err := IsEmailExists(db, u)
	if exists {
		respondWithError(w, http.StatusUnprocessableEntity, "Email already exists")
		return
	}
	if err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	exists, err = IsUsernameExists(db, u)
	if exists {
		respondWithError(w, http.StatusUnprocessableEntity, "Username already exists")
		return
	}
	if err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if err := u.SetPassword(password); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := SaveUser(db, &u); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := GenerateToken(&u); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, u) // add token here to user object
}

// lgout todo
// refresh todo
// dummy resource
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
