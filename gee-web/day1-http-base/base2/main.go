package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "Url.Path=%q\n", "/")
	case "/hello":
		fmt.Fprintf(w, "Url.Path=%q\n", "/hello")
	default:
		fmt.Fprintf(w, "404 NOT FOUND=%q\n", req.URL.Path)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":8000", engine))
}
