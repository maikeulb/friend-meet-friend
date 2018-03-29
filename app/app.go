package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	fmt.Println("Listening on port: 5000")
	fmt.Println("/api/login")
	fmt.Println("/api/status")
	fmt.Println("/api/messages/{id}")
	fmt.Println("/api/messages/sent")
	fmt.Println("/api/messages/received")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/login", a.Login)
	a.Router.HandleFunc("/api/status", a.Status)
	a.Router.HandleFunc("/api/messages/{id:[0-9]+}", a.GetMessage).Methods("GET")
	a.Router.HandleFunc("/api/messages/sent", a.GetSentMessages).Methods("GET")
	a.Router.HandleFunc("/api/messages/recieved", a.GetRecievedMessages).Methods("GET")
	a.Router.HandleFunc("/api/profiles", a.GetProfiles).Methods("GET")
	a.Router.Use(a.AddContextMiddleware)
}

func (a *App) GetSentMessages(w http.ResponseWriter, r *http.Request) {
	messages.GetSentMessages(a.DB, w, r) // consider squashing sent and recieved to one urla nd adding a filter
}

func (a *App) GetRecievedMessages(w http.ResponseWriter, r *http.Request) {
	messages.GetRecievedMessages(a.DB, w, r)
}

func (a *App) GetMessage(w http.ResponseWriter, r *http.Request) {
	messages.GetMessage(a.DB, w, r)
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

func (a *App) GetProfiles(w http.ResponseWriter, r *http.Request) {
	users.GetProfiles(a.DB, w, r)
}

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "username", Value: "demo@gmail.com", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func (a *App) Status(w http.ResponseWriter, r *http.Request) {
	if username := r.Context().Value("Username"); username != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello " + username.(string) + "\n"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Logged in"))
	}
}
