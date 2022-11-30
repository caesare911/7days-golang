package main

import (
	"7days-golang/gee-web/day4-group/gee"
	"net/http"
)

func main() {
	r := gee.New()
	v1 := r.Group("v1")
	v1.GET("/hello", func(c *gee.Context) {
		c.HTML(http.StatusOK, "hello")
	})

	r.Run(":8000")
}
