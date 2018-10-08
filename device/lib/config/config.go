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
	Worknum  int
	Sendtime int
	Sendnum  int
}

/**
 * @Function 获取配置
 * @Auther Nelg
 */
func GetConfig() (config Config) {
	//获取配置
	allConfig, err := fconf.NewFileConf("../device.ini")
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
	config.Sendtime, _ = allConfig.Int("handle.sendtime")
	config.Sendnum, _ = allConfig.Int("handle.sendnum")
	return
}
