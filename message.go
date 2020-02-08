package main

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
