package main

import (
	"github.com/gorilla/websocket"
)

type conversation struct {
	user1    *websocket.Conn
	user2    *websocket.Conn
	incoming chan incomingMsg
}

/*
//Constantly connects users
func newConversation() {
	for {
		for _, channel := range topicQueues {
			conv := &conversation{
				upgrade(<-channel),
				upgrade(<-channel),
				make(chan incomingMsg),
			}
			go conv.read(conv.user1)
			go conv.read(conv.user2)
			go conv.receiver()
			convos[genID()] = conv
		}
	}
}

//Reads from either connection
func (c *conversation) read(conn *websocket.Conn) incomingMsg {
	var msg incomingMsg
	for {
		err := conn.ReadJSON(msg)
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
*/
