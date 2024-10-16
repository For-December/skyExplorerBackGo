package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JobsGetAllHandler(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func JobsGetByIdHandler(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func JobsDeleteByIdHandler(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func JobsCancelHandler(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
