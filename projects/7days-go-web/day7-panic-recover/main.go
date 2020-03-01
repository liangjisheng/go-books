package main

// curl "http://localhost:8080"
// curl "http://localhost:8080/panic"

import (
	"net/http"

	"gee"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run("127.0.0.1:8080")
}
