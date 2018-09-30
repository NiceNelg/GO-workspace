package server

import (
	"fmt"
	"net"
	"os"
	"time"

	"../config"
	"../handle"
	"../protocol808"
)

//服务器
type Server struct {
	ip     string
	port   string
	config config.Config
}

//设备
type Client struct {
	//设备连接号
	conn *net.TCPConn
	//设备编号
	id string
	//未完成数据包
	buffer []byte
	//心跳时间
	heart int
}

/**
 * @Function 初始化服务
 * @Auther Nelg
 */
func Init() (serv Server) {
	//获取配置
	serverConfig := config.GetConfig()
	serv = Server{ip: serverConfig.ServerIp, port: serverConfig.ServerPort, config: serverConfig}
	return
}

/**
 * @Function 开启硬件对接服务器
 * @Auther Nelg
 */
func (this *Server) Start() {

	//建立TCP服务器
	tcpAddr, err := net.ResolveTCPAddr("tcp", this.ip+":"+this.port)
	if err != nil {
		fmt.Println("Start server error：", err)
		os.Exit(-1)
	}
	//监听端口
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		//接受连接请求
		tcpConn, _ := tcpListener.AcceptTCP()
		//记录连接
		cli := Client{conn: tcpConn, heart: this.config.HeartTimeOut}
		//新建设备协程
		go cli.deviceCoroutines()
	}
}

/**
 * @Function 设备协程
 * @Auther Nelg
 */
func (this *Client) deviceCoroutines() {
	//协程结束后关闭连接
	defer this.conn.Close()
	for {
		//设置连接存活时间
		this.conn.SetReadDeadline(time.Now().Add(time.Duration(this.heart) * time.Second))
		//读取数据
		readData := make([]byte, 2048)
		length, err := this.conn.Read(readData)
		if err != nil {
			break
		} else if length <= 0 {
			continue
		}
		//截取数据包
		var dataArray [][]byte
		dataArray, this.buffer = protocol808.Cutpack(readData[0:length], this.buffer)
		//存入处理队列
		go handle.SaveTask(dataArray)
	}

}
