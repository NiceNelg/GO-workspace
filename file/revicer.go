package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func RevicerFile(conn net.Conn, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("create file error", err)
		return
	}
	buf := make([]byte, 1024*4)
	var n int
	for {
		n, err = conn.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("File revice offer")
			return
		} else if err != nil {
			fmt.Println("File revice error ", err)
			return
		}
		if n == 0 {
			fmt.Println("File revice offer")
			return
		}
		f.Write(buf[:n])
	}
}

func main() {
	//监听
	listenner, err := net.Listen("tcp", "127.0.0.1:7000")
	if err != nil {
		return
	}
	defer listenner.Close()

	conn, err1 := listenner.Accept()
	if err != nil {
		fmt.Println("listen Accept error ", err1)
		return
	}

	buf := make([]byte, 1024)
	var n int
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Read file error ", err)
		return
	}

	filename := string(buf[:n])

	conn.Write([]byte("OK"))

	RevicerFile(conn, filename)
}
