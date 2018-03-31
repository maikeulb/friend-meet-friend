package messages

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func GetSentMessages(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	currentUserID := r.Context().Value("userId")
	fmt.Println(currentUserID)
	userID, err := strconv.Atoi(vars["userId"])
	fmt.Println(userID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	fmt.Println(currentUserID)
	fmt.Println(userID)
	// if currentUserID.(int) != userID.(int) {
	// respondWithError(w, http.StatusForbidden, "Forbidden")
	// return
	// }

	var m []Message
	messages, err := GetSentMessagesForUser(db, m, userID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "No messages found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, messages)
}

func GetRecievedMessages(db *sql.DB, w http.ResponseWriter, r *http.Request) {
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

	var m []Message
	messages, err := GetRecievedMessagesForUser(db, m, userID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "No messages found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, messages)
}

func GetMessage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Message ID")
		return
	}

	m := &Message{ID: id}
	if err := GetMessageForUser(db, m, userID); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, m)
	return
}

func SendMessage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
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

	m := &Message{SenderID: userID}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(m); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := SendMessageToUser(db, m); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, m)
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
