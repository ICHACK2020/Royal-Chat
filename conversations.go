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
	msgs := make([]message, 0)
	return &conversation{conn: socket,
		messages: msgs}
}

//receiver constantly listens for messages, and when it receives them
//forwards them to the python api
func (c *conversation) receiver() {

}
