package messages

import (
	"database/sql"
	"encoding/json"
	"net/http"
	// "strconv"

	_ "github.com/lib/pq"
)

//set a flag
func GetSentMessages(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// sid, err := strconv.Atoi(vars["userid"])

	// if err != nil {
	//  respondWithError(w, http.StatusBadRequest, "Invalid message ID")
	//  return
	// }

	var m []*Message
	messages, err := GetSentMessagesForUser(db, m, 2)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, messages)
}

func GetRecievedMessages(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id, err := strconv.Atoi(vars["id"])
	// if err != nil {
	// respondWithError(w, http.StatusBadRequest, "Invalid message ID")
	// return
	// }

	var m []*Message
	messages, err := GetRecievedMessagesForUser(db, m, 2)
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

	// m := *Message{ID: 1}
	var m Message
	if err := GetMessageForUser(db, &m, 1); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, m)
	return
}

// func SendMessage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
//     vars := mux.Vars(r)
//     m := &Message{}
//     defer r.Body.Close()

//     if err := json.NewDecoder(r.Body).Decode(m); err != nil
//     respondWithError(w, http.StatusInternalServerError, err.Error())

//     if err := SendMessage(db, m); err != nil {
//         respondWithError(w, http.StatusInternalServerError, err.Error())
//         return

//         respondWithJSON(w, http.StatusCreated, m)
//     }

//     func DeleteMessage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
//         vars := mux.Vars(r)
//         id, err := strconv.Atoi(vars["id"])
//         if err != nil {
//             respondWithError(w, http.StatusBadRequest, "Invalid Message ID")
//             return
//         }

//         m := Message{ID: id}
//         if err := DeleteMessage(db, m); err != nil {
//             respondWithError(w, http.StatusInternalServerError, err.Error())
//             return
//         }

//         respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
//     }

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
