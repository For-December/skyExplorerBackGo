package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProcessesGetAllHandler(c *gin.Context) {

	// TODO
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func ProcessesGetByIdHandler(c *gin.Context) {

	// TODO
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func ProcessesExecutionHandler(c *gin.Context) {

	// TODO
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
