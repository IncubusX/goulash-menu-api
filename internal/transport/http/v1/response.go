package v1

import (
	"github.com/gin-gonic/gin"
	"log"
)

const (
	ErrServiceFailure = "service failure"
)

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type errorResponse struct {
	Message string `json:"message"`
}
