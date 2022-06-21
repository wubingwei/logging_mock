package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Frequency struct {
	EventTime       int64   `json:"event_time"`        // 展示发生的时间戳
	Adx             string  `json:"adx"`               // adx 名称,"mintegral、oppocn、doubleclick"...
	Os              int32   `json:"os"`                // 操作系统，android/ios
	CountryCode     string  `json:"country_code"`      // 国家标识
	ReqPkgName      string  `json:"req_pkg_name"`      // 流量侧包名
	AdType          int32   `json:"adtype"`            // M 体系下的 int 广告类型，dsp需要做映射。https://gitlab.mobvista.com/voyager/common/-/blob/master/enum/enum_adtype.go#L185
	RequestId       string  `json:"request_id"`        // 请求id
	OneId           string  `json:"one_id"`            // 设备唯一id
	CampaignId      int64   `json:"campaign_id"`       // 单子id
	CampaignPkgName string  `json:"campaign_pkg_name"` // 单子包名
	UserActivation  uint8   `json:"user_activation"`   // 1: 拉活；2:拉新
	BussinessType   uint8   `json:"bussiness_type"`    // 事件类型，1:impression;2:click;3:conversion;4:install
	IsTryNew        bool    `json:"is_try_new"`        // dsp使用，ext3["try_new"] 为1则是试新单子，无或者为0则为其他单子
	Price           float32 `json:"price"`             // dsp使用，实际成交价格,千次价格
	OnlyIpua        bool    `json:"only_ipua"`         // 0:有合法设备id，1:只有ipua
	Region          string  `json:"region"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/frequency", func(c *gin.Context) {
		fc := new(Frequency)
		if err := c.ShouldBindJSON(fc); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
