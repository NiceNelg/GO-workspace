package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("open file error ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*4)

	//读文件内容
	for {
		n, err := f.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("File send offer")
			return
		} else if err != nil {
			return
		}
		//发送内容
		conn.Write(buf[:n])
	}
}

func main() {
	fmt.Println("请输入文件名：")

	var path string
	fmt.Scan(&path)

	//获取文件名
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("error = ", err, "\n")
		return
	}

	//连接服务器
	conn, errorConn := net.Dial("tcp", "127.0.0.1:7000")
	if errorConn != nil {
		fmt.Println("net dial error ", errorConn, "\n")
		return
	}
	defer conn.Close()

	_, errorWrite := conn.Write([]byte(info.Name()))
	if errorWrite != nil {
		fmt.Println("write error ", errorWrite, "\n")
		return
	}

	//接收对方的回复
	var n int
	buf := make([]byte, 1024)

	n, errorRead := conn.Read(buf)
	if errorRead != nil {
		fmt.Println("connect read err ", err)
		return
	}

	if "OK" != string(buf[:n]) {
		fmt.Println("server error")
	} else {
		SendFile(path, conn)
	}
}
