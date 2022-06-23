package main

import (
	"github.com/wubingwei/logging_mock/handler"
	"github.com/wubingwei/logging_mock/mock"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin.DefaultWriter = ioutil.Discard

	region := os.Args[0]
	interval, _ := strconv.Atoi(os.Args[1])

	r := gin.Default()

	r.GET("/ping", handler.Ping)

	r.POST("/frequency", handler.Frequency)

	go func() {
		mock.JsonLog(time.Duration(interval), region)
	}()

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
