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
	r.HandleFunc("/message/send/{resourceId}", handlers.SendMessage)
	r.HandleFunc("/message/receive/{resourceId}", handlers.ReceiveMessages)
	http.Handle("/", r)

	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
