package main

// curl -i http://localhost:8080/index
// curl -i http://localhost:8080/v1/
// curl "http://localhost:8080/v1/hello?name=geektutu"
// curl "http://localhost:8080/v2/hello/geektutu"
// curl "http://localhost:8080/v2/login" -X POST -d 'username=geektutu&password=1234'
// curl "http://localhost:8080/hello"

import (
	"net/http"

	"gee"
)

func main() {
	r := gee.New()
	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run("127.0.0.1:8080")
}
