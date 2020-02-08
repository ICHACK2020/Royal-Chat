package main

import "github.com/gorilla/websocket"

type conversation struct {
	Conn *websocket.Conn
	ID   int
}
