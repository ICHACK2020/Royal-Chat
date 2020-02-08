package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type conversation struct {
	conn     *websocket.Conn
	messages []message
}

func newConversation(w http.ResponseWriter, r *http.Request) *conversation {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer socket.Close()
	msgs := make([]message, 30) //30 is the default length of messages
	return &conversation{conn: socket,
		messages: msgs}
}
