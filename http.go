package main

import (
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func main() {
	// your http.Handle calls here
	http.ListenAndServe("localhost:4000", nil)
}
