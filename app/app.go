package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	// "time"

	"github.com/gorilla/mux"
	"github.com/maikeulb/friend-meet-friend/app/auth"
	"github.com/maikeulb/friend-meet-friend/app/followings"
	"github.com/maikeulb/friend-meet-friend/app/messages"
	"github.com/maikeulb/friend-meet-friend/app/users"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(host, port, user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

func (a *App) Run(addr string) {
	fmt.Println("Listening on port: 5000")
	fmt.Println("/api/login")
	fmt.Println("/api/register")
	fmt.Println("/api/status")
	fmt.Println("/api/users")
	fmt.Println("/api/users/{userId}")
	fmt.Println("/api/users/{userId}/messages")
	fmt.Println("/api/users/{userId}/messages/{id}")
	fmt.Println("/api/users/{userId}/messages/sent")
	fmt.Println("/api/users/{userId}/messages/recieved")
	fmt.Println("/api/users/{userId}/follow")
	fmt.Println("/api/users/{userId}/unfollow")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/api/login", a.LoginUser).Methods("POST")
	a.Router.HandleFunc("/api/register", a.RegisterUser).Methods("POST")
	a.Router.HandleFunc("/api/status", a.Status)
	a.Router.HandleFunc("/api/users", a.GetUsers).Methods("GET")
	a.Router.HandleFunc("/api/users/{userId:[0-9]+}", a.GetUser).Methods("GET")
	a.Router.HandleFunc("/api/users/{userId:[0-9]+}", a.UpdateUser).Methods("PUT")
	a.Router.HandleFunc("/api/users/{userId:[0-9]+}/messages", a.SendUserMessage).Methods("POST")
	a.Router.HandleFunc("/api/users/{userId:[0-9]+}/messages/{id:[0-9]+}", a.GetUserMessage).Methods("GET")
	a.Router.HandleFunc("/api/users/{userId:[0-9]+}/messages/sent", a.GetUserSentMessages).Methods("GET")
	a.Router.HandleFunc("/api/users/{userId:[0-9]+}/messages/recieved", a.GetUserRecievedMessages).Methods("GET")
	a.Router.HandleFunc("/api/users/{userId:[0-9]+}/follow", a.FollowUser).Methods("POST")
	a.Router.HandleFunc("/api/users/{userId:[0-9]+}/unfollow", a.UnFollowUser).Methods("POST")
	a.Router.Use(a.AddContextMiddleware)
}

func (a *App) AddContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)
		cookie, _ := r.Cookie("username")
		if cookie != nil {
			ctx := context.WithValue(r.Context(), "Username", cookie.Value)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func (a *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	users.GetUsers(a.DB, w, r)
}

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	users.GetUser(a.DB, w, r)
}

func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	users.UpdateUser(a.DB, w, r)
}

func (a *App) GetUserMessage(w http.ResponseWriter, r *http.Request) {
	messages.GetMessage(a.DB, w, r)
}

func (a *App) SendUserMessage(w http.ResponseWriter, r *http.Request) {
	messages.SendMessage(a.DB, w, r)
}

func (a *App) GetUserSentMessages(w http.ResponseWriter, r *http.Request) {
	messages.GetSentMessages(a.DB, w, r)
}

func (a *App) GetUserRecievedMessages(w http.ResponseWriter, r *http.Request) {
	messages.GetRecievedMessages(a.DB, w, r)
}

func (a *App) FollowUser(w http.ResponseWriter, r *http.Request) {
	followings.Follow(a.DB, w, r)
}

func (a *App) UnFollowUser(w http.ResponseWriter, r *http.Request) {
	followings.UnFollow(a.DB, w, r)
}

func (a *App) LoginUser(w http.ResponseWriter, r *http.Request) {
	auth.LoginUser(a.DB, w, r)
}

func (a *App) RegisterUser(w http.ResponseWriter, r *http.Request) {
	auth.RegisterUser(a.DB, w, r)
}

// func (a *App) Login(w http.ResponseWriter, r *http.Request) {
// 	expiration := time.Now().Add(365 * 24 * time.Hour)
// 	cookie := http.Cookie{Name: "username", Value: "demo@gmail.com", Expires: expiration}
// 	http.SetCookie(w, &cookie)
// }

func (a *App) Status(w http.ResponseWriter, r *http.Request) {
	if username := r.Context().Value("Username"); username != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello " + username.(string) + "\n"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Logged in"))
	}
}
