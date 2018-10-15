package handle

import (
	"encoding/json"
	"errors"
	"time"

	"../data"
	"./handleunit"
	"github.com/garyburd/redigo/redis"
)

/**
 * @Function 开启数据处理协程
 * @Auther Nelg
 */
func (this *Handle) StartHandle() {
	for i := 0; i < this.worknum; i++ {
		go this.invoke()
	}
}

/**
 * @Function 处理业务
 * @Auther Nelg
 */
func (this *Handle) invoke() {
	for {
		//从redis中取出数据
		redisCli := this.redisPool.Get()
		redisData, err := redis.String(redisCli.Do("rpop", this.handleList))
		//归还redis连接到redis连接池
		redisCli.Close()
		if err != nil || redisData == "" {
			time.Sleep(time.Duration(1) * time.Second)
			continue
		}
		var cmd data.Data
		err = json.Unmarshal([]byte(redisData), &cmd)
		if err != nil {
			continue
		}
		//分发业务
		hand, err := this.dispense(cmd)
		if err != nil {
			continue
		}
		//处理业务
		send := hand.HandleBusiness()
		//若没有返回发送对象则代表此数据包不需要下发
		if send == nil {
			continue
		}
		send.SaveToDatabase(this.db)
	}
}

/**
 * @Function 分发业务
 * @Auther Nelg
 */
func (this *Handle) dispense(cmd data.Data) (unit handleunit.Hand, err error) {
	switch cmd.Sign {
	//TODO 初始化业务结构体并解析协议内容（根据协议不同需要重写标识）
	case "0102":
		unit = handleunit.AuthcheckInit(cmd)
	default:
		errors.New("dispense failr")
	}
	return
}
