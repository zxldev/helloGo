package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}



func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func (h String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, h)
}


func (h Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, h.Greeting+","+h.Who)
}

func main() {
	var h Hello
	http.Handle("/", h)
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})


	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
