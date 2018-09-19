package main

import (
	"fmt"
	"net"
)

//客户端对象
type Client struct {
	Channel chan string
	Name    string
	Address string
}

//客户端map集合
var onlineMap = make(map[string]Client)

//全局数据接收器
var message = make(chan string)

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.Channel {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Address + "]" + cli.Name + "：" + msg
	return
}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	cliAddress := conn.RemoteAddr().String()

	cli := Client{make(chan string), cliAddress, cliAddress}

	onlineMap[cliAddress] = cli

	fmt.Println(onlineMap)

	go WriteMsgToClient(cli, conn)

	message <- MakeMsg(cli, "I am login")

	//	for {
	//	}
}

func Manager() {
	for {
		msg := <-message
		for _, cli := range onlineMap {
			cli.Channel <- msg
		}
	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", ":7000")
	if err != nil {
		fmt.Println("net.Listen err ", err)
		return
	}
	defer listener.Close()

	//新开协程
	go Manager()

	//主协程
	for {
		conn, errConn := listener.Accept()
		if errConn != nil {
			fmt.Println("listener.Accept error ", errConn)
			continue
		}

		go HandleConn(conn)
	}
}
