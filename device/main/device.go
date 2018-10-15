package main

import (
	"../lib/config"
	"../lib/handle"
	"../lib/model"
	"../lib/redispool"
	"../lib/server"
)

func main() {
	//获取配置
	allConfig := config.GetConfig()

	//初始化redis
	redisPool := redispool.NewPool(
		allConfig.RedisIp+":"+allConfig.RedisPort,
		allConfig.RedisPwd,
		allConfig.RedisDB,
	)

	//初始化数据库
	db := model.Init(
		allConfig.MysqlUsername,
		allConfig.MysqlPwd,
		allConfig.MysqlHost,
		allConfig.MysqlPort,
		allConfig.MysqlDatabase,
	)
	defer db.Close()

	//实例化数据处理对象
	handleObj := handle.Init(db, redisPool, allConfig)
	//开启数据处理队列
	handleObj.StartHandle()

	//初始化连接
	serv := server.Init(allConfig, handleObj)
	serv.Start()
}
