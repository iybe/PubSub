package handlers

import (
	"encoding/json"
	"net/http"
	"pubsub/pubsub"

	"github.com/gorilla/mux"
)

type MessageRequest struct {
	Msg string `json:"message"`
}

func (h *Handlers) SendMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resourceID, ok := vars["resourceId"]
	if !ok {
		http.Error(w, "resourceId empty", http.StatusBadRequest)
		return
	}

	var receivedMsg MessageRequest
	err := json.NewDecoder(r.Body).Decode(&receivedMsg)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	sendMessageService(resourceID, h.Manager, receivedMsg.Msg)
	w.WriteHeader(http.StatusOK)
}

func sendMessageService(resourceID string, manager *pubsub.Manager, msg string) {
	manager.SendMsg(resourceID, msg)
}
