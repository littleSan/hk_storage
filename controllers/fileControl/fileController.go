/*
*

	@author:
	@date : 2023/10/8
*/
package fileControl

import (
	"github.com/gin-gonic/gin"
	"hk_storage/common/logger"
	"hk_storage/common/response"
	"hk_storage/services/ipfs"
	"hk_storage/utils/validate"
	"net/http"
)

var _ Controller = (*controller)(nil)

type Controller interface {
	i()
	UploadImage(ctx *gin.Context)
	UploadFile(ctx *gin.Context)
	UploadVideo(ctx *gin.Context)
	UploadIpfs(ctx *gin.Context)
}
type controller struct {
	//AliOssServer aliOss.Service
	IpfsServer ipfs.Service
}

func New() *controller {
	return &controller{
		//AliOssServer: aliOss.New(),
		IpfsServer: ipfs.New(),
	}
}
func (c *controller) i() {
}

func (c *controller) UploadImage(ctx *gin.Context) {
	// 支持多文件上传
	multiForm, err := ctx.MultipartForm()
	if err != nil {
		logger.Info("解析上传文件参数出错")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	//获取文件
	fileSlice := []string{}
	f := multiForm.File["file"]
	if !validate.FilesValidate(f, validate.IMGAllow, int64(validate.ImageMaxSize)) {
		logger.Info("文件格式错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.FileFormatErr))
		return
	}
	fileSize := len(f)
	for _, multiFile := range f {
		logger.Info("上传文件名称", multiFile.Filename)
		reset, err := c.IpfsServer.Upload(multiFile)
		if err != nil {
			continue
		}
		if fileSize == 1 {
			ctx.JSON(http.StatusOK, response.SUCCESS(ctx, reset))
			return
		}
		fileSlice = append(fileSlice, reset)
	}

	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, fileSlice))
}
func (c *controller) UploadFile(ctx *gin.Context) {
	// 支持多文件上传
	multiForm, err := ctx.MultipartForm()
	if err != nil {
		logger.Info("解析上传文件参数出错")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	//获取文件
	fileSlice := []string{}
	f := multiForm.File["file"]
	if !validate.FilesValidate(f, validate.FileAllow, int64(validate.FileMaxSize)) {
		logger.Info("文件格式错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.FileFormatErr))
		return
	}

	fileSize := len(f)

	for _, multiFile := range f {
		logger.Info("上传文件名称", multiFile.Filename)
		reset, err := c.IpfsServer.Upload(multiFile)
		if err != nil {
			continue
		}
		if fileSize == 1 {
			ctx.JSON(http.StatusOK, response.SUCCESS(ctx, reset))
			return
		}
		fileSlice = append(fileSlice, reset)
	}
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, fileSlice))
}
func (c *controller) UploadVideo(ctx *gin.Context) {
	// 支持多文件上传
	multiForm, err := ctx.MultipartForm()
	if err != nil {
		logger.Info("解析上传文件参数出错")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	//获取文件
	fileSlice := []string{}
	f := multiForm.File["file"]
	fileSize := len(f)

	if !validate.FilesValidate(f, validate.VideoAllow, int64(validate.VideoMaxSize)) {
		logger.Info("文件格式错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.FileFormatErr))
		return
	}
	for _, multiFile := range f {
		logger.Info("上传文件名称", multiFile.Filename)
		reset, err := c.IpfsServer.Upload(multiFile)
		if err != nil {
			continue
		}
		if fileSize == 1 {
			ctx.JSON(http.StatusOK, response.SUCCESS(ctx, reset))
			return
		}
		fileSlice = append(fileSlice, reset)
	}
	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, fileSlice))
}

func (c *controller) UploadIpfs(ctx *gin.Context) {
	// 支持多文件上传
	multiForm, err := ctx.MultipartForm()
	if err != nil {
		logger.Info("解析上传文件参数出错")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}
	//获取文件
	fileSlice := []string{}
	f := multiForm.File["file"]
	if !validate.FilesValidate(f, validate.TotalAllow, int64(validate.VideoMaxSize)) {
		logger.Info("文件格式错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.FileFormatErr))
		return
	}
	fileSize := len(f)
	for _, multiFile := range f {
		logger.Info("上传文件名称", multiFile.Filename)
		reset, err := c.IpfsServer.Upload(multiFile)
		if err != nil {
			continue
		}
		if fileSize == 1 {
			ctx.JSON(http.StatusOK, response.SUCCESS(ctx, reset))
			return
		}
		fileSlice = append(fileSlice, reset)
	}

	ctx.JSON(http.StatusOK, response.SUCCESS(ctx, fileSlice))
}
