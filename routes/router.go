package routes

import (
	"github.com/gin-gonic/gin"
	"hk_storage/controllers/fileControl"
	"hk_storage/controllers/payControl"
	sysUser2 "hk_storage/controllers/sysUser"
	"hk_storage/middleware"
)

func InitRouter(engine *gin.Engine) {
	middleware := middleware.New()
	//短信controller
	//smsCon := sms.New()
	//用户
	sysUser := sysUser2.New()
	// 文件
	fileCon := fileControl.New()
	//支付
	payCon := payControl.New()

	//登陆注册不需要权限信息
	sysGroup := engine.Group("/api/sys", middleware.Cors())
	{
		sysGroup.POST("/login", sysUser.Login)
		sysGroup.POST("/saveFile", sysUser.SaveUserFile)
		sysGroup.GET("/fileList", sysUser.UserFile)
		sysGroup.POST("/draw/ylem", sysUser.DrawYlem)
		sysGroup.POST("/coin/pay", payCon.PayBaseCoin)
		sysGroup.DELETE("/delete/file", sysUser.DeleteUserFile)
		//sysGroup.GET("/sendMsg", smsCon.SendMsg)
	}

	fileGroup := engine.Group("/api/file", middleware.Cors())
	{
		fileGroup.POST("/uploadImage", fileCon.UploadImage)
		fileGroup.POST("/uploadFile", fileCon.UploadFile)
		fileGroup.POST("/uploadVideo", fileCon.UploadVideo)
		fileGroup.POST("/uploadIpfs", fileCon.UploadIpfs)
	}

}
