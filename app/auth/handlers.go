package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// const SIGN_KEY = []byte("secret")

func Login(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var u []*users.User
	email := strings.ToLower(r.FormValue("email"))
	password, err := redis.Bytes(conn.Do("GET", email))
	if err == nil {
		err = bcrypt.CompareHashAndPassword(password, []byte(r.FormValue("password")))
		if err != nil {
			http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", "Wrong password"), 401)
			return
		}
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = email
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		tokenString, _ := token.SignedString(SIGN_KEY)
		w.Write([]byte(fmt.Sprintf("{ \"access_token\": \"%s\" }", tokenString)))
	} else {
		http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", "Email not found"), 401)
		return
	}
}

func Register(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var err error
	email := strings.ToLower(r.FormValue("email"))
	exists, err := redis.Bool(conn.Do("EXISTS", email))
	if exists {
		w.Write([]byte(fmt.Sprintf("{ \"error\": \"%s\" }", "Email taken")))
		return
	}
	password := r.FormValue("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password),
		bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	_, err = conn.Do("SET", email, string(hashedPassword[:]))
	if err != nil {
		log.Println(err)
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, _ := token.SignedString(SIGN_KEY)
	w.Write([]byte(fmt.Sprintf("{ \"access_token\": \"%s\" }", tokenString)))
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
