package handler

import (
	"bufio"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
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
	sc := bufio.NewScanner(c.Request.Body)

	num := 0
	t := time.Now()
	for sc.Scan() {
		fc := new(mock.Frequency)
		if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(sc.Bytes(), fc); err != nil {
			continue
		}
		num++
		mock.ContainerObj.Count(fc.Region, time.Now().UnixMilli()-fc.EventTime)
	}

	//mock.ContainerObj.Count(fc.Region, time.Now().Unix()-fc.EventTime)
	c.JSON(http.StatusOK, gin.H{
		"status":       "ok",
		"num":          num,
		"cost_time_ms": time.Since(t).Milliseconds(),
	})
}
