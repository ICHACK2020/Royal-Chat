package main

//protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto api.proto
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
