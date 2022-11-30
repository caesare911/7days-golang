package main

import (
	"fmt"
	"main/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello")
	})
	r.POST("/go", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "go")
	})

	r.Run(":8000")

}
