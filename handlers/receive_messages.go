package handlers

import (
	"fmt"
	"net/http"
	"pubsub/pubsub"
	"strings"

	"github.com/gorilla/websocket"
)

func (h *Handlers) ReceiveMessages(w http.ResponseWriter, r *http.Request) {
	resourceID := strings.TrimPrefix(r.URL.Path, "/message/receive/")

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
