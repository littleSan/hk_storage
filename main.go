package main

import (
	_ "github.com/go-sql-driver/mysql"
	"hk_storage/common/client"
	"hk_storage/common/logger"
	"hk_storage/common/rabbitMq"
	"hk_storage/common/rabbitMq/consumer"
	"hk_storage/common/redisCli"
	"hk_storage/core/app"
	"hk_storage/core/configs"
	"hk_storage/core/database"
)

func init() {
	//初始化数据库配置
	database.Register("mysql", configs.TomlConfig.Db.Dsn)
	//初始化日志信息
	logger.Initialize(configs.TomlConfig.Logs.LogLevel, configs.TomlConfig.Logs.LogPath, configs.TomlConfig.Logs.OutPut)
	//redis 注册
	redisCli.RegisterRedis()
	//eth
	client.GetEthClient()
	//rabbit
	rabbitMq.Connect()
	consumer.ConsumersInit()
}

//主启动类

func main() {
	app.Run()
}
