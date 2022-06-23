package mock

import (
	"github.com/wubingwei/logging_mock/log2file"
	"sync"
)

type Frequency struct {
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
	Region          string  `json:"region"`            // 流量region
}

var ContainerObj Container

func init() {
	ContainerObj = NewContainer()
}

type Container map[string]*Counter

type Counter struct {
	N     int64
	Delta int64
	sync.RWMutex
}

func NewContainer() map[string]*Counter {
	return make(map[string]*Counter)
}

func (c Container) Count(region string, delta int64) {
	counter, ok := c[region]
	if !ok {
		c[region] = &Counter{
			RWMutex: sync.RWMutex{},
		}
	}
	counter.Lock()
	defer counter.Unlock()
	counter.N += 1
	counter.Delta += delta
	if counter.N == 100 {
		log2file.Frequency.Infof("[region: %s]|[avg_time: %.3f]", region, float64(counter.Delta)/float64(counter.N))
		counter.N, counter.Delta = 0, 0
	}
}
