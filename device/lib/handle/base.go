package handle

import (
	"../config"
	"github.com/garyburd/redigo/redis"
)

type Handle struct {
	//redis连接
	redisPool *redis.Pool

	//队列相关
	handleList string
	sendList   string
	resendList string

	//处理数据相关
	worknum  int
	sendtime int
	sendnum  int
}

/**
 * @Function 初始化数据处理操作
 * @Auther Nelg
 */
func Init(redisPool *redis.Pool, allConfig config.Config) (handle Handle) {
	handle = Handle{
		redisPool:  redisPool,
		handleList: allConfig.HandleList,
		sendList:   allConfig.SendList,
		resendList: allConfig.ResendList,
		worknum:    allConfig.Worknum,
		sendtime:   allConfig.Sendtime,
		sendnum:    allConfig.Sendnum,
	}
	return
}
