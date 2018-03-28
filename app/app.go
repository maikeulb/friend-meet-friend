package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maikeulb/friend-meet-friend/app/messages"
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
	fmt.Println("/api/messages/{id}")
	fmt.Println("/api/messages/sent")
	fmt.Println("/api/messages/received")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/messages/{id:[0-9]+}", a.GetMessage).Methods("GET")
	a.Router.HandleFunc("/api/messages/sent", a.GetSentMessages).Methods("GET")
	a.Router.HandleFunc("/api/messages/recieved", a.GetRecievedMessages).Methods("GET")
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
