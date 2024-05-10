package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type error struct {
	Message string `json:"message"`
}

func errorResponse(c *gin.Context, status int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(status, error{message})
}
