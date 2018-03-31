package auth

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func LoginUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var u *User
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

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

	if err := u.CheckPassword(); err != nil {
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
	var u User
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if exists, _ := IsEmailExists(db, u); exists {
		respondWithError(w, http.StatusUnprocessableEntity, "Email already exists")
		return
	}

	if err := u.SetPassword(); err != nil {
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

	respondWithJSON(w, http.StatusOK, u)
}

// TODO: logout
// TODO: refresh

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
