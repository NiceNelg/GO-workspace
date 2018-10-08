package main

import (
	"../lib/config"
	"../lib/handle"
	"../lib/redispool"
	"../lib/server"
)

func main() {
	//获取配置
	allConfig := config.GetConfig()

	//初始化redis
	redisPool := redispool.NewPool(allConfig.RedisIp+":"+allConfig.RedisPort, allConfig.RedisPwd, allConfig.RedisDB)

	//开启数据处理队列
	handleObj := handle.Init(redisPool, allConfig)
	handleObj.StartHandle()

	//初始化连接
	serv := server.Init(allConfig, redisPool)
	serv.Start()
}
