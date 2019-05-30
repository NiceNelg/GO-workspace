package handle

import (
	"core/config"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"core/data"
	"core/handle/handleunit"
	"github.com/garyburd/redigo/redis"
)

/**
 * @Struct 服务操作
 * @Auther Nelg
 * @Date 2019.05.30
 */
type Handle struct {
	/*mysql连接池*/
	db *sql.DB
	/*redis连接池*/
	redisPool *redis.Pool

	/*队列相关redis-key*/
	handleList string
	sendList   string
	resendList string

	/*处理数据相关配置*/
	worknum    int
	resendtime int64
	sendnum    int
}

/**
 * @Function 初始化服务操作对象
 * @Auther Nelg
 * @Date 2019.05.30
 */
func Init(db *sql.DB, redisPool *redis.Pool, allConfig *config.Config) (handle Handle) {
	handle = Handle{
		redisPool:  redisPool,
		db:         db,
		handleList: allConfig.HandleList,
		sendList:   allConfig.SendList,
		resendList: allConfig.ResendList,
		worknum:    allConfig.Worknum,
		resendtime: allConfig.Resendtime,
		sendnum:    allConfig.Sendnum,
	}
	return
}

/**
 * @Function 开启数据处理协程
 * @Auther Nelg
 * @Date 2019.05.30
 */
func (this *Handle) StartHandle() {
	for i := 0; i < this.worknum; i++ {
		go this.invoke()
	}
	return
}

/**
 * @Function 处理业务
 * @Auther Nelg
 * @Date 2019.05.30
 */
func (this *Handle) invoke() {
	for {
		/*从redis中取出数据*/
		redisCli := this.redisPool.Get()
		redisData, err := redis.String(redisCli.Do("rpop", this.handleList))
		/*归还redis连接到redis连接池*/
		redisCli.Close()
		if err != nil || redisData == "" {
			time.Sleep(1 * time.Second)
			continue
		}
		var cmd data.Data
		err = json.Unmarshal([]byte(redisData), &cmd)
		if err != nil {
			continue
		}
		/*分发业务*/
		hand, err := this.dispense(cmd)
		if err != nil {
			continue
		}
		/*处理业务*/
		send := hand.HandleBusiness()
		/*若没有返回发送对象则代表此数据包不需要下发*/
		if send == nil {
			continue
		}
		/*存入数据库*/
		send.SaveToDatabase(this.db)
		/*组成发送数据*/
		send.HandleSend()
		/*存入redis队列*/
		send.SaveToSendList(this.redisPool, this.sendList)
	}
	return
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
		err = errors.New("dispense failr")
	}
	return
}
