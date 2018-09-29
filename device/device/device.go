package device

import (
	"fmt"
	"net"
	"os"
	//"time"
)

//设备
type Device struct {
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
 * @Function 开启硬件对接服务器
 * @Auther Nelg
 */
func Start() {
	//获取配置
	config := getConfig()

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
		//记录连接
		device := Device{conn: tcpConn, heart: config.heart_time_out}
		//新建设备协程
		go device.deviceCoroutines()
	}
}

/**
 * @Function 设备协程
 * @Auther Nelg
 */
func (this *Device) deviceCoroutines() {
	//协程结束后关闭连接
	defer this.conn.Close()
	for {
		//设置连接存活时间
		//this.conn.SetReadDeadline(time.Now().Add(time.Duration(this.heart)))
		//读取数据
		readData := make([]byte, 2048)
		length, err := this.conn.Read(readData)
		if err != nil {
			break
		} else if length <= 0 {
			continue
		}
		//截取数据包
		dataArray := this.cutpack(readData[0:length])
		for _, value := range dataArray {
			fmt.Println(value)
		}
	}

}

/**
 * 数据截取
 */
func (this *Device) cutpack(content []byte) (data [][]byte) {
	//判断buffer是否为空
	if len(this.buffer) <= 0 {
		isContinue := false
		for _, value := range content {
			if value == 0x7e {
				isContinue = true
				break
			}
		}
		if isContinue == false {
			return
		}
	}
	//与上次未完整接收的buffer数据合并
	content = append(this.buffer, content...)
	maxIndex := len(content) - 1
	//数据切割
	dataArray := make([][]byte, 0)
	num := 0
	for index, value := range content {
		if value == 0x7e && index+1 <= maxIndex && content[index+1] != 0x7e {
			dataArray = append(dataArray, make([]byte, 0))
			dataArray[num] = append(dataArray[num], value)
		} else if value == 0x7e && index+1 <= maxIndex && content[index+1] == 0x7e {
			dataArray[num] = append(dataArray[num], value)
			dataArray = append(dataArray, make([]byte, 0))
			num++
		} else {
			dataArray[num] = append(dataArray[num], value)
		}
	}
	//命令筛选
	discard := make([]int, 0)
	arrMaxIndex = len(dataArray) - 1
	for index, cmd := range dataArray {
		signNum := 0
		for index, value := range cmd {
			if value == 0x7e {
				signNum++
			}
		}
		//标识位少于2，代表命令不全，加入buffer
		if signNum < 2 {
			if index == arrMaxIndex {
				this.buffer = cmd
			}
			//需要丢弃的数据
			discard = append(discard, index)
		} else if signNum > 2 {
			//TODO 待完善
			times := 0
			for i, v := range cmd {
				if v == 0x7e {
					times++
				}
				if times > 1 {
					dataArray[index] = cmd[i:]
					break
				}
			}
		}
	}
	//删除无用命令数据
	for _, index := range discard {

	}
	return
}
