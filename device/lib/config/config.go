package config

import (
	"os"

	"github.com/aWildProgrammer/fconf"
)

//硬件服务器配置信息
type Config struct {
	Stdlog       string
	Errlog       string
	RedisIp      string
	RedisPort    string
	HandleList   string
	SendList     string
	ResendList   string
	ServerIp     string
	ServerPort   string
	HeartTimeOut int
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
	config.Stdlog = allConfig.String("log.stdlog")
	config.Errlog = allConfig.String("log.errlog")
	config.RedisIp = allConfig.String("redis.ip")
	config.RedisPort = allConfig.String("redis.port")
	config.HandleList = allConfig.String("list.handle")
	config.SendList = allConfig.String("list.send")
	config.ResendList = allConfig.String("list.resend")
	config.ServerIp = allConfig.String("server.ip")
	config.ServerPort = allConfig.String("server.port")
	config.HeartTimeOut, _ = allConfig.Int("device.heart_time_out")
	return
}
