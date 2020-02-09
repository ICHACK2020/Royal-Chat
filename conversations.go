package main

import (
	"context"
	"fmt"
	"log"

	api "ichack2020/proto"

	"github.com/gorilla/websocket"
)

type conversation struct {
	user1    *websocket.Conn
	user2    *websocket.Conn
	incoming chan incomingMsg
	topic    string
	u1r      int
	u2r      int
}

//Constantly connects users
func newConversation() {
	for k, channel := range topicQueues {
		go func(channel chan *websocket.Conn, k string) {
			for {
				u1 := <-channel
				u2 := <-channel
				conv := &conversation{
					u1,
					u2,
					make(chan incomingMsg),
					k, 10, 10,
				}
				conv.user1.WriteMessage(1, []byte("0"))
				conv.user2.WriteMessage(1, []byte("1"))
				go conv.read(conv.user1)
				go conv.read(conv.user2)
				go conv.receiver()
				convos[genID()] = conv
			}
		}(channel, k)
	}
}

//Reads from either connection
func (c *conversation) read(conn *websocket.Conn) incomingMsg {
	fmt.Println("reading")
	var msg incomingMsg
	for {
		err := conn.ReadJSON(&msg)
		if err != nil {
			c.user1.Close()
			c.user2.Close()
			return incomingMsg{}
		}
		c.incoming <- msg
	}
}

//Write to both connections
func (c *conversation) write(msg outgoingMsg) {
	defer c.user1.Close()
	defer c.user2.Close()
	err := c.user1.WriteJSON(msg)
	if err != nil {
		return
	}
	err = c.user2.WriteJSON(msg)
	if err != nil {
		return
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

		if len(msg.Msg) == 0 {
			return
		}
		//Python stuff
		data := &api.ApiCall{ConvId: c.topic + "/" + msg.ConvID,
			Uid: !(msg.UID == 0),
			Msg: msg.Msg}

		response, err := trollClient.Troll(context.Background(), data)
		var troll float32
		if err != nil {
			log.Fatal(err)
			troll = 0.5
		} else {
			troll = response.GetScore()
		}

		rollingScore := response.GetRollingScore()
		if rollingScore > 0.8 || troll > 0.85 {
			return
		}

		relevanceResponse, err := relevanceClient.Relevance(context.Background(), data)
		var relevance float32
		if err != nil {
			log.Fatal(err)
			relevance = 0
		} else {
			relevance = relevanceResponse.GetScore()
		}

		if msg.UID == 0 {
			c.u1r -= (1 - int(relevance))
			if c.u1r == 0 {
				return
			}
		} else {
			c.u2r -= (1 - int(relevance))
			if c.u1r == 0 {
				return
			}
		}
		//Non python stuff

		outgoing := outgoingMsg{UID: msg.UID,
			Msg:       msg.Msg,
			Troll:     troll,
			Relevance: relevance}
		c.write(outgoing)
	}
}
