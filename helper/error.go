package helper

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Panic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorResponse(message string) gin.H {
	return gin.H{
		"error":   message,
		"success": false,
	}
}
