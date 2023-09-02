package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pubsub/pubsub"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Msg string `json:"msg"`
}

func Run() {
	manager := pubsub.NewManager()

	http.HandleFunc("/send/", func(w http.ResponseWriter, r *http.Request) {
		resourceID := strings.TrimPrefix(r.URL.Path, "/send/")

		var receivedMsg Message
		err := json.NewDecoder(r.Body).Decode(&receivedMsg)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		sendMessage(resourceID, manager, receivedMsg.Msg)
		w.WriteHeader(http.StatusCreated)
	})

	http.HandleFunc("/observer/", func(w http.ResponseWriter, r *http.Request) {
		resourceID := strings.TrimPrefix(r.URL.Path, "/observer/")
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Error upgrading connection:", err)
			return
		}

		handleWebSocket(conn, resourceID, manager)
	})

	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func sendMessage(resourceID string, manager *pubsub.Manager, msg string) {
	manager.SendMsg(resourceID, msg)
}

func handleWebSocket(conn *websocket.Conn, resourceID string, manager *pubsub.Manager) {
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
