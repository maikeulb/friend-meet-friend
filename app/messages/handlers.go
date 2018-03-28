package messages

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/maikeulb/friend-meet-friend/app/models"
)

//set a flag
func GetSentMessages(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid, err := strconv.Atoi(vars["userid"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid message ID")
		return
	}

	var m []models.Message
	m.ID = 1
	messages, err := GetSentMessages(db, m)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, messages)
}

// func GetRecievedMessage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
//     vars := mux.Vars(r)
//     id, err := strconv.Atoi(vars["id"])
//     if err != nil {
//         respondWithError(w, http.StatusBadRequest, "Invalid message ID")
//         return
//     }

//     var m models.Message
//     m = models.Message{ID: id}
//     if err := GetRecievedMessages(db, m); err != nil {
//         switch err {
//         case sql.ErrNoRows:
//             respondWithError(w, http.StatusNotFound, "Message not found")
//         default:
//             respondWithError(w, http.StatusInternalServerError, err.Error())
//         }
//         return
//     }

//     respondWithJSON(w, http.StatusOK, m)
// }

// func messageHandler(w http.ResponseWriter, r *http.Request) {

//     if err := json.NewDecoder(r.Body).Decode(s); err != nil { // decode body to message object
//         respondWithError(w, http.StatusInternalServerError, err.Error())
//     } else {
//         respondWithJSON(w, http.StatusOK, s)
//     }
// }

// func SendMessage(db *sql.DB, w http.ResponseWriter, r *http.Request) {
//     vars := mux.Vars(r)
//     m := &models.Message{}
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

//         m := models.Message{ID: id}
//         if err := DeleteMessage(db, m); err != nil {
//             respondWithError(w, http.StatusInternalServerError, err.Error())
//             return
//         }

//         respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
//     }

//     func respondWithError(w http.ResponseWriter, code int, message string) {
//         respondWithJSON(w, code, map[string]string{"error": message})
//     }

//     func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
//         response, _ := json.Marshal(payload)

//         w.Header().Set("Content-Type", "application/json")
//         w.WriteHeader(code)
//         w.Write(response)
//     }
