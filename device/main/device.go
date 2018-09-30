package main

import (
	"../lib/server"
)

func main() {
	//初始化连接
	serv := server.Init()
	serv.Start()
}
