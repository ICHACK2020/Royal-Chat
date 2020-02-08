package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./html/index.html")
	if err != nil {
		panic(err)
	}
	/*
		r.ParseForm()
		if len(r.FormValue("topic")) > 0 {
			if channel, ok := topicQueues[r.FormValue("topic")]; ok {
				channel <- wr{w: w, r: r}
			} else {
				fmt.Println("someone bypassed the system")
			}
		}*/
	t.Execute(w, nil)
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./html/chat.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws := upgrade(wr{w, r})
	fmt.Println(ws.RemoteAddr().String())
}
