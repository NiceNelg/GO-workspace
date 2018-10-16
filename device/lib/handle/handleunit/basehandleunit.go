package handleunit

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"strconv"

	"../../../tool"
	"../../data"
	"../../model"
	"../../protocol808"
	"github.com/garyburd/redigo/redis"
)

type Hand interface {
	HandleBusiness() Hand
	HandleSend()
	SaveToSendList(redisPool *redis.Pool, sendList string)
	SaveToDatabase(db *sql.DB)
}

type HandUnit struct {
	data.Data
	//是否存储到缓存
	StoredCache bool
	//是否记录到数据库
	StoredDatabase bool
}

/**
 * @Function 数据包存入发送队列	//TODO 可按需重写
 * @Auther Nelg
 */
func (this *HandUnit) SaveToSendList(redisPool *redis.Pool, sendList string) {
	if !this.StoredCache {
		return
	}
	this.HandleSend()
	//从redis中取出连接
	redisCli := redisPool.Get()
	_, err := redisCli.Do("lpush", sendList+this.Device)
	if err != nil {

	}
	//归还redis连接到redis连接池
	redisCli.Close()
}

/**
 * @Function 数据包存入到数据库	//TODO 可按需重写
 * @Auther Nelg
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
	//添加标识头尾
	dataByte = append(dataByte, 0x7e)
	dataByte = append([]byte{0x7e}, dataByte...)
	this.Content = hex.EncodeToString(dataByte)
	return
}
