package server

import (
	"fmt"
	"net/http"
	"pubsub/handlers"

	"github.com/gorilla/mux"
)

func Run() {
	handlers := handlers.NewHandlers()

	r := mux.NewRouter()
	s := r.PathPrefix("/message").Subrouter()
	s.HandleFunc("/send/{resourceId}", handlers.SendMessage).Methods(http.MethodPost)
	s.HandleFunc("/receive/{resourceId}", handlers.ReceiveMessages).Methods(http.MethodGet)
	http.Handle("/", s)

	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
