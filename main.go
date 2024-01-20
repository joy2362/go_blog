package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joy2362/go_blog/helper"
)

type Post struct {
	ID     uint64 `json:"id"`
	POST   string `json:"post"`
	Author string `json:"author"`
}

var dbConfig = mysql.Config{
	User:                 "root",
	Passwd:               "",
	Net:                  "tcp",
	Addr:                 "127.0.0.1:3306",
	DBName:               "blog",
	AllowNativePasswords: true,
}

const routeUrl string = "/post/:id"

func main() {
	log.Printf("Starting web server>>>>")

	r := gin.New()

	r.GET("/", home)
	r.GET("/post", index)
	r.POST("/post", store)
	r.GET(routeUrl, show)
	r.PUT(routeUrl, update)
	r.DELETE(routeUrl, destroy)

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
 * @version	v1.0.1	Saturday, January 20th, 2024.
 * @global
 * @param	c	*gin.Context
 * @return	void
 */
func index(c *gin.Context) {

	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	helper.Panic(err)

	pinngErr := db.Ping()
	helper.Panic(pinngErr)

	var content []Post
	rows, err := db.Query("select * from content")

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, helper.ErrorResponse("Someting went wrong!!"))
		return
	}

	defer db.Close()
	defer rows.Close()

	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.POST, &post.Author); err != nil {
			log.Println(err)
			c.JSON(http.StatusNotFound, helper.ErrorResponse("Someting went wrong!!"))
			return
		}
		content = append(content, post)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, helper.ErrorResponse("Someting went wrong!!"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    content,
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
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	helper.Panic(err)

	pinngErr := db.Ping()
	helper.Panic(pinngErr)

	var post Post
	if err := c.BindJSON(&post); err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, helper.ErrorResponse("Something went wrong!!"))
		return
	}

	result, err := db.Exec("insert into content(post , author) values (?,?)", post.POST, post.Author)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusNotFound, helper.ErrorResponse("Something went wrong!!"))
		return
	}

	defer db.Close()

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusNotFound, helper.ErrorResponse("Something went wrong!!"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    id,
		"success": true,
	})
}

/**
 * show.
 *
 * @author	Joy2362
 * @since	v0.0.1
 * @version	v1.0.0	Thursday, January 4th, 2024.
 * @version	v1.0.1	Saturday, January 20th, 2024.
 * @global
 * @param	c	*gin.Context
 * @return	void
 */
func show(c *gin.Context) {
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	helper.Panic(err)

	pinngErr := db.Ping()
	helper.Panic(pinngErr)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	helper.Panic(err)

	var content Post
	row := db.QueryRow("select * from content where id = ?", id)

	defer db.Close()

	if err := row.Scan(&content.ID, &content.POST, &content.Author); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, helper.ErrorResponse("post not found"))
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    content,
		"success": true,
	})
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

	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	helper.Panic(err)

	pinngErr := db.Ping()
	helper.Panic(pinngErr)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	helper.Panic(err)

	var post Post
	if err := c.BindJSON(&post); err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, helper.ErrorResponse("Something went wrong!!"))
		return
	}

	result, err := db.Exec("update content set post = ? , author = ? where id = ? ", post.POST, post.Author, id)

	defer db.Close()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, helper.ErrorResponse("Something went wrong!!"))
		return
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, helper.ErrorResponse("Something went wrong!!"))
		return
	}

	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "content update successfully",
			"success": true,
		})
	} else {
		c.JSON(http.StatusNotFound, helper.ErrorResponse("content not found or nothing to upate!!"))
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

	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	helper.Panic(err)

	pinngErr := db.Ping()
	helper.Panic(pinngErr)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	helper.Panic(err)

	result, err := db.Exec("delete from content where id = ? ", id)

	defer db.Close()

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, helper.ErrorResponse("Something went wrong!!"))
		return
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, helper.ErrorResponse("Something went wrong!!"))
		return
	}

	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Post remove successfully",
			"success": true,
		})
	} else {
		c.JSON(http.StatusNotFound, helper.ErrorResponse("post not found!!"))
	}
}
