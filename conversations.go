package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type conversation struct {
	user1    *websocket.Conn
	user2    *websocket.Conn
	incoming chan incomingMsg
}

//Constantly connects users
func newConversation() {
	for k, channel := range topicQueues {
		fmt.Println(k)
		go func(channel chan *websocket.Conn) {
			for {
				if len(channel) > 0 {
					fmt.Println("hi")
				} else {
					fmt.Println(len(channel), k)
				}
				u1 := <-channel
				u2 := <-channel
				conv := &conversation{
					u1,
					u2,
					make(chan incomingMsg),
				}
				conv.user1.WriteMessage(1, []byte("0"))
				conv.user2.WriteMessage(1, []byte("1"))
				go conv.read(conv.user1)
				go conv.read(conv.user2)
				go conv.receiver()
				convos[genID()] = conv
			}
		}(channel)
	}
}

//Reads from either connection
func (c *conversation) read(conn *websocket.Conn) incomingMsg {
	fmt.Println("reading")
	var msg incomingMsg
	for {
		err := conn.ReadJSON(&msg)
		if err != nil {
			panic(err)
		}
		c.incoming <- msg
	}
}

//Write to both connections
func (c *conversation) write(msg outgoingMsg) {
	err := c.user1.WriteJSON(msg)
	if err != nil {
		panic(err)
	}
	err = c.user2.WriteJSON(msg)
	if err != nil {
		panic(err)
	}
}

//receiver constantly listens for messages, and when it receives them
//forwards them to the python api
func (c *conversation) receiver() {
	defer c.user1.Close()
	defer c.user2.Close()
	var msg incomingMsg
	for {
		msg = <-c.incoming

		//Python stuff
		var troll float32 = 0.1
		rollingTroll := 0.1
		var relevance float32 = 0.6
		//rollingRelevance := 0.2
		if rollingTroll > 0.9 {

		}
		//Non python stuff

		outgoing := outgoingMsg{UID: msg.UID,
			Msg:       msg.Msg,
			Troll:     troll,
			Relevance: relevance}
		c.write(outgoing)
	}
}
