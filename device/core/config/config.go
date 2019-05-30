package config

import (
	"os"

	"github.com/aWildProgrammer/fconf"
)

//硬件服务器配置信息
type Config struct {
	//日志相关
	Stdlog string
	Errlog string

	//redis相关
	RedisIp   string
	RedisPort string
	RedisPwd  string
	RedisDB   string

	//数据库相关
	MysqlHost     string
	MysqlPort     string
	MysqlUsername string
	MysqlPwd      string
	MysqlDatabase string

	//队列相关
	HandleList string
	SendList   string
	ResendList string

	//服务相关
	ServerIp   string
	ServerPort string

	//设备相关
	HeartTimeOut int

	//处理数据相关
	Worknum    int
	Resendtime int64
	Sendnum    int
}

/**
 * @Function 获取配置
 * @Auther Nelg
 */
func GetConfig() (config Config) {
	//获取配置
	allConfig, err := fconf.NewFileConf("device.ini")
	if err != nil {
		os.Exit(-1)
	}
	//日志相关
	config.Stdlog = allConfig.String("log.stdlog")
	config.Errlog = allConfig.String("log.errlog")

	//redis相关
	config.RedisIp = allConfig.String("redis.ip")
	config.RedisPort = allConfig.String("redis.port")
	config.RedisPwd = allConfig.String("redis.pwd")
	config.RedisDB = allConfig.String("redis.db")

	//mysql相关
	config.MysqlHost = allConfig.String("mysql.host")
	config.MysqlPort = allConfig.String("mysql.port")
	config.MysqlUsername = allConfig.String("mysql.username")
	config.MysqlPwd = allConfig.String("mysql.pwd")
	config.MysqlDatabase = allConfig.String("mysql.database")

	//队列相关
	config.HandleList = allConfig.String("list.handle")
	config.SendList = allConfig.String("list.send")
	config.ResendList = allConfig.String("list.resend")

	//服务相关
	config.ServerIp = allConfig.String("server.ip")
	config.ServerPort = allConfig.String("server.port")

	//设备相关
	config.HeartTimeOut, _ = allConfig.Int("device.heart_time_out")

	//处理数据相关
	config.Worknum, _ = allConfig.Int("handle.worknum")
	config.Resendtime, _ = allConfig.Int64("handle.resendtime")
	config.Sendnum, err = allConfig.Int("handle.sendnum")
	return
}
