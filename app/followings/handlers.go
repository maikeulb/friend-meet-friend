package followings

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Follow(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	followerID := r.Context().Value("userId").(int)
	followeeID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Profile ID")
		return
	}

	if followerID == followeeID {
		respondWithError(w, http.StatusForbidden, "You cannot follow yourself")
		return
	}

	f := &Following{FollowerID: followerID, FolloweeID: followeeID}
	if err := AddFollowing(db, f); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, f)
}

func UnFollow(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	followerID := r.Context().Value("userId").(int)
	followeeID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Profile ID")
		return
	}

	if followerID == followeeID {
		respondWithError(w, http.StatusForbidden, "You cannot unfollow yourself")
		return
	}

	f := &Following{FollowerID: followerID, FolloweeID: followeeID}
	if err := RemoveFollowing(db, f); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, f)
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
