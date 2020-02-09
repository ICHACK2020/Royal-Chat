package main

import (
	"math/rand"
	"net/http"
	"time"

	api "ichack2020/proto"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

const (
	chars = "0123456789ABCDEF"
)

var (
	convos   = make(map[string]*conversation)
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	topicQueues     = initTopicQueues()
	trollClient     api.ProcessClient
	relevanceClient api.ProcessClient
)

func main() {
	rand.Seed(time.Now().UnixNano())
	go newConversation()

	conn, err := grpc.Dial("146.169.157.5:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	trollClient = api.NewProcessClient(conn)

	conn1, err := grpc.Dial("146.169.139.247:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn1.Close()

	relevanceClient = api.NewProcessClient(conn1)

	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("css/"))))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", http.FileServer(http.Dir("js/"))))
	http.Handle("/static/imgs/", http.StripPrefix("/static/imgs/", http.FileServer(http.Dir("html/images"))))

	http.HandleFunc("/", homepageHandler)
	http.HandleFunc("/talk/", chatHandler)
	http.HandleFunc("/ws/", wsHandler)
	http.ListenAndServe("146.169.207.172:8080", nil)
}
