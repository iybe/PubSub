package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) Routes() *mux.Router {
	r := mux.NewRouter()
	messageRoutes := r.PathPrefix("/message").Subrouter()
	messageRoutes.HandleFunc("/send/{resourceId}", s.Handlers.SendMessage).Methods(http.MethodPost)
	messageRoutes.HandleFunc("/receive/{resourceId}", s.Handlers.ReceiveMessages).Methods(http.MethodGet)
	return r
}
