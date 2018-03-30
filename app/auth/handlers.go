package auth

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	"net/http"
	// "strconv"

	// "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// const SIGN_KEY = []byte("secret")

func LoginUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	email := "michael@email.com"
	password := "P@ssw0rd!"
	u := &User{Email: email}
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
	if err == nil {
		err = u.CheckPassword(password)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		// token := jwt.New(jwt.SigningMethodHS256)
		// claims := token.Claims.(jwt.MapClaims)
		// claims["email"] = email
		// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		// tokenString, _ := token.SignedString(SIGN_KEY)
		// w.Write([]byte(fmt.Sprintf("{ \"access_token\": \"%s\" }", tokenString)))
	}
	respondWithJSON(w, http.StatusOK, u)
}

func RegisterUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var err error
	email := "fico@email.com"
	username := "fico"
	password := "P@ssw0rd!"
	u := User{Email: email, Username: username}
	exists, err := IsEmailExists(db, u)
	if exists {
		respondWithError(w, http.StatusInternalServerError, "Email already exists")
	}
	exists, err = IsUsernameExists(db, u)
	if exists {
		respondWithError(w, http.StatusInternalServerError, "Username already exists")
	}
	u.SetPassword(password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	err = SaveUser(db, &u)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["email"] = email
	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// tokenString, _ := token.SignedString(SIGN_KEY)
	// w.Write([]byte(fmt.Sprintf("{ \"access_token\": \"%s\" }", tokenString)))
	respondWithJSON(w, http.StatusOK, u)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
