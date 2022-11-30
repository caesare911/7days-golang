package main

import (
	"7days-golang/gee-web/day5-middlewares/gee"
	"net/http"
)

func main() {
	r := gee.Default()
	v1 := r.Group("/v1")
	v1.Use(gee.Logger())
	v1.GET("/hello", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})
	v1.GET("/hello1", func(c *gee.Context) {
		c.String(http.StatusOK, "6666")
	})

	r.Run(":8000")
}
