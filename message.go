package main

type message struct {
	Msg       string
	Troll     float32
	Relevance float32
}

type apiCall struct {
	ID  string
	Msg *message
}
