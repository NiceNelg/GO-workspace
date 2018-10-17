package handle

import (
	"database/sql"

	"../config"
	"github.com/garyburd/redigo/redis"
)

type Handle struct {
	//mysql连接
	db *sql.DB
	//redis连接
	redisPool *redis.Pool

	//队列相关
	handleList string
	sendList   string
	resendList string

	//处理数据相关
	worknum  int
	sendtime int64
	sendnum  int
}

/**
 * @Function 初始化数据处理操作
 * @Auther Nelg
 */
func Init(db *sql.DB, redisPool *redis.Pool, allConfig config.Config) (handle Handle) {
	handle = Handle{
		redisPool:  redisPool,
		db:         db,
		handleList: allConfig.HandleList,
		sendList:   allConfig.SendList,
		resendList: allConfig.ResendList,
		worknum:    allConfig.Worknum,
		sendtime:   allConfig.Sendtime,
		sendnum:    allConfig.Sendnum,
	}
	return
}
