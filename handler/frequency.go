package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wubingwei/logging_mock/mock"
	"net/http"
	"time"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Frequency(c *gin.Context) {
	fc := new(mock.Frequency)
	if err := c.ShouldBindJSON(fc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mock.ContainerObj.Count(fc.Region, time.Now().Unix()-fc.EventTime)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
