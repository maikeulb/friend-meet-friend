package messages

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func GetSentMessages(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// userID := 2 // get from context
	// if err != nil {
	//  respondWithError(w, http.StatusBadRequest, "Invalid message ID")
	//  return
	// } compare with userID

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
	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// userID := 2 // get from context
	// if err != nil {
	//  respondWithError(w, http.StatusBadRequest, "Invalid message ID")
	//  return
	// } compare with userID

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
	// userID := 2 // get from context

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
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

	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
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
