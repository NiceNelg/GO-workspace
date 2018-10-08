package handle

import (
	"encoding/json"
	"time"

	"../data"
	"github.com/garyburd/redigo/redis"
)

/**
 * @Function 开启数据处理协程
 * @Auther Nelg
 */
func (this *Handle) StartHandle() {
	return
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
		this.dispense(cmd)
	}
}

/**
 * @Function 分发业务（根据协议不同需要重写标识）
 * @Auther Nelg
 */
func (this *Handle) dispense(cmd data.Data) {
	switch cmd.Sign {
	case "0102":

	}
}
