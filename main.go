package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID      uint64 `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var posts = []Post{
	{ID: 1, Content: "This is demo post details", Author: "joy"},
	{ID: 2, Content: "This is demo post details 2", Author: "joy"},
	{ID: 3, Content: "This is demo post details 3", Author: "joy"},
}

func main() {
	fmt.Println("Starting web server>>>>")
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	r.GET("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, posts)
	})

	r.Run()
}
