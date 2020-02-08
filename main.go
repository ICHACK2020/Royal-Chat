package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	convos   = make(map[string]*conversation)
	upgrader = websocket.Upgrader{ //TODO make this smarter
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	//topicQueues = initTopicQueues()
)

func dummyHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer socket.Close()
	_, p, err := socket.ReadMessage()
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
	for i := 0; i < 10; i++ {
		err := socket.WriteMessage(1, []byte("hello"))
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	//go newConversation()

	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("css/"))))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", http.FileServer(http.Dir("js/"))))
	http.Handle("/static/imgs/", http.StripPrefix("/static/imgs/", http.FileServer(http.Dir("html/images"))))

	http.HandleFunc("/", homepageHandler)
	http.HandleFunc("/talk", chatHandler)
	http.HandleFunc("/ws", dummyHandler)
	http.ListenAndServe(":8080", nil)
	//http.ListenAndServe("146.169.207.172:8080", nil)
}
