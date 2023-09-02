package handlers

import (
	"encoding/json"
	"net/http"
	"pubsub/pubsub"
	"strings"
)

type MessageRequest struct {
	Msg string `json:"message"`
}

func (h *Handlers) SendMessage(w http.ResponseWriter, r *http.Request) {
	resourceID := strings.TrimPrefix(r.URL.Path, "/message/send/")

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
