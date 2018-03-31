package auth

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	"net/http"
	"strings"
	// "strconv"

	// "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func LoginUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	email := "michael@email.com"
	password := "P@ssw0rd!"

	u := &User{Email: strings.ToLower(email)} // move ToLower in request marshalling
	err := GetUser(db, u)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, err.Error())
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := u.CheckPassword(password); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	tokenString, err := GenerateToken(*u)
	if err != nil {
		return err
	}

	respondWithJSON(w, http.StatusOK, tokenString)
}

func RegisterUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	email := "fico@email.com"
	username := "fico"
	password := "P@ssw0rd!"

	u := User{Email: strings.ToLower(email), Username: strings.ToLower(username)} // move tolower to request logic

	//make more slick || better status codes
	var err error
	exists, err := IsEmailExists(db, u)
	if exists {
		respondWithError(w, http.StatusInternalServerError, "Email already exists")
	}
	exists, err = IsUsernameExists(db, u)
	if exists {
		respondWithError(w, http.StatusInternalServerError, "Username already exists")
	}

	//combine the two, auto save password with saveuser
	u.SetPassword(password) //figure out how to error handle this
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	err = SaveUser(db, &u)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	tokenString, err := GenerateToken(*u)
	if err != nil {
		return err
	}

	respondWithJSON(w, http.StatusOK, tokenString) // add token here to user object
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
