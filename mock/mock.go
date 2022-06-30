package mock

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/wubingwei/logging_mock/log2file"
	"time"
)

func JsonLog(interval time.Duration, region string) {
	go func() {
		for range time.Tick(interval * time.Millisecond) {
			logJson := Frequency{
				EventTime:   time.Now().UnixMilli(),
				Region:      region,
				Adx:         "mintegral",
				Os:          "ios",
				RequestType: 8,
				CountryCode: "CN",
				ReqPkgName:  "aaa",
				PublisherId: 6028,
				AdType:      "b",
				RequestId:   "aaaaaaaaaa",
				OneId:       "aaaaaa",
				EventType:   1,
				CampaignId:  1212414,
				TryNew:      true,
				Price:       1.2,
			}
			payload, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(logJson)
			log2file.Forward.Info(string(payload))
		}
	}()
}
