package main

import (
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
	topicQueues = initTopicQueues()
)

func main() {
	go newConversation()

	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("css/"))))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", http.FileServer(http.Dir("js/"))))
	http.Handle("/static/imgs/", http.StripPrefix("/static/imgs/", http.FileServer(http.Dir("/html/images"))))

	http.HandleFunc("/", homepageHandler)
	http.HandleFunc("/talk/", chatHandler)
	http.ListenAndServe("146.169.207.172:8080", nil)
}
