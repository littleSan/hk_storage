/*
@author: little
@date : 2023/9/26
*/
package sms

import (
	"github.com/gin-gonic/gin"
	"hk_storage/common/logger"
	"hk_storage/common/redisCli/smsRedis"
	"hk_storage/common/response"
	"hk_storage/services/sms"
	"hk_storage/services/sms/ems"
	"hk_storage/utils/smsCode"
	"hk_storage/utils/validate"
	"net/http"
	"strconv"
)

var _ Controller = (*controller)(nil)

type Controller interface {
	i()
	SendMsg(ctx *gin.Context)
}

type controller struct {
	SmsService sms.Service
	SmsRedis   smsRedis.Service
	EmsService ems.Service
}

func New() *controller {
	return &controller{
		SmsService: sms.New(),
		SmsRedis:   smsRedis.New(),
		EmsService: ems.New(),
	}
}
func (c *controller) i() {

}
func (c *controller) SendMsg(ctx *gin.Context) {
	account := ctx.DefaultQuery("account", "")
	smsType := ctx.DefaultQuery("Type", "0")
	sm, err := strconv.Atoi(smsType)
	if err != nil {
		logger.Info("信息类型错误")
		ctx.JSON(http.StatusOK, response.Failure(ctx, response.ParamError))
		return
	}

	b1 := validate.ValidatePhoneNumber(account)
	if b1 {
		code, err := c.SmsRedis.GetSmsCode(account)
		if err != nil {
			code = smsCode.GetSmsCode()
		}
		res := c.SmsService.SendLoginMsg(sm, account, code)
		logger.Info("发送短信验证码account{},code{},status{}", account, code, res)
		if res {
			c.SmsRedis.SetSmsCode(account, code)
			ctx.JSON(http.StatusOK, response.SUCCESS(ctx, nil))
		} else {
			ctx.JSON(http.StatusOK, response.Failure(ctx, response.Fail))
		}
	}
	b2 := validate.ValidateEmail(account)
	if b2 {
		code, err := c.SmsRedis.GetSmsCode(account)
		if err != nil {
			code = smsCode.GetSmsCode()
		}
		res := c.EmsService.SendMail(code, account)
		logger.Info("邮箱发送验证码account{},code{},status{}", account, code, res)
		if res {
			c.SmsRedis.SetSmsCode(account, code)
			ctx.JSON(http.StatusOK, response.SUCCESS(ctx, nil))
		} else {
			ctx.JSON(http.StatusOK, response.Failure(ctx, response.Fail))
		}
	}

}
