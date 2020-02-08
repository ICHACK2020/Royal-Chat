package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	convos   = make(map[int]*conversation)
	upgrader = websocket.Upgrader{ //Make this smarter
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func dummyHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer socket.Close()
	messageType, msg, err := socket.ReadMessage()
	fmt.Println(messageType, msg)
	for i := 0; i < 10; i++ {
		err := socket.WriteMessage(messageType, []byte("hello world"))
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	http.HandleFunc("/", dummyHandler)
	http.ListenAndServe("146.169.207.172:80800", nil)
}
