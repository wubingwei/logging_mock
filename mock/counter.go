package mock

import (
	"github.com/wubingwei/logging_mock/log2file"
	"sync"
)

type Frequency struct {
	Region      string  `json:"rg"`  // 日志生产的region
	EventTime   int64   `json:"t"`   // 展示发生的时间戳
	Adx         string  `json:"adx"` // adx 名称,"mintegral、oppocn、doubleclick"...
	Os          string  `json:"os"`  // 操作系统，android/ios
	RequestType int32   `json:"rt"`  // m 流量请求类型,见下面枚举
	CountryCode string  `json:"cc"`  // 国家标识
	ReqPkgName  string  `json:"rpn"` // 流量侧包名
	PublisherId int64   `json:"pid"` // 开发者id，adn日志中publsherId为"6028"的是dsp数据
	AdType      string  `json:"at"`  // 广告类型
	RequestId   string  `json:"rid"` // 请求id
	OneId       string  `json:"oid"` // 设备唯一id
	EventType   int32   `json:"et"`  // 展示、点击、安装/转化。1:impression;2:click;3:conversion;4:install
	CampaignId  int64   `json:"cid"` // 单子id
	TryNew      bool    `json:"tn"`  // dsp使用，ext3["try_new"] 为1则是试新单子，无或者为0则为其他单子
	Price       float64 `json:"p"`   // dsp使用，实际成交价格,千次价格
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
	c[region].Lock()
	defer c[region].Unlock()
	c[region].N += 1
	c[region].Delta += delta
	if c[region].N == 1000 {
		log2file.Frequency.Infof("[region: %s]|[avg_time: %.3f]", region, float64(counter.Delta)/float64(counter.N))
		c[region].N, c[region].Delta = 0, 0
	}
}
