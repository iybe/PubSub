package server

import (
	"fmt"
	"net/http"
	"pubsub/handlers"
)

type Server struct {
	Handlers *handlers.Handlers
}

func Run() {
	serve := Server{
		Handlers: handlers.NewHandlers(),
	}

	http.Handle("/", serve.Routes())

	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
