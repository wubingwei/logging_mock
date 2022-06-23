package mock

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/wubingwei/logging_mock/log2file"
	"log"
	"time"
)

func JsonLog(interval time.Duration, region string) {
	log.Printf("mock JsonLog start, interval: %d ms", 1)
	for range time.Tick(interval * time.Millisecond) {
		logJson := Frequency{
			EventTime: time.Now().Unix(),
			Region:    region,
		}
		payload, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(logJson)
		log2file.Forward.Info(string(payload))
	}
}
