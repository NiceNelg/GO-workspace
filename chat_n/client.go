package main

import (
	"fmt"
	"net"
)

func main() {
	block := make(chan string)
	//连接服务器
	conn, errorConn := net.Dial("tcp", "127.0.0.1:7000")
	if errorConn != nil {
		fmt.Println("net dial error ", errorConn, "\n")
		return
	}
	defer conn.Close()

	go func(conn net.Conn) {
		for {
			var n int
			buf := make([]byte, 1024)
			n, errorRead := conn.Read(buf)
			if errorRead != nil {
				fmt.Println("connect read error ", errorRead)
				return
			}
			fmt.Println(string(buf[:n]))
			block <- string(buf[:n])
		}
	}(conn)

	for {
		_, errorWrite := conn.Write([]byte("test"))
		if errorWrite != nil {
			fmt.Println("write error ", errorWrite, "\n")
			return
		}
		<-block
	}
}
