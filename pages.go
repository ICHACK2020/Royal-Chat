package main

import (
	"html/template"
	"http"
)

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./html/index.html")
	if err != nil {
		panic(err)
	}

}
