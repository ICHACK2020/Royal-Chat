package main

import (
	"math/rand"

	"github.com/gorilla/websocket"
)

func initTopicQueues() map[string](chan *websocket.Conn) {
	topics := make(map[string](chan *websocket.Conn))

	topics["Brexit"] = make(chan *websocket.Conn)
	topics["US Election"] = make(chan *websocket.Conn)
	topics["Coronavirus"] = make(chan *websocket.Conn)

	return topics
}

func genID() string {
	id := make([]byte, 8)
	for i := 0; i < 8; i++ {
		id[i] = chars[rand.Intn(16)]
	}
	return string(id)
}
