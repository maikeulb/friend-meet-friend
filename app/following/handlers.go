package followings

func Follow(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Profile ID")
		return
	}

	if err := u.addFollowing(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, u)
}

func UnFollow(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Profile ID")
		return
	}

	if err := u.removeFollowing(db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, u)
}
