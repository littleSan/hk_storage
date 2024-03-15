package app

import (
	"github.com/gin-gonic/gin"
	"hk_storage/common/constants"
	"hk_storage/common/task/orderTask"
	"hk_storage/core/configs"
	"hk_storage/routes"
)

func Run() {
	app := gin.Default()
	//设置路由信息
	routes.InitRouter(app)
	//启动定时任务
	go orderTask.NewCronJob(constants.TaskTimeUint, orderTask.ConfirmationOrderBack)
	app.Run(configs.TomlConfig.HttpAddress)
}
