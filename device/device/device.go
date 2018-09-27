package device

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/aWildProgrammer/fconf"
)

//硬件服务器配置信息
type DeviceConfig struct {
	stdlog     string
	errlog     string
	redisIp    string
	redisPort  string
	handleList string
	sendList   string
	resendList string
	serverIp   string
	serverPort string
	heart_time_out int
}

//设备
type Device struct {
	//设备连接号
	conn *net.TCPConn
	//设备编号
	id string
	//未完成数据包
	buffer byte
	//心跳时间
	heart chan int
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
	config.heart_time_out = allConfig.String("device.heart_time_out")
	config.redisIp = allConfig.String("redis.ip")
	config.redisPort = allConfig.String("redis.port")
	config.handleList = allConfig.String("list.handle")
	config.sendList = allConfig.String("list.send")
	config.resendList = allConfig.String("list.resend")
	config.serverIp = allConfig.String("server.ip")
	config.serverPort = allConfig.String("server.port")
	return
}

/**
 * @Function 开始硬件对接服务器
 * @Auther Nelg
 */
func Start() {
	//获取配置
	config := getConfig()
	fmt.Println(stdconfig)

	//建立TCP服务器
	tcpAddr, err := net.ResolveTCPAddr("tcp", config.serverIp+":"+config.serverPort)
	if err != nil {
		os.Exit(-1)
	}
	//监听端口
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	for {
		//接受连接请求
		tcpConn, _ := tcpListener.AcceptTCP()
		//新建设备协程
		go deviceCoroutines(tcpConn)
	}
}

/**
 * @Function 设备协程
 * @Auther Nelg
 */
func deviceCoroutines(deviceConn *net.TCPConn) {
	//获取配置
	config := getConfig()
	//记录连接
	device := Device{conn: deviceConn,heart:config.heart_time_out}
	//协程结束后关闭连接
	defer device.conn.Close()
	for {
		select {
			case <-time.After(device.heart*time.Second) {
				return
			}
		}
		//读取数据
		data := make([]byte)
		length, err := device.conn.Read(data)
		if err != nil {
			break
		}
		//截取数据包
		//device.buffer, data :=
	}

}
