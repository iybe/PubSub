package server

import (
	"fmt"
	"net/http"
	"pubsub/handlers"
)

func Run() {
	handlers := handlers.NewHandlers()

	http.HandleFunc("/message/send/", handlers.SendMessage)
	http.HandleFunc("/message/receive/", handlers.ReceiveMessages)

	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
