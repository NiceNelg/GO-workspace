package handleunit

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"strconv"

	"core/data"
	"core/model"
	"core/protocol808"
	"github.com/garyburd/redigo/redis"
	"tool"
)

/**
 * @Interface 对数据进行操作的接口
 * @Auther Nelg
 * @Date 2019.05.30
 */
type Hand interface {
	HandleBusiness() Hand
	HandleSend()
	SaveToSendList(*redis.Pool, string)
	SaveToDatabase(db *sql.DB)
}

/**
 * @Struct 数据单元
 * @Auther Nelg
 * @Date 2019.05.30
 */
type HandUnit struct {
	data.Data
	/*是否存储到缓存*/
	StoredCache bool `json:"-"`
	/*是否记录到数据库*/
	StoredDatabase bool `json:"-"`
}

/**
 * @Function 数据包存入发送队列	//TODO 可按需重写
 * @Auther Nelg
 * @Date 2019.05.30
 */
func (this *HandUnit) SaveToSendList(redisPool *redis.Pool, sendList string) {
	if !this.StoredCache {
		return
	}
	cmd, _ := json.Marshal(this)
	//从redis中取出连接
	redisCli := redisPool.Get()
	//归还redis连接到redis连接池
	defer redisCli.Close()
	redisCli.Do("lpush", sendList+"_"+this.Device, string(cmd))
}

/**
 * @Function 数据包存入到数据库	//TODO 可按需重写
 * @Auther Nelg
 * @Date 2019.05.30
 */
func (this *HandUnit) SaveToDatabase(db *sql.DB) {
	if !this.StoredDatabase {
		return
	}
	commandModel := model.CommandModelInit(db)
	//保存指令
	id := commandModel.SaveCommand(this.Data)
	sn := strconv.FormatInt(id, 16)
	this.Sn = tool.StrPad(sn, "0", 4, "LEFT")
	//更新指令流水号
	commandModel.SaveSn(id, this.Sn)
	return
}

/**
 * @Function 组成发送数据	//TODO 可按需重写
 * @Auther Nelg
 * @Date 2019.05.30
 */
func (this *HandUnit) HandleSend() {
	//组成发送数据
	content := this.Body["ack_sign"] + this.Body["ack_sn"] + this.Body["result"]
	this.Attribute = tool.StrPad(strconv.Itoa(len(content)/2), "0", 4, "LEFT")
	data := this.Sign + this.Attribute + this.Device + this.Sn + content
	//转换成[]byte
	dataByte, _ := hex.DecodeString(data)
	//异或计算
	dataByte = append(dataByte, protocol808.BuildBCC(dataByte))
	//转义
	dataByte = protocol808.Escape(dataByte)
	this.Content = hex.EncodeToString(dataByte)
	return
}
