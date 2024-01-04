package main

import (
	"log"
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joy2362/go_blog/helper"
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
	log.Printf("Starting web server>>>>")

	r := gin.New()

	r.GET("/", home)
	r.GET("/post", index)
	r.POST("/post", store)
	r.GET("/post/:id", show)
	r.PUT("/post/:id", update)
	r.DELETE("/post/:id", destroy)

	r.Run()
}

/**
 * home.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Thursday, January 4th, 2024.
 * @global
 * @param	c	*gin.Context
 * @return	void
 */
func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

/**
 * index.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Thursday, January 4th, 2024.
 * @global
 * @param	c	*gin.Context
 * @return	void
 */
func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":    posts,
		"success": true,
	})
}

/**
 * store.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Thursday, January 4th, 2024.
 * @global
 * @param	c	*gin.Context
 * @return	void
 */
func store(c *gin.Context) {
	var post Post
	if err := c.BindJSON(&post); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusNotFound, helper.ErrorResponse("Something went wrong!!"))
		return
	}
	posts = append(posts, post)

	c.JSON(http.StatusOK, gin.H{
		"data":    post,
		"success": true,
	})
}

/**
 * show.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Thursday, January 4th, 2024.
 * @global
 * @param	c	*gin.Context
 * @return	void
 */
func show(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	helper.Panic(err)
	index := slices.IndexFunc(posts, func(p Post) bool { return int64(p.ID) == id })

	if index == -1 {
		c.JSON(http.StatusNotFound, helper.ErrorResponse("post not found"))
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    posts[index],
			"success": true,
		})
	}
}

/**
 * update.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Thursday, January 4th, 2024.
 * @global
 * @param	c	*gin.Context
 * @return	void
 */
func update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	helper.Panic(err)
	index := slices.IndexFunc(posts, func(p Post) bool { return int64(p.ID) == id })
	if index == -1 {
		c.JSON(http.StatusNotFound, helper.ErrorResponse("post not found"))
	} else {
		var post Post
		if err := c.BindJSON(&post); err != nil {
			log.Fatal(err)
			c.JSON(http.StatusNotFound, helper.ErrorResponse("Something went wrong!!"))
			return
		}
		posts[index] = post
		c.JSON(http.StatusOK, gin.H{
			"post":    post,
			"success": true,
		})
	}
}

/**
 * destroy.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Thursday, January 4th, 2024.
 * @global
 * @param	c	*gin.Context
 * @return	void
 */
func destroy(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	helper.Panic(err)
	index := slices.IndexFunc(posts, func(p Post) bool { return int64(p.ID) == id })
	if index == -1 {
		c.JSON(http.StatusNotFound, helper.ErrorResponse("post not found"))
	} else {
		posts = append(posts[:index], posts[index+1:]...)
		c.JSON(http.StatusOK, gin.H{
			"message": "Post remove successfully",
			"success": true,
		})
	}
}
