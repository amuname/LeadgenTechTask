package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if c.Errors != nil {
		log.Println(c.Errors.String())

		c.JSON(-1, c.Errors)
	}
}
