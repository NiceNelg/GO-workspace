package device

import (
	"github.com/aWildProgrammer/fconf"
)

//硬件服务器配置信息
type DeviceConfig struct {
	stdlog         string
	errlog         string
	redisIp        string
	redisPort      string
	handleList     string
	sendList       string
	resendList     string
	serverIp       string
	serverPort     string
	heart_time_out int
}

/**
 * @Function 获取配置
 * @Auther Nelg
 */
func getConfig() (config DeviceConfig) {
	//获取配置
	allConfig, err := fconf.NewFileConf("../device.ini")
	if err != nil {

	}
	config.stdlog = allConfig.String("log.stdlog")
	config.errlog = allConfig.String("log.errlog")
	config.heart_time_out, _ = allConfig.Int("device.heart_time_out")
	config.redisIp = allConfig.String("redis.ip")
	config.redisPort = allConfig.String("redis.port")
	config.handleList = allConfig.String("list.handle")
	config.sendList = allConfig.String("list.send")
	config.resendList = allConfig.String("list.resend")
	config.serverIp = allConfig.String("server.ip")
	config.serverPort = allConfig.String("server.port")
	return
}
