package users

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func GetUsers(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var u []User
	users, err := GetUserProfiles(db, u)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "No messages found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

func GetUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Profile ID")
		return
	}

	u := User{ID: userID}
	user, err := GetUserProfile(db, u)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "No messages found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func UpdateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currentUserID := r.Context().Value("userId")
	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	if currentUserID != userID {
		respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}

	u := &User{ID: userID}
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := UpdateUserProfile(db, u); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, u)
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
