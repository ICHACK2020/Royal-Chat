package main

import (
	"fmt"
	"gorilla/websocket"
	"sync"
)

type incomingMsg struct {
	UID    int
	ConvID string
	Msg    string
}

func test() {
	addr := "ws://146.169.207.172:8080/ws/Brexit/TESTER12"
	ws, _, err := websocket.Dialer().Dial(addr)
	if err != nil {
		panic(err)
	}
	ws.WriteJSON(incomingMsg{UID: 0, ConvID: "TESTER12", Msg: "Hello there"})
	var msg incomingMsg
	ws.ReadJSON(&msg)
	fmt.Println(*msg)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go test()
		wg.Add(1)
	}
	wg.Wait()
}
