package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type wr struct {
	w http.ResponseWriter
	r *http.Request
}

func initTopicQueues() map[string](chan wr) {
	topics := make(map[string](chan wr))

	topics["Brexit"] = make(chan wr, 10)
	topics["Abortion"] = make(chan wr, 10)
	topics["Religion"] = make(chan wr, 10)

	return topics
}

func upgrade(obj wr) *websocket.Conn {
	socket, err := upgrader.Upgrade(obj.w, obj.r, nil)
	if err != nil {
		panic(err)
	}
	return socket
}

func genID() string {
	return "anvir"
}
