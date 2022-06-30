package main

import (
	"github.com/wubingwei/logging_mock/handler"
	"github.com/wubingwei/logging_mock/mock"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin.DefaultWriter = ioutil.Discard

	region := os.Args[1]
	goroutine, _ := strconv.Atoi(os.Args[2])
	interval, _ := strconv.Atoi(os.Args[3])

	r := gin.Default()

	r.GET("/ping", handler.Ping)

	r.POST("/frequency", handler.Frequency)

	log.Printf("mock JsonLog start, goroutine: %d,interval: %d ms", goroutine, interval)
	for i := goroutine; i > 0; i-- {
		mock.JsonLog(time.Duration(interval), region)
	}
	log.Printf("========================================")
	if err := r.Run(); err != nil {
		os.Exit(1)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
