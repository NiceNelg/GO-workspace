package src

import (
	"core/config"
	"core/handle"
	"core/model"
	"core/redispool"
	"core/server"
)

func main() {

	/*获取配置*/
	allConfig := config.GetConfig()

	/*初始化redis*/
	redisPool := redispool.NewPool(
		allConfig.RedisIp+":"+allConfig.RedisPort,
		allConfig.RedisPwd,
		allConfig.RedisDB,
	)

	/*初始化数据库*/
	db := model.Init(
		allConfig.MysqlUsername,
		allConfig.MysqlPwd,
		allConfig.MysqlHost,
		allConfig.MysqlPort,
		allConfig.MysqlDatabase,
	)
	defer db.Close()

	/*实例化数据处理对象*/
	handler := handle.Init(db, redisPool, &allConfig)

	/*开启数据处理队列*/
	handler.StartHandle()

	/*初始化服务*/
	serv := server.Init(&allConfig, handler)

	/*服务开始*/
	serv.Start()
}
