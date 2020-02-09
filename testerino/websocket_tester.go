package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var wg sync.WaitGroup

type incomingMsg struct {
	UID    int
	ConvID string
	Msg    string
}

type outgoingMsg struct {
	UID       int
	Msg       string
	Troll     float32
	Relevance float32
}

func test() {
	addr := "ws://146.169.207.172:8080/ws/Brexit/TESTER12"
	w := websocket.Dialer{}
	ws, _, err := w.Dial(addr, nil)
	if err != nil {
		panic(err)
	}
	ws.WriteJSON(incomingMsg{UID: 0, ConvID: "TESTER12", Msg: "Hello there"})
	var msg outgoingMsg
	ws.ReadJSON(&msg)
	fmt.Println(msg)
	wg.Done()
}

func main() {

	for i := 0; i < 1000; i++ {
		go test()
		time.Sleep(time.Millisecond * 50)
		wg.Add(1)
	}
	wg.Wait()
}
