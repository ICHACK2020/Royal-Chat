package main

import (
	"html/template"
	"net/http"
	"strings"
)

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./html/index.html")
	if err != nil {
		panic(err)
	}
	r.ParseForm()
	topic := r.FormValue("topic")
	if len(topic) > 0 {
		http.Redirect(w, r, "/talk/"+topic+"/"+genID(), http.StatusFound)
	}
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
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	topicQueues[strings.Split(r.URL.Path, "/")[2]] <- socket
}
