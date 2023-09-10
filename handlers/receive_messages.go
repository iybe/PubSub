package handlers

import (
	"fmt"
	"net/http"
	"pubsub/pubsub"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func (h *Handlers) ReceiveMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resourceID, ok := vars["resourceId"]
	if !ok {
		http.Error(w, "resourceId empty", http.StatusBadRequest)
		return
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}

	receiveMessagesService(conn, resourceID, h.Manager)
}

func receiveMessagesService(conn *websocket.Conn, resourceID string, manager *pubsub.Manager) {
	defer conn.Close()

	sub := pubsub.NewSub()
	manager.Register(resourceID, sub)
	defer manager.DeregisterSafe(resourceID, sub)

	go func() {
		for msg := range sub.Chan {
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg.(string)))
			if err != nil {
				fmt.Println("Error writing message:", err)
				return
			}
		}
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}
	}
}
