package handler

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/wubingwei/logging_mock/mock"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Frequency(c *gin.Context) {
	sc := bufio.NewScanner(c.Request.Body)

	for sc.Scan() {
		fc := new(mock.Frequency)
		if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(sc.Bytes(), fc); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(*fc)
		}

	}

	//mock.ContainerObj.Count(fc.Region, time.Now().Unix()-fc.EventTime)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
